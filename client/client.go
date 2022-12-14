package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	token "github.com/luczito/auctionhouse/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

// client struct.
type Client struct {
	token.UnimplementedClientServer
	address       string              //own address
	primaryClient token.ManagerClient //pointer to the primary manager
	connection    *grpc.ClientConn
	ctx           context.Context
}

// heartbeat function to send replies to the primary managers heartbeats
func (c *Client) Heartbeat(ctx context.Context, prim *token.Primary) (*token.Reply, error) {
	var addr string

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		addr = md.Get("address")[0]
	}

	log.Printf("Recieved heartbeat from primary %v", addr)

	if addr != c.address {
		c.address = addr

		conn, err := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
		if err != nil {
			log.Fatalf("Did not connect %v", err)
		}

		c.connection = conn
		c.primaryClient = token.NewManagerClient(conn)
	}

	return &token.Reply{}, nil
}

func main() {
	//command line flags
	address := flag.String("primary", "127.0.0.1:5002", "primary ip")
	startPort := flag.Int("port", 5000, "port")

	flag.Parse()

	//setup ip
	ip := fmt.Sprintf("127.0.0.1:%d", *startPort)
	ctx, _ := context.WithCancel(context.Background())
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("address", ip))

	//setup logger for client
	log_, err := os.OpenFile(fmt.Sprintf("client-log-%s.txt", ip), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer log_.Close()

	log.SetOutput(log_)

	//dial up the primary manager
	conn, err := grpc.DialContext(ctx, *address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect to client: %v", err)
	}

	//create the client struct
	c := &Client{
		address:       *address,
		primaryClient: token.NewManagerClient(conn),
		connection:    conn,
		ctx:           ctx,
	}

	//listen on the ip
	list, err := net.Listen("tcp", ip)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	token.RegisterClientServer(grpcServer, c)

	go func() {
		log.Printf("\nlistening on %s\n", ip)
		// fmt.Printf("\nlistening on %s\n", ip)

		if err := grpcServer.Serve(list); err != nil {
			log.Fatalf("failed to listen on: %v", err)
		}
	}()

	//print the commands to the terminal / user
	fmt.Println("\"bid <price>\" - to bid on an item")
	fmt.Println("\"result\" - to get the result of the auction")
	fmt.Println("\"exit\" - to exit the program")
	fmt.Printf("Enter command: ")

	scanner := bufio.NewScanner(os.Stdin)
	//scanner loop, infinite
	for scanner.Scan() {
		command := scanner.Text()

		//split command
		args := strings.Split(command, " ")

		//check which command has been entered
		switch args[0] {
		//case bid, answer accordingly
		case "bid":
			bid, err := strconv.Atoi(args[1])
			if err != nil {
				log.Println("Invalid bid price")
				fmt.Println("Invalid bid price")
			}

			response, err := c.primaryClient.Bid(ctx, &token.Amount{Value: int32(bid)})
			if err != nil {
				log.Println("Error sending bid")
				fmt.Println("Error sending bid")
			}

			if response.Status == token.Status_SUCCESS {
				log.Printf("Your bid: %d was successful\n", bid)
				fmt.Printf("Your bid: %d was successful\n", bid)
			} else {
				log.Println("Bid unsuccessful - bid was too low")
				fmt.Println("Bid unsuccessful - bid was too low")
			}
			//case result answer accordingly
		case "result":
			response, err := c.primaryClient.Result(ctx, &token.Void{})
			if err != nil {
				log.Println("Error getting result")
			}

			switch v := response.Outcome.(type) {
			case *token.Outcome_AuctionResult:
				log.Printf("Auction is over! The winning bid was %d and made by %s\n", v.AuctionResult.Amount, v.AuctionResult.Id)
				fmt.Printf("Auction is over! The winning bid was %d and made by %s\n", v.AuctionResult.Amount, v.AuctionResult.Id)
			case *token.Outcome_CurrentBid:
				log.Printf("Auction is still ongoing, %d seconds remaining. Currently, the highest bid is %d\n", response.TimeRemaining, v.CurrentBid.Amount)
				fmt.Printf("Auction is still ongoing, %d seconds remaining. Currently, the highest bid is %d\n", response.TimeRemaining, v.CurrentBid.Amount)
			}
			//exit terminates the client.
		case "exit":
			os.Exit(0)
		}

		fmt.Print("Enter command: ")
		log.Print("Enter command: ")
	}
}

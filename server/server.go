package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	token "github.com/luczito/auctionhouse/proto"
	"google.golang.org/grpc"
)

var port = flag.Int("port", 5000, "port") //port for the node default 5000

// node struct.
type Server struct {
	token.UnimplementedAuctionServer
	id            int32 //port
	primaryId     int32 //primary server port
	clients       map[int32]token.AuctionClient
	servers       map[int32]token.AuctionClient
	ctx           context.Context
	currentBid    Bid
	timeRemaining int32
}

type Bid struct {
	id     int32
	amount int32
}

func main() {
	f, err := os.OpenFile("log.server", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v\n", err)
	}
	defer f.Close()
	log.SetOutput(f)

	flag.Parse() //set port with -port in the commandline when running the program

	ctx_, cancel := context.WithCancel(context.Background())
	defer cancel()

	var port = int32(*port)

	//create a node for this proccess
	n := &Server{
		id:      port,
		clients: make(map[int32]token.AuctionClient),
		servers: make(map[int32]token.AuctionClient),
		ctx:     ctx_,
	}

	//creates listener on port
	list, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("Failed to listen on port: %v\n", err)
	}

	log.Printf("Node created on port: %d\n", n.id)
	fmt.Printf("Node created on port %v\n", n.id)

	grpcServer := grpc.NewServer()
	token.RegisterAuctionServer(grpcServer, n)

	//serve on the listener
	go func() {
		if err := grpcServer.Serve(list); err != nil {
			log.Fatalf("failed to server %v\n", err)
		}
	}()

	//for loop to dial all other nodes in the network, if this loop is increased the number of nodes in the network is aswell
	for i := 0; i < 3; i++ {
		nodePort := int32(5000 + i)
		if nodePort == n.id {
			continue
		}
		if nodePort < 5003 {
			var conn *grpc.ClientConn
			log.Printf("Trying to dial: %v\n", nodePort)
			conn, err := grpc.Dial(fmt.Sprintf(":%v", nodePort), grpc.WithInsecure(), grpc.WithBlock())

			if err != nil {
				log.Fatalf("Could not connect: %v\n", err)
			}

			defer conn.Close()
			log.Printf("Succes connecting to: %v\n", nodePort)
			c := token.NewAuctionClient(conn)
			n.servers[nodePort] = c
		} // else {
		// 	var conn *grpc.ClientConn
		// 	log.Printf("Trying to dial: %v\n", nodePort)
		// 	conn, err := grpc.Dial(fmt.Sprintf(":%v", nodePort), grpc.WithInsecure(), grpc.WithBlock())

		// 	if err != nil {
		// 		log.Fatalf("Could not connect: %v\n", err)
		// 	}

		// 	defer conn.Close()
		// 	log.Printf("Succes connecting to: %v\n", nodePort)
		// 	c := token.NewAuctionClient(conn)
		// 	n.clients[nodePort] = c
		// }
	}

	log.Printf("%v is connected to %v other nodes\n", n.id, len(n.clients))
	fmt.Printf("%v is connected to %v other nodes\n", n.id, len(n.clients))

	scanner := bufio.NewScanner(os.Stdin)

	//scanner that requests access from the given node when something is written in the terminal.
	for scanner.Scan() {

	}
}
func (s *Server) Election(ctx context.Context, input *token.ElectionRequest) (*token.Ack_, error) {
	s.CallElection()

	return &token.Ack_{}, nil
}

func (s *Server) CallElection() {
	request := &token.ElectionRequest{
		Id: s.id,
	}
	votes := len(s.servers)

	for id, server := range s.servers {
		if id > s.id {
			response, err := server.Election(s.ctx, request)

			time.Sleep(time.Second * 5)

			if response == nil {
				log.Printf("Server %v has crashed\n", id)
				votes--
			}
			if err != nil {
				log.Printf("Something went wrong with server: %v, %v\n", id, err)
				votes--
			}
		}
	}
	if votes == 0 {
		s.primaryId = s.id
		s.SendCoordination()
	}
}

func (s *Server) reply() {

}

func (s *Server) SendCoordination() {
	coordination := &token.Coord{
		Id: s.id,
	}

	for id, server := range s.servers {
		_, err := server.Coordination(s.ctx, coordination)

		if err != nil {
			log.Printf("Something went wrong with server: %v, %v\n", id, err)
			delete(s.servers, id)
		}
	}
}

func (s *Server) Coordinate(ctx context.Context, input *token.Coord) (*token.Ack_, error) {
	port := input.Id

	s.primaryId = port

	wasLeader := (s.id == s.primaryId && s.id != port)

	if _, ok := s.servers[port]; !ok {
		log.Println("Saving new client")

		conn, err := grpc.Dial(fmt.Sprintf(":%v", port), grpc.WithInsecure(), grpc.WithBlock())

		if err != nil {
			log.Fatalf("could not reconnect: %v", err)
		}

		defer conn.Close()

		c := token.NewAuctionClient(conn)
		s.servers[port] = c
	}

	if wasLeader {
		ports := make([]int32, len(s.clients))

		for port := range s.clients {
			ports = append(ports, port)
		}

		s.servers[port].Update(ctx, &token.Data{
			CurrentBid: &token.CurrentBid{
				Id:     s.currentBid.id,
				Amount: s.currentBid.amount,
			},
			TimeRemaining: s.timeRemaining,
			Clients:       ports,
		})
	}

	return &token.Ack_{}, nil
}

func (s *Server) Update(ctx context.Context, data *token.Data) (*token.Ack_, error) {
	log.Printf("Recieved an update for internal data")

	s.currentBid.id = data.CurrentBid.Id
	s.currentBid.amount = data.CurrentBid.Amount
	s.timeRemaining = data.TimeRemaining

	for _, port := range data.Clients {
		conn, err := grpc.Dial(fmt.Sprintf(":%v", port), grpc.WithInsecure(), grpc.WithBlock())

		if err != nil {
			continue
		}

		// defer conn.Close()

		c := token.NewAuctionClient(conn)

		s.clients[port] = c
	}

	return &token.Ack_{}, nil
}

func (s *Server) UpdateServers() error {
	ports := make([]int32, 0, len(s.clients))

	for port := range s.clients {
		ports = append(ports, port)
	}

	for _, server := range s.servers {
		server.Update(s.ctx, &token.Data{
			CurrentBid: &token.CurrentBid{
				Id:     s.currentBid.id,
				Amount: s.currentBid.amount,
			},
			TimeRemaining: s.timeRemaining,
			Clients:       ports,
		})
	}
	return nil
}

func (s *Server) Bid(ctx context.Context, input *token.Amount) (*token.Ack, error) {
	var message string
	port := input.Id

	if _, ok := s.clients[port]; !ok {
		conn, err := grpc.Dial(fmt.Sprintf(":%v", port), grpc.WithInsecure(), grpc.WithBlock())

		if err != nil {
			log.Fatalf("could not reconnect: %v", err)
		}

		// defer conn.Close()

		c := token.NewAuctionClient(conn)
		s.clients[port] = c
	}

	if input.Amount > s.currentBid.amount {
		s.currentBid.amount = input.Amount
		s.currentBid.id = input.Id
		message = "Bid accepted as the new highest bid."
		if err := s.UpdateServers(); err != nil {
			message = "Database failure"
		}
	} else {
		message = "Bid rejected, lower than the current highest bid."
	}

	reply := &token.Ack{
		Status: message,
	}

	return reply, nil
}

func (s *Server) Result(ctx context.Context, input *token.AuctionResult) (*token.Response, error) {
	reply := &token.Response{}
	if s.timeRemaining == 0 {
		reply.Outcome = &token.Response_AuctionResult{
			AuctionResult: &token.AuctionResult{
				Id:     s.currentBid.id,
				Amount: s.currentBid.amount,
			},
		}
	} else {
		reply.Outcome = &token.Response_CurrentBid{
			CurrentBid: &token.CurrentBid{
				Id:     s.currentBid.id,
				Amount: s.currentBid.amount,
			},
		}
	}

	return reply, nil
}

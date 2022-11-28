package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	token "github.com/luczito/auctionhouse/proto"
	"google.golang.org/grpc"
)

var port = flag.Int("port", 5000, "port") //port for the node default 5000

type STATE int32 //state struct to see wether the node has access, wants access or not.

const (
	RELEASED STATE = iota
	HELD
	REQUESTED
)

// node struct.
type Client struct {
	token.UnimplementedAuctionServer
	id       int32 //port
	clients  map[int32]token.AuctionClient
	server   token.AuctionServer
	ctx      context.Context
	queue    []int32
	state    STATE
	requests int
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
	n := &Client{
		id:       port,
		clients:  make(map[int32]token.AuctionClient),
		queue:    make([]int32, 0),
		ctx:      ctx_,
		state:    RELEASED,
		requests: 0,
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
	for i := 0; i < 6; i++ {
		nodePort := int32(5000 + i)

		if nodePort == n.id {
			continue
		}

		var conn *grpc.ClientConn
		log.Printf("Trying to dial: %v\n", nodePort)
		conn, err := grpc.Dial(fmt.Sprintf(":%v", nodePort), grpc.WithInsecure(), grpc.WithBlock())

		if err != nil {
			log.Fatalf("Could not connect: %v\n", err)
		}

		defer conn.Close()
		log.Printf("Succes connecting to: %v\n", nodePort)
		c := token.NewAuctionClient(conn)
		n.clients[nodePort] = c
	}

	log.Printf("%v is connected to %v other nodes\n", n.id, len(n.clients))
	fmt.Printf("%v is connected to %v other nodes\n", n.id, len(n.clients))

	scanner := bufio.NewScanner(os.Stdin)

	//scanner that requests access from the given node when something is written in the terminal.
	for scanner.Scan() {

	}
}

package server

import (
	"flag"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	token "github.com/luczito/auctionhouse/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

//server file, holds the server struct and methods for the servers (multiple).

var name = flag.String("name", "localhost", "the server name")
var port = flag.String("port", "8080", "the server port")

type Server struct {
	token.UnimplementedAuctionServer
	isPrimary   bool
	port        string
	name        string
	clients     map[string]token.Auction_ConnectServer
	clientNames map[int32]int32
	mutex       sync.Mutex
	auctions    map[string]string
}

func createAuction() {
	//creates a new aution from commandline.
}

func checkAuctions() {
	//returns a list of all auctions running.
}

func updateAution() {
	//updates an aution with the new bid
}

func endAuction() {
	//ends aution after time runs out
}

func (s *Server) Connect(stream token.Auction_ConnectServer) error {
	log.Printf("Participant connected.")

	name, err := stream.Recv()
	if err != nil {
		log.Printf("Unable to recieve the username %v", err)
	}

	peer_, _ := peer.FromContext(stream.Context())
	s.clients[peer_.Addr.String()] = stream
	log.Printf("User: %v, connected to the chat via %v\n", len(s.clientNames)+1, peer_.Addr.String())

	s.clientNames[int32(len(s.clientNames)-1)] = int32(len(s.clientNames) + 1)

	message := &token.ServerMsg{
		Name: "Server message",
		Msg:  fmt.Sprintf("%v just joined!", name.Msg),
	}

	log.Printf("Sending new participant connection to all participants.")

	for _, client := range s.clients {
		client.Send(message)
	}

	for {
		log.Printf("Listening for messages")
		msg, err := stream.Recv()
		if err != nil {
			if status.Code(err).String() == "Cancelled" || status.Code(err).String() == "EOF" {
				log.Printf("%v disconnected from the chat via %v\n", name.Msg, peer_.Addr.String())
				break
			} else {
				log.Printf("Unable to recieve the message %v", err)
				break
			}
		}

		message := &token.ServerMsg{
			Name: string(s.clientNames[s.clientNames[int32(len(s.clientNames)-1)]]),
			Msg:  msg.Msg,
		}

		log.Printf("Send message to all other participants.")
		for _, client := range s.clients {
			client.Send(message)
		}
	}
	leaveMessage := &token.ServerMsg{
		Name: "Server",
		Msg:  fmt.Sprintf("%v left the chat", string(s.clientNames[s.clientNames[int32(len(s.clientNames)-1)]])),
	}

	log.Printf("Send leave msg to all other participants\n")

	for _, client := range s.clients {
		client.Send(leaveMessage)
	}

	delete(s.clients, peer_.Addr.String())
	delete(s.clientNames, s.clientNames[int32(len(s.clientNames))-1])
	return nil
}

func start(s *Server) {
	listener, err := net.Listen("tcp", s.name+":"+s.port)
	if err != nil {
		log.Fatalf("Unable to create server %v", err)
	}

	grpcServer := grpc.NewServer()

	log.Printf("Creating server on port %v", s.port)

	token.RegisterAuctionServer(grpcServer, s)
	serveErr := grpcServer.Serve(listener)
	if serveErr != nil {
		log.Fatalf("Unable to start the server %v", serveErr)
	}
}

func main() {
	flag.Parse()

	server := &Server{
		name:        *name,
		port:        *port,
		clients:     make(map[string]token.Auction_ConnectServer),
		clientNames: make(map[int32]int32),
	}

	go start(server)
	for {
		time.Sleep(500 * time.Millisecond)
	}
}

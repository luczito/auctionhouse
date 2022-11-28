package main

import (
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

// server struct.
type Server struct {
	token.UnimplementedAuctionServer
	id                  int32 //port
	primaryId           int32 //primary server port
	clients             map[int32]token.AuctionClient
	servers             map[int32]token.AuctionClient
	ctx                 context.Context
	currentBid          Bid
	timeRemaining       int32
	heartbeatTimeout    chan bool
	coordinationTimeout chan bool
	expectingAnswer     bool
}

// bid struct
type Bid struct {
	id     int32
	amount int32
}

func main() {

	flag.Parse() //set port with -port in the commandline when running the program

	ctx_, cancel := context.WithCancel(context.Background())
	defer cancel()

	var port = int32(*port)

	//creates a log file for each server
	f, err := os.OpenFile(fmt.Sprintf("log-%d.txt", port), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v\n", err)
	}
	defer f.Close()
	log.SetOutput(f)

	//create a server for this proccess
	s := &Server{
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

	log.Printf("Node created on port: %d\n", s.id)
	fmt.Printf("Node created on port %v\n", s.id)

	grpcServer := grpc.NewServer()
	token.RegisterAuctionServer(grpcServer, s)

	//serve on the listener
	go func() {
		if err := grpcServer.Serve(list); err != nil {
			log.Fatalf("failed to server %v\n", err)
		}
	}()

	//for loop to dial all other servers in the network, if this loop is increased the number of servers have to be increased aswell
	for i := 0; i < 3; i++ {
		nodePort := int32(5000 + i)
		if nodePort == s.id {
			continue
		}
		if nodePort < int32(5003) {
			var conn *grpc.ClientConn
			log.Printf("Trying to dial: %v\n", nodePort)
			conn, err := grpc.Dial(fmt.Sprintf(":%v", nodePort), grpc.WithInsecure(), grpc.WithBlock())

			if err != nil {
				log.Fatalf("Could not connect: %v\n", err)
			}

			defer conn.Close()
			log.Printf("Succes connecting to: %v\n", nodePort)
			c := token.NewAuctionClient(conn)
			s.servers[nodePort] = c
		}
	}

	log.Printf("%v is connected to %v other servers\n", s.id, len(s.servers))
	fmt.Printf("%v is connected to %v other servers\n", s.id, len(s.servers))

	s.CallElection()
	//heartbeat loop
	for {
		time.Sleep(time.Second * 5)
		s.loop()
	}
}

// election listener.
func (s *Server) Election(ctx context.Context, input *token.ElectionRequest) (*token.Ack_, error) {
	log.Printf("Recieved election call, calling election\n")
	s.CallElection()

	return &token.Ack_{}, nil
}

// calls an election to all nodes with a higher id than oneself. then calls coordination if the node is the highest id
func (s *Server) CallElection() {
	log.Printf("Sending election to higher nodes(servers)\n")
	request := &token.ElectionRequest{
		Id: s.id,
	}
	votes := len(s.servers)

	for id, server := range s.servers {
		if id > s.id {
			response, err := server.Election(s.ctx, request)

			if response == nil {
				log.Printf("Server %v has crashed\n", id)
				votes--
			}
			if err != nil {
				log.Printf("Something went wrong with server: %v, %v\n", id, err)
				votes--
			}
		} else if id < s.id {
			votes--
		}
	}
	if votes == 0 {
		s.primaryId = s.id
		s.SendCoordination()
	} else {
		s.CoordinationTimeout(10)
	}
}

// func to timeout the wait for a coordination from another node.
func (s *Server) CoordinationTimeout(seconds int) {
	s.expectingAnswer = true

	select {
	case <-s.coordinationTimeout:
		log.Println("New primary lives")
	case <-time.After(time.Duration(seconds) * time.Second):
		// primary dead
		log.Println("Coordination timeout, call new election")
		s.CallElection()
	}

	s.expectingAnswer = false
}

// sends coordination to other servers in the network
func (s *Server) SendCoordination() {
	log.Printf("Sending coordination to backups\n")
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

// listener for sendcoordination function
func (s *Server) Coordination(ctx context.Context, input *token.Coord) (*token.Ack_, error) {
	port := input.Id

	s.primaryId = port

	wasLeader := (s.id == s.primaryId && s.id != port)

	if _, ok := s.servers[port]; !ok {
		log.Println("Saving new client")

		conn, err := grpc.Dial(fmt.Sprintf(":%v", port), grpc.WithInsecure(), grpc.WithBlock())

		if err != nil {
			log.Fatalf("could not reconnect: %v\n", err)
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

// listens and recieves updates from primary server.
func (s *Server) Update(ctx context.Context, data *token.Data) (*token.Ack_, error) {
	log.Printf("Recieved an update for internal data\n")

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

// updates the other servers in the network
func (s *Server) UpdateServers() error {
	log.Printf("Sending update to backups\n")
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

// listens and registers a bid from a client.
func (s *Server) Bid(ctx context.Context, input *token.Amount) (*token.Ack, error) {
	var message string
	port := input.Id

	if _, ok := s.clients[port]; !ok {
		conn, err := grpc.Dial(fmt.Sprintf(":%v", port), grpc.WithInsecure(), grpc.WithBlock())

		if err != nil {
			log.Fatalf("could not reconnect: %v\n", err)
		}

		// defer conn.Close()

		c := token.NewAuctionClient(conn)
		s.clients[port] = c
	}

	if input.Amount > s.currentBid.amount {
		s.currentBid.amount = input.Amount
		s.currentBid.id = input.Id
		message = "Bid accepted as the new highest bid."
		log.Printf("Bid accepted as the new highest bid.\n")
		if err := s.UpdateServers(); err != nil {
			message = "Database failure"
			log.Fatalf("Database failure\n")
		}
	} else {
		log.Printf("Bid rejected, lower than the current highest bid.\n")
		message = "Bid rejected, lower than the current highest bid."
	}

	reply := &token.Ack{
		Status: message,
	}

	return reply, nil
}

// when the auction is over this method will return the winner and the amount
func (s *Server) Result(ctx context.Context, input *token.AuctionResult) (*token.Response, error) {
	reply := &token.Response{}
	if s.timeRemaining == 0 {
		log.Printf("Auction over.\n")
		reply.Outcome = &token.Response_AuctionResult{
			AuctionResult: &token.AuctionResult{
				Id:     s.currentBid.id,
				Amount: s.currentBid.amount,
			},
		}
	} else {
		log.Printf("Auction not over resuming.\n")
		reply.Outcome = &token.Response_CurrentBid{
			CurrentBid: &token.CurrentBid{
				Id:     s.currentBid.id,
				Amount: s.currentBid.amount,
			},
		}
	}

	return reply, nil
}

// func to listen for a heartbeat from the primary server.
func (s *Server) Heartbeat(ctx context.Context, beat *token.Beat) (*token.Ack_, error) {
	log.Printf("Recieved heartbeat from primary, sending reply\n")

	s.heartbeatTimeout <- true

	reply := &token.Ack_{}

	return reply, nil
}

// func to send a heartbeat to all servers AND clients in the network.
func (s *Server) SendHeartbeat() {
	time.Sleep(time.Second * 5)

	log.Printf("Sending heartbeat to backups\n")
	for _, server := range s.servers {
		_, err := server.Heartbeat(s.ctx, &token.Beat{})
		if err != nil {
			log.Printf("%v", err)
		}
	}
	for _, client := range s.clients {
		_, err := client.Heartbeat(s.ctx, &token.Beat{})
		if err != nil {
			log.Printf("%v", err)
		}
	}

	time.Sleep(time.Second * 2)
}

// timeout func for the heartbeat, will timeout after 5 seconds with no answer.
func (s *Server) HeartbeatTimeout(reset <-chan bool, seconds int) {
	select {
	case <-reset:
		fmt.Println("Received heartbeat from primary")
		log.Println("Received heartbeat from primary")
	case <-time.After(time.Duration(seconds) * time.Second):
		// primary crashed
		fmt.Println("primary is dead recalling election")
		log.Println("primary is dead recalling election")
		delete(s.servers, s.primaryId)
		s.CallElection()
	}
}

// main loop.
func (s *Server) loop() {
	for {
		if s.primaryId == s.id {
			s.LeaderLoop()
		} else {
			s.BackupLoop()
		}
	}
}

// loop for the primary server.
func (s *Server) LeaderLoop() {
	log.Printf("Leader loop running\n")

	s.SendCoordination()

	for {
		if s.id != s.primaryId {
			log.Printf("No longer primary id breaking.\n")
			return
		}

		s.SendHeartbeat()
	}
}

// loop for the backup servers.
func (s *Server) BackupLoop() {
	log.Printf("Backup loop running\n")

	for {
		if s.id == s.primaryId {
			log.Printf("id is primary id, breaking out of the loop\n")
			s.heartbeatTimeout <- true
			return
		}
		s.HeartbeatTimeout(s.heartbeatTimeout, 5)
	}
}

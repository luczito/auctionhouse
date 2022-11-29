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
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

// command line args
var startPort = flag.Int("port", 5000, "port")
var maxManagers = flag.Int("managers", 3, "Max manager count")

func main() {

	//init
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var list net.Listener
	var port = int32(*startPort)
	var err error

	//loop to find a free port
	for i := 0; i < *maxManagers; i++ {
		list, err = net.Listen("tcp", fmt.Sprintf(":%v", port))

		if err == nil {
			break
		}

		port++
	}

	if err != nil {
		log.Fatalln("Could not find port(20 tries).")
	}

	//setup ip
	ip := fmt.Sprintf("127.0.0.1:%d", port)
	fmt.Println(ip)
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("address", ip))

	// set Log output, one log for each manager
	f, err := os.OpenFile(fmt.Sprintf("manager-log-%s.txt", ip), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v\n", err)
	}
	defer f.Close()
	log.SetOutput(f)

	//create a manager for this proccess
	m := &Manager{
		Id:                  ip,
		Managers:            make(map[string]token.ManagerClient),
		Clients:             make(map[string]token.ClientClient),
		Ctx:                 ctx,
		PrimaryId:           ip,
		TimeRemaining:       0,
		TimeoutHeartbeat:    make(chan bool, 1),
		TimeoutCoordination: make(chan bool, 1),
		CurrentBid: Bid{
			id:     "",
			amount: 0,
		},
		ExpectingAnswer: false,
	}

	grpcServer := grpc.NewServer()
	token.RegisterManagerServer(grpcServer, m)

	//discover other managers in the network
	var managerPort = int32(*startPort)

	log.Printf("start on port: %v\n", port)
	go func() {
		if err := grpcServer.Serve(list); err != nil {
			log.Fatalf("failed to server %v\n", err)
		}
	}()

	log.Printf("my ip is: %s\n", ip)

	//dial the other managers
	for i := 0; i < *maxManagers; i++ {
		if port == managerPort {
			managerPort++
			continue
		}

		log.Printf("dialing: %v\n", managerPort)

		ip := fmt.Sprintf("127.0.0.1:%v", managerPort)

		conn, err := grpc.DialContext(ctx, ip, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())

		if err != nil {
			managerPort++
			continue
		}

		defer conn.Close()

		log.Printf("[main]: ip %s\n", ip)

		c := token.NewManagerClient(conn)

		m.Managers[ip] = c

		managerPort++
	}

	//sleep then run the main loop.
	<-time.After(3 * time.Second)

	m.MainLoop()
}

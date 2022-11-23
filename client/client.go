package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	token "github.com/luczito/auctionhouse/proto"
	"google.golang.org/grpc"
)

//client file, holds the client struct and the methods of client.

type Client struct {
	name string
	port string
}

func parse() (string, error) {
	var input string
	sc := bufio.NewScanner(os.Stdin)
	if sc.Scan() {
		input = sc.Text()
	}
	return input, nil
}

func listen(clientConnection token.Auction_ConnectClient, c *Client) {
	for {
		msg, err := clientConnection.Recv()
		if err != nil {
			log.Fatalf("Unable to recieve the message %v", err)
		}
		fmt.Printf("%v: %v\n", msg.Name, msg.Msg)
	}
}

func start(c *Client) {
	var clientConnection token.Auction_ConnectClient

	for {
		connection, err := grpc.Dial(c.name+":"+c.port, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Unable to connect to the server %v", err)
		}

		client := token.NewAuctionClient(connection)

		clientConnection, err = client.Connect(context.Background())
		if err != nil {
			log.Fatalf("Unable to connect to the server %v", err)
		}

		clientConnection.Send(&token.ClientMsg{
			Name: c.name,
		})

		message, err := clientConnection.Recv()
		if err == nil {
			fmt.Printf("%v: %v\n", message.Name, message.Msg)
			break
		}
		fmt.Println("Username already in use, please try again.")
	}

	go listen(clientConnection, c)

	for {
		input, err := parse()
		if err != nil {
			log.Fatalf("Unable to retrieve the user input")
		}

		if input == "exit" {
			break
		}

		clientConnection.Send(&token.ClientMsg{
			Name: c.name,
			Msg:  input,
		})
	}
	clientConnection.CloseSend()
}

func seeAuctions() {
	//func to get a list of all autions listed
}

func bidOnAuction() {
	//func to bid on a certain auction
}

func main() {
	//set ip/port for server
	fmt.Print("Enter Server IP:Port, press enter for default IP:Port")
	reader := bufio.NewReader(os.Stdin)
	server, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("Unable to read from console: %v", err)
	}
	server = strings.Trim(server, "\r\n")
	var serverName string
	var serverPort string
	//if ip:port is not set then default value.
	if server == "" {
		serverName = "localhost"
		serverPort = "8080"
	} else {
		serverInfo := strings.Split(server, ":")
		serverName = serverInfo[0]
		serverPort = serverInfo[1]
	}

	log.Printf("Connecting to the server: %v:%v", serverName, serverPort)

	client := &Client{
		name: serverName,
		port: serverPort,
	}
	start(client)
}

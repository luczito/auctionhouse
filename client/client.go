package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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

func listen() {

}

func start(c *Client) {

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

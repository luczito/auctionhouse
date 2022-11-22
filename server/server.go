package main

import (
	"sync"

	token "github.com/luczito/auctionhouse/proto"
)

//server file, holds the server struct and methods for the servers (multiple).

type Server struct {
	isPrimary   bool
	port        string
	name        string
	clients     map[string]token.AuctionClient
	clientNames map[string]string
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

func (s *Server) connect() {

}

func start(s *Server) {

}

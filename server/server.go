package main

import (
	"sync"
)

//server file, holds the server struct and methods for the servers (multiple).

type Server struct {
	isPrimary   bool
	port        string
	name        string
	clients     map[string]
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

func (s *server) connect() {

}

func start(s *server) {

}

package main

import (
	"context"
	"log"
	"time"

	token "github.com/luczito/auctionhouse/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

// election listener.
func (m *Manager) Election(ctx context.Context, input *token.ElectionRequest) (*token.Ack, error) {
	log.Printf("Recieved election call, calling election\n")
	m.CallElection()

	return &token.Ack{}, nil
}

// calls an election to all nodes with a higher id than oneself. then calls coordination if the node is the highest id
func (m *Manager) CallElection() {
	log.Printf("Sending election to higher Managers\n")
	votes := len(m.Managers)

	for id, Manager := range m.Managers {
		if id < m.Id {
			votes--
			continue
		}

		_, err := Manager.Election(m.Ctx, &token.ElectionRequest{})

		if err != nil {
			log.Fatalf("Unable to obtain response from %v, %v\n", Manager, err)
			votes--
			delete(m.Managers, id)
		}
	}

	log.Printf("need %v votes to be elected\n", votes)

	if votes == 0 {
		m.PrimaryId = m.Id
		log.Printf("im the captain now\n")
		//s.SendCoordination()
	} else {
		log.Printf("Waiting for coordination from primary")
		m.CoordinationTimeout(5000)
	}
}

// func to timeout the wait for a coordination from another node.
func (m *Manager) CoordinationTimeout(ms int) {
	m.ExpectingAnswer = true

	select {
	case <-m.TimeoutCoordination:
		log.Println("New primary lives")
	case <-time.After(time.Duration(ms) * time.Millisecond):
		// primary dead
		log.Println("Coordination timeout, call new election")
		m.CallElection()
	}

	m.ExpectingAnswer = false
}

// sends coordination to other Managers in the network
func (m *Manager) SendCoordination() {
	log.Printf("Sending coordination to backups\n")

	coordination := &token.Coord{
		Id: m.Id,
	}

	for id, Manager := range m.Managers {
		_, err := Manager.Coordination(m.Ctx, coordination)

		if err != nil {
			log.Printf("Something went wrong with Manager: %v, %v\n", id, err)
			//delete(s.Managers, id)
		}
	}
}

// listener for sendcoordination function
func (m *Manager) Coordination(ctx context.Context, input *token.Coord) (*token.Ack, error) {
	var addr string
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		// log.Println(md)
		// log.Printf("[Coordinate]: address %v\n", addr)
		addr = md.Get("address")[0]
	}
	wasLeader := (m.Id == m.PrimaryId && m.Id != addr)

	m.PrimaryId = addr

	select {
	case m.TimeoutCoordination <- true:
	default:
	}

	if _, ok := m.Managers[addr]; !ok {
		log.Println("Saving new client")

		conn, err := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())

		if err != nil {
			log.Fatalf("could not reconnect: %v\n", err)
		}

		defer conn.Close()

		c := token.NewManagerClient(conn)
		m.Managers[addr] = c
	}

	if wasLeader {
		addrs := make([]string, len(m.Clients))

		for addr := range m.Clients {
			addrs = append(addrs, addr)
		}

		m.Managers[addr].Update(ctx, &token.Data{
			CurrentBid: &token.CurrentBid{
				Id:     m.CurrentBid.id,
				Amount: m.CurrentBid.amount,
			},
			TimeRemaining: m.TimeRemaining,
			Clients:       addrs,
		})
	}

	return &token.Ack{}, nil
}

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	token "github.com/luczito/auctionhouse/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

// Manager struct.
type Manager struct {
	token.UnimplementedManagerServer
	Id                  string //own manager ip
	PrimaryId           string //primary Manager ip
	Clients             map[string]token.ClientClient
	Managers            map[string]token.ManagerClient
	Ctx                 context.Context
	CurrentBid          Bid
	TimeRemaining       int32
	TimeoutHeartbeat    chan bool
	TimeoutCoordination chan bool
	ExpectingAnswer     bool
}

// bid struct
type Bid struct {
	id     string
	amount int32
}

// listens and recieves updates from primary Manager.
func (m *Manager) Update(ctx context.Context, data *token.Data) (*token.Ack, error) {
	log.Printf("Recieved an update for internal data\n")

	//update own info to the new updated info.
	m.CurrentBid.id = data.CurrentBid.Id
	m.CurrentBid.amount = data.CurrentBid.Amount
	m.TimeRemaining = data.TimeRemaining

	for _, addr := range data.Clients {
		conn, err := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())

		if err != nil {
			continue
		}

		c := token.NewClientClient(conn)
		m.Clients[addr] = c
	}
	return &token.Ack{}, nil
}

// updates the other Managers in the network
func (m *Manager) UpdateManagers() error {
	log.Printf("Sending update to backups\n")
	addrs := make([]string, 0, len(m.Clients))

	for addr := range m.Clients {
		addrs = append(addrs, addr)
	}
	//loop to send the update to all other managers
	for _, manag := range m.Managers {
		manag.Update(m.Ctx, &token.Data{
			CurrentBid: &token.CurrentBid{
				Id:     m.CurrentBid.id,
				Amount: m.CurrentBid.amount,
			},
			TimeRemaining: m.TimeRemaining,
			Clients:       addrs,
		})
	}
	return nil
}

// func to listen for a heartbeat from the primary Manager.
func (m *Manager) Heartbeat(ctx context.Context, beat *token.Beat) (*token.Ack, error) {
	log.Printf("Recieved heartbeat from primary, sending reply\n")

	m.TimeoutHeartbeat <- true

	reply := &token.Ack{}

	return reply, nil
}

// func to send a heartbeat to all Managers AND clients in the network.
func (m *Manager) SendHeartbeat() {

	log.Printf("Sending heartbeat to everyone\n")
	//send heartbeat to all managers
	for id, manag := range m.Managers {
		_, err := manag.Heartbeat(m.Ctx, &token.Beat{})
		if err != nil {
			delete(m.Managers, id)
			log.Printf("%v", err)
		}
	}
	//send heartbeat to all clients
	for id, client := range m.Clients {
		_, err := client.Heartbeat(m.Ctx, &token.Primary{})
		if err != nil {
			delete(m.Clients, id)
			log.Printf("%v", err)
		}
	}

	//sleep 2 sec so that the heartbeat isnt spammed.
	time.Sleep(time.Second * 2)
}

// timeout func for the heartbeat, will timeout after 5 seconds with no answer.
func (m *Manager) HeartbeatTimeout(reset <-chan bool, seconds int) {

	select {
	case <-reset:
		//primary lives
		fmt.Println("Received heartbeat from primary")
		log.Println("Received heartbeat from primary")

	case <-time.After(time.Duration(seconds) * time.Second):
		// primary dead
		fmt.Println("primary is dead recalling election")
		log.Println("primary is dead recalling election")
		delete(m.Managers, m.PrimaryId)
		m.CallElection()
	}
}

// listens and registers a bid from a client. Also registers the client if its the first bid from that client.
func (m *Manager) Bid(ctx context.Context, input *token.Amount) (*token.Ack, error) {
	var status token.Status
	var message string

	var addr string

	//get md from client
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		addr = md.Get("address")[0]
	}

	//dial up the new client and add to the client map
	if _, ok := m.Clients[addr]; !ok {
		conn, err := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
		if err != nil {
			log.Fatalf("could not reconnect: %v\n", err)
		}

		c := token.NewClientClient(conn)
		m.Clients[addr] = c
	}

	//check the bid, is it higher than the current bid?
	if input.Value > m.CurrentBid.amount {
		m.CurrentBid.amount = input.Value
		m.CurrentBid.id = addr
		status = *token.Status_SUCCESS.Enum()
		message = "Bid accepted as the new highest bid."
		log.Printf("Bid accepted as the new highest bid.\n")
		if err := m.UpdateManagers(); err != nil {
			message = "Database failure"
			log.Fatalf("Database failure\n")
		}
	} else {
		log.Printf("Bid rejected, lower than the current highest bid.\n")
		status = *token.Status_EXCEPTION.Enum()
		message = "Bid rejected, lower than the current highest bid."
	}

	//reply the client.
	reply := &token.Ack{
		Status:  status,
		Message: message,
	}

	return reply, nil
}

// returns the current state of the auction.
func (m *Manager) Result(ctx context.Context, input *token.Void) (*token.Outcome, error) {
	reply := &token.Outcome{
		TimeRemaining: m.TimeRemaining,
	}

	//if the auction is over then respond accordingly.
	//WILL ALWAYS HIT THIS IF STATEMENT IN THE CURRENT PROGRAM SINCE THE COUNTER ISNT SETUP
	if m.TimeRemaining == 0 {
		log.Printf("Auction over.\n")
		reply.Outcome = &token.Outcome_AuctionResult{
			AuctionResult: &token.AuctionResult{
				Id:     m.CurrentBid.id,
				Amount: m.CurrentBid.amount,
			},
		}
		//else respon accordingly. WILL NEVER RETURN THIS IN THE CURRENT PROGRAM
	} else {
		log.Printf("Auction not over resuming.\n")
		reply.Outcome = &token.Outcome_CurrentBid{
			CurrentBid: &token.CurrentBid{
				Id:     m.CurrentBid.id,
				Amount: m.CurrentBid.amount,
			},
		}
	}
	return reply, nil
}

// main loop.
func (m *Manager) MainLoop() {
	for {
		//check if the manager is the primary
		if m.PrimaryId == m.Id {
			m.LeaderLoop()
		} else {
			m.BackupLoop()
		}
	}
}

// loop for the primary Manager.
func (m *Manager) LeaderLoop() {
	log.Printf("Leader loop running\n")
	//send coordination first time this is ran.
	m.SendCoordination()

	for {
		//check if this manager is still the primary
		if m.Id != m.PrimaryId {
			log.Printf("No longer primary id breaking.\n")
			return
		}
		//send heartbeat to all connections.
		m.SendHeartbeat()
	}
}

// loop for the backup Managers.
func (m *Manager) BackupLoop() {
	log.Printf("Backup loop running\n")

	for {
		//check if this manager is suddenly the primary
		if m.Id == m.PrimaryId {
			log.Printf("id is primary id, breaking out of the loop\n")
			m.TimeoutHeartbeat <- true
			return
		}
		//wait for a heartbeat
		m.HeartbeatTimeout(m.TimeoutHeartbeat, 5)
	}
}

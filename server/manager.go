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
	PrimaryId           string //primary Manager port
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
func (m *Manager) Update(ctx context.Context, data *token.Data) (*token.Ack_, error) {
	log.Printf("Recieved an update for internal data\n")

	m.CurrentBid.id = data.CurrentBid.Id
	m.CurrentBid.amount = data.CurrentBid.Amount
	m.TimeRemaining = data.TimeRemaining

	for _, addr := range data.Clients {
		conn, err := grpc.Dial(fmt.Sprintf(":%v", addr), grpc.WithInsecure(), grpc.WithBlock())

		if err != nil {
			continue
		}

		// defer conn.Close()

		c := token.NewClientClient(conn)

		m.Clients[addr] = c
	}

	return &token.Ack_{}, nil
}

// updates the other Managers in the network
func (m *Manager) UpdateManagers() error {
	log.Printf("Sending update to backups\n")
	addrs := make([]int32, 0, len(m.Clients))

	for addr := range m.Clients {
		addrs = append(addrs, addr)
	}

	for _, manag := range m.Managers {
		manag.Update(s.Ctx, &token.Data{
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
	time.Sleep(time.Second * 5)

	log.Printf("Sending heartbeat to everyone\n")
	for _, manag := range m.Managers {
		_, err := manag.Heartbeat(s.Ctx, &token.Beat{})
		if err != nil {
			log.Printf("%v", err)
		}
	}
	for _, client := range m.Clients {
		_, err := client.Heartbeat(s.Ctx, &token.Primary{})
		if err != nil {
			log.Printf("%v", err)
		}
	}

	time.Sleep(time.Second * 2)
}

// timeout func for the heartbeat, will timeout after 5 seconds with no answer.
func (m *Manager) HeartbeatTimeout(reset <-chan bool, seconds int) {

	select {

	case <-reset:
		fmt.Println("Received heartbeat from primary")
		log.Println("Received heartbeat from primary")

	case <-time.After(time.Duration(seconds) * time.Second):
		// primary crashed
		fmt.Println("primary is dead recalling election")
		log.Println("primary is dead recalling election")
		delete(s.Managers, s.PrimaryId)
		s.CallElection()
	}
}

// main loop.
func (m *Manager) MainLoop() {
	for {
		fmt.Println("loop running")
		if s.PrimaryId == s.Id {
			fmt.Println("leader loop")
			s.LeaderLoop()
		} else {
			fmt.Println("backup loop")
			s.BackupLoop()
		}
	}
}

// loop for the primary Manager.
func (m *Manager) LeaderLoop() {
	log.Printf("Leader loop running\n")

	s.SendCoordination()

	for {
		if s.Id != s.PrimaryId {
			log.Printf("No longer primary id breaking.\n")
			return
		}

		s.SendHeartbeat()
	}
}

// loop for the backup Managers.
func (m *Manager) BackupLoop() {
	log.Printf("Backup loop running\n")

	for {
		if m.Id == m.PrimaryId {
			log.Printf("id is primary id, breaking out of the loop\n")
			m.TimeoutHeartbeat <- true
			return
		}
		m.HeartbeatTimeout(m.TimeoutHeartbeat, 5)
	}
}

// listens and registers a bid from a client.
func (m *Manager) Bid(ctx context.Context, input *token.Amount) (*token.Ack, error) {
	var status token.Status
	var message string

	var addr string

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		addr = md.Get("address")[0]
	}

	if _, ok := m.Clients[addr]; !ok {
		conn, err := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
		if err != nil {
			log.Fatalf("could not reconnect: %v\n", err)
		}

		c := token.NewClientClient(conn)
		m.Clients[addr] = c
	}

	if input.Amount > m.CurrentBid.amount {
		m.CurrentBid.amount = input.Amount
		m.CurrentBid.id = input.Id
		message = "Bid accepted as the new highest bid."
		log.Printf("Bid accepted as the new highest bid.\n")
		if err := m.UpdateManagers(); err != nil {
			message = "Database failure"
			log.Fatalf("Database failure\n")
		}
	} else {
		log.Printf("Bid rejected, lower than the current highest bid.\n")
		message = "Bid rejected, lower than the current highest bid."
	}

	reply := &token.Ack{
		Status:  status,
		Message: message,
	}

	return reply, nil
}

// when the auction is over this method will return the winner and the amount
func (m *Manager) Result(ctx context.Context, input *token.Void) (*token.Outcome, error) {
	reply := &token.Outcome{
		TimeRemaining: m.TimeRemaining,
	}

	if m.TimeRemaining == 0 {
		log.Printf("Auction over.\n")
		reply.Outcome = &token.Outcome_Result{
			AuctionResult: &token.AuctionResult{
				Id:     m.CurrentBid.id,ActionResult{
				Id:     m.CurrentBid.id,
				Amount: m.CurrentBid.amount,
			,
			}
		}
	}
	} else {
		log.Printf("Auction not over resuming.\n")
		reply.Outcome = &token.ResponseCurrentBid{
			CurrentBid: &token.CurrenBid{
				Id:     m.CurrentBid.id,
				Aount: m.CurrentBid.amount,
			,
			}
		}
	}
	return reply, nil
}

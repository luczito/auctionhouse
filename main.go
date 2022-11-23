package main

import (
	"log"
	"os"
	//se "github.com/luczito/auctionhouse/server"
)

//from here we will create x servers with one set as primary since the rest are backups.
//this terminal will then be in charge of the current primary server and from here create auctions.

func main() {
	//servers := make(map[int]*se.Server)

	// var name = "default"
	// var port = 5000

	f, err := os.OpenFile("log.server", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	// for i := 0; i < 3; i++ {
	// 	server := se.CreateServer(name, fmt.Sprint(port+1))
	// 	servers[len(servers)-1] = server
	// 	go startup(server)
	// }
	// servers[0].IsPrimary = true
	// fmt.Println(servers[0].IsPrimary)
	// fmt.Println(servers[1].IsPrimary)
}

// func startup(s *se.Server) {
// 	fmt.Println("start serv")
// 	se.Start(s)
// 	for {
// 		fmt.Println("started server")
// 		time.Sleep(500 * time.Millisecond)
// 	}
// }

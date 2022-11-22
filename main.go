package server

import (
	"flag"
	"log"
	"os"
)

//from here we will create x servers with one set as primary since the rest are backups.
//this terminal will then be in charge of the current primary server and from here create auctions.

func main() {
	f, err := os.OpenFile("log.server", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	flag.Parse()

	//for x start servers.
}

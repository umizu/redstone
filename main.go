package main

import (
	"fmt"
	"log"
	"os"

	"github.com/umizu/redstone/p2p"
)

func main() {
	port := os.Getenv("REDSTONE_PORT")
	if port == "" {
		port = "7000"
	}
	
	destinationAddr := os.Getenv("REDSTONE_DESTINATION_ADDR")
	if destinationAddr == "" {
		log.Fatal("$REDSTONE_DESTINATION_ADDR is required")
	}

	opts := p2p.TCPTransportOpts{
		ListenPort:      fmt.Sprintf(":%s", port),
		DestinationAddr: destinationAddr,
	}

	tr := p2p.NewTCPTransport(opts)

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf(
		"tcp server listening on port %s\n",
		opts.ListenPort)

	select {}
}

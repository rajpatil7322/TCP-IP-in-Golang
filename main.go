package main

import (
	"CAS/p2p"
	"fmt"
	"log"
)

func main() {
	tr := p2p.NewTCPTransport(":9000")

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Hello world")

	select {}
}

package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":9000")

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(conn)
	}

	data := []byte("Hello, Server!")
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

package p2p

import (
	"fmt"
	"net"
	"sync"
)

type TCPPeer struct {
	conn     net.Conn
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

type TCPTransport struct {
	listenAddress string
	listner       net.Listener
	mu            sync.RWMutex
	peers         map[net.Addr]Peer
}

func NewTCPTransport(listnenAddr string) *TCPTransport {
	return &TCPTransport{
		listenAddress: listnenAddr,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.listner, err = net.Listen("tcp", t.listenAddress)

	if err != nil {
		return err
	}

	go t.startAcceptLoop()
	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listner.Accept()
		if err != nil {
			fmt.Printf("TCP accept error %s\n", err)
		}

		go t.handleConn(conn)
	}
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)

	// Create a buffer to read data into
	buffer := make([]byte, 1024)
	fmt.Printf("New incoming connection %+v\n", peer)

	for {
		// Read data from the client
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("Received: %s\n", buffer[:n])
	}

}

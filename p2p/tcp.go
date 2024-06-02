package p2p

import (
	"fmt"
	"io"
	"log"
	"net"
)

type TCPTransportOpts struct {
	ListenPort      string
	DestinationAddr string
}

type TCPTransport struct {
	TCPTransportOpts
	listener net.Listener
}

func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.listener, err = net.Listen("tcp", t.ListenPort)
	if err != nil {
		return err
	}

	go t.startAcceptLoop()
	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept() // handshake
		if err != nil {
			fmt.Printf("tcp accept error: %+v\n", err)
		}
		fmt.Printf("new incoming connection: %+v\n", conn)

		go t.handleConn(conn)
	}
}

func (t *TCPTransport) handleConn(clientConn net.Conn) {
	defer clientConn.Close()

	serverConn, err := net.Dial("tcp", t.DestinationAddr)
	if err != nil {
		log.Printf("error connecting to destination server: %v", err)
		return
	}
	defer serverConn.Close()

	go io.Copy(serverConn, clientConn)
	io.Copy(clientConn, serverConn)
}

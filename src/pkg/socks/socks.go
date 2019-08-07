package socks

// Concurrent socket connection handler.

import (
	"bufio"
	"fmt"
	"net"

	handlers "github.com/arush15june/revshellz/src/handlers"
	chanstore "github.com/arush15june/revshellz/src/pkg/chanstore"
)

var (
	handler handlers.ConnHandler
)

// TCPListener initializes TCP server and connection handlers.
func TCPListener(port string) {
	port = ":" + port

	listener, err := net.Listen("tcp4", port)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		connChan := chanstore.AddChannel(conn.RemoteAddr().String())

		handler.HandleNewConnection(conn.RemoteAddr().String())
		go connectionHandler(connChan.WChannel, connChan.RChannel, conn)
	}

}

// InitTCPListener runs a TCP Listener goroutine on `port`
func InitTCPListener(port string, connHandler handlers.ConnHandler) {
	handler = connHandler
	go TCPListener(port)
}

// connectionHandler handles connections and RW channels of the socket.
func connectionHandler(writechan chan []byte, readchan chan []byte, conn net.Conn) error {
	defer func() {
		handler.HandleCloseConnection(conn.RemoteAddr().String())
		chanstore.RemoveChannel(conn.RemoteAddr().String())
		conn.Close()
		conn = nil
	}()

	r := bufio.NewReader(conn)
	scanner := bufio.NewScanner(r)

	w := bufio.NewWriter(conn)

	connStatus := make(chan bool)

	go readHandler(readchan, connStatus, scanner, conn)

	for {
		select {
		case msg := <-writechan:
			w.Write(msg)
			w.Flush()
		// Verify TCP Connection Status.
		case status := <-connStatus:
			if !status {
				break
			}
		}
	}

	return nil
}

func readHandler(readchan chan []byte, status chan bool, scanner *bufio.Scanner, conn net.Conn) {

	for {
		connected := scanAndVerifyConnection(scanner)

		// Notify handler to close connection.
		status <- connected
		if !connected {
			break
		}

		msg := scanner.Bytes()
		handler.HandleReadMessage(conn.RemoteAddr().String(), msg)
	}
}

func scanAndVerifyConnection(scanner *bufio.Scanner) bool {
	connected := true
	if scanned := scanner.Scan(); !scanned {
		connected = false
	}

	return connected
}

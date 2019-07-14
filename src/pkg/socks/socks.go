package socks

// Concurrent socket connection handler.

import (
	"bufio"
	"fmt"
	"net"

	chanstore "../chanstore"
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

		fmt.Printf("[*] Connection from! %s\n", conn.RemoteAddr().String())
		go connectionHandler(connChan.WChannel, connChan.RChannel, conn)
	}

}

// connectionHandler handles connections and RW channels of the socket.
func connectionHandler(writechan chan []byte, readchan chan []byte, conn net.Conn) error {
	defer func() {
		fmt.Printf("[*] Connection from %s closed\n", conn.RemoteAddr().String())
		chanstore.RemoveChannel(conn.RemoteAddr().String())
		conn.Close()
		conn = nil
	}()

	r := bufio.NewReader(conn)
	scanner := bufio.NewScanner(r)

	w := bufio.NewWriter(conn)

	var status bool
	connStatus := make(chan bool)

	go readHandler(readchan, connStatus, scanner)

	for {
		// Verify TCP Connection Status.
		status = true
		select {
		case status = <-connStatus:
		default:
		}

		if !status {
			break
		}

		select {
		case msg := <-writechan:
			w.Write(msg)
			w.Flush()
		default:
		}
	}

	return nil
}

func readHandler(readchan chan []byte, status chan bool, scanner *bufio.Scanner) {

	for {
		connected := scanAndVerifyConnection(scanner)
		// Notify handler to close connection.
		status <- connected
		if !connected {
			break
		}

		msg := scanner.Bytes()
		select {
		case readchan <- msg:
		}
	}

	return
}

func scanAndVerifyConnection(scanner *bufio.Scanner) bool {
	connected := true
	if scanned := scanner.Scan(); !scanned {
		connected = false
	}

	return connected
}

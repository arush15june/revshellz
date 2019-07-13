package socks

// Concurrent socket connection handler.

import (
	"fmt"
	"io"
	"net"
	"time"

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
		go connectionHandler(connChan.Channel, conn)
	}

}

func connectionHandler(msgchan chan string, conn net.Conn) {
	defer conn.Close()
	var msg []byte

	for {
		conn.SetReadDeadline(time.Now().Add(10 * time.Millisecond))
		_, err := conn.Read(msg)
		if err == io.EOF {
			fmt.Println("Connection Closed")
			chanstore.RemoveChannel(conn.RemoteAddr().String())
			break
		}

		msg = []byte(<-msgchan)
		conn.Write(msg)
	}
	conn = nil
}

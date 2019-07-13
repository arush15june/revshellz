package main

import (
	"fmt"
	"time"

	chanstore "./pkg/chanstore"
	socks "./pkg/socks"
)

func main() {
	if !chanstore.IsStoreInitialized() {
		chanstore.InitStore()
	}

	go socks.TCPListener("18000")

	for {
		chans := chanstore.GetChans()

		for k, v := range chans {
			fmt.Printf("Writing to %v\r\n", k)
			v.Channel <- fmt.Sprintf("Test Message to %v\r\n", v)
		}
		time.Sleep(500 * time.Millisecond)
	}

}

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

			// Channel writer.
			go chanstore.writeChannel(v.WChannel, []byte(fmt.Sprintf("Test Message to %v %d\r\n", v, uint64(time.Now().Unix()))))
			fmt.Printf("main: writing to %v\r\n", k)

			// Channel reader.
			msg, err := chanstore.readChannel(v.RChannel)
			if err == nil {
				fmt.Println("main:", string(msg))
			}
		}
		time.Sleep(1000 * time.Millisecond)
	}

}

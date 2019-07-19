package main

// TODO: Implement REST API.
// TODO: Implement WebSocket API.

import (
	"fmt"
	"time"

	api "./api"
	chanstore "./pkg/chanstore"
	socks "./pkg/socks"
)

func main() {
	if !chanstore.IsStoreInitialized() {
		chanstore.InitStore()
	}

	socks.InitTCPListener("18000")
	api.InitRestApi("5000")

	go func() {
		for {
			chans := chanstore.GetChans()
			for _, v := range chans {
				// Channel reader.
				msg, err := v.ReadChannel()
				if err == nil {
					fmt.Println(v, ":", string(msg))
				}
			}
		}
	}()

	for {
		chans := chanstore.GetChans()
		for k, v := range chans {

			// Channel writer.
			go v.WriteChannel([]byte(fmt.Sprintf("Test Message to %v %d\r\n", v, uint64(time.Now().Unix()))))
			fmt.Printf("main: writing to %v\r\n", k)
		}
		time.Sleep(500 * time.Millisecond)
	}

}

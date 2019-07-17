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

	for {
		chans := chanstore.GetChans()
		for _, v := range chans {

			// Channel writer.
			// go chanstore.WriteChannel(v.WChannel, []byte(fmt.Sprintf("Test Message to %v %d\r\n", v, uint64(time.Now().Unix()))))
			// fmt.Printf("main: writing to %v\r\n", k)

			// Channel reader.
			msg, err := chanstore.ReadChannel(v.RChannel)
			if err == nil {
				fmt.Println("main:", string(msg))
			}
		}
		time.Sleep(1000 * time.Millisecond)
	}

}

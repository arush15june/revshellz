package main

// TODO: Implement REST API.
// TODO: Implement WebSocket API.

import (
	api "./api"
	flags "./flags"
	handlers "./handlers"
	chanstore "./pkg/chanstore"
	socks "./pkg/socks"
	tui "./tui"
)

func main() {
	flags.InitFlags()
	if !chanstore.IsStoreInitialized() {
		chanstore.InitStore()
	}

	var handlerInterface handlers.ConnHandler

	if *flags.RestApi && !*flags.Tui {
		handlerInterface = handlers.RestApiHandler{}
	} else {
		handlerInterface = handlers.TuiHandler{}
	}

	socks.InitTCPListener("18000", handlerInterface)

	if *flags.RestApi {
		api.InitRestApi("5000")
	} else if *flags.Tui {
		tui.InitTUI()
	}

	// go func() {
	// 	for {
	// 		chans := chanstore.GetChans()
	// 		for _, v := range chans {
	// 			// Channel reader.
	// 			msg, err := v.ReadChannel()
	// 			if err == nil {
	// 				tui.WriteLogView(fmt.Sprintln(v, ":", string(msg)))
	// 			}
	// 		}
	// 	}
	// }()

	// for {
	// 	chans := chanstore.GetChans()
	// 	for _, v := range chans {

	// 		// Channel writer.
	// 		// go handlerInterface.HandleWriteMessage(v.IPAddr, []byte(fmt.Sprintf("Test Message to %v %d\r\n", v, uint64(time.Now().Unix()))))
	// 		go handlerInterface.HandleWriteMessage(v.IPAddr, []byte("dir"))
	// 		// fmt.Printf("main: writing to %v\r\n", k)
	// 	}
	// 	time.Sleep(500 * time.Millisecond)
	// }

}

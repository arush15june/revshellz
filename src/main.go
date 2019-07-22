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
}

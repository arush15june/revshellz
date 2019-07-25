package main

// TODO: Implement REST API.
// TODO: Implement WebSocket API.

import (
	"fmt"
	
	api "github.com/arush15june/revshellz/src/api"
	flags "github.com/arush15june/revshellz/src/flags"
	handlers "github.com/arush15june/revshellz/src/handlers"
	chanstore "github.com/arush15june/revshellz/src/pkg/chanstore"
	socks "github.com/arush15june/revshellz/src/pkg/socks"
	tui "github.com/arush15june/revshellz/src/tui"
)

func main() {
	flags.InitFlags()
	if !chanstore.IsStoreInitialized() {
		chanstore.InitStore()
	}

	var handlerInterface handlers.ConnHandler

	if *flags.RestApi {
		handlerInterface = handlers.RestApiHandler{}
	} else if *flags.Line {
		handlerInterface = handlers.LineHandler{}
	} else if *flags.Tui {
		handlerInterface = handlers.TuiHandler{}
	}

	socks.InitTCPListener(*flags.Port, handlerInterface)

	if *flags.RestApi {
		api.InitRestApi("5000")
	} else if *flags.Tui {
		tui.InitTUI()
	} else if *flags.Line {
		fmt.Printf("Listening on :%v\n", *flags.Port)
		for {
			
		}
	}
}

package flags

import (
	"flag"
)

var (
	// RestApi flag.
	RestApi = flag.Bool("rest", false, "Enable Rest API (Disables Terminal UI)")

	// Tui flag.
	Tui = flag.Bool("tui", true, "Enable terminal UI (default)")

	// Port flag.
	Port = flag.String("port", "18000", "Port for TCP Listener")
)

func InitFlags() {
	flag.Parse()
}

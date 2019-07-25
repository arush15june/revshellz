package flags

import (
	"flag"
)

var (
	// RestApi flag.
	RestApi = flag.Bool("rest", false, "Enable Rest API (Disables Terminal UI)")

	// Tui flag.
	Tui = flag.Bool("tui", false, "Enable terminal UI (default)")

	// Port flag.
	Port = flag.String("port", "18000", "Port for TCP Listener")

	// LineUI flag.
	Line = flag.Bool("line", false, "Enable line UI.")
)

func InitFlags() {
	flag.Parse()
}

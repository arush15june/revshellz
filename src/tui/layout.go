package tui

// Absolute layout and handlers for the Terminal UI.
import (
	"fmt"

	chanstore "github.com/arush15june/revshellz/src/pkg/chanstore"
	"github.com/rivo/tview"
)

var (
	// ConnectionView is the TextView for the connection list.
	ConnectionView *tview.TextView

	// LogView is the TextView for the log box.
	LogView *tview.TextView
)

// InitTopFlex initializes the top flex bar.
func InitTopFlex() {
	TopFlex = tview.NewFlex().SetDirection(tview.FlexRow)
	TopFlex.SetBorder(true)
}

// InitBottomFlex initialzes the bottom flex bar.
func InitBottomFlex() {
	ConnectionView = NewTextView("Connections")
	LogView = NewTextView("Logs")

	BottomFlex = tview.NewFlex().
		AddItem(ConnectionView, 0, 1, false).
		AddItem(LogView, 0, 1, false)
}

// Handlers initializes the ViewHandler.
func Handlers() {
	go connectionBoxHandler()
}

func connectionBoxHandler() {
	nchans := 0
	var chans map[string]*chanstore.Messenger
	clear := true

	for {
		chans = chanstore.GetChans()

		if nchans == len(chans) {
			if len(chans) == 0 && !clear {
				ConnectionView.Clear()
				clear = true
			}
			continue
		}

		ConnectionView.Clear()
		for k := range chans {
			fmt.Fprintf(ConnectionView, "[red]%v ", k)
		}

		nchans = len(chans)
		clear = false
	}
}

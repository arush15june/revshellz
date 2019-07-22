package tui

// Terminal UI

import (
	"fmt"

	flags "../flags"
	"github.com/rivo/tview"
)

var (
	// The tview application.
	app *tview.Application

	// RootFlex is the root flex manager.
	RootFlex *tview.Flex

	// BottomFlex is the bottom bar flex manager.
	BottomFlex *tview.Flex

	// TopFlex is the the top bar flex manger.
	TopFlex *tview.Flex
)

// InitTUI sets up the terminal based UI for the application.
func InitTUI() {
	app = tview.NewApplication()

	InitTopFlex()
	InitBottomFlex()

	RootFlex = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(TopFlex, 0, 8, false).
		AddItem(BottomFlex, 7, 1, false)

	Handlers()
	WriteTextView(LogView, fmt.Sprintf("[+] Started listener on :%v\n", *flags.Port))
	if err := app.SetRoot(RootFlex, true).SetFocus(LogView).Run(); err != nil {
		panic(err)
	}
}

package tui

// tview creation utilties.

import (
	"fmt"
	"sync"

	"github.com/rivo/tview"
)

var (
	// ConnViews map the IP Addresses to the connection TextViews.
	ConnViews map[string]*tview.TextView
	mutex     sync.Mutex

	// CurrentFocusedView points to the current focused primitive.
	CurrentFocusedView tview.Primitive
)

// NewTitledBox returns a titled Box.
func NewTitledBox(title string) *tview.Box {
	return tview.NewBox().SetBorder(true).SetTitle(title)
}

// NewTextView returns a new TextView.
func NewTextView(title string) *tview.TextView {
	textView := tview.NewTextView().SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true).
		SetScrollable(true).
		SetChangedFunc(func() {
			app.Draw()
		})
	textView.SetBorder(true).SetTitle(title)

	return textView
}

// Handle input to ShellView.
// func ShellViewEventHandler(evt *tcell.EventKey) *tcell.EventKey {
// 	key := evt.Key()
// 	switch key {
// 	case tcell.KeyRune:
// 		{
// 			WriteTextView(CurrentFocusedView.(*tview.TextView), string(evt.Rune()))
// 		}
// 	case tcell.KeyEscape:
// 		{
// 			app.SetFocus(RootFlex)
// 		}
// 	}

// 	return evt
// }

// RootFlexHandler handles input for the RootFlex.
// func RootFlexHandler(evt *tcell.EventKey) *tcell.EventKey {
// 	key := evt.Key()
// 	nShells := len(ConnViews)

// 	switch key {
// 	case tcell.KeyRune:
// 		if strconv.ParseUint(event.Rune()) <= nShells && strconv.ParseUint(event.Rune()) > 0:
// 			CurrentFocusedView = ConnViews[]
// 	case tcell.KeyEscape:
// 		{
// 			app.SetFocus(RootFlex)
// 		}
// 	}

// 	return evt
// }

// NewShellView creates a new box for a connection.
func NewShellView(title string) *tview.TextView {
	return NewTextView(title)
}

// AddViewToMap adds a TextView to the map.
func AddViewToMap(view *tview.TextView, ip string) {
	if ConnViews == nil {
		ConnViews = make(map[string]*tview.TextView)
	}
	mutex.Lock()
	defer mutex.Unlock()

	ConnViews[ip] = view
}

// RemoveViewFromMap removes the TextView from the map.
func RemoveViewFromMap(ip string) {
	view := GetViewFromIp(ip)
	view.Clear()
	app.QueueUpdateDraw(func() {
		TopFlex.RemoveItem(view)
		delete(ConnViews, ip)
	})
}

// GetViewFromIp returns the view from the ConnViews map.
func GetViewFromIp(ip string) *tview.TextView {
	mutex.Lock()
	defer mutex.Unlock()
	return ConnViews[ip]
}

// AddConnection adds a new shell connection.
func AddConnection(ip string) *tview.TextView {
	WriteLogView(fmt.Sprintf("[*] Connection from! %s\n", ip))

	ShellView := NewShellView(fmt.Sprintf("[red]%v", ip))
	// ShellView.SetInputCapture(ShellViewEventHandler)
	AddViewToMap(ShellView, ip)

	app.QueueUpdateDraw(func() {
		TopFlex.AddItem(ShellView, 0, 1, false)
		// app.SetFocus(ShellView)
		CurrentFocusedView = ShellView
	})

	return ShellView
}

// RemoveConnection removes connection from shell list.
func RemoveConnection(ip string) {
	WriteLogView(fmt.Sprintf("[*] Connection from %s closed\n", ip))
	RemoveViewFromMap(ip)
}

// WriteTextView writes data to a text view.
func WriteTextView(view *tview.TextView, data string) {
	fmt.Fprintf(view, data)
}

// WriteLogView writes data to the LogView.
func WriteLogView(data string) {
	WriteTextView(LogView, data)
}

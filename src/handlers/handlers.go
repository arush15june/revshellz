package handlers

import (
	"fmt"

	chanstore "github.com/arush15june/revshellz/src/pkg/chanstore"
	tui "github.com/arush15june/revshellz/src/tui"
)

// ConnHandler is the interface to handling information for connections.
// - New Connections
// - Close Connection
// - When a message is received.
// - When a message is to be written.
type ConnHandler interface {
	HandleNewConnection(ip string)
	HandleCloseConnection(ip string)
	HandleReadMessage(ip string, data []byte)
	HandleWriteMessage(ip string, data []byte)
}

type TuiHandler struct{}

type RestApiHandler struct{}

type LineHandler struct{}

func (t TuiHandler) HandleNewConnection(ip string) {
	tui.AddConnection(ip)
}

func (t TuiHandler) HandleCloseConnection(ip string) {
	go func() {
		tui.RemoveConnection(ip)
	}()
}

func (t TuiHandler) HandleReadMessage(ip string, data []byte) {
	view := tui.GetViewFromIp(ip)
	tui.WriteTextView(view, fmt.Sprintf("[green]>> [white]%v\n", string(data)))
	view.ScrollToEnd()
}

func (t TuiHandler) HandleWriteMessage(ip string, data []byte) {
	view := tui.GetViewFromIp(ip)
	channel := chanstore.GetChannel(ip)
	channel.WriteChannel(data)

	tui.WriteTextView(view, fmt.Sprintf("[red]$ [white]%v\n", string(data)))
}

func (r RestApiHandler) HandleNewConnection(ip string) {

}

func (r RestApiHandler) HandleCloseConnection(ip string) {

}

func (r RestApiHandler) HandleReadMessage(ip string, data []byte) {

}

func (r RestApiHandler) HandleWriteMessage(ip string, data []byte) {

}

func (l LineHandler) HandleNewConnection(ip string) {
	fmt.Printf("[+] New Connection: %v\n", ip)
}

func (l LineHandler) HandleCloseConnection(ip string) {
	fmt.Printf("[-] Closed Connection %v\n", ip)
}

func (l LineHandler) HandleReadMessage(ip string, data []byte) {
	fmt.Printf("[%v] %v\n", ip, string(data))
}

func (l LineHandler) HandleWriteMessage(ip string, data []byte) {
	channel := chanstore.GetChannel(ip)
	channel.WriteChannel(data)

	fmt.Printf(">>[%v] %v\n", ip, string(data))
}

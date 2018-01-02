package chrome

import (
	"github.com/mkenney/go-chrome/socket"
)

/*
AddEventHandler implements Socketer
*/
func (tab *Tab) AddEventHandler(handler socket.EventHandler) {
	tab.Socket().AddEventHandler(handler)
}

/*
RemoveEventHandler implements Socketer
*/
func (tab *Tab) RemoveEventHandler(handler socket.EventHandler) {
	tab.Socket().RemoveEventHandler(handler)
}

/*
SendCommand implements Socketer
*/
func (tab *Tab) SendCommand(command socket.Commander) *socket.Payload {
	return tab.Socket().SendCommand(command)
}

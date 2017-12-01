package chrome

import (
	domstorage "app/chrome/domstorage"
	"app/chrome/protocol"
)

/*
DOMStorage queries and modifies DOM storage.
*/
type DOMStorage struct{}

/*
Enable enables storage tracking, storage events will now be delivered to the client.
*/
func (DOMStorage) Enable(socket *Socket) error {
	command := new(protocol.Command)
	command.method = "DOMStorage.enable"
	socket.SendCommand(command)
	return command.Err
}

/*
Disable disables storage tracking, prevents storage events from being sent to the client.
*/
func (DOMStorage) Disable(socket *Socket) error {
	command := new(protocol.Command)
	command.method = "DOMStorage.disable"
	socket.SendCommand(command)
	return command.Err
}

/*
Clear clears  a stored item.
*/
func (DOMStorage) Clear(socket *Socket, params *domstorage.ClearParams) error {
	command := new(protocol.Command)
	command.method = "DOMStorage.clear"
	socket.SendCommand(command)
	return command.Err
}

/*
GetDOMStorageItems gets a stored item.
*/
func (DOMStorage) GetDOMStorageItems(socket *Socket, params *domstorage.GetDOMStorageItemsParams) error {
	command := new(protocol.Command)
	command.method = "DOMStorage.getDOMStorageItems"
	socket.SendCommand(command)
	return command.Err
}

/*
SetDOMStorageItem sets a stored item.
*/
func (DOMStorage) SetDOMStorageItem(socket *Socket, params *domstorage.SetDOMStorageItemParams) error {
	command := new(protocol.Command)
	command.method = "DOMStorage.setDOMStorageItem"
	socket.SendCommand(command)
	return command.Err
}

/*
RemoveDOMStorageItem removes  a stored item.
*/
func (DOMStorage) RemoveDOMStorageItem(socket *Socket, params *domstorage.RemoveDOMStorageItemParams) error {
	command := new(protocol.Command)
	command.method = "DOMStorage.removeDOMStorageItem"
	socket.SendCommand(command)
	return command.Err
}

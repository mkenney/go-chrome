package chrome

import "app/chrome/protocol"

/*
DOMStorage - https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/
Queries and modifies DOM storage.
*/
type DOMStorage struct{}

/*
Enable enables storage tracking, storage events will now be delivered to the client.
*/
func (DOMStorage) Enable(socket *Socket) error {
	command := &protocol.Command{
		method: "DOMStorage.enable",
		params: nil,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Disable disables storage tracking, prevents storage events from being sent to the client.
*/
func (DOMStorage) Disable(socket *Socket) error {
	command := &protocol.Command{
		method: "DOMStorage.disable",
		params: nil,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Clear clears  a stored item.
*/
func (DOMStorage) Clear(socket *Socket, params *dom_storage.ClearParams) error {
	command := &protocol.Command{
		method: "DOMStorage.clear",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
GetDOMStorageItems gets a stored item.
*/
func (DOMStorage) GetDOMStorageItems(socket *Socket, params *dom_storage.GetDOMStorageItemsParams) error {
	command := &protocol.Command{
		method: "DOMStorage.getDOMStorageItems",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetDOMStorageItem sets a stored item.
*/
func (DOMStorage) SetDOMStorageItem(socket *Socket, params *dom_storage.SetDOMStorageItemParams) error {
	command := &protocol.Command{
		method: "DOMStorage.setDOMStorageItem",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
RemoveDOMStorageItem removes  a stored item.
*/
func (DOMStorage) RemoveDOMStorageItem(socket *Socket, params *dom_storage.RemoveDOMStorageItemParams) error {
	command := &protocol.Command{
		method: "DOMStorage.removeDOMStorageItem",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

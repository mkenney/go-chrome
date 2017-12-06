package chrome

import "app/chrome/protocol"

/*
DOMStorage - https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/
Queries and modifies DOM storage.
*/
type DOMStorage struct{}

/*
Clear clears  a stored item.
*/
func (DOMStorage) Clear(
	socket *Socket,
	params *dom_storage.ClearParams,
) (nil, error) {
	command := &protocol.Command{
		method: "DOMStorage.clear",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
Disable disables storage tracking, prevents storage events from being sent to the client.
*/
func (DOMStorage) Disable(
	socket *Socket,
) (nil, error) {
	command := &protocol.Command{
		method: "DOMStorage.disable",
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
Enable enables storage tracking, storage events will now be delivered to the client.
*/
func (DOMStorage) Enable(
	socket *Socket,
) (nil, error) {
	command := &protocol.Command{
		method: "DOMStorage.enable",
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
GetDOMStorageItems gets a stored item.
*/
func (DOMStorage) GetDOMStorageItems(
	socket *Socket,
	params *dom_storage.GetDOMStorageItemsParams,
) (dom_storage.GetDOMStorageItemsResult, error) {
	command := &protocol.Command{
		method: "DOMStorage.getDOMStorageItems",
		params: params,
	}
	socket.SendCommand(command)
	return command.Result.(dom_storage.GetDOMStorageItemsResult), command.Err
}

/*
RemoveDOMStorageItem removes  a stored item.
*/
func (DOMStorage) RemoveDOMStorageItem(
	socket *Socket,
	params *dom_storage.RemoveDOMStorageItemParams,
) (nil, error) {
	command := &protocol.Command{
		method: "DOMStorage.removeDOMStorageItem",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
SetDOMStorageItem sets a stored item.
*/
func (DOMStorage) SetDOMStorageItem(
	socket *Socket,
	params *dom_storage.SetDOMStorageItemParams,
) (nil, error) {
	command := &protocol.Command{
		method: "DOMStorage.setDOMStorageItem",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

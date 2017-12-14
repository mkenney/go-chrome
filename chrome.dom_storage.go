package chrome

import (
	dom_storage "app/chrome/dom_storage"
	"app/chrome/protocol"
	"encoding/json"
)

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
) error {
	command := &protocol.Command{
		Method: "DOMStorage.clear",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Disable disables storage tracking, prevents storage events from being sent to the client.
*/
func (DOMStorage) Disable(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "DOMStorage.disable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Enable enables storage tracking, storage events will now be delivered to the client.
*/
func (DOMStorage) Enable(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "DOMStorage.enable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
GetDOMStorageItems gets a stored item.
*/
func (DOMStorage) GetDOMStorageItems(
	socket *Socket,
	params *dom_storage.GetDOMStorageItemsParams,
) (dom_storage.GetDOMStorageItemsResult, error) {
	command := &protocol.Command{
		Method: "DOMStorage.getDOMStorageItems",
		Params: params,
	}
	result := dom_storage.GetDOMStorageItemsResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
RemoveDOMStorageItem removes  a stored item.
*/
func (DOMStorage) RemoveDOMStorageItem(
	socket *Socket,
	params *dom_storage.RemoveDOMStorageItemParams,
) error {
	command := &protocol.Command{
		Method: "DOMStorage.removeDOMStorageItem",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetDOMStorageItem sets a stored item.
*/
func (DOMStorage) SetDOMStorageItem(
	socket *Socket,
	params *dom_storage.SetDOMStorageItemParams,
) error {
	command := &protocol.Command{
		Method: "DOMStorage.setDOMStorageItem",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

package Chrome

import (
	"encoding/json"

	domStorage "github.com/mkenney/go-chrome/dom_storage"
	"github.com/mkenney/go-chrome/protocol"
	log "github.com/sirupsen/logrus"
)

/*
DOMStorage queries and modifies DOM storage.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/
*/
type DOMStorage struct{}

/*
Clear clears  a stored item.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-clear
*/
func (DOMStorage) Clear(
	socket *Socket,
	params *domStorage.ClearParams,
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

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-disable
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

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-enable
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
GetItems gets a stored item.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-getDOMStorageItems
*/
func (DOMStorage) GetItems(
	socket *Socket,
	params *domStorage.GetItemsParams,
) (domStorage.GetItemsResult, error) {
	command := &protocol.Command{
		Method: "DOMStorage.getDOMStorageItems",
		Params: params,
	}
	result := domStorage.GetItemsResult{}
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
RemoveItem removes  a stored item.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-removeDOMStorageItem
*/
func (DOMStorage) RemoveItem(
	socket *Socket,
	params *domStorage.RemoveItemParams,
) error {
	command := &protocol.Command{
		Method: "DOMStorage.removeDOMStorageItem",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetItem sets a stored item.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-setDOMStorageItem
*/
func (DOMStorage) SetItem(
	socket *Socket,
	params *domStorage.SetItemParams,
) error {
	command := &protocol.Command{
		Method: "DOMStorage.setDOMStorageItem",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
OnItemAdded adds a handler to the DOMStorage.domStorageItemAdded event.
DOMStorage.domStorageItemAdded fires when an item is added to DOM storage.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#event-domStorageItemAdded
*/
func (DOM) OnItemAdded(
	socket *Socket,
	callback func(event *domStorage.ItemAddedEvent),
) {
	handler := protocol.NewEventHandler(
		"DOMStorage.domStorageItemAdded",
		func(name string, params []byte) {
			event := &domStorage.ItemAddedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnItemRemoved adds a handler to the DOMStorage.domStorageItemRemoved event.
DOMStorage.domStorageItemRemoved fires when an item is removed from DOM storage.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#event-domStorageItemRemoved
*/
func (DOM) OnItemRemoved(
	socket *Socket,
	callback func(event *domStorage.ItemRemovedEvent),
) {
	handler := protocol.NewEventHandler(
		"DOMStorage.domStorageItemRemoved",
		func(name string, params []byte) {
			event := &domStorage.ItemRemovedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnItemUpdated adds a handler to the DOMStorage.domStorageItemUpdated event.
DOMStorage.domStorageItemUpdated fires when an item in DOM storage is updated.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#event-domStorageItemUpdated
*/
func (DOM) OnItemUpdated(
	socket *Socket,
	callback func(event *domStorage.ItemUpdatedEvent),
) {
	handler := protocol.NewEventHandler(
		"DOMStorage.domStorageItemUpdated",
		func(name string, params []byte) {
			event := &domStorage.ItemUpdatedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnItemsCleared adds a handler to the DOMStorage.domStorageItemsCleared event.
DOMStorage.domStorageItemsCleared fires when items in DOM storage are cleared.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#event-domStorageItemsCleared
*/
func (DOM) OnItemsCleared(
	socket *Socket,
	callback func(event *domStorage.ItemsClearedEvent),
) {
	handler := protocol.NewEventHandler(
		"DOMStorage.domStorageItemsCleared",
		func(name string, params []byte) {
			event := &domStorage.ItemsClearedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

package protocol

import (
	"encoding/json"

	domStorage "github.com/mkenney/go-chrome/protocol/dom_storage"
	sock "github.com/mkenney/go-chrome/socket"
	log "github.com/sirupsen/logrus"
)

/*
DOMStorage is a struct that provides a namespace for the Chrome DOMStorage protocol methods.

The DOMStorage protocol queries and modifies DOM storage.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/
*/
type DOMStorage struct{}

/*
Clear clears  a stored item.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-clear
*/
func (DOMStorage) Clear(
	socket *sock.Socket,
	params *domStorage.ClearParams,
) error {
	command := &sock.Command{
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
	socket *sock.Socket,
) error {
	command := &sock.Command{
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
	socket *sock.Socket,
) error {
	command := &sock.Command{
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
	socket *sock.Socket,
	params *domStorage.GetItemsParams,
) (domStorage.GetItemsResult, error) {
	command := &sock.Command{
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
	socket *sock.Socket,
	params *domStorage.RemoveItemParams,
) error {
	command := &sock.Command{
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
	socket *sock.Socket,
	params *domStorage.SetItemParams,
) error {
	command := &sock.Command{
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
	socket *sock.Socket,
	callback func(event *domStorage.ItemAddedEvent),
) {
	handler := sock.NewEventHandler(
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
	socket *sock.Socket,
	callback func(event *domStorage.ItemRemovedEvent),
) {
	handler := sock.NewEventHandler(
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
	socket *sock.Socket,
	callback func(event *domStorage.ItemUpdatedEvent),
) {
	handler := sock.NewEventHandler(
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
	socket *sock.Socket,
	callback func(event *domStorage.ItemsClearedEvent),
) {
	handler := sock.NewEventHandler(
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

package protocol

import (
	"encoding/json"

	domStorage "github.com/mkenney/go-chrome/protocol/dom_storage"
	sock "github.com/mkenney/go-chrome/socket"
	log "github.com/sirupsen/logrus"
)

/*
DOMStorage provides a namespace for the Chrome DOMStorage protocol methods. The DOMStorage protocol
queries and modifies DOM storage.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/
*/
var DOMStorage = DOMStorageProtocol{}

/*
DOMStorageProtocol defines the DOMStorage protocol methods.
*/
type DOMStorageProtocol struct{}

/*
Clear clears  a stored item.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-clear
*/
func (DOMStorageProtocol) Clear(
	socket sock.Socketer,
	params *domStorage.ClearParams,
) error {
	command := sock.NewCommand("DOMStorage.clear", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
Disable disables storage tracking, prevents storage events from being sent to the client.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-disable
*/
func (DOMStorageProtocol) Disable(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("DOMStorage.disable", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
Enable enables storage tracking, storage events will now be delivered to the client.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-enable
*/
func (DOMStorageProtocol) Enable(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("DOMStorage.enable", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
GetItems gets a stored item.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-getDOMStorageItems
*/
func (DOMStorageProtocol) GetItems(
	socket sock.Socketer,
	params *domStorage.GetItemsParams,
) (*domStorage.GetItemsResult, error) {
	command := sock.NewCommand("DOMStorage.getDOMStorageItems", params)
	result := &domStorage.GetItemsResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
RemoveItem removes  a stored item.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-removeDOMStorageItem
*/
func (DOMStorageProtocol) RemoveItem(
	socket sock.Socketer,
	params *domStorage.RemoveItemParams,
) error {
	command := sock.NewCommand("DOMStorage.removeDOMStorageItem", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetItem sets a stored item.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-setDOMStorageItem
*/
func (DOMStorageProtocol) SetItem(
	socket sock.Socketer,
	params *domStorage.SetItemParams,
) error {
	command := sock.NewCommand("DOMStorage.setDOMStorageItem", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
OnItemAdded adds a handler to the DOMStorage.domStorageItemAdded event.
DOMStorage.domStorageItemAdded fires when an item is added to DOM storage.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#event-domStorageItemAdded
*/
func (DOMStorageProtocol) OnItemAdded(
	socket sock.Socketer,
	callback func(event *domStorage.ItemAddedEvent),
) {
	handler := sock.NewEventHandler(
		"DOMStorage.domStorageItemAdded",
		func(response *sock.Response) {
			event := &domStorage.ItemAddedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
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
func (DOMStorageProtocol) OnItemRemoved(
	socket sock.Socketer,
	callback func(event *domStorage.ItemRemovedEvent),
) {
	handler := sock.NewEventHandler(
		"DOMStorage.domStorageItemRemoved",
		func(response *sock.Response) {
			event := &domStorage.ItemRemovedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
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
func (DOMStorageProtocol) OnItemUpdated(
	socket sock.Socketer,
	callback func(event *domStorage.ItemUpdatedEvent),
) {
	handler := sock.NewEventHandler(
		"DOMStorage.domStorageItemUpdated",
		func(response *sock.Response) {
			event := &domStorage.ItemUpdatedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
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
func (DOMStorageProtocol) OnItemsCleared(
	socket sock.Socketer,
	callback func(event *domStorage.ItemsClearedEvent),
) {
	handler := sock.NewEventHandler(
		"DOMStorage.domStorageItemsCleared",
		func(response *sock.Response) {
			event := &domStorage.ItemsClearedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

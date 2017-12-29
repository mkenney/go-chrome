package socket

import (
	"encoding/json"

	domStorage "github.com/mkenney/go-chrome/protocol/dom_storage"
	log "github.com/sirupsen/logrus"
)

/*
DOMStorageProtocol provides a namespace for the Chrome DOMStorage protocol methods. The DOMStorage
protocol queries and modifies DOM storage.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/
*/
type DOMStorageProtocol struct {
	Socket Socketer
}

/*
Clear clears  a stored item.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-clear
*/
func (protocol *DOMStorageProtocol) Clear(
	params *domStorage.ClearParams,
) error {
	command := NewCommand("DOMStorage.clear", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
Disable disables storage tracking, prevents storage events from being sent to the client.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-disable
*/
func (protocol *DOMStorageProtocol) Disable() error {
	command := NewCommand("DOMStorage.disable", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
Enable enables storage tracking, storage events will now be delivered to the client.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-enable
*/
func (protocol *DOMStorageProtocol) Enable() error {
	command := NewCommand("DOMStorage.enable", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
GetItems gets a stored item.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-getDOMStorageItems
*/
func (protocol *DOMStorageProtocol) GetItems(
	params *domStorage.GetItemsParams,
) (*domStorage.GetItemsResult, error) {
	command := NewCommand("DOMStorage.getDOMStorageItems", params)
	result := &domStorage.GetItemsResult{}
	protocol.Socket.SendCommand(command)

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
func (protocol *DOMStorageProtocol) RemoveItem(
	params *domStorage.RemoveItemParams,
) error {
	command := NewCommand("DOMStorage.removeDOMStorageItem", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetItem sets a stored item.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-setDOMStorageItem
*/
func (protocol *DOMStorageProtocol) SetItem(
	params *domStorage.SetItemParams,
) error {
	command := NewCommand("DOMStorage.setDOMStorageItem", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
OnItemAdded adds a handler to the DOMStorage.domStorageItemAdded event.
DOMStorage.domStorageItemAdded fires when an item is added to DOM storage.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#event-domStorageItemAdded
*/
func (protocol *DOMStorageProtocol) OnItemAdded(
	callback func(event *domStorage.ItemAddedEvent),
) {
	handler := NewEventHandler(
		"DOMStorage.domStorageItemAdded",
		func(response *Response) {
			event := &domStorage.ItemAddedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnItemRemoved adds a handler to the DOMStorage.domStorageItemRemoved event.
DOMStorage.domStorageItemRemoved fires when an item is removed from DOM storage.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#event-domStorageItemRemoved
*/
func (protocol *DOMStorageProtocol) OnItemRemoved(
	callback func(event *domStorage.ItemRemovedEvent),
) {
	handler := NewEventHandler(
		"DOMStorage.domStorageItemRemoved",
		func(response *Response) {
			event := &domStorage.ItemRemovedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnItemUpdated adds a handler to the DOMStorage.domStorageItemUpdated event.
DOMStorage.domStorageItemUpdated fires when an item in DOM storage is updated.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#event-domStorageItemUpdated
*/
func (protocol *DOMStorageProtocol) OnItemUpdated(
	callback func(event *domStorage.ItemUpdatedEvent),
) {
	handler := NewEventHandler(
		"DOMStorage.domStorageItemUpdated",
		func(response *Response) {
			event := &domStorage.ItemUpdatedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnItemsCleared adds a handler to the DOMStorage.domStorageItemsCleared event.
DOMStorage.domStorageItemsCleared fires when items in DOM storage are cleared.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#event-domStorageItemsCleared
*/
func (protocol *DOMStorageProtocol) OnItemsCleared(
	callback func(event *domStorage.ItemsClearedEvent),
) {
	handler := NewEventHandler(
		"DOMStorage.domStorageItemsCleared",
		func(response *Response) {
			event := &domStorage.ItemsClearedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

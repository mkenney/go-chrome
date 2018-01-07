package socket

import (
	"encoding/json"

	domStorage "github.com/mkenney/go-chrome/cdtp/dom_storage"
)

/*
DOMStorageProtocol provides a namespace for the Chrome DOMStorage protocol
methods. The DOMStorage protocol queries and modifies DOM storage.

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
) chan *domStorage.ClearResult {
	resultChan := make(chan *domStorage.ClearResult)
	command := NewCommand(protocol.Socket, "DOMStorage.clear", params)
	result := &domStorage.ClearResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
Disable disables storage tracking, prevents storage events from being sent to
the client.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-disable
*/
func (protocol *DOMStorageProtocol) Disable() chan *domStorage.DisableResult {
	resultChan := make(chan *domStorage.DisableResult)
	command := NewCommand(protocol.Socket, "DOMStorage.disable", nil)
	result := &domStorage.DisableResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
Enable enables storage tracking, storage events will now be delivered to the
client.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-enable
*/
func (protocol *DOMStorageProtocol) Enable() chan *domStorage.EnableResult {
	resultChan := make(chan *domStorage.EnableResult)
	command := NewCommand(protocol.Socket, "DOMStorage.enable", nil)
	result := &domStorage.EnableResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
GetItems gets a stored item.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-getDOMStorageItems
*/
func (protocol *DOMStorageProtocol) GetItems(
	params *domStorage.GetItemsParams,
) chan *domStorage.GetItemsResult {
	resultChan := make(chan *domStorage.GetItemsResult)
	command := NewCommand(protocol.Socket, "DOMStorage.getDOMStorageItems", params)
	result := &domStorage.GetItemsResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
}

/*
RemoveItem removes  a stored item.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-removeDOMStorageItem
*/
func (protocol *DOMStorageProtocol) RemoveItem(
	params *domStorage.RemoveItemParams,
) chan *domStorage.RemoveItemResult {
	resultChan := make(chan *domStorage.RemoveItemResult)
	command := NewCommand(protocol.Socket, "DOMStorage.removeDOMStorageItem", params)
	result := &domStorage.RemoveItemResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
SetItem sets a stored item.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-setDOMStorageItem
*/
func (protocol *DOMStorageProtocol) SetItem(
	params *domStorage.SetItemParams,
) chan *domStorage.SetItemResult {
	resultChan := make(chan *domStorage.SetItemResult)
	command := NewCommand(protocol.Socket, "DOMStorage.setDOMStorageItem", params)
	result := &domStorage.SetItemResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
	}()

	return resultChan
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
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
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
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
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
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
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
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

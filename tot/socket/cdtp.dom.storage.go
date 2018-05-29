package socket

import (
	"encoding/json"

	"github.com/mkenney/go-chrome/tot/cdtp/dom/storage"
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
	params *storage.ClearParams,
) <-chan *storage.ClearResult {
	resultChan := make(chan *storage.ClearResult)
	command := NewCommand(protocol.Socket, "DOMStorage.clear", params)
	result := &storage.ClearResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
Disable disables storage tracking, prevents storage events from being sent to
the client.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-disable
*/
func (protocol *DOMStorageProtocol) Disable() <-chan *storage.DisableResult {
	resultChan := make(chan *storage.DisableResult)
	command := NewCommand(protocol.Socket, "DOMStorage.disable", nil)
	result := &storage.DisableResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
Enable enables storage tracking, storage events will now be delivered to the
client.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-enable
*/
func (protocol *DOMStorageProtocol) Enable() <-chan *storage.EnableResult {
	resultChan := make(chan *storage.EnableResult)
	command := NewCommand(protocol.Socket, "DOMStorage.enable", nil)
	result := &storage.EnableResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
GetItems gets a stored item.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-getDOMStorageItems
*/
func (protocol *DOMStorageProtocol) GetItems(
	params *storage.GetItemsParams,
) <-chan *storage.GetItemsResult {
	resultChan := make(chan *storage.GetItemsResult)
	command := NewCommand(protocol.Socket, "DOMStorage.getDOMStorageItems", params)
	result := &storage.GetItemsResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
RemoveItem removes  a stored item.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-removeDOMStorageItem
*/
func (protocol *DOMStorageProtocol) RemoveItem(
	params *storage.RemoveItemParams,
) <-chan *storage.RemoveItemResult {
	resultChan := make(chan *storage.RemoveItemResult)
	command := NewCommand(protocol.Socket, "DOMStorage.removeDOMStorageItem", params)
	result := &storage.RemoveItemResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
SetItem sets a stored item.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-setDOMStorageItem
*/
func (protocol *DOMStorageProtocol) SetItem(
	params *storage.SetItemParams,
) <-chan *storage.SetItemResult {
	resultChan := make(chan *storage.SetItemResult)
	command := NewCommand(protocol.Socket, "DOMStorage.setDOMStorageItem", params)
	result := &storage.SetItemResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
OnItemAdded adds a handler to the DOMStorage.domStorageItemAdded event.
DOMStorage.domStorageItemAdded fires when an item is added to DOM storage.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#event-domStorageItemAdded
*/
func (protocol *DOMStorageProtocol) OnItemAdded(
	callback func(event *storage.ItemAddedEvent),
) {
	handler := NewEventHandler(
		"DOMStorage.domStorageItemAdded",
		func(response *Response) {
			event := &storage.ItemAddedEvent{}
			json.Unmarshal([]byte(response.Params), event)
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
	callback func(event *storage.ItemRemovedEvent),
) {
	handler := NewEventHandler(
		"DOMStorage.domStorageItemRemoved",
		func(response *Response) {
			event := &storage.ItemRemovedEvent{}
			json.Unmarshal([]byte(response.Params), event)
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
	callback func(event *storage.ItemUpdatedEvent),
) {
	handler := NewEventHandler(
		"DOMStorage.domStorageItemUpdated",
		func(response *Response) {
			event := &storage.ItemUpdatedEvent{}
			json.Unmarshal([]byte(response.Params), event)
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
	callback func(event *storage.ItemsClearedEvent),
) {
	handler := NewEventHandler(
		"DOMStorage.domStorageItemsCleared",
		func(response *Response) {
			event := &storage.ItemsClearedEvent{}
			json.Unmarshal([]byte(response.Params), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

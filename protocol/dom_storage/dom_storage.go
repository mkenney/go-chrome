/*
Package domStorage provides type definitions for use with the Chrome DOMStorage protocol

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/
*/
package domStorage

/*
ClearParams represents DOMStorage.clear parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-clear
*/
type ClearParams struct {
	StorageID StorageID `json:"storageId"`
}

/*
GetItemsParams represents DOMStorage.getDOMStorageItems parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-getDOMStorageItems
*/
type GetItemsParams struct {
	StorageID StorageID `json:"storageId"`
}

/*
GetItemsResult represents the result of calls to DOMStorage.getDOMStorageItems.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-getDOMStorageItems
*/
type GetItemsResult struct {
	Entries []Item `json:"entries"`
}

/*
RemoveItemParams represents DOMStorage.removeDOMStorageItem parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-removeDOMStorageItem
*/
type RemoveItemParams struct {
	StorageID StorageID `json:"storageId"`
	Key       string    `json:"key"`
}

/*
SetItemParams represents DOMStorage.setDOMStorageItem parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-setDOMStorageItem
*/
type SetItemParams struct {
	StorageID StorageID `json:"storageId"`
	Key       string    `json:"key"`
	Value     string    `json:"value"`
}

/*
ItemAddedEvent represents DOM.domStorageItemAdded event data.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#event-domStorageItemAdded
*/
type ItemAddedEvent struct {
	StorageID StorageID `json:"storageId"`
	Key       string    `json:"key"`
	NewValue  string    `json:"newValue"`
}

/*
ItemRemovedEvent represents DOM.domStorageItemRemoved event data.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#event-domStorageItemRemoved
*/
type ItemRemovedEvent struct {
	StorageID StorageID `json:"storageId"`
	Key       string    `json:"key"`
}

/*
ItemUpdatedEvent represents DOM.domStorageItemUpdated event data.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#event-domStorageItemUpdated
*/
type ItemUpdatedEvent struct {
	// Storage ID.
	StorageID StorageID `json:"storageId"`

	// Key.
	Key string `json:"key"`

	// Old value.
	OldValue string `json:"oldValue"`

	// New value.
	NewValue string `json:"newValue"`
}

/*
ItemsClearedEvent represents DOM.domStorageItemsCleared event data.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#event-domStorageItemsCleared
*/
type ItemsClearedEvent struct {
	// Storage ID.
	StorageID StorageID `json:"storageId"`
}

/*
StorageID is a DOM Storage identifier.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#type-StorageId
*/
type StorageID struct {
	// Security origin for the storage.
	SecurityOrigin string `json:"securityOrigin"`

	// Whether the storage is local storage (not session storage).
	IsLocalStorage bool `json:"isLocalStorage"`
}

/*
Item is a DOM Storage item.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#type-Item
*/
type Item []interface{}

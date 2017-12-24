package DOMStorage

/*
ClearParams represents DOMStorage.clear parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-clear
*/
type ClearParams struct {
	StorageID StorageID `json:"storageId"`
}

/*
GetDOMStorageItemsParams represents DOMStorage.getDOMStorageItems parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-getDOMStorageItems
*/
type GetDOMStorageItemsParams struct {
	StorageID StorageID `json:"storageId"`
}

/*
GetDOMStorageItemsResult represents the result of calls to DOMStorage.getDOMStorageItems.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-getDOMStorageItems
*/
type GetDOMStorageItemsResult struct {
	Entries []Item `json:"entries"`
}

/*
RemoveDOMStorageItemParams represents DOMStorage.removeDOMStorageItem parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-removeDOMStorageItem
*/
type RemoveDOMStorageItemParams struct {
	StorageID StorageID `json:"storageId"`
	Key       string    `json:"key"`
}

/*
SetDOMStorageItemParams represents DOMStorage.setDOMStorageItem parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-setDOMStorageItem
*/
type SetDOMStorageItemParams struct {
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

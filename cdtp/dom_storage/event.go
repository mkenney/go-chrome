package domStorage

/*
ItemAddedEvent represents DOM.domStorageItemAdded event data.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#event-domStorageItemAdded
*/
type ItemAddedEvent struct {
	StorageID *StorageID `json:"storageId"`
	Key       string     `json:"key"`
	NewValue  string     `json:"newValue"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
ItemRemovedEvent represents DOM.domStorageItemRemoved event data.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#event-domStorageItemRemoved
*/
type ItemRemovedEvent struct {
	StorageID *StorageID `json:"storageId"`
	Key       string     `json:"key"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
ItemUpdatedEvent represents DOM.domStorageItemUpdated event data.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#event-domStorageItemUpdated
*/
type ItemUpdatedEvent struct {
	// Storage ID.
	StorageID *StorageID `json:"storageId"`

	// Key.
	Key string `json:"key"`

	// Old value.
	OldValue string `json:"oldValue"`

	// New value.
	NewValue string `json:"newValue"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
ItemsClearedEvent represents DOM.domStorageItemsCleared event data.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#event-domStorageItemsCleared
*/
type ItemsClearedEvent struct {
	// Storage ID.
	StorageID *StorageID `json:"storageId"`

	// Error information related to this event
	Err error `json:"-"`
}

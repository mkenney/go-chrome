package DOMStorage

/*
ClearParams represents DOMStorage.clear parameters.
*/
type ClearParams struct {
	StorageID StorageID `json:"storageId"`
}

/*
GetDOMStorageItemsParams represents DOMStorage.getDOMStorageItems parameters.
*/
type GetDOMStorageItemsParams struct {
	StorageID StorageID `json:"storageId"`
}

/*
GetDOMStorageItemsResult represents the result of calls to DOMStorage.getDOMStorageItems.
*/
type GetDOMStorageItemsResult struct {
	Entries []Item `json:"entries"`
}

/*
RemoveDOMStorageItemParams represents DOMStorage.removeDOMStorageItem parameters.
*/
type RemoveDOMStorageItemParams struct {
	StorageID StorageID `json:"storageId"`
	Key       string    `json:"key"`
}

/*
SetDOMStorageItemParams represents DOMStorage.setDOMStorageItem parameters.
*/
type SetDOMStorageItemParams struct {
	StorageID StorageID `json:"storageId"`
	Key       string    `json:"key"`
	Value     string    `json:"value"`
}

/*
StorageID is a DOM Storage identifier.
*/
type StorageID struct {
	// Security origin for the storage.
	SecurityOrigin string `json:"securityOrigin"`

	// Whether the storage is local storage (not session storage).
	IsLocalStorage bool `json:"isLocalStorage"`
}

/*
Item is a DOM Storage item.
*/
type Item []interface{}

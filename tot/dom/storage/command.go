package storage

/*
ClearParams represents DOMStorage.clear parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-clear
*/
type ClearParams struct {
	StorageID *ID `json:"storageId"`
}

/*
ClearResult represents the result of calls to DOMStorage.clear.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-clear
*/
type ClearResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
DisableResult represents the result of calls to DOMStorage.disable.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-disable
*/
type DisableResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
EnableResult represents the result of calls to DOMStorage.enable.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-enable
*/
type EnableResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
GetItemsParams represents DOMStorage.getDOMStorageItems parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-getDOMStorageItems
*/
type GetItemsParams struct {
	StorageID *ID `json:"storageId"`
}

/*
GetItemsResult represents the result of calls to DOMStorage.getDOMStorageItems.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-getDOMStorageItems
*/
type GetItemsResult struct {
	Entries []Item `json:"entries"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
RemoveItemParams represents DOMStorage.removeDOMStorageItem parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-removeDOMStorageItem
*/
type RemoveItemParams struct {
	StorageID *ID    `json:"storageId"`
	Key       string `json:"key"`
}

/*
RemoveItemResult represents the result of calls to DOMStorage.removeItem.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-removeItem
*/
type RemoveItemResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SetItemParams represents DOMStorage.setDOMStorageItem parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-setDOMStorageItem
*/
type SetItemParams struct {
	StorageID *ID    `json:"storageId"`
	Key       string `json:"key"`
	Value     string `json:"value"`
}

/*
SetItemResult represents the result of calls to DOMStorage.setItem.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-setItem
*/
type SetItemResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

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

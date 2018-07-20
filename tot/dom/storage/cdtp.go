/*
Package storage provides type definitions for use with the Chrome DOMStorage protocol

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/
*/
package storage

/*
ID is a DOM Storage identifier.

https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#type-StorageId
*/
type ID struct {
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

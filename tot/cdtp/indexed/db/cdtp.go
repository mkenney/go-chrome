/*
Package db provides type definitions for use with the Chrome IndexedDB protocol

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/
*/
package db

import (
	"github.com/mkenney/go-chrome/tot/cdtp/runtime"
)

/*
DatabaseWithObjectStores is a database with an array of object stores.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#type-DatabaseWithObjectStores
*/
type DatabaseWithObjectStores struct {
	// Database name.
	Name string `json:"name"`

	// Database version.
	Version int `json:"version"`

	// Object stores in this database.
	ObjectStores []*ObjectStore `json:"objectStores"`
}

/*
ObjectStore is a object store

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#type-ObjectStore
*/
type ObjectStore struct {
	// Object store name.
	Name string `json:"name"`

	// Object store key path.
	KeyPath *KeyPath `json:"keyPath"`

	// If true, object store has auto increment flag set.
	AutoIncrement bool `json:"autoIncrement"`

	// Indexes in this object store.
	Indexes []*ObjectStoreIndex `json:"indexes"`
}

/*
ObjectStoreIndex is an object store index.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#type-ObjectStoreIndex
*/
type ObjectStoreIndex struct {
	// Index name.
	Name string `json:"name"`

	// Index key path.
	KeyPath *KeyPath `json:"keyPath"`

	// If true, index is unique.
	Unique bool `json:"unique"`

	// If true, index allows multiple entries for a key.
	MultiEntry bool `json:"multiEntry"`
}

/*
Key is a key.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#type-Key
*/
type Key struct {
	// Key type. Allowed values:
	//	- number
	//	- string
	//	- date
	//	- array
	Type KeyTypeEnum `json:"type"`

	// Optional. Number value.
	Number float64 `json:"number,omitempty"`

	// Optional. String value.
	String string `json:"string,omitempty"`

	// Optional. Date value.
	Date float64 `json:"date,omitempty"`

	// Optional. Array value.
	Array []*Key `json:"array,omitempty"`
}

/*
KeyRange is a key range.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#type-KeyRange
*/
type KeyRange struct {
	// Optional. Lower bound.
	Lower *Key `json:"lower,omitempty"`

	// Optional. Upper bound.
	Upper *Key `json:"upper,omitempty"`

	// If true lower bound is open.
	LowerOpen bool `json:"lowerOpen"`

	// If true upper bound is open.
	UpperOpen bool `json:"upperOpen"`
}

/*
DataEntry is a data entry.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#type-DataEntry
*/
type DataEntry struct {
	// Key object.
	Key *runtime.RemoteObject `json:"key"`

	// Primary key object.
	PrimaryKey *runtime.RemoteObject `json:"primaryKey"`

	// Value object.
	Value *runtime.RemoteObject `json:"value"`
}

/*
KeyPath is a key path.

https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#type-KeyPath
*/
type KeyPath struct {
	// Key path type. Allowed values:
	//	- null
	//	- string
	//	- array
	Type KeyPathTypeEnum `json:"type"`

	// Optional. String value.
	String string `json:"string,omitempty"`

	// Optional. Array value.
	Array []string `json:"array,omitempty"`
}

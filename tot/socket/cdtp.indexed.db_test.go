package socket

import (
	"testing"

	"github.com/mkenney/go-chrome/tot/indexed/db"
)

func TestIndexedDBClearObjectStore(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &db.ClearObjectStoreParams{
		SecurityOrigin:  "security-origin",
		DatabaseName:    "database-name",
		ObjectStoreName: "object-store-name",
	}
	mockResult := &db.ClearObjectStoreResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.IndexedDB().ClearObjectStore(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.IndexedDB().ClearObjectStore(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestIndexedDBDeleteDatabase(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &db.DeleteDatabaseParams{
		SecurityOrigin:  "security-origin",
		DatabaseName:    "database-name",
		ObjectStoreName: "object-store-name",
	}
	mockResult := &db.DeleteDatabaseResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.IndexedDB().DeleteDatabase(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.IndexedDB().DeleteDatabase(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestIndexedDBDeleteObjectStoreEntries(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &db.DeleteObjectStoreEntriesParams{
		SecurityOrigin: "security-origin",
		DatabaseName:   "database-name",
	}
	mockResult := &db.DeleteObjectStoreEntriesResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.IndexedDB().DeleteObjectStoreEntries(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.IndexedDB().DeleteObjectStoreEntries(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestIndexedDBDisable(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &db.DisableResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.IndexedDB().Disable()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.IndexedDB().Disable()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestIndexedDBEnable(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &db.EnableResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.IndexedDB().Enable()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.IndexedDB().Enable()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestIndexedDBRequestData(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &db.RequestDataParams{
		SecurityOrigin:  "security-origin",
		DatabaseName:    "database-name",
		ObjectStoreName: "object-store-name",
		IndexName:       "index-name",
		SkipCount:       1,
		PageSize:        1,
		KeyRange: &db.KeyRange{
			Lower: &db.Key{
				Type:   db.KeyType.Number,
				Number: 1,
				String: "string",
				Date:   1,
				Array:  []*db.Key{},
			},
			Upper: &db.Key{
				Type:   db.KeyType.Number,
				Number: 1,
				String: "string",
				Date:   1,
				Array:  []*db.Key{},
			},
			LowerOpen: true,
			UpperOpen: true,
		},
	}
	mockResult := &db.RequestDataResult{
		ObjectStoreDataEntries: []*db.DataEntry{{
			Key:        nil,
			PrimaryKey: nil,
			Value:      nil,
		}},
		HasMore: true,
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.IndexedDB().RequestData(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.ObjectStoreDataEntries[0].Key != result.ObjectStoreDataEntries[0].Key {
		t.Errorf("Expected %v, got %v", mockResult.ObjectStoreDataEntries[0].Key, result.ObjectStoreDataEntries[0].Key)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.IndexedDB().RequestData(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestIndexedDBRequestDatabase(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &db.RequestDatabaseParams{
		SecurityOrigin: "security-origin",
		DatabaseName:   "database-name",
	}
	mockResult := &db.RequestDatabaseResult{
		DatabaseWithObjectStores: &db.DatabaseWithObjectStores{
			Name:    "name",
			Version: 1,
			ObjectStores: []*db.ObjectStore{{
				Name: "name",
				KeyPath: &db.KeyPath{
					Type:   db.KeyPathType.Null,
					String: "string",
					Array:  []string{"string1", "string2"},
				},
				AutoIncrement: true,
				Indexes: []*db.ObjectStoreIndex{{
					Name: "name",
					KeyPath: &db.KeyPath{
						Type:   db.KeyPathType.Null,
						String: "string",
						Array:  []string{"string1", "string2"},
					},
					Unique:     true,
					MultiEntry: true,
				}},
			}},
		},
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.IndexedDB().RequestDatabase(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.DatabaseWithObjectStores.ObjectStores[0].KeyPath.Type != result.DatabaseWithObjectStores.ObjectStores[0].KeyPath.Type {
		t.Errorf("Expected '%s', got '%s'", mockResult.DatabaseWithObjectStores.ObjectStores[0].KeyPath.Type, result.DatabaseWithObjectStores.ObjectStores[0].KeyPath.Type)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.IndexedDB().RequestDatabase(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestIndexedDBRequestDatabaseNames(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &db.RequestDatabaseNamesParams{
		SecurityOrigin: "security-origin",
	}
	mockResult := &db.RequestDatabaseNamesResult{
		DatabaseNames: []string{"db1", "db2"},
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.IndexedDB().RequestDatabaseNames(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.DatabaseNames[0] != result.DatabaseNames[0] {
		t.Errorf("Expected '%s', got '%s'", mockResult.DatabaseNames[0], result.DatabaseNames[0])
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.IndexedDB().RequestDatabaseNames(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

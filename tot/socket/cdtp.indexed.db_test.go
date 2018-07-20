package socket

import (
	"encoding/json"
	"net/url"
	"testing"

	"github.com/mkenney/go-chrome/tot/indexed/db"
)

func TestIndexedDBClearObjectStore(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestIndexedDBClearObjectStore")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	params := &db.ClearObjectStoreParams{
		SecurityOrigin:  "security-origin",
		DatabaseName:    "database-name",
		ObjectStoreName: "object-store-name",
	}
	resultChan := mockSocket.IndexedDB().ClearObjectStore(params)
	mockResult := &db.ClearObjectStoreResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.IndexedDB().ClearObjectStore(params)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestIndexedDBDeleteDatabase(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestIndexedDBDeleteDatabase")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	params := &db.DeleteDatabaseParams{
		SecurityOrigin:  "security-origin",
		DatabaseName:    "database-name",
		ObjectStoreName: "object-store-name",
	}
	resultChan := mockSocket.IndexedDB().DeleteDatabase(params)
	mockResult := &db.DeleteDatabaseResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.IndexedDB().DeleteDatabase(params)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestIndexedDBDeleteObjectStoreEntries(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestIndexedDBDeleteObjectStoreEntries")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	params := &db.DeleteObjectStoreEntriesParams{
		SecurityOrigin: "security-origin",
		DatabaseName:   "database-name",
	}
	resultChan := mockSocket.IndexedDB().DeleteObjectStoreEntries(params)
	mockResult := &db.DeleteObjectStoreEntriesResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.IndexedDB().DeleteObjectStoreEntries(params)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestIndexedDBDisable(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestIndexedDBDisable")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.IndexedDB().Disable()
	mockResult := &db.DisableResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.IndexedDB().Disable()
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestIndexedDBEnable(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestIndexedDBEnable")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.IndexedDB().Enable()
	mockResult := &db.EnableResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.IndexedDB().Enable()
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestIndexedDBRequestData(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestIndexedDBRequestData")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

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
	resultChan := mockSocket.IndexedDB().RequestData(params)
	mockResult := &db.RequestDataResult{
		ObjectStoreDataEntries: []*db.DataEntry{{
			Key:        nil,
			PrimaryKey: nil,
			Value:      nil,
		}},
		HasMore: true,
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.ObjectStoreDataEntries[0].Key != result.ObjectStoreDataEntries[0].Key {
		t.Errorf("Expected %v, got %v", mockResult.ObjectStoreDataEntries[0].Key, result.ObjectStoreDataEntries[0].Key)
	}

	resultChan = mockSocket.IndexedDB().RequestData(params)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestIndexedDBRequestDatabase(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestIndexedDBRequestDatabase")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	params := &db.RequestDatabaseParams{
		SecurityOrigin: "security-origin",
		DatabaseName:   "database-name",
	}
	resultChan := mockSocket.IndexedDB().RequestDatabase(params)
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
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.DatabaseWithObjectStores.ObjectStores[0].KeyPath.Type != result.DatabaseWithObjectStores.ObjectStores[0].KeyPath.Type {
		t.Errorf("Expected '%s', got '%s'", mockResult.DatabaseWithObjectStores.ObjectStores[0].KeyPath.Type, result.DatabaseWithObjectStores.ObjectStores[0].KeyPath.Type)
	}

	resultChan = mockSocket.IndexedDB().RequestDatabase(params)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestIndexedDBRequestDatabaseNames(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestIndexedDBRequestDatabaseNames")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	params := &db.RequestDatabaseNamesParams{
		SecurityOrigin: "security-origin",
	}
	resultChan := mockSocket.IndexedDB().RequestDatabaseNames(params)
	mockResult := &db.RequestDatabaseNamesResult{
		DatabaseNames: []string{"db1", "db2"},
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.DatabaseNames[0] != result.DatabaseNames[0] {
		t.Errorf("Expected '%s', got '%s'", mockResult.DatabaseNames[0], result.DatabaseNames[0])
	}

	resultChan = mockSocket.IndexedDB().RequestDatabaseNames(params)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

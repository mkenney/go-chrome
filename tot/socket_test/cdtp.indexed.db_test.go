package socket

import (
	"encoding/json"
	"net/url"
	"testing"

	indexedDB "github.com/mkenney/go-chrome/tot/cdtp/indexed/db"
	"github.com/mkenney/go-chrome/tot/socket"
)

func TestIndexedDBClearObjectStore(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &indexedDB.ClearObjectStoreParams{
		SecurityOrigin:  "security-origin",
		DatabaseName:    "database-name",
		ObjectStoreName: "object-store-name",
	}
	resultChan := mockSocket.IndexedDB().ClearObjectStore(params)
	mockResult := &indexedDB.ClearObjectStoreResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&socket.Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &socket.Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.IndexedDB().ClearObjectStore(params)
	mockSocket.Conn().AddMockData(&socket.Response{
		ID: mockSocket.CurCommandID(),
		Error: &socket.Error{
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
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &indexedDB.DeleteDatabaseParams{
		SecurityOrigin:  "security-origin",
		DatabaseName:    "database-name",
		ObjectStoreName: "object-store-name",
	}
	resultChan := mockSocket.IndexedDB().DeleteDatabase(params)
	mockResult := &indexedDB.DeleteDatabaseResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&socket.Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &socket.Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.IndexedDB().DeleteDatabase(params)
	mockSocket.Conn().AddMockData(&socket.Response{
		ID: mockSocket.CurCommandID(),
		Error: &socket.Error{
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
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &indexedDB.DeleteObjectStoreEntriesParams{
		SecurityOrigin: "security-origin",
		DatabaseName:   "database-name",
	}
	resultChan := mockSocket.IndexedDB().DeleteObjectStoreEntries(params)
	mockResult := &indexedDB.DeleteObjectStoreEntriesResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&socket.Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &socket.Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.IndexedDB().DeleteObjectStoreEntries(params)
	mockSocket.Conn().AddMockData(&socket.Response{
		ID: mockSocket.CurCommandID(),
		Error: &socket.Error{
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
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.IndexedDB().Disable()
	mockResult := &indexedDB.DisableResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&socket.Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &socket.Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.IndexedDB().Disable()
	mockSocket.Conn().AddMockData(&socket.Response{
		ID: mockSocket.CurCommandID(),
		Error: &socket.Error{
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
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.IndexedDB().Enable()
	mockResult := &indexedDB.EnableResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&socket.Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &socket.Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.IndexedDB().Enable()
	mockSocket.Conn().AddMockData(&socket.Response{
		ID: mockSocket.CurCommandID(),
		Error: &socket.Error{
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
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &indexedDB.RequestDataParams{
		SecurityOrigin:  "security-origin",
		DatabaseName:    "database-name",
		ObjectStoreName: "object-store-name",
		IndexName:       "index-name",
		SkipCount:       1,
		PageSize:        1,
		KeyRange: &indexedDB.KeyRange{
			Lower: &indexedDB.Key{
				Type:   indexedDB.KeyType.Number,
				Number: 1,
				String: "string",
				Date:   1,
				Array:  []*indexedDB.Key{},
			},
			Upper: &indexedDB.Key{
				Type:   indexedDB.KeyType.Number,
				Number: 1,
				String: "string",
				Date:   1,
				Array:  []*indexedDB.Key{},
			},
			LowerOpen: true,
			UpperOpen: true,
		},
	}
	resultChan := mockSocket.IndexedDB().RequestData(params)
	mockResult := &indexedDB.RequestDataResult{
		ObjectStoreDataEntries: []*indexedDB.DataEntry{{
			Key:        nil,
			PrimaryKey: nil,
			Value:      nil,
		}},
		HasMore: true,
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&socket.Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &socket.Error{},
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
	mockSocket.Conn().AddMockData(&socket.Response{
		ID: mockSocket.CurCommandID(),
		Error: &socket.Error{
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
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &indexedDB.RequestDatabaseParams{
		SecurityOrigin: "security-origin",
		DatabaseName:   "database-name",
	}
	resultChan := mockSocket.IndexedDB().RequestDatabase(params)
	mockResult := &indexedDB.RequestDatabaseResult{
		DatabaseWithObjectStores: &indexedDB.DatabaseWithObjectStores{
			Name:    "name",
			Version: 1,
			ObjectStores: []*indexedDB.ObjectStore{{
				Name: "name",
				KeyPath: &indexedDB.KeyPath{
					Type:   indexedDB.KeyPathType.Null,
					String: "string",
					Array:  []string{"string1", "string2"},
				},
				AutoIncrement: true,
				Indexes: []*indexedDB.ObjectStoreIndex{{
					Name: "name",
					KeyPath: &indexedDB.KeyPath{
						Type:   indexedDB.KeyPathType.Null,
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
	mockSocket.Conn().AddMockData(&socket.Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &socket.Error{},
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
	mockSocket.Conn().AddMockData(&socket.Response{
		ID: mockSocket.CurCommandID(),
		Error: &socket.Error{
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
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &indexedDB.RequestDatabaseNamesParams{
		SecurityOrigin: "security-origin",
	}
	resultChan := mockSocket.IndexedDB().RequestDatabaseNames(params)
	mockResult := &indexedDB.RequestDatabaseNamesResult{
		DatabaseNames: []string{"db1", "db2"},
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&socket.Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &socket.Error{},
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
	mockSocket.Conn().AddMockData(&socket.Response{
		ID: mockSocket.CurCommandID(),
		Error: &socket.Error{
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

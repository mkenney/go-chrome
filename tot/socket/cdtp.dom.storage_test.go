package socket

import (
	"encoding/json"
	"net/url"
	"testing"

	"github.com/mkenney/go-chrome/tot/dom/storage"
)

func TestDOMStorageClear(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestDOMStorageClear")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.DOMStorage().Clear(&storage.ClearParams{
		StorageID: &storage.ID{
			SecurityOrigin: "security-origin",
			IsLocalStorage: true,
		},
	})
	mockResult := &storage.ClearResult{}
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

	resultChan = mockSocket.DOMStorage().Clear(&storage.ClearParams{
		StorageID: &storage.ID{
			SecurityOrigin: "security-origin",
			IsLocalStorage: true,
		},
	})
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

func TestDOMStorageDisable(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestDOMStorageDisable")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.DOMStorage().Disable()
	mockResult := &storage.DisableResult{}
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

	resultChan = mockSocket.DOMStorage().Disable()
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

func TestDOMStorageEnable(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestDOMStorageEnable")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.DOMStorage().Enable()
	mockResult := &storage.EnableResult{}
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

	resultChan = mockSocket.DOMStorage().Enable()
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

func TestDOMStorageGetItems(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestDOMStorageGetItems")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.DOMStorage().GetItems(&storage.GetItemsParams{
		StorageID: &storage.ID{
			SecurityOrigin: "security-origin",
			IsLocalStorage: true,
		},
	})
	mockResult := &storage.GetItemsResult{
		Entries: []storage.Item{{float64(1), "two"}},
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
	if mockResult.Entries[0][0].(float64) != result.Entries[0][0].(float64) {
		t.Errorf("Expected '%d', got '%d'", mockResult.Entries[0][0], result.Entries[0][0])
	}
	if mockResult.Entries[0][1].(string) != result.Entries[0][1].(string) {
		t.Errorf("Expected '%s', got '%s'", mockResult.Entries[0][0], result.Entries[0][0])
	}

	resultChan = mockSocket.DOMStorage().GetItems(&storage.GetItemsParams{
		StorageID: &storage.ID{
			SecurityOrigin: "security-origin",
			IsLocalStorage: true,
		},
	})
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

func TestDOMStorageRemoveItem(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestDOMStorageRemoveItem")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.DOMStorage().RemoveItem(&storage.RemoveItemParams{
		StorageID: &storage.ID{
			SecurityOrigin: "security-origin",
			IsLocalStorage: true,
		},
		Key: "item-key",
	})
	mockResult := &storage.RemoveItemResult{}
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

	resultChan = mockSocket.DOMStorage().RemoveItem(&storage.RemoveItemParams{
		StorageID: &storage.ID{
			SecurityOrigin: "security-origin",
			IsLocalStorage: true,
		},
		Key: "item-key",
	})
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

func TestDOMStorageSetItem(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestDOMStorageSetItem")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.DOMStorage().SetItem(&storage.SetItemParams{
		StorageID: &storage.ID{
			SecurityOrigin: "security-origin",
			IsLocalStorage: true,
		},
		Key:   "item-key",
		Value: "item value",
	})
	mockResult := &storage.SetItemResult{}
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

	resultChan = mockSocket.DOMStorage().SetItem(&storage.SetItemParams{
		StorageID: &storage.ID{
			SecurityOrigin: "security-origin",
			IsLocalStorage: true,
		},
		Key: "item-key",
	})
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

func TestDOMStorageOnItemAdded(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestDOMStorageOnItemAdded")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *storage.ItemAddedEvent)
	mockSocket.DOMStorage().OnItemAdded(func(eventData *storage.ItemAddedEvent) {
		resultChan <- eventData
	})
	mockResult := &storage.ItemAddedEvent{
		StorageID: &storage.ID{
			SecurityOrigin: "security-origin",
			IsLocalStorage: true,
		},
		Key:      "item-key",
		NewValue: "item value",
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "DOMStorage.domStorageItemAdded",
		Params: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	resultChan = make(chan *storage.ItemAddedEvent)
	mockSocket.DOMStorage().OnItemAdded(func(eventData *storage.ItemAddedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "DOMStorage.domStorageItemAdded",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMStorageOnItemRemoved(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestDOMStorageOnItemRemoved")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *storage.ItemRemovedEvent)
	mockSocket.DOMStorage().OnItemRemoved(func(eventData *storage.ItemRemovedEvent) {
		resultChan <- eventData
	})
	mockResult := &storage.ItemRemovedEvent{
		StorageID: &storage.ID{
			SecurityOrigin: "security-origin",
			IsLocalStorage: true,
		},
		Key: "item-key",
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "DOMStorage.domStorageItemRemoved",
		Params: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	resultChan = make(chan *storage.ItemRemovedEvent)
	mockSocket.DOMStorage().OnItemRemoved(func(eventData *storage.ItemRemovedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "DOMStorage.domStorageItemRemoved",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMStorageOnItemUpdated(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestDOMStorageOnItemUpdated")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *storage.ItemUpdatedEvent)
	mockSocket.DOMStorage().OnItemUpdated(func(eventData *storage.ItemUpdatedEvent) {
		resultChan <- eventData
	})
	mockResult := &storage.ItemUpdatedEvent{
		StorageID: &storage.ID{
			SecurityOrigin: "security-origin",
			IsLocalStorage: true,
		},
		Key:      "item-key",
		OldValue: "old-value",
		NewValue: "new-value",
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "DOMStorage.domStorageItemUpdated",
		Params: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	resultChan = make(chan *storage.ItemUpdatedEvent)
	mockSocket.DOMStorage().OnItemUpdated(func(eventData *storage.ItemUpdatedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "DOMStorage.domStorageItemUpdated",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMStorageOnItemsCleared(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestDOMStorageOnItemsCleared")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *storage.ItemsClearedEvent)
	mockSocket.DOMStorage().OnItemsCleared(func(eventData *storage.ItemsClearedEvent) {
		resultChan <- eventData
	})
	mockResult := &storage.ItemsClearedEvent{
		StorageID: &storage.ID{
			SecurityOrigin: "security-origin",
			IsLocalStorage: true,
		},
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "DOMStorage.domStorageItemsCleared",
		Params: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	resultChan = make(chan *storage.ItemsClearedEvent)
	mockSocket.DOMStorage().OnItemsCleared(func(eventData *storage.ItemsClearedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "DOMStorage.domStorageItemsCleared",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

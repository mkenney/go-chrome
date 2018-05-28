package socket

import (
	"encoding/json"
	"net/url"
	"testing"

	domStorage "github.com/mkenney/go-chrome/tot/cdtp/dom/storage"
)

func TestDOMStorageClear(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.DOMStorage().Clear(&domStorage.ClearParams{
		StorageID: &domStorage.ID{
			SecurityOrigin: "security-origin",
			IsLocalStorage: true,
		},
	})
	mockResult := &domStorage.ClearResult{}
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

	resultChan = mockSocket.DOMStorage().Clear(&domStorage.ClearParams{
		StorageID: &domStorage.ID{
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
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.DOMStorage().Disable()
	mockResult := &domStorage.DisableResult{}
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
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.DOMStorage().Enable()
	mockResult := &domStorage.EnableResult{}
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
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.DOMStorage().GetItems(&domStorage.GetItemsParams{
		StorageID: &domStorage.ID{
			SecurityOrigin: "security-origin",
			IsLocalStorage: true,
		},
	})
	mockResult := &domStorage.GetItemsResult{
		Entries: []domStorage.Item{{float64(1), "two"}},
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

	resultChan = mockSocket.DOMStorage().GetItems(&domStorage.GetItemsParams{
		StorageID: &domStorage.ID{
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
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.DOMStorage().RemoveItem(&domStorage.RemoveItemParams{
		StorageID: &domStorage.ID{
			SecurityOrigin: "security-origin",
			IsLocalStorage: true,
		},
		Key: "item-key",
	})
	mockResult := &domStorage.RemoveItemResult{}
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

	resultChan = mockSocket.DOMStorage().RemoveItem(&domStorage.RemoveItemParams{
		StorageID: &domStorage.ID{
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
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.DOMStorage().SetItem(&domStorage.SetItemParams{
		StorageID: &domStorage.ID{
			SecurityOrigin: "security-origin",
			IsLocalStorage: true,
		},
		Key:   "item-key",
		Value: "item value",
	})
	mockResult := &domStorage.SetItemResult{}
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

	resultChan = mockSocket.DOMStorage().SetItem(&domStorage.SetItemParams{
		StorageID: &domStorage.ID{
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
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *domStorage.ItemAddedEvent)
	mockSocket.DOMStorage().OnItemAdded(func(eventData *domStorage.ItemAddedEvent) {
		resultChan <- eventData
	})
	mockResult := &domStorage.ItemAddedEvent{
		StorageID: &domStorage.ID{
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

	resultChan = make(chan *domStorage.ItemAddedEvent)
	mockSocket.DOMStorage().OnItemAdded(func(eventData *domStorage.ItemAddedEvent) {
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
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *domStorage.ItemRemovedEvent)
	mockSocket.DOMStorage().OnItemRemoved(func(eventData *domStorage.ItemRemovedEvent) {
		resultChan <- eventData
	})
	mockResult := &domStorage.ItemRemovedEvent{
		StorageID: &domStorage.ID{
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

	resultChan = make(chan *domStorage.ItemRemovedEvent)
	mockSocket.DOMStorage().OnItemRemoved(func(eventData *domStorage.ItemRemovedEvent) {
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
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *domStorage.ItemUpdatedEvent)
	mockSocket.DOMStorage().OnItemUpdated(func(eventData *domStorage.ItemUpdatedEvent) {
		resultChan <- eventData
	})
	mockResult := &domStorage.ItemUpdatedEvent{
		StorageID: &domStorage.ID{
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

	resultChan = make(chan *domStorage.ItemUpdatedEvent)
	mockSocket.DOMStorage().OnItemUpdated(func(eventData *domStorage.ItemUpdatedEvent) {
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
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *domStorage.ItemsClearedEvent)
	mockSocket.DOMStorage().OnItemsCleared(func(eventData *domStorage.ItemsClearedEvent) {
		resultChan <- eventData
	})
	mockResult := &domStorage.ItemsClearedEvent{
		StorageID: &domStorage.ID{
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

	resultChan = make(chan *domStorage.ItemsClearedEvent)
	mockSocket.DOMStorage().OnItemsCleared(func(eventData *domStorage.ItemsClearedEvent) {
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

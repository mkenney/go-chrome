package socket

import (
	"testing"

	"github.com/mkenney/go-chrome/tot/dom/storage"
)

func TestDOMStorageClear(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &storage.ClearResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DOMStorage().Clear(&storage.ClearParams{
		StorageID: &storage.ID{
			SecurityOrigin: "security-origin",
			IsLocalStorage: true,
		},
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DOMStorage().Clear(&storage.ClearParams{
		StorageID: &storage.ID{
			SecurityOrigin: "security-origin",
			IsLocalStorage: true,
		},
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMStorageDisable(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &storage.DisableResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DOMStorage().Disable()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DOMStorage().Disable()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMStorageEnable(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &storage.EnableResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DOMStorage().Enable()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DOMStorage().Enable()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMStorageGetItems(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &storage.GetItemsResult{
		Entries: []storage.Item{{float64(1), "two"}},
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DOMStorage().GetItems(&storage.GetItemsParams{
		StorageID: &storage.ID{
			SecurityOrigin: "security-origin",
			IsLocalStorage: true,
		},
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Entries[0][0].(float64) != result.Entries[0][0].(float64) {
		t.Errorf("Expected '%d', got '%d'", mockResult.Entries[0][0], result.Entries[0][0])
	}
	if mockResult.Entries[0][1].(string) != result.Entries[0][1].(string) {
		t.Errorf("Expected '%s', got '%s'", mockResult.Entries[0][0], result.Entries[0][0])
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DOMStorage().GetItems(&storage.GetItemsParams{
		StorageID: &storage.ID{
			SecurityOrigin: "security-origin",
			IsLocalStorage: true,
		},
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMStorageRemoveItem(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &storage.RemoveItemResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DOMStorage().RemoveItem(&storage.RemoveItemParams{
		StorageID: &storage.ID{
			SecurityOrigin: "security-origin",
			IsLocalStorage: true,
		},
		Key: "item-key",
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DOMStorage().RemoveItem(&storage.RemoveItemParams{
		StorageID: &storage.ID{
			SecurityOrigin: "security-origin",
			IsLocalStorage: true,
		},
		Key: "item-key",
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMStorageSetItem(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &storage.SetItemResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DOMStorage().SetItem(&storage.SetItemParams{
		StorageID: &storage.ID{
			SecurityOrigin: "security-origin",
			IsLocalStorage: true,
		},
		Key:   "item-key",
		Value: "item value",
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DOMStorage().SetItem(&storage.SetItemParams{
		StorageID: &storage.ID{
			SecurityOrigin: "security-origin",
			IsLocalStorage: true,
		},
		Key: "item-key",
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMStorageOnItemAdded(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()

	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *storage.ItemAddedEvent)
	soc.DOMStorage().OnItemAdded(func(eventData *storage.ItemAddedEvent) {
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
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "DOMStorage.domStorageItemAdded",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "DOMStorage.domStorageItemAdded",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMStorageOnItemRemoved(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *storage.ItemRemovedEvent)
	soc.DOMStorage().OnItemRemoved(func(eventData *storage.ItemRemovedEvent) {
		resultChan <- eventData
	})

	mockResult := &storage.ItemRemovedEvent{
		StorageID: &storage.ID{
			SecurityOrigin: "security-origin",
			IsLocalStorage: true,
		},
		Key: "item-key",
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "DOMStorage.domStorageItemRemoved",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "DOMStorage.domStorageItemRemoved",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMStorageOnItemUpdated(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *storage.ItemUpdatedEvent)
	soc.DOMStorage().OnItemUpdated(func(eventData *storage.ItemUpdatedEvent) {
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
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "DOMStorage.domStorageItemUpdated",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "DOMStorage.domStorageItemUpdated",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMStorageOnItemsCleared(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *storage.ItemsClearedEvent)
	soc.DOMStorage().OnItemsCleared(func(eventData *storage.ItemsClearedEvent) {
		resultChan <- eventData
	})

	mockResult := &storage.ItemsClearedEvent{
		StorageID: &storage.ID{
			SecurityOrigin: "security-origin",
			IsLocalStorage: true,
		},
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "DOMStorage.domStorageItemsCleared",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "DOMStorage.domStorageItemsCleared",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

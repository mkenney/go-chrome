package socket

import (
	"testing"

	"github.com/mkenney/go-chrome/tot/io"
	"github.com/mkenney/go-chrome/tot/runtime"
)

func TestIOClose(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &io.CloseParams{
		Handle: io.StreamHandle("stream-handle"),
	}
	mockResult := &io.CloseResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.IO().Close(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.IO().Close(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestIORead(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &io.ReadParams{
		Handle: io.StreamHandle("stream-handle"),
		Offset: 1,
		Size:   1,
	}
	mockResult := &io.ReadResult{
		Base64Encoded: true,
		Data:          "data",
		EOF:           true,
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.IO().Read(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Base64Encoded != result.Base64Encoded {
		t.Errorf("Expected %v, got %v", mockResult.Base64Encoded, result.Base64Encoded)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.IO().Read(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestIOResolveBlob(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &io.ResolveBlobParams{
		ObjectID: runtime.RemoteObjectID("remote-object-id"),
	}
	mockResult := &io.ResolveBlobResult{
		UUID: "uuid",
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.IO().ResolveBlob(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.UUID != result.UUID {
		t.Errorf("Expected %v, got %v", mockResult.UUID, result.UUID)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.IO().ResolveBlob(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

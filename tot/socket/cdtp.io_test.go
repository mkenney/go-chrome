package socket

import (
	"encoding/json"
	"net/url"
	"testing"

	io "github.com/mkenney/go-chrome/tot/cdtp/io"
	runtime "github.com/mkenney/go-chrome/tot/cdtp/runtime"
)

func TestIOClose(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &io.CloseParams{
		Handle: io.StreamHandle("stream-handle"),
	}
	resultChan := mockSocket.IO().Close(params)
	mockResult := &io.CloseResult{}
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

	resultChan = mockSocket.IO().Close(params)
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

func TestIORead(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &io.ReadParams{
		Handle: io.StreamHandle("stream-handle"),
		Offset: 1,
		Size:   1,
	}
	resultChan := mockSocket.IO().Read(params)
	mockResult := &io.ReadResult{
		Base64Encoded: true,
		Data:          "data",
		EOF:           true,
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
	if mockResult.Base64Encoded != result.Base64Encoded {
		t.Errorf("Expected %v, got %v", mockResult.Base64Encoded, result.Base64Encoded)
	}

	resultChan = mockSocket.IO().Read(params)
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

func TestIOResolveBlob(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &io.ResolveBlobParams{
		ObjectID: runtime.RemoteObjectID("remote-object-id"),
	}
	resultChan := mockSocket.IO().ResolveBlob(params)
	mockResult := &io.ResolveBlobResult{
		UUID: "uuid",
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
	if mockResult.UUID != result.UUID {
		t.Errorf("Expected %v, got %v", mockResult.UUID, result.UUID)
	}

	resultChan = mockSocket.IO().ResolveBlob(params)
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

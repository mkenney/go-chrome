package socket

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"sync"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/mkenney/go-chrome/tot/tracing"
)

func NewMockChrome() *MockChrome {
	return &MockChrome{
		mux: &sync.Mutex{},
	}
}

var upgrader = websocket.Upgrader{}

var genericError = &Error{
	Code:    1,
	Data:    []byte(`"error data"`),
	Message: "error message",
}

type MockChrome struct {
	IgnoreInput bool
	Server      *httptest.Server
	URL         *url.URL

	mux  *sync.Mutex
	Data []MockData
}

type MockData struct {
	Err    *Error
	Result interface{}
	Method string
}

func (chrome *MockChrome) AddData(data MockData) {
	chrome.mux.Lock()
	chrome.Data = append(chrome.Data, data)
	chrome.mux.Unlock()
}

func (chrome *MockChrome) Close() {
	chrome.Server.Close()
}

func (chrome *MockChrome) handle(writer http.ResponseWriter, request *http.Request) {
	socket, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		return
	}
	defer socket.Close()
	for {
		payload := Payload{}
		response := &Response{}

		if !chrome.IgnoreInput {
			err := socket.ReadJSON(&payload)
			if err != nil {
				break
			}
			response.ID = payload.ID
		}
		chrome.mux.Lock()
		if 0 == len(chrome.Data) {
			chrome.mux.Unlock()
			break
		}
		fmt.Println(len(chrome.Data))
		response.Error = chrome.Data[0].Err
		response.Method = chrome.Data[0].Method

		var mockResultBytes []byte
		if nil != chrome.Data[0].Result {
			mockResultBytes, _ = json.Marshal(chrome.Data[0].Result)
		}

		if response.ID > 0 {
			response.Result = mockResultBytes
		} else {
			response.Params = mockResultBytes
		}

		chrome.Data = chrome.Data[1:]
		chrome.mux.Unlock()

		err = socket.WriteJSON(response)
		if err != nil {
			break
		}
	}
}

func (chrome *MockChrome) ListenAndServe() {
	chrome.Server = httptest.NewServer(http.HandlerFunc(chrome.handle))
	chrome.URL, _ = url.Parse("ws" + strings.TrimPrefix(chrome.Server.URL, "http"))
}

func TestMockSocket(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)

	chrome.AddData(MockData{
		&Error{},
		&tracing.GetCategoriesResult{
			Categories: []string{"cat1", "cat2"},
		},
		"",
	})
	resultChan := soc.Tracing().GetCategories()
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if "cat1" != result.Categories[0] {
		t.Errorf("Expected %s, got %s", "cat1", result.Categories[0])
	}

	chrome.AddData(MockData{
		&Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		nil,
		"",
	})
	resultChan = soc.Tracing().GetCategories()
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

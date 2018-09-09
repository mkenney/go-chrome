package socket

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/bdlm/log"
	"github.com/gorilla/websocket"
	"github.com/mkenney/go-chrome/tot/tracing"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		ForceTTY: true,
		//EnableTrace: true,
	})
	level, err := log.ParseLevel("debug")
	if nil == err {
		log.SetLevel(level)
	}
}

func NewMockChrome() *MockChrome {
	return &MockChrome{
		mux: &sync.Mutex{},
		logger: log.WithFields(log.Fields{
			"mock": "chrome",
		}),
	}
}

var upgrader = websocket.Upgrader{}

var genericError = &Error{
	Code:    1,
	Data:    []byte(`"error data"`),
	Message: "error message",
}

type MockChrome struct {
	breaker chan bool
	logger  *log.Entry
	sleep   time.Duration

	IgnoreInput bool
	Server      *httptest.Server
	URL         *url.URL

	mux  *sync.Mutex
	Data []MockData
}

var socketMux = &sync.Mutex{}

type MockData struct {
	ID     int
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

func (chrome *MockChrome) popMockData(socket *websocket.Conn) MockData {
	data := MockData{}
	for {
		chrome.mux.Lock()

		// no data, break
		if 0 == len(chrome.Data) {
			chrome.mux.Unlock()
			time.Sleep(1 * time.Millisecond)
			continue
		}

		data = chrome.Data[0]
		chrome.Data = chrome.Data[1:]
		chrome.mux.Unlock()

		break
	}
	chrome.logger.WithFields(log.Fields{"data": data}).
		Info("Fetching mock chrome socket response data")
	return data
}

func (chrome *MockChrome) handle(writer http.ResponseWriter, request *http.Request) {
	var err error
	var response *Response
	type ReadLoopData struct {
		Err     error
		Payload Payload
	}
	readLoopBreaker := make(chan bool)
	readLoopChan := make(chan ReadLoopData, 3)
	writeLoopBreaker := make(chan bool)
	writeLoopChan := make(chan MockData, 3)

	// upgrade connection to a websocket
	socket, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		return
	}

	// socket handler cleanup
	defer func() {
		chrome.logger.Warn("Shutting down mock chrome")
		readLoopBreaker <- true
		writeLoopBreaker <- true
		socket.Close()
	}()

	// launch the socket input read loop for this connection.
	go func() {
		if chrome.IgnoreInput {
			<-readLoopBreaker
			return
		}

		socketChan := make(chan ReadLoopData)
		go func() {
			for {
				var data ReadLoopData
				err := socket.ReadJSON(&data.Payload)
				if nil != err {
					data.Err = err
					// If the socket closes kill the loop
					if strings.Contains(data.Err.Error(), "abnormal closure") {
						chrome.logger.WithFields(log.Fields{"data": data}).
							Warn("socket closed, ending read loop")
						return
					}
				}
				chrome.logger.WithFields(log.Fields{"data": data}).
					Debug("returning socket input")
				socketChan <- data
			}
		}()

		for {
			select {
			case <-readLoopBreaker:
				chrome.logger.Debug("shutting down mock chrome socket input read loop")
				return
			case readLoopChan <- <-socketChan:
			}
		}
	}()

	// launch the mock data loop for this connection.
	go func() {
		if !chrome.IgnoreInput {
			<-writeLoopBreaker
			return
		}
		for {
			if chrome.sleep > 0 {
				time.Sleep(chrome.sleep)
			}
			select {
			case <-writeLoopBreaker:
				chrome.logger.Debug("shutting down mock chrome socket write loop")
				return
			case writeLoopChan <- chrome.popMockData(socket):
				chrome.logger.Debug("mock data written to mock chrome socket")
			}
		}
	}()

	// main event loop
	errCnt := 0
	for {
		response = nil

		select {
		case <-chrome.breaker:
			chrome.logger.Warn("Shutting down mock chrome event loop")
			return

		// Manage any data input.
		case data := <-readLoopChan:
			if nil != data.Err {
				errCnt++
				if errCnt > 10 {
					chrome.breaker <- true
				}
				chrome.logger.WithField("error", data.Err).Error("Mock chrome read loop returned an error")
			} else {
				errCnt = 0
			}

			mockData := chrome.popMockData(socket)
			mockData.ID = data.Payload.ID
			mockData.Method = data.Payload.Method
			if mockData.ID == 0 {
				mockData.Result = data.Payload.Params
			}
			chrome.logger.WithField("data", mockData).Info("Handling mock chrome socket input")
			writeLoopChan <- mockData

			// Return responses.
		case data := <-writeLoopChan:
			chrome.logger.WithField("data", data).Info("Building mock chrome socket response")

			response = &Response{}
			response.Error = data.Err
			response.Method = data.Method
			response.ID = data.ID

			var mockResultBytes []byte
			if nil != data.Result {
				mockResultBytes, _ = json.Marshal(data.Result)
			}
			if response.ID > 0 {
				response.Result = mockResultBytes
			} else {
				response.Params = mockResultBytes
			}

			chrome.logger.WithFields(log.Fields{"data": data, "response": response}).
				Info("Writing response to mock chrome socket connection")

			socketMux.Lock()
			err = socket.WriteJSON(response)
			socketMux.Unlock()

			if nil != err {
				chrome.logger.Errorf("%-v", err)
				break
			}
		}
	}
}

func (chrome *MockChrome) ListenAndServe() {
	chrome.Server = httptest.NewServer(http.HandlerFunc(chrome.handle))
	chrome.URL, _ = url.Parse("ws" + strings.TrimPrefix(chrome.Server.URL, "http"))
}

func (chrome *MockChrome) Sleep(duration time.Duration) {
	chrome.sleep = duration
}

func TestMockSocket(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer func() { chrome.Close() }()
	soc := New(chrome.URL)
	defer func() { soc.Stop() }()

	chrome.AddData(MockData{
		Err: &Error{},
		Result: &tracing.GetCategoriesResult{
			Categories: []string{"cat1", "cat2"},
		},
		Method: "Tracing.getCategories",
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
		Err: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Result: nil,
		Method: "Tracing.getCategories",
	})
	resultChan = soc.Tracing().GetCategories()
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

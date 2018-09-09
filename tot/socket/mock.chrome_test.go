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
		Info("fetching mock chrome socket response data")
	return data
}

type ReadLoopData struct {
	Err     error
	Payload Payload
}

// launchListener starts the websocket input read loop.
func (chrome *MockChrome) launchListener(
	websoc *websocket.Conn,
	breaker chan bool,
	dataChan chan ReadLoopData,
) {

	if chrome.IgnoreInput {
		<-breaker
		return
	}

	dataFeed := make(chan ReadLoopData)
	go func() {
		for {
			var data ReadLoopData
			err := websoc.ReadJSON(&data.Payload)
			if nil != err {
				data.Err = err
				// If the websocket closes kill the loop
				if strings.Contains(data.Err.Error(), "abnormal closure") {
					chrome.logger.WithFields(log.Fields{"data": data}).
						Warn("websocket closed, ending read loop")
					return
				}
			}
			chrome.logger.WithFields(log.Fields{"data": data}).
				Debug("returning websocket input")
			dataFeed <- data
		}
	}()

	for {
		select {
		case <-breaker:
			chrome.logger.Debug("shutting down mock chrome websocket input read loop")
			breaker <- true
			return
		case dataChan <- <-dataFeed:
		}
	}
}

// launchMockDataFeed starts the mock input data feeder.
func (chrome *MockChrome) launchMockDataFeed(
	websoc *websocket.Conn,
	breaker chan bool,
	dataChan chan MockData,
) {
	if !chrome.IgnoreInput {
		<-breaker
		return
	}
	for {
		if chrome.sleep > 0 {
			time.Sleep(chrome.sleep)
		}
		select {
		case <-breaker:
			chrome.logger.Debug("shutting down mock chrome websocket write loop")
			breaker <- true
			return
		case dataChan <- chrome.popMockData(websoc):
			chrome.logger.Debug("mock data written to mock chrome websocket")
		}
	}
}

// handle is the request data handler.
func (chrome *MockChrome) handle(writer http.ResponseWriter, request *http.Request) {
	var err error
	var response *Response
	readLoopBreaker := make(chan bool, 1)
	readLoopChan := make(chan ReadLoopData, 3)
	writeLoopBreaker := make(chan bool, 1)
	writeLoopChan := make(chan MockData, 3)

	// upgrade connection to a websocket
	websoc, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		return
	}

	// websocket handler cleanup
	defer func() {
		chrome.logger.Warn("shutting down mock chrome")
		readLoopBreaker <- true
		writeLoopBreaker <- true
		<-readLoopBreaker
		<-writeLoopBreaker
		websoc.Close()
	}()

	// launch the websocket input read loop for this connection.
	go chrome.launchListener(websoc, readLoopBreaker, readLoopChan)

	// launch the mock data loop for this connection.
	go chrome.launchMockDataFeed(websoc, writeLoopBreaker, writeLoopChan)

	// main event loop
	errCnt := 0
	for {
		response = nil

		select {
		case <-chrome.breaker:
			chrome.logger.Warn("shutting down mock chrome event loop")
			return

		// Manage any data input.
		case data := <-readLoopChan:
			if nil != data.Err {
				errCnt++
				if errCnt > 10 {
					chrome.breaker <- true
				}
				chrome.logger.WithField("error", data.Err).Error("mock chrome read loop returned an error")
			} else {
				errCnt = 0
			}

			mockData := chrome.popMockData(websoc)
			mockData.ID = data.Payload.ID
			mockData.Method = data.Payload.Method
			if mockData.ID == 0 {
				mockData.Result = data.Payload.Params
			}
			chrome.logger.WithField("data", mockData).Info("handling mock chrome websocket input")
			writeLoopChan <- mockData

			// Return responses.
		case data := <-writeLoopChan:
			chrome.logger.WithField("data", data).Info("building mock chrome websocket response")

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
				Info("writing response to mock chrome websocket connection")

			socketMux.Lock()
			err = websoc.WriteJSON(response)
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
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

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

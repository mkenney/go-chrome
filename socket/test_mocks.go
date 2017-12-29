package socket

import (
	"encoding/json"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

/*
NewMock returns a mock websocket for unit testing
*/
func NewMock(socketURL string) (Socketer, error) {
	socket := &socket{
		conn:     &mockConn{},
		commands: NewCommandMap(),
		handlers: NewEventHandlerMap(),
		url:      socketURL,
	}

	socket.accessibility = &AccessibilityProtocol{Socket: socket}
	socket.animation = &AnimationProtocol{Socket: socket}
	socket.applicationCache = &ApplicationCacheProtocol{Socket: socket}
	socket.audits = &AuditsProtocol{Socket: socket}
	socket.browser = &BrowserProtocol{Socket: socket}
	socket.cacheStorage = &CacheStorageProtocol{Socket: socket}
	socket.console = &ConsoleProtocol{Socket: socket}
	socket.css = &CSSProtocol{Socket: socket}
	socket.database = &DatabaseProtocol{Socket: socket}
	socket.debugger = &DebuggerProtocol{Socket: socket}
	socket.deviceOrientation = &DeviceOrientationProtocol{Socket: socket}
	socket.domDebugger = &DOMDebuggerProtocol{Socket: socket}
	socket.domSnapshot = &DOMSnapshotProtocol{Socket: socket}
	socket.domStorage = &DOMStorageProtocol{Socket: socket}
	socket.dom = &DOMProtocol{Socket: socket}
	socket.emulation = &EmulationProtocol{Socket: socket}
	socket.headlessExperimental = &HeadlessExperimentalProtocol{Socket: socket}
	socket.heapProfiler = &HeapProfilerProtocol{Socket: socket}
	socket.indexedDB = &IndexedDBProtocol{Socket: socket}
	socket.input = &InputProtocol{Socket: socket}
	socket.io = &IOProtocol{Socket: socket}
	socket.layerTree = &LayerTreeProtocol{Socket: socket}
	socket.log = &LogProtocol{Socket: socket}
	socket.memory = &MemoryProtocol{Socket: socket}
	socket.network = &NetworkProtocol{Socket: socket}
	socket.overlay = &OverlayProtocol{Socket: socket}
	socket.page = &PageProtocol{Socket: socket}
	socket.performance = &PerformanceProtocol{Socket: socket}
	socket.profiler = &ProfilerProtocol{Socket: socket}
	socket.runtime = &RuntimeProtocol{Socket: socket}
	socket.schema = &SchemaProtocol{Socket: socket}
	socket.security = &SecurityProtocol{Socket: socket}
	socket.serviceWorker = &ServiceWorkerProtocol{Socket: socket}
	socket.storage = &StorageProtocol{Socket: socket}
	socket.systemInfo = &SystemInfoProtocol{Socket: socket}
	socket.target = &TargetProtocol{Socket: socket}
	socket.tethering = &TetheringProtocol{Socket: socket}
	socket.tracing = &TracingProtocol{Socket: socket}

	return socket, nil
}

/*
mockConn implements Conner for mocking
*/
type mockConn struct{}

/*
Close implements Conner
*/
func (mockConn) Close() error {
	return nil
}

/*
ReadJSON implements Conner
*/
func (mockConn) ReadJSON(v interface{}) error {
	var err Error
	var data []byte

	log.Infof("MockJSONData: %s", MockJSONData)
	log.Infof("MockJSONType: %s", MockJSONType)
	log.Infof("MockJSONCommandID: %d", MockJSONCommandID)
	log.Infof("MockJSONRead: %v", MockJSONRead)
	log.Infof("MockJSONError: %v", MockJSONError)
	log.Infof("MockJSONThrowError: %v", MockJSONThrowError)

	time.Sleep(time.Millisecond * 100)
	if !MockJSONRead {
		if MockJSONThrowError {
			return fmt.Errorf("Mock Read Error")
		}

		MockJSONRead = true
		if MockJSONError {
			err.Code = 1
			err.Data = []byte(`{"data": "Error data"}`)
			err.Message = "Mock Error"
		}

		if "command" == MockJSONType {
			data, _ = json.Marshal(struct {
				Error  Error
				ID     int
				Method string
				Params json.RawMessage
				Result json.RawMessage
			}{
				Error:  err,
				ID:     MockJSONCommandID,
				Method: "Some.method",
				Params: nil,
				Result: MockJSONData,
			})
		} else if "event" == MockJSONType {
			data, _ = json.Marshal(struct {
				Error  Error
				ID     int
				Method string
				Params json.RawMessage
				Result json.RawMessage
			}{
				Error:  err,
				ID:     0,
				Method: "Some.event",
				Params: nil,
				Result: MockJSONData,
			})
		}
		json.Unmarshal(data, &v)
	}
	return nil
}

// MockJSONData flags for mocking ReadJSON()
var MockJSONData []byte

// MockJSONRead flags for mocking ReadJSON()
var MockJSONRead = false

// MockJSONType flags for mocking ReadJSON()
var MockJSONType = "command"

// MockJSONError flags for mocking ReadJSON()
var MockJSONError = true

// MockJSONCommandID flags for mocking ReadJSON()
var MockJSONCommandID = 1

// MockJSONThrowError flags for mocking ReadJSON()
var MockJSONThrowError = false

/*
WriteJSON implements Conner
*/
func (mockConn) WriteJSON(v interface{}) error {
	return nil
}

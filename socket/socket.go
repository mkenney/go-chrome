/*
Package socket allows for tools to instrument, inspect, debug and profile Chromium, Chrome and other
Blink-based browsers. Many existing projects currently use the protocol. The Chrome DevTools uses
this protocol and the team maintains its API.

See https://chromedevtools.github.io/devtools-protocol/ for more information.
*/
package socket

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

/*
New returns a new Socketer websocket connection listening to the specified URL.
*/
func New(socketURL string) (Socketer, error) {

	dialer := &websocket.Dialer{EnableCompression: false}
	header := http.Header{"Origin": []string{socketURL}}

	webSocket, response, err := dialer.Dial(socketURL, header)
	if err != nil {
		return nil, fmt.Errorf("Could not create websocket connection. %s responded with '%s'", socketURL, response.Status)
	}

	socket := &socket{
		conn:     webSocket,
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

	go socket.Listen()

	log.Infof("New socket connection listening on %s: %s", socket.url, response.Status)
	return socket, nil
}

//////////////////////////////////////////////////
// Socketer
//////////////////////////////////////////////////

/*
socket implements Socketer.
*/
type socket struct {
	commandID     int
	commands      CommandMapper
	handlers      EventHandlerMapper
	url           string
	conn          Conner
	stopListening bool

	//////////////////////////////////////////////////
	// Protocol interfaces
	//////////////////////////////////////////////////

	accessibility        *AccessibilityProtocol
	animation            *AnimationProtocol
	applicationCache     *ApplicationCacheProtocol
	audits               *AuditsProtocol
	browser              *BrowserProtocol
	cacheStorage         *CacheStorageProtocol
	console              *ConsoleProtocol
	css                  *CSSProtocol
	database             *DatabaseProtocol
	debugger             *DebuggerProtocol
	deviceOrientation    *DeviceOrientationProtocol
	domDebugger          *DOMDebuggerProtocol
	domSnapshot          *DOMSnapshotProtocol
	domStorage           *DOMStorageProtocol
	dom                  *DOMProtocol
	emulation            *EmulationProtocol
	headlessExperimental *HeadlessExperimentalProtocol
	heapProfiler         *HeapProfilerProtocol
	indexedDB            *IndexedDBProtocol
	input                *InputProtocol
	io                   *IOProtocol
	layerTree            *LayerTreeProtocol
	log                  *LogProtocol
	memory               *MemoryProtocol
	network              *NetworkProtocol
	overlay              *OverlayProtocol
	page                 *PageProtocol
	performance          *PerformanceProtocol
	profiler             *ProfilerProtocol
	runtime              *RuntimeProtocol
	schema               *SchemaProtocol
	security             *SecurityProtocol
	serviceWorker        *ServiceWorkerProtocol
	storage              *StorageProtocol
	systemInfo           *SystemInfoProtocol
	target               *TargetProtocol
	tethering            *TetheringProtocol
	tracing              *TracingProtocol
}

/*
Close implements Socketer.
*/
func (socket *socket) Close() error {
	return socket.conn.Close()
}

/*
Conn implements Socketer.
*/
func (socket *socket) Conn() Conner {
	return socket.conn
}

/*
Listen implements Socketer.
*/
func (socket *socket) Listen() error {
	var err error
	socket.stopListening = false
	for {
		if socket.stopListening {
			break
		}

		response := &Response{}
		err = socket.conn.ReadJSON(&response)
		if nil != err {
			log.Error(err)
			socket.Stop() // This will end the loop after handling the current response (if any)
		}

		if response.ID > 0 {
			socket.HandleCommand(response)

		} else {
			socket.HandleEvent(response)
		}
	}
	log.Infof("%s: Socket shutting down", socket.URL())

	return err
}

/*
Stop implements Socketer.
*/
func (socket *socket) Stop() {
	socket.stopListening = true
}

/*
URL implements Socketer.
*/
func (socket *socket) URL() string {
	return socket.url
}

//////////////////////////////////////////////////
// Socket response data
//////////////////////////////////////////////////

/*
Error represents a socket response error.
*/
type Error struct {
	Code    int             `json:"code"`
	Data    json.RawMessage `json:"data"`
	Message string          `json:"message"`
}

/*
Response represents a socket message.
*/
type Response struct {
	Error  *Error          `json:"error"`
	ID     int             `json:"id"`
	Method string          `json:"method"`
	Params json.RawMessage `json:"params"`
	Result json.RawMessage `json:"result"`
}

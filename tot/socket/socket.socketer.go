package socket

import (
	"encoding/json"
	"fmt"
	"net/url"
	"sync"

	errs "github.com/bdlm/errors"
	"github.com/bdlm/log"
)

/*
New returns a pointer to a websocket struct that implements Socketer interface
listening to the specified URL.
*/
func New(url *url.URL) *Socket {
	socket := &Socket{
		commandIDMux: &sync.Mutex{},
		commands:     NewCommandMap(),
		handlers:     NewEventHandlerMap(),
		mux:          &sync.Mutex{},
		newSocket:    NewWebsocket,
		socketID:     NextSocketID(),
		url:          url,
		stop:         make(chan bool),
	}
	socket.logger = log.WithFields(log.Fields{
		"socket.id":  socket.socketID,
		"socket.url": socket.url.String(),
	})

	// Init the protocol interfaces for the API.
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

	socket.Listen()
	socket.logger.Info("New socket connection listening")

	return socket
}

var _socketCounterMux = &sync.Mutex{}
var _socketCounter = 0

/*
NextSocketID increments and returns the socket ID for mapping Commander structs
to socket responses.
*/
func NextSocketID() int {
	_socketCounterMux.Lock()
	_socketCounter++
	id := _socketCounter
	_socketCounterMux.Unlock()
	return id
}

/*
Socket is a Socketer implementation.
*/
type Socket struct {
	commandID    int
	commandIDMux *sync.Mutex
	commands     CommandMapper
	handlers     EventHandlerMapper
	listenErr    errs.Err
	logger       *log.Entry
	newSocket    func(socketURL *url.URL) (WebSocketer, error)
	socketID     int
	url          *url.URL

	stop chan bool

	mux  *sync.Mutex
	conn WebSocketer

	// Protocol interfaces for the API.
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
AddEventHandler adds an event handler to the stack of listeners for an event.

AddEventHandler is a Socketer implementation.
*/
func (socket *Socket) AddEventHandler(
	handler EventHandler,
) {
	socket.handlers.Add(handler)
}

/*
CurCommandID returns the latest command ID.

CurCommandID is a Socketer implementation.
*/
func (socket *Socket) CurCommandID() int {
	socket.commandIDMux.Lock()
	id := socket.commandID
	socket.commandIDMux.Unlock()
	return id
}

/*
handleResponse receives the responses to requests sent to the websocket
connection.
*/
func (socket *Socket) handleResponse(response *Response) {
	if command, err := socket.commands.Get(response.ID); nil != err {
		socket.logger.WithFields(log.Fields{"error": err, "response": response}).
			Debugf("%-v", errs.Wrap(err, 0, fmt.Sprintf("command #%d not found", response.ID)))

	} else {
		socket.logger.WithFields(log.Fields{"command.id": command.ID(), "command.method": command.Method()}).
			Debug("executing handler")
		command.Respond(response)
		socket.commands.Delete(command.ID())
		socket.logger.WithFields(log.Fields{"response": response, "command.id": command.ID(), "command.method": command.Method()}).
			Debug("Command complete")
	}
}

/*
handleEvent receives all events and associated data read from the websocket
connection.
*/
func (socket *Socket) handleEvent(response *Response) {
	socket.logger.WithFields(log.Fields{"response": response}).
		Debug("handling event")

	if response.Method == "Inspector.targetCrashed" {
		socket.logger.Warn("Chrome has crashed!")
	}

	if handlers, err := socket.handlers.Get(response.Method); nil != err {
		socket.logger.WithFields(log.Fields{"error": err}).
			Debug(err)
	} else {
		for a, event := range handlers {
			socket.logger.WithFields(log.Fields{"event": response.Method, "handler#": a}).
				Info("Executing handler")
			go event.Handle(response)
		}
	}
}

/*
handleUnknown receives all other socket responses.
*/
func (socket *Socket) handleUnknown(
	response *Response,
) {
	var command Commander
	var err error

	socket.logger.WithFields(log.Fields{"response": response}).
		Debug("handling unexpected data")

	// Check for a command listening for ID 0.
	if command, err = socket.commands.Get(response.ID); nil != err {
		err = errs.Wrap(err, 0, fmt.Sprintf("command #%d not found", response.ID))
		if nil != response.Error && 0 != response.Error.Code {
			err = err.(*errs.Err).With(response.Error, err.Error())
		}
		socket.logger.WithFields(log.Fields{"error": err, "response": response}).
			Debugf("%-v", err)
		return
	}

	command.Respond(response)
	socket.logger.WithFields(log.Fields{"command.id": command.ID(), "error": response.Error, "command.method": command.Method()}).
		Debug("Unrecognised socket message")
}

/*
Listen starts the socket read loop and delivers messages to handleResponse() and
handleEvent() as appropriate.

Listen is a Socketer implementation.
*/
func (socket *Socket) Listen() {
	socket.Connect()
	go socket.listen()
}

func (socket *Socket) listen() error {
	var failures int
	breaker := make(chan bool)
	type ReadLoopData struct {
		Err      error
		Response *Response
	}
	readLoopChan := make(chan ReadLoopData)

	// launch the main
	go func() {
		for {
			select {
			// stop signal, stop the socket read loop and end the listner loop.
			case <-socket.stop:
				socket.logger.Info("Socket shutting down")
				breaker <- true
				return

			case data := <-readLoopChan:
				// read failure.
				if nil != data.Err {
					socket.logger.WithFields(log.Fields{"error": data.Err, "data": data}).
						Warnf("%-v", data.Err)
					socket.listenErr.With(data.Err, fmt.Sprintf("socket #%d - socket read failed", socket.socketID))
					failures++
				} else {
					failures = 0
				}

				// too many read failures.
				if failures > 10 {
					socket.logger.Error("too many read failures, shutting down")
					socket.Stop()
					continue
				}

				switch true {
				// command responses.
				case data.Response.ID > 0:
					socket.logger.WithFields(log.Fields{"responseID": data.Response.ID}).Debug("sending socket message to command handler")
					socket.handleResponse(data.Response)

				// event responses.
				case "" != data.Response.Method:
					socket.logger.WithFields(log.Fields{"command.method": data.Response.Method}).Debug("sending socket message to event handler")
					socket.handleEvent(data.Response)

				// unknown responses.
				default:
					tmp, _ := json.Marshal(data.Response)
					if nil == data.Response.Error {
						data.Response.Error = &Error{
							Message: "Unknown response from web socket",
						}
					}

					socket.logger.WithFields(log.Fields{"data": data, "error": data.Response.Error, "response": string(tmp)}).
						Debugf("%-v", data.Response.Error)
					socket.handleUnknown(data.Response)
				}
			}
		}
	}()

	// launch the socket reader.
	for {
		select {
		case readLoopChan <- func() ReadLoopData {
			data := ReadLoopData{
				Response: &Response{},
			}
			data.Err = socket.ReadJSON(&data.Response)
			return data
		}():
		case <-breaker:
			return nil
		}
	}
}

/*
NextCommandID generates and returns the next command ID.

NextCommandID is a Socketer implementation.
*/
func (socket *Socket) NextCommandID() int {
	socket.commandIDMux.Lock()
	socket.commandID++
	id := socket.commandID
	socket.commandIDMux.Unlock()
	return id
}

/*
RemoveEventHandler removes a handler from the stack of listeners for an event.

RemoveEventHandler is a Socketer implementation.
*/
func (socket *Socket) RemoveEventHandler(
	handler EventHandler,
) error {

	handlers, err := socket.handlers.Get(handler.Name())
	if nil != err {
		socket.logger.WithFields(log.Fields{"error": err}).
			Warn("Could not remove handler")
		return errs.Wrap(err, 0, fmt.Sprintf("failed to remove event handler '%s'", handler.Name()))
	}

	for i, hndlr := range handlers {
		if hndlr == handler {
			handlers = append(handlers[:i], handlers[i+1:]...)
			socket.handlers.Set(handler.Name(), handlers)
			socket.logger.WithFields(log.Fields{"handler": handler.Name(), "handlerID": i}).
				Info("Removed event handler")
			return nil
		}
	}

	socket.logger.Warn("handler not found")
	return nil
}

/*
SendCommand delivers a command payload to the websocket connection.

SendCommand is a Socketer implementation.

Workflow:
	1. The socket's command mutex is locked.
	2. The command counter is incremented.
	3. The payload is sent to the socket connection and the mutex is unlocked.
	4. The command is stored using the generated ID.
	5. When the command has been executed and the socket responds,
	socket.HandleCmd() is triggered from the command instance to generate the
	response and the command unlocks itself.
*/
func (socket *Socket) SendCommand(command Commander) chan *Response {
	log.WithFields(log.Fields{"command.ID": command.ID(), "command.Method": command.Method(), "command.Params": command.Params()}).
		Debug("sending command payload to socket")

	socket.commands.Set(command)
	go func() {
		payload := &Payload{
			ID:     command.ID(),
			Method: command.Method(),
			Params: command.Params(),
		}

		if err := socket.WriteJSON(payload); err != nil {
			err = errs.Wrap(err, 0, "write failed: could not write data to websocket")
			command.Respond(&Response{Error: &Error{
				Code:    1,
				Data:    []byte(fmt.Sprintf(`"%#v"`, err)),
				Message: "Failed to send command payload to socket connection",
			}})
			return
		}
	}()

	return command.Response()
}

/*
Stop signals the socket read loop to stop listening for data and close the
websocket connection.

Stop is a Socketer implementation.
*/
func (socket *Socket) Stop() error {
	go func() { socket.stop <- true }()
	return socket.Disconnect()
}

/*
URL returns the URL of the websocket connection.

URL is a Socketer implementation.
*/
func (socket *Socket) URL() *url.URL {
	return socket.url
}

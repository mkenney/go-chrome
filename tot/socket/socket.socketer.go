package socket

import (
	"encoding/json"
	"fmt"
	"net/url"
	"sync"
	"time"

	errs "github.com/mkenney/go-errors"
	log "github.com/sirupsen/logrus"
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
	}

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
	log.Infof("socket #%d - New socket connection listening on %s", socket.socketID, socket.url)

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
	commands     CommandMapper
	commandID    int
	commandIDMux *sync.Mutex
	conn         WebSocketer
	connected    bool
	handlers     EventHandlerMapper
	listenCh     chan bool
	newSocket    func(socketURL *url.URL) (WebSocketer, error)
	url          *url.URL
	socketID     int
	listening    bool
	mux          *sync.Mutex

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
	// Log a message on error
	if command, err := socket.commands.Get(response.ID); nil != err {
		err = errs.Wrap(err, fmt.Sprintf("command #%d not found", response.ID))
		errorMessage := ""
		if nil != response.Error && 0 != response.Error.Code {
			errorMessage = response.Error.Error()
		}
		log.Debugf(
			"socket #%d - socket.handleResponse(): %v - result=%s err='%s'",
			socket.socketID,
			err,
			response.Result,
			errorMessage,
		)

	} else {
		log.Debugf(
			"socket #%d - socket.handleResponse(): executing handler for command #%d - %s",
			socket.socketID,
			command.ID(),
			command.Method(),
		)
		command.Respond(response)
		socket.commands.Delete(command.ID())
		log.Debugf(
			"socket #%d - Command #%d complete: %s{%s}",
			socket.socketID,
			command.ID(),
			socket.URL().String(),
			command.Method(),
		)
	}
}

/*
handleEvent receives all events and associated data read from the websocket
connection.
*/
func (socket *Socket) handleEvent(
	response *Response,
) {
	log.Debugf(
		"socket #%d - socket.handleEvent(): handling event %s:{%s}",
		socket.socketID,
		socket.URL(),
		response.Method,
	)

	if response.Method == "Inspector.targetCrashed" {
		log.Errorf("socket #%d - Chrome has crashed!", socket.socketID)
	}

	if handlers, err := socket.handlers.Get(response.Method); nil != err {
		log.Errorf("socket #%d - %v", socket.socketID, err)

	} else {
		for a, event := range handlers {
			log.Infof("socket #%d - Executing handler #%d for event %s", socket.socketID, a, response.Method)
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
	log.Debugf(
		"socket #%d - socket.handleUnknown(): handling unexpected data from %s",
		socket.socketID,
		socket.URL(),
	)
	var command Commander
	var err error

	// Check for a command listening for ID 0
	if command, err = socket.commands.Get(response.ID); nil != err {
		err = errs.Wrap(err, fmt.Sprintf("command #%d not found", response.ID))
		if nil != response.Error && 0 != response.Error.Code {
			err = err.(errs.Err).With(fmt.Errorf("%s: %s", response.Error.Error(), err.Error()))
		}
		log.Debugf(
			"socket #%d - socket.handleUnknown(): result='%s' err='%v'",
			socket.socketID,
			response.Result,
			err,
		)
		return
	}

	command.Respond(response)
	log.Debugf(
		"socket #%d - socket.handleUnknown(): %v - Unrecognised socket message sent to command #%d - method: '%s' - Error: %v",
		socket.socketID,
		err,
		command.ID(),
		command.Method(),
		response.Error,
	)
}

/*
Listen starts the socket read loop and delivers messages to handleResponse() and
handleEvent() as appropriate.

Listen is a Socketer implementation.
*/
func (socket *Socket) Listen() {
	socket.listenCh = make(chan bool)
	socket.listening = true
	errCh := make(chan error)
	go socket.listen(errCh)
}

func (socket *Socket) listen(errCh chan error) {
	var err error

	err = socket.Connect()
	if nil != err {
		errCh <- errs.Wrap(err, "socket connection failed")
	}
	defer socket.Disconnect()

	for {
		response := &Response{}
		err = socket.ReadJSON(&response)
		if nil != err {
			log.Errorf("socket #%d - %v", socket.socketID, err)
		}

		if response.ID > 0 {
			log.Debugf(
				"socket #%d - socket.Listen(): Response ID #%d, sending to command handler",
				socket.socketID,
				response.ID,
			)
			socket.handleResponse(response)

		} else if "" != response.Method {
			log.Debugf(
				"socket #%d - socket.Listen(): Response method %s, sending to event handler",
				socket.socketID,
				response.Method,
			)
			socket.handleEvent(response)

		} else {
			tmp, _ := json.Marshal(response)
			log.WithFields(log.Fields{
				"socket": socket.socketID,
				"method": "socket.handleUnknown()",
				"data":   string(tmp),
			}).Errorf(
				"socket #%d - Unknown response from web socket: id='%d', method='%s'",
				socket.socketID,
				response.ID,
				response.Method,
			)

			if nil == response.Error {
				response.Error = &Error{
					Message: "Unknown response from web socket",
				}
			}
			socket.handleUnknown(response)
		}

		if !socket.listening {
			log.Infof("socket #%d - %s: Socket shutting down", socket.socketID, socket.URL().String())
			go func() {
				select {
				case socket.listenCh <- true:
				case <-time.After(10 * time.Second):
				}
			}()
			break
		}
	}

	if nil != err {
		err = errs.Wrap(err, "socket read failed")
	}

	errCh <- err
	socket.listening = false
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
	socket.handlers.Lock()
	defer socket.handlers.Unlock()

	handlers, err := socket.handlers.Get(handler.Name())
	if nil != err {
		log.Warnf("socket #%d - RemoveEventHandler(): Could not remove handler: %s", socket.socketID, err.Error())
		return errs.Wrap(err, fmt.Sprintf("failed to remove event handler '%s'", handler.Name()))
	}

	for i, hndlr := range handlers {
		if hndlr == handler {
			handlers = append(handlers[:i], handlers[i+1:]...)
			socket.handlers.Set(handler.Name(), handlers)
			log.Infof("socket #%d - RemoveEventHandler(): Removed event handler %d from %s", socket.socketID, i, handler.Name())
			return nil
		}
	}

	log.Warnf("socket #%d - RemoveEventHandler(): handler not found", socket.socketID)
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
	log.Debugf(
		"socket #%d - socket.SendCommand(): sending command #%d ('%+v') payload to socket",
		socket.socketID,
		command.ID(),
		command.Method(),
	)
	go func() {
		payload := &Payload{
			ID:     command.ID(),
			Method: command.Method(),
			Params: command.Params(),
		}

		if err := socket.WriteJSON(payload); err != nil {
			err = errs.Wrap(err, "write failed: could not write data to websocket")
			command.Respond(&Response{Error: &Error{
				Code:    1,
				Data:    []byte(fmt.Sprintf(`"%#v"`, err)),
				Message: "Failed to send command payload to socket connection",
			}})
			return
		}

		socket.commands.Set(command)
	}()

	return command.Response()
}

/*
Stop signals the socket read loop to stop listening for data and close the
websocket connection.

Stop is a Socketer implementation.
*/
func (socket *Socket) Stop() {
	if socket.listening {
		socket.listening = false
		select {
		case <-socket.listenCh:
		case <-time.After(1 * time.Second):
			socket.conn.Close()
		}
		log.Debugf("socket #%d: socket stopped", socket.socketID)
	}
}

/*
URL returns the URL of the websocket connection.

URL is a Socketer implementation.
*/
func (socket *Socket) URL() *url.URL {
	return socket.url
}

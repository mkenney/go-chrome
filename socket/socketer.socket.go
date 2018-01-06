package socket

import (
	"fmt"
	"net/url"
	"sync"

	log "github.com/sirupsen/logrus"
)

/*
New returns a new Socketer websocket connection listening to the specified URL.
*/
func New(url *url.URL) *Socket {
	socket := &Socket{
		commands:     NewCommandMap(),
		commandIDMux: &sync.Mutex{},
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

	go socket.Listen()
	log.Infof("socket #%d - New socket connection listening on %s", socket.socketID, socket.url)

	return socket
}

/*
NextSocketID increments and returns the socket ID.
*/
func NextSocketID() int {
	socketIDMux.Lock()
	defer socketIDMux.Unlock()
	_socketCounter++
	id := _socketCounter
	return id
}

var socketIDMux = &sync.Mutex{}
var _socketCounter = 0

/*
Socket implements Socketer.
*/
type Socket struct {
	commands      CommandMapper
	commandID     int
	commandIDMux  *sync.Mutex
	conn          WebSocketer
	connected     bool
	handlers      EventHandlerMapper
	newSocket     func(socketURL *url.URL) (WebSocketer, error)
	url           *url.URL
	socketID      int
	stopListening bool
	mux           *sync.Mutex

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
AddEventHandler implements Socketer.
*/
func (socket *Socket) AddEventHandler(
	handler EventHandler,
) {
	socket.handlers.Add(handler)
}

/*
CurCommandID implements Socketer.
*/
func (socket *Socket) CurCommandID() int {
	socket.commandIDMux.Lock()
	defer socket.commandIDMux.Unlock()
	id := socket.commandID
	return id
}

/*
HandleCommand implements Socketer.
*/
func (socket *Socket) HandleCommand(response *Response) {
	socket.commands.Lock()
	defer socket.commands.Unlock()

	if command, err := socket.commands.Get(response.ID); nil != err {
		errorMessage := ""
		if nil != response.Error && 0 != response.Error.Code {
			errorMessage = response.Error.Error()
		}
		log.Debugf(
			"socket #%d - socket.HandleCommand(): %s - result=%s err='%s'",
			socket.socketID,
			err.Error(),
			response.Result,
			errorMessage,
		)

	} else {
		log.Debugf(
			"socket #%d - socket.HandleCommand(): executing handler for command #%d - %s",
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
HandleEvent implements Socketer.
*/
func (socket *Socket) HandleEvent(
	response *Response,
) {
	log.Debugf(
		"socket #%d - socket.HandleEvent(): handling event %s:{%s}",
		socket.socketID,
		socket.URL(),
		response.Method,
	)

	if response.Method == "Inspector.targetCrashed" {
		log.Errorf("socket #%d - Chrome has crashed!", socket.socketID)
	}

	if handlers, err := socket.handlers.Get(response.Method); nil != err {
		log.Debugf("socket #%d - %s", socket.socketID, err.Error())

	} else {
		for a, event := range handlers {
			log.Infof("socket #%d - Executing handler #%d for event %s", socket.socketID, a, response.Method)
			go event.Handle(response)
		}
	}
}

/*
Listen implements Socketer.
*/
func (socket *Socket) Listen() error {
	var err error

	err = socket.Connect()
	if nil != err {
		return err
	}
	defer socket.Disconnect()

	socket.stopListening = false
	for {
		response := &Response{}
		err = socket.ReadJSON(&response)
		if nil != err {
			log.Errorf("socket #%d - %s", err.Error())
			socket.Stop() // This will end the loop after handling the current response (if any)
		}

		if response.ID > 0 {
			log.Debugf("socket #%d - socket.Listen(): Response ID #%d, sending to command handler", socket.socketID, response.ID)
			socket.HandleCommand(response)

		} else if "" != response.Method {
			log.Debugf("socket #%d - socket.Listen(): Response method %s, sending to event handler", socket.socketID, response.Method)
			socket.HandleEvent(response)

		} else {
			log.Error(fmt.Errorf(
				"socket #%d - Unknown response from web socket: id=%d, method=%s",
				socket.socketID,
				response.ID,
				response.Method,
			))
		}

		if socket.stopListening {
			log.Infof("socket #%d - %s: Socket shutting down", socket.socketID, socket.URL().String())
			break
		}
	}

	return err
}

/*
NextCommandID implements Socketer.
*/
func (socket *Socket) NextCommandID() int {
	socket.commandIDMux.Lock()
	defer socket.commandIDMux.Unlock()
	socket.commandID++
	id := socket.commandID
	return id
}

/*
RemoveEventHandler implements Socketer.
*/
func (socket *Socket) RemoveEventHandler(
	handler EventHandler,
) {
	socket.handlers.Lock()
	defer socket.handlers.Unlock()

	if handlers, err := socket.handlers.Get(handler.Name()); nil != err {
		log.Warnf("socket #%d - Could not remove handler: %s", socket.socketID, err.Error())
	} else {
		for i, hndlr := range handlers {
			if hndlr == handler {
				handlers = append(handlers[:i], handlers[i+1:]...)
				socket.handlers.Set(handler.Name(), handlers)
				return
			}
		}
	}
}

/*
SendCommand implements Socketer.

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
		"socket #%d - socket.SendCommand(): sending command #%d (%s) payload to socket",
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
			command.Respond(&Response{Error: &Error{
				Code:    1,
				Data:    []byte(fmt.Sprintf(`"%s"`, err.Error())),
				Message: "Failed to send command payload to socket connection",
			}})
			return
		}

		socket.commands.Set(command)
	}()

	return command.Response()
}

/*
Stop implements Socketer.
*/
func (socket *Socket) Stop() {
	socket.stopListening = true
}

/*
URL implements Socketer.
*/
func (socket *Socket) URL() *url.URL {
	return socket.url
}

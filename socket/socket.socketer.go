package socket

import (
	"fmt"
	"net/url"
	"sync"

	"github.com/mkenney/go-chrome/errors"
	log "github.com/sirupsen/logrus"
)

/*
New returns a new Socketer websocket connection listening to the specified URL.
*/
func New(url *url.URL) *Socket {
	socket := &Socket{
		commands:  NewCommandMap(),
		handlers:  NewEventHandlerMap(),
		url:       url,
		newSocket: NewWebsocket,
		mux:       &sync.Mutex{},
	}

	// Init the protocol interfaces for the API.
	socket.accessibility = &AccessibilityProtocol{
		Socket: socket,
	}
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
	log.Infof("New socket connection listening on %s", socket.url)

	return socket
}

/*
Socket implements Socketer.
*/
type Socket struct {
	commands      CommandMapper
	conn          WebSocketer
	connected     bool
	handlers      EventHandlerMapper
	newSocket     func(socketURL *url.URL) (WebSocketer, error)
	url           *url.URL
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
	socket.handlers.Lock()
	defer socket.handlers.Unlock()

	handlers, err := socket.handlers.Get(handler.Name())
	if nil != err {
		handlers = make([]EventHandler, 0)
	}
	for _, hndl := range handlers {
		if hndl == handler {
			log.Warnf("Attempted to add a duplicate handler for event '%s', skipping...", handler.Name())
			return
		}
	}

	log.Debugf("Adding handler for event '%s'", handler.Name())
	handlers = append(handlers, handler)
	socket.handlers.Set(handler.Name(), handlers)
}

/*
HandleCommand implements Socketer.
*/
func (socket *Socket) HandleCommand(response *Response) {
	socket.commands.Lock()
	defer socket.commands.Unlock()

	log.Debugf("socket.HandleCommand(): executing handler for command #%d", response.ID)
	if command, err := socket.commands.Get(response.ID); nil != err {
		errorMessage := ""
		if nil != response.Error && 0 != response.Error.Code {
			errorMessage = response.Error.Message
		}
		log.Debugf(
			"%s: result=%s err=%s",
			err.Error(),
			response.Result,
			errorMessage,
		)

	} else {
		socket.commands.Delete(command.ID())
		if nil != response.Error && "" != response.Error.Message {
			errorMessage := ""
			if nil != response.Error && 0 != response.Error.Code {
				errorMessage = response.Error.Message
			}
			log.Errorf(
				"Socket command responded with error: result=%s err=%s",
				response.Result,
				errorMessage,
			)
			err = errors.SocketErrorResponse{Type: errors.Type{
				Caller: errors.GetCaller(),
				Err: fmt.Errorf(
					"%d %s: %s",
					response.Error.Code,
					response.Error.Message,
					response.Error.Data,
				),
				Msg: "The socket command failed",
			}}
		}

		command.Done(response.Result, err)
		log.Debugf(
			"Command #%d complete: %s {%s}",
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
	log.Debugf("socket.HandleEvent(): handling event %s:{%s}", socket.URL(), response.Method)
	if response.Method == "Inspector.targetCrashed" {
		log.Errorf("Chrome has crashed!")
	}

	socket.handlers.Lock()
	defer socket.handlers.Unlock()

	if handlers, err := socket.handlers.Get(response.Method); nil != err {
		log.Error(err)

	} else {
		for a, event := range handlers {
			log.Infof("Executing handler #%d for event %s", a, response.Method)
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
			log.Error(err)
			socket.Stop() // This will end the loop after handling the current response (if any)
		} else if "" != response.Method {
			if response.ID > 0 {
				log.Debugf("socket.Listen(): Handling command #%d - %s", response.ID, response.Method)
				socket.HandleCommand(response)
			} else {
				log.Debugf("socket.Listen(): Handling event %s", response.Method)
				socket.HandleEvent(response)
			}
		}

		if socket.stopListening {
			log.Infof("%s: Socket shutting down", socket.URL().String())
			break
		}
	}

	return err
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
		log.Warnf("Could not remove handler: %s", err.Error())
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
func (socket *Socket) SendCommand(command Commander) *Payload {
	log.Debugf("socket.SendCommand(): sending command #%d - %s", command.ID(), command.Method())

	// Add the command to the internal stack and create the payload
	socket.commands.Set(command)
	payload := &Payload{
		ID:     command.ID(),
		Method: command.Method(),
		Params: command.Params(),
	}

	command.WaitGroup().Add(1)
	if err := socket.WriteJSON(payload); err != nil {
		command.Done(nil, errors.SocketWriteFailed{Type: errors.Type{
			Caller: errors.GetCaller(),
			Err:    err,
			Msg:    "Failed to send command payload to socket connection",
		}})
	}
	log.Debugf("socket.SendCommand(): waiting for response to command #%d - %s", command.ID(), command.Method())
	command.WaitGroup().Wait()

	return payload
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

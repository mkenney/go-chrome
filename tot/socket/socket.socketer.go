package socket

import (
	"encoding/json"
	"fmt"
	"net/url"
	"sync"
	"time"

	errs "github.com/bdlm/errors"
	"github.com/bdlm/log"
	"github.com/mkenney/go-chrome/codes"
)

/*
New returns a pointer to a websocket struct that implements Socketer interface
listening to the specified URL.
*/
func New(url *url.URL) *Socket {
	socket := &Socket{
		commandIDMux: &sync.Mutex{},
		commands:     NewCommandMap(),
		errCh:        make(chan error, 3),
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

	log.WithFields(log.Fields{
		"socketID": socket.socketID,
		"url":      socket.url.String(),
	}).Info("New socket connection listening")

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
	conn         WebSocketer
	connected    bool
	errCh        chan error
	handlers     EventHandlerMapper
	listenCh     chan bool
	listening    bool
	mux          *sync.Mutex
	newSocket    func(socketURL *url.URL) (WebSocketer, error)
	socketID     int
	url          *url.URL

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
		err = errs.Wrap(err, 0, fmt.Sprintf("command #%d not found", response.ID))
		log.WithFields(log.Fields{
			"error":    err,
			"result":   response.Result,
			"socketID": socket.socketID,
		}).Debug(response.Error)

	} else {
		log.WithFields(log.Fields{
			"commandID": command.ID(),
			"method":    command.Method(),
			"socketID":  socket.socketID,
		}).Debug("executing handler")
		command.Respond(response)
		socket.commands.Delete(command.ID())
		log.WithFields(log.Fields{
			"commandID": command.ID(),
			"method":    command.Method(),
			"socketID":  socket.socketID,
			"url":       socket.url.String(),
		}).Debug("Command complete")
	}
}

/*
handleEvent receives all events and associated data read from the websocket
connection.
*/
func (socket *Socket) handleEvent(
	response *Response,
) {
	log.WithFields(log.Fields{
		"event":    response.Method,
		"socketID": socket.socketID,
		"url":      socket.url.String(),
	}).Debug("handling event")

	if response.Method == "Inspector.targetCrashed" {
		log.WithFields(log.Fields{
			"socketID": socket.socketID,
		}).Error("Chrome has crashed!")
	}

	if handlers, err := socket.handlers.Get(response.Method); nil != err {
		log.WithFields(log.Fields{
			"error":    err,
			"socketID": socket.socketID,
		}).Debug(err)
	} else {
		for a, event := range handlers {
			log.WithFields(log.Fields{
				"event":    response.Method,
				"handler#": a,
				"socketID": socket.socketID,
			}).Info("Executing handler")
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
	log.WithFields(log.Fields{
		"socketID": socket.socketID,
		"url":      socket.url.String(),
	}).Debug("handling unexpected data")
	var command Commander
	var err error

	// Check for a command listening for ID 0
	if command, err = socket.commands.Get(response.ID); nil != err {
		err = errs.Wrap(err, 0, fmt.Sprintf("command #%d not found", response.ID))
		if nil != response.Error && 0 != response.Error.Code {
			err = err.(errs.Err).With(response.Error, err.Error())
		}
		log.WithFields(log.Fields{
			"error":    err,
			"result":   response.Result,
			"socketID": socket.socketID,
		}).Debug(err)
		return
	}

	command.Respond(response)
	log.WithFields(log.Fields{
		"commandID": command.ID(),
		"error":     response.Error,
		"method":    command.Method(),
		"socketID":  socket.socketID,
	}).Debug("Unrecognised socket message")
}

/*
Errors returns the last error returned by the listen routine, if any.
*/
func (socket *Socket) Errors() chan error {
	return socket.errCh
}

/*
Listen starts the socket read loop and delivers messages to handleResponse() and
handleEvent() as appropriate.

Listen is a Socketer implementation.
*/
func (socket *Socket) Listen() {
	socket.listenCh = make(chan bool)
	socket.listening = true
	go socket.listen(socket.errCh)
}

func (socket *Socket) listen(errCh chan error) {
	var err error

	// recover socket panics caused by defunct or dead connections. This
	// also returns the error signal.
	defer func() {
		if r := recover(); r != nil {
			err = errs.New(codes.SocketPanic, "recovered from panic in Socket.listen()")
			if e, ok := r.(error); ok {
				err = errs.Wrap(e, codes.SocketPanic, "recovered from panic in Socket.listen()")
			}
			log.WithFields(log.Fields{
				"error": err,
			}).Error(err)
		}
		errCh <- err
	}()

	err = socket.Connect()
	if nil != err {
		errCh <- errs.Wrap(err, 0, "socket connection failed")
		return
	}
	defer socket.Disconnect()

	for {
		response := &Response{}
		err = socket.ReadJSON(&response)
		if nil != err {
			err = errs.Wrap(err, codes.SocketReadFailed, fmt.Sprintf("socket #%d - socket read failed", socket.socketID))
			log.WithFields(log.Fields{
				"socketID": socket.socketID,
			}).Error(err)
		}
		if 0 == response.ID &&
			"" == response.Method &&
			0 == len(response.Params) &&
			0 == len(response.Result) {
			log.WithFields(log.Fields{
				"socketID": socket.socketID,
			}).Error("nil response from socket")
		}

		if response.ID > 0 {
			log.WithFields(log.Fields{
				"responseID": response.ID,
				"socketID":   socket.socketID,
			}).Debug("sending to command handler")
			socket.handleResponse(response)

		} else if "" != response.Method {
			log.WithFields(log.Fields{
				"method":   response.Method,
				"socketID": socket.socketID,
			}).Debug("sending to event handler")
			socket.handleEvent(response)

		} else {
			tmp, _ := json.Marshal(response)
			log.WithFields(log.Fields{
				"data":       string(tmp),
				"method":     response.Method,
				"responseID": response.ID,
				"socketID":   socket.socketID,
			}).Error("Unknown response from web socket")

			if nil == response.Error {
				response.Error = &Error{
					Message: "Unknown response from web socket",
				}
			}
			socket.handleUnknown(response)
		}

		if !socket.listening {
			log.WithFields(log.Fields{
				"socketID": socket.socketID,
				"url":      socket.url.String(),
			}).Info("Socket shutting down")
			go func() {
				select {
				case socket.listenCh <- true:
				case <-time.After(10 * time.Second):
				}
			}()
			break
		}
	}

	socket.listening = false
	if nil != err {
		errCh <- errs.Wrap(err, 0, "socket closed")
		return
	}
	errCh <- nil
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
		log.WithFields(log.Fields{
			"error":    err,
			"socketID": socket.socketID,
		}).Warn("Could not remove handler")
		return errs.Wrap(err, 0, fmt.Sprintf("failed to remove event handler '%s'", handler.Name()))
	}

	for i, hndlr := range handlers {
		if hndlr == handler {
			handlers = append(handlers[:i], handlers[i+1:]...)
			socket.handlers.Set(handler.Name(), handlers)
			log.WithFields(log.Fields{
				"handler":   handler.Name(),
				"handlerID": i,
				"socketID":  socket.socketID,
			}).Info("Removed event handler")
			return nil
		}
	}

	log.WithFields(log.Fields{
		"socketID": socket.socketID,
	}).Warn("handler not found")
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
	log.WithFields(log.Fields{
		"commandID": command.ID(),
		"method":    command.Method(),
		"socketID":  socket.socketID,
	}).Debug("sending command payload to socket")
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
		log.WithFields(log.Fields{
			"socketID": socket.socketID,
		}).Debug("socket stopped")
	}
}

/*
URL returns the URL of the websocket connection.

URL is a Socketer implementation.
*/
func (socket *Socket) URL() *url.URL {
	return socket.url
}

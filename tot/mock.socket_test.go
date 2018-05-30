package chrome

import (
	"net/url"

	"github.com/mkenney/go-chrome/tot/socket"
)

func NewMockSocket(url *url.URL) *MockSocket {
	return &MockSocket{
		url: url,
	}
}

/*
Socket is a Socketer implementation.
*/
type MockSocket struct {
	url       *url.URL
	commandID int

	// Protocol interfaces for the API.
	accessibility        *socket.AccessibilityProtocol
	animation            *socket.AnimationProtocol
	applicationCache     *socket.ApplicationCacheProtocol
	audits               *socket.AuditsProtocol
	browser              *socket.BrowserProtocol
	cacheStorage         *socket.CacheStorageProtocol
	console              *socket.ConsoleProtocol
	css                  *socket.CSSProtocol
	database             *socket.DatabaseProtocol
	debugger             *socket.DebuggerProtocol
	deviceOrientation    *socket.DeviceOrientationProtocol
	domDebugger          *socket.DOMDebuggerProtocol
	domSnapshot          *socket.DOMSnapshotProtocol
	domStorage           *socket.DOMStorageProtocol
	dom                  *socket.DOMProtocol
	emulation            *socket.EmulationProtocol
	headlessExperimental *socket.HeadlessExperimentalProtocol
	heapProfiler         *socket.HeapProfilerProtocol
	indexedDB            *socket.IndexedDBProtocol
	input                *socket.InputProtocol
	io                   *socket.IOProtocol
	layerTree            *socket.LayerTreeProtocol
	log                  *socket.LogProtocol
	memory               *socket.MemoryProtocol
	network              *socket.NetworkProtocol
	overlay              *socket.OverlayProtocol
	page                 *socket.PageProtocol
	performance          *socket.PerformanceProtocol
	profiler             *socket.ProfilerProtocol
	runtime              *socket.RuntimeProtocol
	schema               *socket.SchemaProtocol
	security             *socket.SecurityProtocol
	serviceWorker        *socket.ServiceWorkerProtocol
	storage              *socket.StorageProtocol
	systemInfo           *socket.SystemInfoProtocol
	target               *socket.TargetProtocol
	tethering            *socket.TetheringProtocol
	tracing              *socket.TracingProtocol
}

/*
AddEventHandler is a Socketer implementation.
*/
func (socket *MockSocket) AddEventHandler(
	handler socket.EventHandler,
) {
}

/*
CurCommandID is a Socketer implementation.
*/
func (socket *MockSocket) CurCommandID() int {
	id := socket.commandID
	return id
}

/*
Listen starts the socket read loop and delivers messages to handleResponse() and
handleEvent() as appropriate.

Listen is a Socketer implementation.
*/
func (socket *MockSocket) Listen() {
}

/*
NextCommandID generates and returns the next command ID.
*/
func (socket *MockSocket) NextCommandID() int {
	socket.commandID++
	return socket.commandID
}

/*
RemoveEventHandler is a Socketer implementation.
*/
func (socket *MockSocket) RemoveEventHandler(
	handler socket.EventHandler,
) error {
	return nil
}

/*
SendCommand is a Socketer implementation.
*/
func (socket *MockSocket) SendCommand(command socket.Commander) chan *socket.Response {
	return command.Response()
}

/*
Stop is a Socketer implementation.
*/
func (socket *MockSocket) Stop() {
}

/*
URL returns the URL of the websocket connection.
*/
func (socket *MockSocket) URL() *url.URL {
	return socket.url
}

/*
Accessibility is a Protocoller implementation.
*/
func (socket *MockSocket) Accessibility() *socket.AccessibilityProtocol {
	return socket.accessibility
}

/*
Animation is a Protocoller implementation.
*/
func (socket *MockSocket) Animation() *socket.AnimationProtocol {
	return socket.animation
}

/*
ApplicationCache is a Protocoller implementation.
*/
func (socket *MockSocket) ApplicationCache() *socket.ApplicationCacheProtocol {
	return socket.applicationCache
}

/*
Audits is a Protocoller implementation.
*/
func (socket *MockSocket) Audits() *socket.AuditsProtocol {
	return socket.audits
}

/*
Browser is a Protocoller implementation.
*/
func (socket *MockSocket) Browser() *socket.BrowserProtocol {
	return socket.browser
}

/*
CacheStorage is a Protocoller implementation.
*/
func (socket *MockSocket) CacheStorage() *socket.CacheStorageProtocol {
	return socket.cacheStorage
}

/*
Console is a Protocoller implementation.
*/
func (socket *MockSocket) Console() *socket.ConsoleProtocol {
	return socket.console
}

/*
CSS is a Protocoller implementation.
*/
func (socket *MockSocket) CSS() *socket.CSSProtocol {
	return socket.css
}

/*
Database is a Protocoller implementation.
*/
func (socket *MockSocket) Database() *socket.DatabaseProtocol {
	return socket.database
}

/*
Debugger is a Protocoller implementation.
*/
func (socket *MockSocket) Debugger() *socket.DebuggerProtocol {
	return socket.debugger
}

/*
DeviceOrientation is a Protocoller implementation.
*/
func (socket *MockSocket) DeviceOrientation() *socket.DeviceOrientationProtocol {
	return socket.deviceOrientation
}

/*
DOMDebugger is a Protocoller implementation.
*/
func (socket *MockSocket) DOMDebugger() *socket.DOMDebuggerProtocol {
	return socket.domDebugger
}

/*
DOMSnapshot is a Protocoller implementation.
*/
func (socket *MockSocket) DOMSnapshot() *socket.DOMSnapshotProtocol {
	return socket.domSnapshot
}

/*
DOMStorage is a Protocoller implementation.
*/
func (socket *MockSocket) DOMStorage() *socket.DOMStorageProtocol {
	return socket.domStorage
}

/*
DOM is a Protocoller implementation.
*/
func (socket *MockSocket) DOM() *socket.DOMProtocol {
	return socket.dom
}

/*
Emulation is a Protocoller implementation.
*/
func (socket *MockSocket) Emulation() *socket.EmulationProtocol {
	return socket.emulation
}

/*
HeadlessExperimental is a Protocoller implementation.
*/
func (socket *MockSocket) HeadlessExperimental() *socket.HeadlessExperimentalProtocol {
	return socket.headlessExperimental
}

/*
HeapProfiler is a Protocoller implementation.
*/
func (socket *MockSocket) HeapProfiler() *socket.HeapProfilerProtocol {
	return socket.heapProfiler
}

/*
IndexedDB is a Protocoller implementation.
*/
func (socket *MockSocket) IndexedDB() *socket.IndexedDBProtocol {
	return socket.indexedDB
}

/*
Input is a Protocoller implementation.
*/
func (socket *MockSocket) Input() *socket.InputProtocol {
	return socket.input
}

/*
IO is a Protocoller implementation.
*/
func (socket *MockSocket) IO() *socket.IOProtocol {
	return socket.io
}

/*
LayerTree is a Protocoller implementation.
*/
func (socket *MockSocket) LayerTree() *socket.LayerTreeProtocol {
	return socket.layerTree
}

/*
Log is a Protocoller implementation.
*/
func (socket *MockSocket) Log() *socket.LogProtocol {
	return socket.log
}

/*
Memory is a Protocoller implementation.
*/
func (socket *MockSocket) Memory() *socket.MemoryProtocol {
	return socket.memory
}

/*
Network is a Protocoller implementation.
*/
func (socket *MockSocket) Network() *socket.NetworkProtocol {
	return socket.network
}

/*
Overlay is a Protocoller implementation.
*/
func (socket *MockSocket) Overlay() *socket.OverlayProtocol {
	return socket.overlay
}

/*
Page is a Protocoller implementation.
*/
func (socket *MockSocket) Page() *socket.PageProtocol {
	return socket.page
}

/*
Performance is a Protocoller implementation.
*/
func (socket *MockSocket) Performance() *socket.PerformanceProtocol {
	return socket.performance
}

/*
Profiler is a Protocoller implementation.
*/
func (socket *MockSocket) Profiler() *socket.ProfilerProtocol {
	return socket.profiler
}

/*
Runtime is a Protocoller implementation.
*/
func (socket *MockSocket) Runtime() *socket.RuntimeProtocol {
	return socket.runtime
}

/*
Schema is a Protocoller implementation.
*/
func (socket *MockSocket) Schema() *socket.SchemaProtocol {
	return socket.schema
}

/*
Security is a Protocoller implementation.
*/
func (socket *MockSocket) Security() *socket.SecurityProtocol {
	return socket.security
}

/*
ServiceWorker is a Protocoller implementation.
*/
func (socket *MockSocket) ServiceWorker() *socket.ServiceWorkerProtocol {
	return socket.serviceWorker
}

/*
Storage is a Protocoller implementation.
*/
func (socket *MockSocket) Storage() *socket.StorageProtocol {
	return socket.storage
}

/*
SystemInfo is a Protocoller implementation.
*/
func (socket *MockSocket) SystemInfo() *socket.SystemInfoProtocol {
	return socket.systemInfo
}

/*
Target is a Protocoller implementation.
*/
func (socket *MockSocket) Target() *socket.TargetProtocol {
	return socket.target
}

/*
Tethering is a Protocoller implementation.
*/
func (socket *MockSocket) Tethering() *socket.TetheringProtocol {
	return socket.tethering
}

/*
Tracing is a Protocoller implementation.
*/
func (socket *MockSocket) Tracing() *socket.TracingProtocol {
	return socket.tracing
}

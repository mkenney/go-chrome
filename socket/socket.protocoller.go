package socket

/*
Accessibility implements Protocoller
*/
func (socket *Socket) Accessibility() *AccessibilityProtocol {
	return socket.accessibility
}

/*
Animation implements Protocoller
*/
func (socket *Socket) Animation() *AnimationProtocol {
	return socket.animation
}

/*
ApplicationCache implements Protocoller
*/
func (socket *Socket) ApplicationCache() *ApplicationCacheProtocol {
	return socket.applicationCache
}

/*
Audits implements Protocoller
*/
func (socket *Socket) Audits() *AuditsProtocol {
	return socket.audits
}

/*
Browser implements Protocoller
*/
func (socket *Socket) Browser() *BrowserProtocol {
	return socket.browser
}

/*
CacheStorage implements Protocoller
*/
func (socket *Socket) CacheStorage() *CacheStorageProtocol {
	return socket.cacheStorage
}

/*
Console implements Protocoller
*/
func (socket *Socket) Console() *ConsoleProtocol {
	return socket.console
}

/*
CSS implements Protocoller
*/
func (socket *Socket) CSS() *CSSProtocol {
	return socket.css
}

/*
Database implements Protocoller
*/
func (socket *Socket) Database() *DatabaseProtocol {
	return socket.database
}

/*
Debugger implements Protocoller
*/
func (socket *Socket) Debugger() *DebuggerProtocol {
	return socket.debugger
}

/*
DeviceOrientation implements Protocoller
*/
func (socket *Socket) DeviceOrientation() *DeviceOrientationProtocol {
	return socket.deviceOrientation
}

/*
DOMDebugger implements Protocoller
*/
func (socket *Socket) DOMDebugger() *DOMDebuggerProtocol {
	return socket.domDebugger
}

/*
DOMSnapshot implements Protocoller
*/
func (socket *Socket) DOMSnapshot() *DOMSnapshotProtocol {
	return socket.domSnapshot
}

/*
DOMStorage implements Protocoller
*/
func (socket *Socket) DOMStorage() *DOMStorageProtocol {
	return socket.domStorage
}

/*
DOM implements Protocoller
*/
func (socket *Socket) DOM() *DOMProtocol {
	return socket.dom
}

/*
Emulation implements Protocoller
*/
func (socket *Socket) Emulation() *EmulationProtocol {
	return socket.emulation
}

/*
HeadlessExperimental implements Protocoller
*/
func (socket *Socket) HeadlessExperimental() *HeadlessExperimentalProtocol {
	return socket.headlessExperimental
}

/*
HeapProfiler implements Protocoller
*/
func (socket *Socket) HeapProfiler() *HeapProfilerProtocol {
	return socket.heapProfiler
}

/*
IndexedDB implements Protocoller
*/
func (socket *Socket) IndexedDB() *IndexedDBProtocol {
	return socket.indexedDB
}

/*
Input implements Protocoller
*/
func (socket *Socket) Input() *InputProtocol {
	return socket.input
}

/*
IO implements Protocoller
*/
func (socket *Socket) IO() *IOProtocol {
	return socket.io
}

/*
LayerTree implements Protocoller
*/
func (socket *Socket) LayerTree() *LayerTreeProtocol {
	return socket.layerTree
}

/*
Log implements Protocoller
*/
func (socket *Socket) Log() *LogProtocol {
	return socket.log
}

/*
Memory implements Protocoller
*/
func (socket *Socket) Memory() *MemoryProtocol {
	return socket.memory
}

/*
Network implements Protocoller
*/
func (socket *Socket) Network() *NetworkProtocol {
	return socket.network
}

/*
Overlay implements Protocoller
*/
func (socket *Socket) Overlay() *OverlayProtocol {
	return socket.overlay
}

/*
Page implements Protocoller
*/
func (socket *Socket) Page() *PageProtocol {
	return socket.page
}

/*
Performance implements Protocoller
*/
func (socket *Socket) Performance() *PerformanceProtocol {
	return socket.performance
}

/*
Profiler implements Protocoller
*/
func (socket *Socket) Profiler() *ProfilerProtocol {
	return socket.profiler
}

/*
Runtime implements Protocoller
*/
func (socket *Socket) Runtime() *RuntimeProtocol {
	return socket.runtime
}

/*
Schema implements Protocoller
*/
func (socket *Socket) Schema() *SchemaProtocol {
	return socket.schema
}

/*
Security implements Protocoller
*/
func (socket *Socket) Security() *SecurityProtocol {
	return socket.security
}

/*
ServiceWorker implements Protocoller
*/
func (socket *Socket) ServiceWorker() *ServiceWorkerProtocol {
	return socket.serviceWorker
}

/*
Storage implements Protocoller
*/
func (socket *Socket) Storage() *StorageProtocol {
	return socket.storage
}

/*
SystemInfo implements Protocoller
*/
func (socket *Socket) SystemInfo() *SystemInfoProtocol {
	return socket.systemInfo
}

/*
Target implements Protocoller
*/
func (socket *Socket) Target() *TargetProtocol {
	return socket.target
}

/*
Tethering implements Protocoller
*/
func (socket *Socket) Tethering() *TetheringProtocol {
	return socket.tethering
}

/*
Tracing implements Protocoller
*/
func (socket *Socket) Tracing() *TracingProtocol {
	return socket.tracing
}

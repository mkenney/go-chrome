package socket

/*
Accessibility returns the AccessibilityProtocol instance.

Accessibility is a Protocoller implementation.
*/
func (socket *Socket) Accessibility() *AccessibilityProtocol {
	return socket.accessibility
}

/*
Animation returns the AnimationProtocol instance.

Animation is a Protocoller implementation.
*/
func (socket *Socket) Animation() *AnimationProtocol {
	return socket.animation
}

/*
ApplicationCache returns the ApplicationCacheProtocol instance.

ApplicationCache is a Protocoller implementation.
*/
func (socket *Socket) ApplicationCache() *ApplicationCacheProtocol {
	return socket.applicationCache
}

/*
Audits returns the AuditsProtocol instance.

Audits is a Protocoller implementation.
*/
func (socket *Socket) Audits() *AuditsProtocol {
	return socket.audits
}

/*
Browser returns the BrowserProtocol instance.

Browser is a Protocoller implementation.
*/
func (socket *Socket) Browser() *BrowserProtocol {
	return socket.browser
}

/*
CacheStorage returns the CacheStorageProtocol instance.

CacheStorage is a Protocoller implementation.
*/
func (socket *Socket) CacheStorage() *CacheStorageProtocol {
	return socket.cacheStorage
}

/*
Console returns the ConsoleProtocol instance.

Console is a Protocoller implementation.
*/
func (socket *Socket) Console() *ConsoleProtocol {
	return socket.console
}

/*
CSS returns the CSSProtocol instance.

CSS is a Protocoller implementation.
*/
func (socket *Socket) CSS() *CSSProtocol {
	return socket.css
}

/*
Database returns the DatabaseProtocol instance.

Database is a Protocoller implementation.
*/
func (socket *Socket) Database() *DatabaseProtocol {
	return socket.database
}

/*
Debugger returns the DebuggerProtocol instance.

Debugger is a Protocoller implementation.
*/
func (socket *Socket) Debugger() *DebuggerProtocol {
	return socket.debugger
}

/*
DeviceOrientation returns the DeviceOrientationProtocol instance.

DeviceOrientation is a Protocoller implementation.
*/
func (socket *Socket) DeviceOrientation() *DeviceOrientationProtocol {
	return socket.deviceOrientation
}

/*
DOMDebugger returns the DOMDebuggerProtocol instance.

DOMDebugger is a Protocoller implementation.
*/
func (socket *Socket) DOMDebugger() *DOMDebuggerProtocol {
	return socket.domDebugger
}

/*
DOMSnapshot returns the DOMSnapshotProtocol instance.

DOMSnapshot is a Protocoller implementation.
*/
func (socket *Socket) DOMSnapshot() *DOMSnapshotProtocol {
	return socket.domSnapshot
}

/*
DOMStorage returns the DOMStorageProtocol instance.

DOMStorage is a Protocoller implementation.
*/
func (socket *Socket) DOMStorage() *DOMStorageProtocol {
	return socket.domStorage
}

/*
DOM returns the DOMProtocol instance.

DOM is a Protocoller implementation.
*/
func (socket *Socket) DOM() *DOMProtocol {
	return socket.dom
}

/*
Emulation returns the EmulationProtocol instance.

Emulation is a Protocoller implementation.
*/
func (socket *Socket) Emulation() *EmulationProtocol {
	return socket.emulation
}

/*
HeadlessExperimental returns the HeadlessExperimentalProtocol instance.

HeadlessExperimental is a Protocoller implementation.
*/
func (socket *Socket) HeadlessExperimental() *HeadlessExperimentalProtocol {
	return socket.headlessExperimental
}

/*
HeapProfiler returns the HeapProfilerProtocol instance.

HeapProfiler is a Protocoller implementation.
*/
func (socket *Socket) HeapProfiler() *HeapProfilerProtocol {
	return socket.heapProfiler
}

/*
IndexedDB returns the IndexedDBProtocol instance.

IndexedDB is a Protocoller implementation.
*/
func (socket *Socket) IndexedDB() *IndexedDBProtocol {
	return socket.indexedDB
}

/*
Input returns the InputProtocol instance.

Input is a Protocoller implementation.
*/
func (socket *Socket) Input() *InputProtocol {
	return socket.input
}

/*
IO returns the IOProtocol instance.

IO is a Protocoller implementation.
*/
func (socket *Socket) IO() *IOProtocol {
	return socket.io
}

/*
LayerTree returns the LayerTreeProtocol instance.

LayerTree is a Protocoller implementation.
*/
func (socket *Socket) LayerTree() *LayerTreeProtocol {
	return socket.layerTree
}

/*
Log returns the LogProtocol instance.

Log is a Protocoller implementation.
*/
func (socket *Socket) Log() *LogProtocol {
	return socket.log
}

/*
Memory returns the MemoryProtocol instance.

Memory is a Protocoller implementation.
*/
func (socket *Socket) Memory() *MemoryProtocol {
	return socket.memory
}

/*
Network returns the NetworkProtocol instance.

Network is a Protocoller implementation.
*/
func (socket *Socket) Network() *NetworkProtocol {
	return socket.network
}

/*
Overlay returns the OverlayProtocol instance.

Overlay is a Protocoller implementation.
*/
func (socket *Socket) Overlay() *OverlayProtocol {
	return socket.overlay
}

/*
Page returns the PageProtocol instance.

Page is a Protocoller implementation.
*/
func (socket *Socket) Page() *PageProtocol {
	return socket.page
}

/*
Performance returns the PerformanceProtocol instance.

Performance is a Protocoller implementation.
*/
func (socket *Socket) Performance() *PerformanceProtocol {
	return socket.performance
}

/*
Profiler returns the ProfilerProtocol instance.

Profiler is a Protocoller implementation.
*/
func (socket *Socket) Profiler() *ProfilerProtocol {
	return socket.profiler
}

/*
Runtime returns the RuntimeProtocol instance.

Runtime is a Protocoller implementation.
*/
func (socket *Socket) Runtime() *RuntimeProtocol {
	return socket.runtime
}

/*
Schema returns the SchemaProtocol instance.

Schema is a Protocoller implementation.
*/
func (socket *Socket) Schema() *SchemaProtocol {
	return socket.schema
}

/*
Security returns the SecurityProtocol instance.

Security is a Protocoller implementation.
*/
func (socket *Socket) Security() *SecurityProtocol {
	return socket.security
}

/*
ServiceWorker returns the ServiceWorkerProtocol instance.

ServiceWorker is a Protocoller implementation.
*/
func (socket *Socket) ServiceWorker() *ServiceWorkerProtocol {
	return socket.serviceWorker
}

/*
Storage returns the StorageProtocol instance.

Storage is a Protocoller implementation.
*/
func (socket *Socket) Storage() *StorageProtocol {
	return socket.storage
}

/*
SystemInfo returns the SystemInfoProtocol instance.

SystemInfo is a Protocoller implementation.
*/
func (socket *Socket) SystemInfo() *SystemInfoProtocol {
	return socket.systemInfo
}

/*
Target returns the TargetProtocol instance.

Target is a Protocoller implementation.
*/
func (socket *Socket) Target() *TargetProtocol {
	return socket.target
}

/*
Tethering returns the TetheringProtocol instance.

Tethering is a Protocoller implementation.
*/
func (socket *Socket) Tethering() *TetheringProtocol {
	return socket.tethering
}

/*
Tracing returns the TracingProtocol instance.

Tracing is a Protocoller implementation.
*/
func (socket *Socket) Tracing() *TracingProtocol {
	return socket.tracing
}

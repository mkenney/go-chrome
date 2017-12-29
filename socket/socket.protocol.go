package socket

import (
	"encoding/json"
)

/*
MarshalResult abstracts marshalling Commander results into protocol data structs
*/
func MarshalResult(command Commander, result interface{}) error {
	if nil != command.Result() {
		err := json.Unmarshal(command.Result(), &result)
		if nil != err {
			return err
		}
	}
	return nil
}

/*
Accessibility implements Socketer
*/
func (socket *socket) Accessibility() *AccessibilityProtocol {
	return socket.accessibility
}

/*
Animation implements Socketer
*/
func (socket *socket) Animation() *AnimationProtocol {
	return socket.animation
}

/*
ApplicationCache implements Socketer
*/
func (socket *socket) ApplicationCache() *ApplicationCacheProtocol {
	return socket.applicationCache
}

/*
Audits implements Socketer
*/
func (socket *socket) Audits() *AuditsProtocol {
	return socket.audits
}

/*
Browser implements Socketer
*/
func (socket *socket) Browser() *BrowserProtocol {
	return socket.browser
}

/*
CacheStorage implements Socketer
*/
func (socket *socket) CacheStorage() *CacheStorageProtocol {
	return socket.cacheStorage
}

/*
Console implements Socketer
*/
func (socket *socket) Console() *ConsoleProtocol {
	return socket.console
}

/*
CSS implements Socketer
*/
func (socket *socket) CSS() *CSSProtocol {
	return socket.css
}

/*
Database implements Socketer
*/
func (socket *socket) Database() *DatabaseProtocol {
	return socket.database
}

/*
Debugger implements Socketer
*/
func (socket *socket) Debugger() *DebuggerProtocol {
	return socket.debugger
}

/*
DeviceOrientation implements Socketer
*/
func (socket *socket) DeviceOrientation() *DeviceOrientationProtocol {
	return socket.deviceOrientation
}

/*
DOMDebugger implements Socketer
*/
func (socket *socket) DOMDebugger() *DOMDebuggerProtocol {
	return socket.domDebugger
}

/*
DOMSnapshot implements Socketer
*/
func (socket *socket) DOMSnapshot() *DOMSnapshotProtocol {
	return socket.domSnapshot
}

/*
DOMStorage implements Socketer
*/
func (socket *socket) DOMStorage() *DOMStorageProtocol {
	return socket.domStorage
}

/*
DOM implements Socketer
*/
func (socket *socket) DOM() *DOMProtocol {
	return socket.dom
}

/*
Emulation implements Socketer
*/
func (socket *socket) Emulation() *EmulationProtocol {
	return socket.emulation
}

/*
HeadlessExperimental implements Socketer
*/
func (socket *socket) HeadlessExperimental() *HeadlessExperimentalProtocol {
	return socket.headlessExperimental
}

/*
HeapProfiler implements Socketer
*/
func (socket *socket) HeapProfiler() *HeapProfilerProtocol {
	return socket.heapProfiler
}

/*
IndexedDB implements Socketer
*/
func (socket *socket) IndexedDB() *IndexedDBProtocol {
	return socket.indexedDB
}

/*
Input implements Socketer
*/
func (socket *socket) Input() *InputProtocol {
	return socket.input
}

/*
IO implements Socketer
*/
func (socket *socket) IO() *IOProtocol {
	return socket.io
}

/*
LayerTree implements Socketer
*/
func (socket *socket) LayerTree() *LayerTreeProtocol {
	return socket.layerTree
}

/*
Log implements Socketer
*/
func (socket *socket) Log() *LogProtocol {
	return socket.log
}

/*
Memory implements Socketer
*/
func (socket *socket) Memory() *MemoryProtocol {
	return socket.memory
}

/*
Network implements Socketer
*/
func (socket *socket) Network() *NetworkProtocol {
	return socket.network
}

/*
Overlay implements Socketer
*/
func (socket *socket) Overlay() *OverlayProtocol {
	return socket.overlay
}

/*
Page implements Socketer
*/
func (socket *socket) Page() *PageProtocol {
	return socket.page
}

/*
Performance implements Socketer
*/
func (socket *socket) Performance() *PerformanceProtocol {
	return socket.performance
}

/*
Profiler implements Socketer
*/
func (socket *socket) Profiler() *ProfilerProtocol {
	return socket.profiler
}

/*
Runtime implements Socketer
*/
func (socket *socket) Runtime() *RuntimeProtocol {
	return socket.runtime
}

/*
Schema implements Socketer
*/
func (socket *socket) Schema() *SchemaProtocol {
	return socket.schema
}

/*
Security implements Socketer
*/
func (socket *socket) Security() *SecurityProtocol {
	return socket.security
}

/*
ServiceWorker implements Socketer
*/
func (socket *socket) ServiceWorker() *ServiceWorkerProtocol {
	return socket.serviceWorker
}

/*
Storage implements Socketer
*/
func (socket *socket) Storage() *StorageProtocol {
	return socket.storage
}

/*
SystemInfo implements Socketer
*/
func (socket *socket) SystemInfo() *SystemInfoProtocol {
	return socket.systemInfo
}

/*
Target implements Socketer
*/
func (socket *socket) Target() *TargetProtocol {
	return socket.target
}

/*
Tethering implements Socketer
*/
func (socket *socket) Tethering() *TetheringProtocol {
	return socket.tethering
}

/*
Tracing implements Socketer
*/
func (socket *socket) Tracing() *TracingProtocol {
	return socket.tracing
}

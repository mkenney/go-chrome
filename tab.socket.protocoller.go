package chrome

import (
	"github.com/mkenney/go-chrome/socket"
)

/*
Accessibility implements socket.Protocoller
*/
func (tab *Tab) Accessibility() *socket.AccessibilityProtocol {
	return tab.protocol.Accessibility()
}

/*
Animation implements socket.Protocoller
*/
func (tab *Tab) Animation() *socket.AnimationProtocol {
	return tab.protocol.Animation()
}

/*
ApplicationCache implements socket.Protocoller
*/
func (tab *Tab) ApplicationCache() *socket.ApplicationCacheProtocol {
	return tab.protocol.ApplicationCache()
}

/*
Audits implements socket.Protocoller
*/
func (tab *Tab) Audits() *socket.AuditsProtocol {
	return tab.protocol.Audits()
}

/*
Browser implements socket.Protocoller
*/
func (tab *Tab) Browser() *socket.BrowserProtocol {
	return tab.protocol.Browser()
}

/*
CacheStorage implements socket.Protocoller
*/
func (tab *Tab) CacheStorage() *socket.CacheStorageProtocol {
	return tab.protocol.CacheStorage()
}

/*
Console implements socket.Protocoller
*/
func (tab *Tab) Console() *socket.ConsoleProtocol {
	return tab.protocol.Console()
}

/*
CSS implements socket.Protocoller
*/
func (tab *Tab) CSS() *socket.CSSProtocol {
	return tab.protocol.CSS()
}

/*
Database implements socket.Protocoller
*/
func (tab *Tab) Database() *socket.DatabaseProtocol {
	return tab.protocol.Database()
}

/*
Debugger implements socket.Protocoller
*/
func (tab *Tab) Debugger() *socket.DebuggerProtocol {
	return tab.protocol.Debugger()
}

/*
DeviceOrientation implements socket.Protocoller
*/
func (tab *Tab) DeviceOrientation() *socket.DeviceOrientationProtocol {
	return tab.protocol.DeviceOrientation()
}

/*
DOMDebugger implements socket.Protocoller
*/
func (tab *Tab) DOMDebugger() *socket.DOMDebuggerProtocol {
	return tab.protocol.DOMDebugger()
}

/*
DOMSnapshot implements socket.Protocoller
*/
func (tab *Tab) DOMSnapshot() *socket.DOMSnapshotProtocol {
	return tab.protocol.DOMSnapshot()
}

/*
DOMStorage implements socket.Protocoller
*/
func (tab *Tab) DOMStorage() *socket.DOMStorageProtocol {
	return tab.protocol.DOMStorage()
}

/*
DOM implements socket.Protocoller
*/
func (tab *Tab) DOM() *socket.DOMProtocol {
	return tab.protocol.DOM()
}

/*
Emulation implements socket.Protocoller
*/
func (tab *Tab) Emulation() *socket.EmulationProtocol {
	return tab.protocol.Emulation()
}

/*
HeadlessExperimental implements socket.Protocoller
*/
func (tab *Tab) HeadlessExperimental() *socket.HeadlessExperimentalProtocol {
	return tab.protocol.HeadlessExperimental()
}

/*
HeapProfiler implements socket.Protocoller
*/
func (tab *Tab) HeapProfiler() *socket.HeapProfilerProtocol {
	return tab.protocol.HeapProfiler()
}

/*
IndexedDB implements socket.Protocoller
*/
func (tab *Tab) IndexedDB() *socket.IndexedDBProtocol {
	return tab.protocol.IndexedDB()
}

/*
Input implements socket.Protocoller
*/
func (tab *Tab) Input() *socket.InputProtocol {
	return tab.protocol.Input()
}

/*
IO implements socket.Protocoller
*/
func (tab *Tab) IO() *socket.IOProtocol {
	return tab.protocol.IO()
}

/*
LayerTree implements socket.Protocoller
*/
func (tab *Tab) LayerTree() *socket.LayerTreeProtocol {
	return tab.protocol.LayerTree()
}

/*
Log implements socket.Protocoller
*/
func (tab *Tab) Log() *socket.LogProtocol {
	return tab.protocol.Log()
}

/*
Memory implements socket.Protocoller
*/
func (tab *Tab) Memory() *socket.MemoryProtocol {
	return tab.protocol.Memory()
}

/*
Network implements socket.Protocoller
*/
func (tab *Tab) Network() *socket.NetworkProtocol {
	return tab.protocol.Network()
}

/*
Overlay implements socket.Protocoller
*/
func (tab *Tab) Overlay() *socket.OverlayProtocol {
	return tab.protocol.Overlay()
}

/*
Page implements socket.Protocoller
*/
func (tab *Tab) Page() *socket.PageProtocol {
	return tab.protocol.Page()
}

/*
Performance implements socket.Protocoller
*/
func (tab *Tab) Performance() *socket.PerformanceProtocol {
	return tab.protocol.Performance()
}

/*
Profiler implements socket.Protocoller
*/
func (tab *Tab) Profiler() *socket.ProfilerProtocol {
	return tab.protocol.Profiler()
}

/*
Runtime implements socket.Protocoller
*/
func (tab *Tab) Runtime() *socket.RuntimeProtocol {
	return tab.protocol.Runtime()
}

/*
Schema implements socket.Protocoller
*/
func (tab *Tab) Schema() *socket.SchemaProtocol {
	return tab.protocol.Schema()
}

/*
Security implements socket.Protocoller
*/
func (tab *Tab) Security() *socket.SecurityProtocol {
	return tab.protocol.Security()
}

/*
ServiceWorker implements socket.Protocoller
*/
func (tab *Tab) ServiceWorker() *socket.ServiceWorkerProtocol {
	return tab.protocol.ServiceWorker()
}

/*
Storage implements socket.Protocoller
*/
func (tab *Tab) Storage() *socket.StorageProtocol {
	return tab.protocol.Storage()
}

/*
SystemInfo implements socket.Protocoller
*/
func (tab *Tab) SystemInfo() *socket.SystemInfoProtocol {
	return tab.protocol.SystemInfo()
}

/*
Target implements socket.Protocoller
*/
func (tab *Tab) Target() *socket.TargetProtocol {
	return tab.protocol.Target()
}

/*
Tethering implements socket.Protocoller
*/
func (tab *Tab) Tethering() *socket.TetheringProtocol {
	return tab.protocol.Tethering()
}

/*
Tracing implements socket.Protocoller
*/
func (tab *Tab) Tracing() *socket.TracingProtocol {
	return tab.protocol.Tracing()
}

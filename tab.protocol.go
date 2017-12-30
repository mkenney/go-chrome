package chrome

/*
Accessibility implements Socketer
*/
func (tab *Tab) Accessibility() *AccessibilityProtocol {
	return tab.Socket().Accessibility()
}

/*
Animation implements Socketer
*/
func (tab *Tab) Animation() *AnimationProtocol {
	return tab.Socket().Animation()
}

/*
ApplicationCache implements Socketer
*/
func (tab *Tab) ApplicationCache() *ApplicationCacheProtocol {
	return tab.Socket().ApplicationCache()
}

/*
Audits implements Socketer
*/
func (tab *Tab) Audits() *AuditsProtocol {
	return tab.Socket().Audits()
}

/*
Browser implements Socketer
*/
func (tab *Tab) Browser() *BrowserProtocol {
	return tab.Socket().Browser()
}

/*
CacheStorage implements Socketer
*/
func (tab *Tab) CacheStorage() *CacheStorageProtocol {
	return tab.Socket().CacheStorage()
}

/*
Console implements Socketer
*/
func (tab *Tab) Console() *ConsoleProtocol {
	return tab.Socket().Console()
}

/*
CSS implements Socketer
*/
func (tab *Tab) CSS() *CSSProtocol {
	return tab.Socket().CSS()
}

/*
Database implements Socketer
*/
func (tab *Tab) Database() *DatabaseProtocol {
	return tab.Socket().Database()
}

/*
Debugger implements Socketer
*/
func (tab *Tab) Debugger() *DebuggerProtocol {
	return tab.Socket().Debugger()
}

/*
DeviceOrientation implements Socketer
*/
func (tab *Tab) DeviceOrientation() *DeviceOrientationProtocol {
	return tab.Socket().DeviceOrientation()
}

/*
DOMDebugger implements Socketer
*/
func (tab *Tab) DOMDebugger() *DOMDebuggerProtocol {
	return tab.Socket().DOMDebugger()
}

/*
DOMSnapshot implements Socketer
*/
func (tab *Tab) DOMSnapshot() *DOMSnapshotProtocol {
	return tab.Socket().DOMSnapshot()
}

/*
DOMStorage implements Socketer
*/
func (tab *Tab) DOMStorage() *DOMStorageProtocol {
	return tab.Socket().DOMStorage()
}

/*
DOM implements Socketer
*/
func (tab *Tab) DOM() *DOMProtocol {
	return tab.Socket().DOM()
}

/*
Emulation implements Socketer
*/
func (tab *Tab) Emulation() *EmulationProtocol {
	return tab.Socket().Emulation()
}

/*
HeadlessExperimental implements Socketer
*/
func (tab *Tab) HeadlessExperimental() *HeadlessExperimentalProtocol {
	return tab.Socket().HeadlessExperimental()
}

/*
HeapProfiler implements Socketer
*/
func (tab *Tab) HeapProfiler() *HeapProfilerProtocol {
	return tab.Socket().HeapProfiler()
}

/*
IndexedDB implements Socketer
*/
func (tab *Tab) IndexedDB() *IndexedDBProtocol {
	return tab.Socket().IndexedDB()
}

/*
Input implements Socketer
*/
func (tab *Tab) Input() *InputProtocol {
	return tab.Socket().Input()
}

/*
IO implements Socketer
*/
func (tab *Tab) IO() *IOProtocol {
	return tab.Socket().IO()
}

/*
LayerTree implements Socketer
*/
func (tab *Tab) LayerTree() *LayerTreeProtocol {
	return tab.Socket().LayerTree()
}

/*
Log implements Socketer
*/
func (tab *Tab) Log() *LogProtocol {
	return tab.Socket().Log()
}

/*
Memory implements Socketer
*/
func (tab *Tab) Memory() *MemoryProtocol {
	return tab.Socket().Memory()
}

/*
Network implements Socketer
*/
func (tab *Tab) Network() *NetworkProtocol {
	return tab.Socket().Network()
}

/*
Overlay implements Socketer
*/
func (tab *Tab) Overlay() *OverlayProtocol {
	return tab.Socket().Overlay()
}

/*
Page implements Socketer
*/
func (tab *Tab) Page() *PageProtocol {
	return tab.Socket().Page()
}

/*
Performance implements Socketer
*/
func (tab *Tab) Performance() *PerformanceProtocol {
	return tab.Socket().Performance()
}

/*
Profiler implements Socketer
*/
func (tab *Tab) Profiler() *ProfilerProtocol {
	return tab.Socket().Profiler()
}

/*
Runtime implements Socketer
*/
func (tab *Tab) Runtime() *RuntimeProtocol {
	return tab.Socket().Runtime()
}

/*
Schema implements Socketer
*/
func (tab *Tab) Schema() *SchemaProtocol {
	return tab.Socket().Schema()
}

/*
Security implements Socketer
*/
func (tab *Tab) Security() *SecurityProtocol {
	return tab.Socket().Security()
}

/*
ServiceWorker implements Socketer
*/
func (tab *Tab) ServiceWorker() *ServiceWorkerProtocol {
	return tab.Socket().ServiceWorker()
}

/*
Storage implements Socketer
*/
func (tab *Tab) Storage() *StorageProtocol {
	return tab.Socket().Storage()
}

/*
SystemInfo implements Socketer
*/
func (tab *Tab) SystemInfo() *SystemInfoProtocol {
	return tab.Socket().SystemInfo()
}

/*
Target implements Socketer
*/
func (tab *Tab) Target() *TargetProtocol {
	return tab.Socket().Target()
}

/*
Tethering implements Socketer
*/
func (tab *Tab) Tethering() *TetheringProtocol {
	return tab.Socket().Tethering()
}

/*
Tracing implements Socketer
*/
func (tab *Tab) Tracing() *TracingProtocol {
	return tab.Socket().Tracing()
}

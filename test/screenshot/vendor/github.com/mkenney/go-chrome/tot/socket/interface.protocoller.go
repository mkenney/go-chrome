package socket

/*
Protocoller defines the Chrome DevTools Protocol API methods

https://chromedevtools.github.io/devtools-protocol/
*/
type Protocoller interface {
	// Accessibility returns the AccessibilityProtocol instance.
	Accessibility() *AccessibilityProtocol

	// Animation returns the AnimationProtocol instance.
	Animation() *AnimationProtocol

	// ApplicationCache returns the ApplicationCacheProtocol instance.
	ApplicationCache() *ApplicationCacheProtocol

	// Audits returns the AuditsProtocol instance.
	Audits() *AuditsProtocol

	// Browser returns the BrowserProtocol instance.
	Browser() *BrowserProtocol

	// CacheStorage returns the CacheStorageProtocol instance.
	CacheStorage() *CacheStorageProtocol

	// Console returns the ConsoleProtocol instance.
	Console() *ConsoleProtocol

	// CSS returns the CSSProtocol instance.
	CSS() *CSSProtocol

	// Database returns the DatabaseProtocol instance.
	Database() *DatabaseProtocol

	// Debugger returns the DebuggerProtocol instance.
	Debugger() *DebuggerProtocol

	// DeviceOrientation returns the DeviceOrientationProtocol instance.
	DeviceOrientation() *DeviceOrientationProtocol

	// DOMDebugger returns the DOMDebuggerProtocol instance.
	DOMDebugger() *DOMDebuggerProtocol

	// DOMSnapshot returns the DOMSnapshotProtocol instance.
	DOMSnapshot() *DOMSnapshotProtocol

	// DOMStorage returns the DOMStorageProtocol instance.
	DOMStorage() *DOMStorageProtocol

	// DOM returns the DOMProtocol instance.
	DOM() *DOMProtocol

	// Emulation returns the EmulationProtocol instance.
	Emulation() *EmulationProtocol

	// HeadlessExperimental returns the HeadlessExperimentalProtocol instance.
	HeadlessExperimental() *HeadlessExperimentalProtocol

	// HeapProfiler returns the HeapProfilerProtocol instance.
	HeapProfiler() *HeapProfilerProtocol

	// IndexedDB returns the IndexedDBProtocol instance.
	IndexedDB() *IndexedDBProtocol

	// Input returns the InputProtocol instance.
	Input() *InputProtocol

	// IO returns the IOProtocol instance.
	IO() *IOProtocol

	// LayerTree returns the LayerTreeProtocol instance.
	LayerTree() *LayerTreeProtocol

	// Log returns the LogProtocol instance.
	Log() *LogProtocol

	// Memory returns the MemoryProtocol instance.
	Memory() *MemoryProtocol

	// Network returns the NetworkProtocol instance.
	Network() *NetworkProtocol

	// Overlay returns the OverlayProtocol instance.
	Overlay() *OverlayProtocol

	// Page returns the PageProtocol instance.
	Page() *PageProtocol

	// Performance returns the PerformanceProtocol instance.
	Performance() *PerformanceProtocol

	// Profiler returns the ProfilerProtocol instance.
	Profiler() *ProfilerProtocol

	// Runtime returns the RuntimeProtocol instance.
	Runtime() *RuntimeProtocol

	// Schema returns the SchemaProtocol instance.
	Schema() *SchemaProtocol

	// Security returns the SecurityProtocol instance.
	Security() *SecurityProtocol

	// ServiceWorker returns the ServiceWorkerProtocol instance.
	ServiceWorker() *ServiceWorkerProtocol

	// Storage returns the StorageProtocol instance.
	Storage() *StorageProtocol

	// SystemInfo returns the SystemInfoProtocol instance.
	SystemInfo() *SystemInfoProtocol

	// Target returns the TargetProtocol instance.
	Target() *TargetProtocol

	// Tethering returns the TetheringProtocol instance.
	Tethering() *TetheringProtocol

	// Tracing returns the TracingProtocol instance.
	Tracing() *TracingProtocol
}

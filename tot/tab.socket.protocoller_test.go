package chrome

import (
	"testing"
)

func TestProtocoller(t *testing.T) {
	var err error
	browser := NewMock(
		&Flags{},
		"", //"path/to/chrome",
		"", //"path/to/stderr",
		"", //"path/to/stdout",
		"", //"path/to/workdir",
	)
	tab, err := browser.NewTab("https://TestProtocoller")
	if nil != err {
		t.Errorf("Expected nil, received error: %v", err)
	}

	if testVal := tab.Accessibility(); nil == testVal {
		t.Errorf("Expected struct, received nil")
	}

	if testVal := tab.Animation(); nil == testVal {
		t.Errorf("Expected struct, received nil")
	}

	if testVal := tab.ApplicationCache(); nil == testVal {
		t.Errorf("Expected struct, received nil")
	}

	if testVal := tab.Audits(); nil == testVal {
		t.Errorf("Expected struct, received nil")
	}

	if testVal := tab.Browser(); nil == testVal {
		t.Errorf("Expected struct, received nil")
	}

	if testVal := tab.CacheStorage(); nil == testVal {
		t.Errorf("Expected struct, received nil")
	}

	if testVal := tab.Console(); nil == testVal {
		t.Errorf("Expected struct, received nil")
	}

	if testVal := tab.CSS(); nil == testVal {
		t.Errorf("Expected struct, received nil")
	}

	if testVal := tab.Database(); nil == testVal {
		t.Errorf("Expected struct, received nil")
	}

	if testVal := tab.Debugger(); nil == testVal {
		t.Errorf("Expected struct, received nil")
	}

	if testVal := tab.DeviceOrientation(); nil == testVal {
		t.Errorf("Expected struct, received nil")
	}

	if testVal := tab.DOMDebugger(); nil == testVal {
		t.Errorf("Expected struct, received nil")
	}

	if testVal := tab.DOMSnapshot(); nil == testVal {
		t.Errorf("Expected struct, received nil")
	}

	if testVal := tab.DOMStorage(); nil == testVal {
		t.Errorf("Expected struct, received nil")
	}

	if testVal := tab.DOM(); nil == testVal {
		t.Errorf("Expected struct, received nil")
	}

	if testVal := tab.Emulation(); nil == testVal {
		t.Errorf("Expected struct, received nil")
	}

	if testVal := tab.HeadlessExperimental(); nil == testVal {
		t.Errorf("Expected struct, received nil")
	}

	if testVal := tab.HeapProfiler(); nil == testVal {
		t.Errorf("Expected struct, received nil")
	}

	if testVal := tab.IndexedDB(); nil == testVal {
		t.Errorf("Expected struct, received nil")
	}

	if testVal := tab.Input(); nil == testVal {
		t.Errorf("Expected struct, received nil")
	}

	if testVal := tab.IO(); nil == testVal {
		t.Errorf("Expected struct, received nil")
	}

	if testVal := tab.LayerTree(); nil == testVal {
		t.Errorf("Expected struct, received nil")
	}

	if testVal := tab.Log(); nil == testVal {
		t.Errorf("Expected struct, received nil")
	}

	if testVal := tab.Memory(); nil == testVal {
		t.Errorf("Expected struct, received nil")
	}

	if testVal := tab.Network(); nil == testVal {
		t.Errorf("Expected struct, received nil")
	}

	if testVal := tab.Overlay(); nil == testVal {
		t.Errorf("Expected struct, received nil")
	}

	if testVal := tab.Page(); nil == testVal {
		t.Errorf("Expected struct, received nil")
	}

	if testVal := tab.Performance(); nil == testVal {
		t.Errorf("Expected struct, received nil")
	}

	if testVal := tab.Profiler(); nil == testVal {
		t.Errorf("Expected struct, received nil")
	}

	if testVal := tab.Runtime(); nil == testVal {
		t.Errorf("Expected struct, received nil")
	}

	if testVal := tab.Schema(); nil == testVal {
		t.Errorf("Expected struct, received nil")
	}

	if testVal := tab.Security(); nil == testVal {
		t.Errorf("Expected struct, received nil")
	}

	if testVal := tab.ServiceWorker(); nil == testVal {
		t.Errorf("Expected struct, received nil")
	}

	if testVal := tab.Storage(); nil == testVal {
		t.Errorf("Expected struct, received nil")
	}

	if testVal := tab.SystemInfo(); nil == testVal {
		t.Errorf("Expected struct, received nil")
	}

	if testVal := tab.Target(); nil == testVal {
		t.Errorf("Expected struct, received nil")
	}

	if testVal := tab.Tethering(); nil == testVal {
		t.Errorf("Expected struct, received nil")
	}

	if testVal := tab.Tracing(); nil == testVal {
		t.Errorf("Expected struct, received nil")
	}
}

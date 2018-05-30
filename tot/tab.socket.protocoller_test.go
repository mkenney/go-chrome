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
	go func() {
		testVal := tab.Accessibility()
		if nil != testVal {
			t.Errorf("Expected nil, received %v", testVal)
		}
	}()
	go func() {
		testVal := tab.Animation()
		if nil != testVal {
			t.Errorf("Expected nil, received %v", testVal)
		}
	}()
	go func() {
		testVal := tab.ApplicationCache()
		if nil != testVal {
			t.Errorf("Expected nil, received %v", testVal)
		}
	}()
	go func() {
		testVal := tab.Audits()
		if nil != testVal {
			t.Errorf("Expected nil, received %v", testVal)
		}
	}()
	go func() {
		testVal := tab.Browser()
		if nil != testVal {
			t.Errorf("Expected nil, received %v", testVal)
		}
	}()
	go func() {
		testVal := tab.CacheStorage()
		if nil != testVal {
			t.Errorf("Expected nil, received %v", testVal)
		}
	}()
	go func() {
		testVal := tab.Console()
		if nil != testVal {
			t.Errorf("Expected nil, received %v", testVal)
		}
	}()
	go func() {
		testVal := tab.CSS()
		if nil != testVal {
			t.Errorf("Expected nil, received %v", testVal)
		}
	}()
	go func() {
		testVal := tab.Database()
		if nil != testVal {
			t.Errorf("Expected nil, received %v", testVal)
		}
	}()
	go func() {
		testVal := tab.Debugger()
		if nil != testVal {
			t.Errorf("Expected nil, received %v", testVal)
		}
	}()
	go func() {
		testVal := tab.DeviceOrientation()
		if nil != testVal {
			t.Errorf("Expected nil, received %v", testVal)
		}
	}()
	go func() {
		testVal := tab.DOMDebugger()
		if nil != testVal {
			t.Errorf("Expected nil, received %v", testVal)
		}
	}()
	go func() {
		testVal := tab.DOMSnapshot()
		if nil != testVal {
			t.Errorf("Expected nil, received %v", testVal)
		}
	}()
	go func() {
		testVal := tab.DOMStorage()
		if nil != testVal {
			t.Errorf("Expected nil, received %v", testVal)
		}
	}()
	go func() {
		testVal := tab.DOM()
		if nil != testVal {
			t.Errorf("Expected nil, received %v", testVal)
		}
	}()
	go func() {
		testVal := tab.Emulation()
		if nil != testVal {
			t.Errorf("Expected nil, received %v", testVal)
		}
	}()
	go func() {
		testVal := tab.HeadlessExperimental()
		if nil != testVal {
			t.Errorf("Expected nil, received %v", testVal)
		}
	}()
	go func() {
		testVal := tab.HeapProfiler()
		if nil != testVal {
			t.Errorf("Expected nil, received %v", testVal)
		}
	}()
	go func() {
		testVal := tab.IndexedDB()
		if nil != testVal {
			t.Errorf("Expected nil, received %v", testVal)
		}
	}()
	go func() {
		testVal := tab.Input()
		if nil != testVal {
			t.Errorf("Expected nil, received %v", testVal)
		}
	}()
	go func() {
		testVal := tab.IO()
		if nil != testVal {
			t.Errorf("Expected nil, received %v", testVal)
		}
	}()
	go func() {
		testVal := tab.LayerTree()
		if nil != testVal {
			t.Errorf("Expected nil, received %v", testVal)
		}
	}()
	go func() {
		testVal := tab.Log()
		if nil != testVal {
			t.Errorf("Expected nil, received %v", testVal)
		}
	}()
	go func() {
		testVal := tab.Memory()
		if nil != testVal {
			t.Errorf("Expected nil, received %v", testVal)
		}
	}()
	go func() {
		testVal := tab.Network()
		if nil != testVal {
			t.Errorf("Expected nil, received %v", testVal)
		}
	}()
	go func() {
		testVal := tab.Overlay()
		if nil != testVal {
			t.Errorf("Expected nil, received %v", testVal)
		}
	}()
	go func() {
		testVal := tab.Page()
		if nil != testVal {
			t.Errorf("Expected nil, received %v", testVal)
		}
	}()
	go func() {
		testVal := tab.Performance()
		if nil != testVal {
			t.Errorf("Expected nil, received %v", testVal)
		}
	}()
	go func() {
		testVal := tab.Profiler()
		if nil != testVal {
			t.Errorf("Expected nil, received %v", testVal)
		}
	}()
	go func() {
		testVal := tab.Runtime()
		if nil != testVal {
			t.Errorf("Expected nil, received %v", testVal)
		}
	}()
	go func() {
		testVal := tab.Schema()
		if nil != testVal {
			t.Errorf("Expected nil, received %v", testVal)
		}
	}()
	go func() {
		testVal := tab.Security()
		if nil != testVal {
			t.Errorf("Expected nil, received %v", testVal)
		}
	}()
	go func() {
		testVal := tab.ServiceWorker()
		if nil != testVal {
			t.Errorf("Expected nil, received %v", testVal)
		}
	}()
	go func() {
		testVal := tab.Storage()
		if nil != testVal {
			t.Errorf("Expected nil, received %v", testVal)
		}
	}()
	go func() {
		testVal := tab.SystemInfo()
		if nil != testVal {
			t.Errorf("Expected nil, received %v", testVal)
		}
	}()
	go func() {
		testVal := tab.Target()
		if nil != testVal {
			t.Errorf("Expected nil, received %v", testVal)
		}
	}()
	go func() {
		testVal := tab.Tethering()
		if nil != testVal {
			t.Errorf("Expected nil, received %v", testVal)
		}
	}()
	go func() {
		testVal := tab.Tracing()
		if nil != testVal {
			t.Errorf("Expected nil, received %v", testVal)
		}
	}()

}

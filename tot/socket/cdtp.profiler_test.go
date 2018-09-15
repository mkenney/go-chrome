package socket

import (
	"testing"

	"github.com/mkenney/go-chrome/tot/debugger"
	"github.com/mkenney/go-chrome/tot/profiler"
	"github.com/mkenney/go-chrome/tot/runtime"
)

func TestProfilerDisable(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &profiler.DisableResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Profiler().Disable()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Profiler().Disable()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestProfilerEnable(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &profiler.EnableResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Profiler().Enable()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Profiler().Enable()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestProfilerGetBestEffortCoverage(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &profiler.GetBestEffortCoverageResult{
		Result: []*profiler.ScriptCoverage{{
			ScriptID: runtime.ScriptID("script-id"),
			URL:      "http://some.url",
			Functions: []*profiler.FunctionCoverage{{
				FunctionName: "functionName",
				Ranges: []*profiler.CoverageRange{{
					StartOffset: 1,
					EndOffset:   1,
					Count:       1,
				}},
				IsBlockCoverage: true,
			}},
		}},
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Profiler().GetBestEffortCoverage()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Result[0].ScriptID != result.Result[0].ScriptID {
		t.Errorf("Expected %s, got %s", mockResult.Result[0].ScriptID, result.Result[0].ScriptID)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Profiler().GetBestEffortCoverage()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestProfilerSetSamplingInterval(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &profiler.SetSamplingIntervalParams{
		Interval: 1,
	}
	mockResult := &profiler.SetSamplingIntervalResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Profiler().SetSamplingInterval(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Profiler().SetSamplingInterval(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestProfilerStart(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &profiler.StartResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Profiler().Start()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Profiler().Start()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestProfilerStartPreciseCoverage(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &profiler.StartPreciseCoverageParams{
		CallCount: true,
		Detailed:  true,
	}
	mockResult := &profiler.StartPreciseCoverageResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Profiler().StartPreciseCoverage(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Profiler().StartPreciseCoverage(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestProfilerStartTypeProfile(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &profiler.StartTypeProfileResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Profiler().StartTypeProfile()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Profiler().StartTypeProfile()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestProfilerStop(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &profiler.StopResult{
		Profile: &profiler.Profile{
			Nodes: []*profiler.ProfileNode{{
				ID:          1,
				CallFrame:   &runtime.CallFrame{},
				HitCount:    1,
				Children:    []int{1, 2},
				DeoptReason: "reason",
				PositionTicks: []*profiler.PositionTickInfo{{
					Line:  1,
					Ticks: 1,
				}},
			}},
			StartTime:  1,
			EndTime:    1,
			Samples:    []int{1, 2},
			TimeDeltas: []int{1, 2},
		},
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Profiler().Stop()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Profile.Nodes[0].ID != result.Profile.Nodes[0].ID {
		t.Errorf("Expected %d, got %d", mockResult.Profile.Nodes[0].ID, result.Profile.Nodes[0].ID)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Profiler().Stop()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestProfilerStopPreciseCoverage(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &profiler.StopPreciseCoverageResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Profiler().StopPreciseCoverage()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Profiler().StopPreciseCoverage()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestProfilerStopTypeProfile(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &profiler.StopTypeProfileResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Profiler().StopTypeProfile()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Profiler().StopTypeProfile()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestProfilerTakePreciseCoverage(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &profiler.TakePreciseCoverageResult{
		Result: []*profiler.ScriptCoverage{{
			ScriptID: runtime.ScriptID("script-id"),
			URL:      "http://some.url",
			Functions: []*profiler.FunctionCoverage{{
				FunctionName: "functionName",
				Ranges: []*profiler.CoverageRange{{
					StartOffset: 1,
					EndOffset:   1,
					Count:       1,
				}},
				IsBlockCoverage: true,
			}},
		}},
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Profiler().TakePreciseCoverage()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Result[0].ScriptID != result.Result[0].ScriptID {
		t.Errorf("Expected %s, got %s", mockResult.Result[0].ScriptID, result.Result[0].ScriptID)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Profiler().TakePreciseCoverage()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestProfilerTakeTypeProfile(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &profiler.TakeTypeProfileResult{
		Result: []*profiler.ScriptTypeProfile{{
			ScriptID: runtime.ScriptID("script-id"),
			URL:      "http://some.url",
			Entries: []*profiler.TypeProfileEntry{{
				Offset: 1,
				Types: []*profiler.TypeObject{{
					Name: "object name",
				}},
			}},
		}},
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Profiler().TakeTypeProfile()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Result[0].ScriptID != result.Result[0].ScriptID {
		t.Errorf("Expected %s, got %s", mockResult.Result[0].ScriptID, result.Result[0].ScriptID)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Profiler().TakeTypeProfile()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestProfilerOnConsoleProfileFinished(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *profiler.ConsoleProfileFinishedEvent)
	soc.Profiler().OnConsoleProfileFinished(func(eventData *profiler.ConsoleProfileFinishedEvent) {
		resultChan <- eventData
	})

	mockResult := &profiler.ConsoleProfileFinishedEvent{
		ID:       "finished-id",
		Location: &debugger.Location{},
		Profile: &profiler.Profile{
			Nodes: []*profiler.ProfileNode{{
				ID:          1,
				CallFrame:   &runtime.CallFrame{},
				HitCount:    1,
				Children:    []int{1, 2},
				DeoptReason: "reason",
				PositionTicks: []*profiler.PositionTickInfo{{
					Line:  1,
					Ticks: 1,
				}},
			}},
			StartTime:  1,
			EndTime:    1,
			Samples:    []int{1, 2},
			TimeDeltas: []int{1, 2},
		},
		Title: "title",
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Profiler.consoleProfileFinished",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.ID != result.ID {
		t.Errorf("Expected %s, got %s", mockResult.ID, result.ID)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Profiler.consoleProfileFinished",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestProfilerOnConsoleProfileStarted(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *profiler.ConsoleProfileStartedEvent)
	soc.Profiler().OnConsoleProfileStarted(func(eventData *profiler.ConsoleProfileStartedEvent) {
		resultChan <- eventData
	})

	mockResult := &profiler.ConsoleProfileStartedEvent{
		ID:       "event-id",
		Location: &debugger.Location{},
		Title:    "title",
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Profiler.consoleProfileStarted",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.ID != result.ID {
		t.Errorf("Expected %s, got %s", mockResult.ID, result.ID)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Profiler.consoleProfileStarted",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

package socket

import (
	"encoding/json"
	"net/url"
	"testing"

	"github.com/mkenney/go-chrome/tot/debugger"
	"github.com/mkenney/go-chrome/tot/profiler"
	"github.com/mkenney/go-chrome/tot/runtime"
)

func TestProfilerDisable(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestProfilerDisable")
	mockSocket := NewMock(socketURL)
	go func() {_ = mockSocket.Listen()}()
	defer mockSocket.Stop()

	resultChan := mockSocket.Profiler().Disable()
	mockResult := &profiler.DisableResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.Profiler().Disable()
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestProfilerEnable(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestProfilerEnable")
	mockSocket := NewMock(socketURL)
	go func() {_ = mockSocket.Listen()}()
	defer mockSocket.Stop()

	resultChan := mockSocket.Profiler().Enable()
	mockResult := &profiler.EnableResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.Profiler().Enable()
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestProfilerGetBestEffortCoverage(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestProfilerGetBestEffortCoverage")
	mockSocket := NewMock(socketURL)
	go func() {_ = mockSocket.Listen()}()
	defer mockSocket.Stop()

	resultChan := mockSocket.Profiler().GetBestEffortCoverage()
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
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Result[0].ScriptID != result.Result[0].ScriptID {
		t.Errorf("Expected %s, got %s", mockResult.Result[0].ScriptID, result.Result[0].ScriptID)
	}

	resultChan = mockSocket.Profiler().GetBestEffortCoverage()
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestProfilerSetSamplingInterval(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestProfilerSetSamplingInterval")
	mockSocket := NewMock(socketURL)
	go func() {_ = mockSocket.Listen()}()
	defer mockSocket.Stop()

	params := &profiler.SetSamplingIntervalParams{
		Interval: 1,
	}
	resultChan := mockSocket.Profiler().SetSamplingInterval(params)
	mockResult := &profiler.SetSamplingIntervalResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.Profiler().SetSamplingInterval(params)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestProfilerStart(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestProfilerStart")
	mockSocket := NewMock(socketURL)
	go func() {_ = mockSocket.Listen()}()
	defer mockSocket.Stop()

	resultChan := mockSocket.Profiler().Start()
	mockResult := &profiler.StartResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.Profiler().Start()
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestProfilerStartPreciseCoverage(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestProfilerStartPreciseCoverage")
	mockSocket := NewMock(socketURL)
	go func() {_ = mockSocket.Listen()}()
	defer mockSocket.Stop()

	params := &profiler.StartPreciseCoverageParams{
		CallCount: true,
		Detailed:  true,
	}
	resultChan := mockSocket.Profiler().StartPreciseCoverage(params)
	mockResult := &profiler.StartPreciseCoverageResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.Profiler().StartPreciseCoverage(params)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestProfilerStartTypeProfile(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestProfilerStartTypeProfile")
	mockSocket := NewMock(socketURL)
	go func() {_ = mockSocket.Listen()}()
	defer mockSocket.Stop()

	resultChan := mockSocket.Profiler().StartTypeProfile()
	mockResult := &profiler.StartTypeProfileResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.Profiler().StartTypeProfile()
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestProfilerStop(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestProfilerStop")
	mockSocket := NewMock(socketURL)
	go func() {_ = mockSocket.Listen()}()
	defer mockSocket.Stop()

	resultChan := mockSocket.Profiler().Stop()
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
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Profile.Nodes[0].ID != result.Profile.Nodes[0].ID {
		t.Errorf("Expected %d, got %d", mockResult.Profile.Nodes[0].ID, result.Profile.Nodes[0].ID)
	}

	resultChan = mockSocket.Profiler().Stop()
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestProfilerStopPreciseCoverage(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestProfilerStopPreciseCoverage")
	mockSocket := NewMock(socketURL)
	go func() {_ = mockSocket.Listen()}()
	defer mockSocket.Stop()

	resultChan := mockSocket.Profiler().StopPreciseCoverage()
	mockResult := &profiler.StopPreciseCoverageResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.Profiler().StopPreciseCoverage()
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestProfilerStopTypeProfile(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestProfilerStopTypeProfile")
	mockSocket := NewMock(socketURL)
	go func() {_ = mockSocket.Listen()}()
	defer mockSocket.Stop()

	resultChan := mockSocket.Profiler().StopTypeProfile()
	mockResult := &profiler.StopTypeProfileResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.Profiler().StopTypeProfile()
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestProfilerTakePreciseCoverage(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestProfilerTakePreciseCoverage")
	mockSocket := NewMock(socketURL)
	go func() {_ = mockSocket.Listen()}()
	defer mockSocket.Stop()

	resultChan := mockSocket.Profiler().TakePreciseCoverage()
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
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Result[0].ScriptID != result.Result[0].ScriptID {
		t.Errorf("Expected %s, got %s", mockResult.Result[0].ScriptID, result.Result[0].ScriptID)
	}

	resultChan = mockSocket.Profiler().TakePreciseCoverage()
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestProfilerTakeTypeProfile(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestProfilerTakeTypeProfile")
	mockSocket := NewMock(socketURL)
	go func() {_ = mockSocket.Listen()}()
	defer mockSocket.Stop()

	resultChan := mockSocket.Profiler().TakeTypeProfile()
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
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Result[0].ScriptID != result.Result[0].ScriptID {
		t.Errorf("Expected %s, got %s", mockResult.Result[0].ScriptID, result.Result[0].ScriptID)
	}

	resultChan = mockSocket.Profiler().TakeTypeProfile()
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestProfilerOnConsoleProfileFinished(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestProfilerOnConsoleProfileFinished")
	mockSocket := NewMock(socketURL)
	go func() {_ = mockSocket.Listen()}()
	defer mockSocket.Stop()

	resultChan := make(chan *profiler.ConsoleProfileFinishedEvent)
	mockSocket.Profiler().OnConsoleProfileFinished(func(eventData *profiler.ConsoleProfileFinishedEvent) {
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
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Profiler.consoleProfileFinished",
		Params: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.ID != result.ID {
		t.Errorf("Expected %s, got %s", mockResult.ID, result.ID)
	}

	resultChan = make(chan *profiler.ConsoleProfileFinishedEvent)
	mockSocket.Profiler().OnConsoleProfileFinished(func(eventData *profiler.ConsoleProfileFinishedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Profiler.consoleProfileFinished",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestProfilerOnConsoleProfileStarted(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestProfilerOnConsoleProfileStarted")
	mockSocket := NewMock(socketURL)
	go func() {_ = mockSocket.Listen()}()
	defer mockSocket.Stop()

	resultChan := make(chan *profiler.ConsoleProfileStartedEvent)
	mockSocket.Profiler().OnConsoleProfileStarted(func(eventData *profiler.ConsoleProfileStartedEvent) {
		resultChan <- eventData
	})
	mockResult := &profiler.ConsoleProfileStartedEvent{
		ID:       "event-id",
		Location: &debugger.Location{},
		Title:    "title",
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Profiler.consoleProfileStarted",
		Params: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.ID != result.ID {
		t.Errorf("Expected %s, got %s", mockResult.ID, result.ID)
	}

	resultChan = make(chan *profiler.ConsoleProfileStartedEvent)
	mockSocket.Profiler().OnConsoleProfileStarted(func(eventData *profiler.ConsoleProfileStartedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Profiler.consoleProfileStarted",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

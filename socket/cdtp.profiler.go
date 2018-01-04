package socket

import (
	"encoding/json"

	profiler "github.com/mkenney/go-chrome/cdtp/profiler"
	log "github.com/sirupsen/logrus"
)

/*
ProfilerProtocol provides a namespace for the Chrome Profiler protocol methods.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/
*/
type ProfilerProtocol struct {
	Socket Socketer
}

/*
Disable disables profiling.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-disable
*/
func (protocol *ProfilerProtocol) Disable() chan *profiler.DisableResult {
	resultChan := make(chan *profiler.DisableResult)
	command := NewCommand(protocol.Socket, "Profiler.disable", nil)
	result := &profiler.DisableResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
Enable enables profiling.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-enable
*/
func (protocol *ProfilerProtocol) Enable() chan *profiler.EnableResult {
	resultChan := make(chan *profiler.EnableResult)
	command := NewCommand(protocol.Socket, "Profiler.enable", nil)
	result := &profiler.EnableResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
GetBestEffortCoverage collects coverage data for the current isolate. The
coverage data may be incomplete due to garbage collection.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-getBestEffortCoverage
*/
func (protocol *ProfilerProtocol) GetBestEffortCoverage() chan *profiler.GetBestEffortCoverageResult {
	resultChan := make(chan *profiler.GetBestEffortCoverageResult)
	command := NewCommand(protocol.Socket, "Profiler.getBestEffortCoverage", nil)
	result := &profiler.GetBestEffortCoverageResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		} else {
			result.CDTPError = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
}

/*
SetSamplingInterval changes CPU profiler sampling interval. Must be called
before CPU profiles recording started.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-setSamplingInterval
*/
func (protocol *ProfilerProtocol) SetSamplingInterval(
	params *profiler.SetSamplingIntervalParams,
) chan *profiler.SetSamplingIntervalResult {
	resultChan := make(chan *profiler.SetSamplingIntervalResult)
	command := NewCommand(protocol.Socket, "Profiler.setSamplingInterval", params)
	result := &profiler.SetSamplingIntervalResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
Start starts profiling.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-start
*/
func (protocol *ProfilerProtocol) Start() chan *profiler.StartResult {
	resultChan := make(chan *profiler.StartResult)
	command := NewCommand(protocol.Socket, "Profiler.start", nil)
	result := &profiler.StartResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
StartPreciseCoverage enable precise code coverage. Coverage data for JavaScript
executed before enabling precise code coverage may be incomplete. Enabling
prevents running optimized code and resets execution counters.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-startPreciseCoverage
*/
func (protocol *ProfilerProtocol) StartPreciseCoverage(
	params *profiler.StartPreciseCoverageParams,
) chan *profiler.StartPreciseCoverageResult {
	resultChan := make(chan *profiler.StartPreciseCoverageResult)
	command := NewCommand(protocol.Socket, "Profiler.startPreciseCoverage", params)
	result := &profiler.StartPreciseCoverageResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
StartTypeProfile enables type profile.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-startTypeProfile
EXPERIMENTAL.
*/
func (protocol *ProfilerProtocol) StartTypeProfile() chan *profiler.StartTypeProfileResult {
	resultChan := make(chan *profiler.StartTypeProfileResult)
	command := NewCommand(protocol.Socket, "Profiler.startTypeProfile", nil)
	result := &profiler.StartTypeProfileResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
Stop stops profiling.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-stop
*/
func (protocol *ProfilerProtocol) Stop() chan *profiler.StopResult {
	resultChan := make(chan *profiler.StopResult)
	command := NewCommand(protocol.Socket, "Profiler.stop", nil)
	result := &profiler.StopResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		} else {
			result.CDTPError = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
}

/*
StopPreciseCoverage disable precise code coverage. Disabling releases
unnecessary execution count records and allows executing optimized code.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-stopPreciseCoverage
*/
func (protocol *ProfilerProtocol) StopPreciseCoverage() chan *profiler.StopPreciseCoverageResult {
	resultChan := make(chan *profiler.StopPreciseCoverageResult)
	command := NewCommand(protocol.Socket, "Profiler.stopPreciseCoverage", nil)
	result := &profiler.StopPreciseCoverageResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
StopTypeProfile disables type profile. Disabling releases type profile data
collected so far.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-stopTypeProfile
EXPERIMENTAL.
*/
func (protocol *ProfilerProtocol) StopTypeProfile() chan *profiler.StopTypeProfileResult {
	resultChan := make(chan *profiler.StopTypeProfileResult)
	command := NewCommand(protocol.Socket, "Profiler.stopTypeProfile", nil)
	result := &profiler.StopTypeProfileResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
TakePreciseCoverage collects coverage data for the current isolate, and resets
execution counters. Precise code coverage needs to have started.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-takePreciseCoverage
*/
func (protocol *ProfilerProtocol) TakePreciseCoverage() chan *profiler.TakePreciseCoverageResult {
	resultChan := make(chan *profiler.TakePreciseCoverageResult)
	command := NewCommand(protocol.Socket, "Profiler.takePreciseCoverage", nil)
	result := &profiler.TakePreciseCoverageResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		} else {
			result.CDTPError = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
}

/*
TakeTypeProfile collect type profile.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-takeTypeProfile
EXPERIMENTAL.
*/
func (protocol *ProfilerProtocol) TakeTypeProfile() chan *profiler.TakeTypeProfileResult {
	resultChan := make(chan *profiler.TakeTypeProfileResult)
	command := NewCommand(protocol.Socket, "Profiler.takeTypeProfile", nil)
	result := &profiler.TakeTypeProfileResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		} else {
			result.CDTPError = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
}

/*
OnConsoleProfileFinished adds a handler to the Profiler.consoleProfileFinished
event. Profiler.consoleProfileFinished fires when profile recording finishes.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#event-consoleProfileFinished
*/
func (protocol *ProfilerProtocol) OnConsoleProfileFinished(
	callback func(event *profiler.ConsoleProfileFinishedEvent),
) {
	handler := NewEventHandler(
		"Profiler.consoleProfileFinished",
		func(response *Response) {
			event := &profiler.ConsoleProfileFinishedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnConsoleProfileStarted adds a handler to the Profiler.consoleProfileStarted
event. Profiler.consoleProfileStarted fires when new profile recording is
started using console.profile() call.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#event-consoleProfileStarted
*/
func (protocol *ProfilerProtocol) OnConsoleProfileStarted(
	callback func(event *profiler.ConsoleProfileStartedEvent),
) {
	handler := NewEventHandler(
		"Profiler.consoleProfileStarted",
		func(response *Response) {
			event := &profiler.ConsoleProfileStartedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

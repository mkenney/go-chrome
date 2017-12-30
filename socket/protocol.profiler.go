package socket

import (
	"encoding/json"

	profiler "github.com/mkenney/go-chrome/protocol/profiler"
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
func (protocol *ProfilerProtocol) Disable() error {
	command := NewCommand("Profiler.disable", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
Enable enables profiling.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-enable
*/
func (protocol *ProfilerProtocol) Enable() error {
	command := NewCommand("Profiler.enable", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
GetBestEffortCoverage collects coverage data for the current isolate. The
coverage data may be incomplete due to garbage collection.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-getBestEffortCoverage
*/
func (protocol *ProfilerProtocol) GetBestEffortCoverage() (*profiler.GetBestEffortCoverageResult, error) {
	command := NewCommand("Profiler.getBestEffortCoverage", nil)
	result := &profiler.GetBestEffortCoverageResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
SetSamplingInterval changes CPU profiler sampling interval. Must be called
before CPU profiles recording started.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-setSamplingInterval
*/
func (protocol *ProfilerProtocol) SetSamplingInterval(
	params *profiler.SetSamplingIntervalParams,
) error {
	command := NewCommand("Profiler.setSamplingInterval", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
Start starts profiling.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-start
*/
func (protocol *ProfilerProtocol) Start() error {
	command := NewCommand("Profiler.start", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
StartPreciseCoverage enable precise code coverage. Coverage data for JavaScript
executed before enabling precise code coverage may be incomplete. Enabling
prevents running optimized code and resets execution counters.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-startPreciseCoverage
*/
func (protocol *ProfilerProtocol) StartPreciseCoverage(
	params *profiler.StartPreciseCoverageParams,
) error {
	command := NewCommand("Profiler.startPreciseCoverage", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
StartTypeProfile enables type profile.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-startTypeProfile
EXPERIMENTAL.
*/
func (protocol *ProfilerProtocol) StartTypeProfile() error {
	command := NewCommand("Profiler.startTypeProfile", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
Stop stops profiling.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-stop
*/
func (protocol *ProfilerProtocol) Stop() (*profiler.StopResult, error) {
	command := NewCommand("Profiler.stop", nil)
	result := &profiler.StopResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
StopPreciseCoverage disable precise code coverage. Disabling releases
unnecessary execution count records and allows executing optimized code.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-stopPreciseCoverage
*/
func (protocol *ProfilerProtocol) StopPreciseCoverage() error {
	command := NewCommand("Profiler.stopPreciseCoverage", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
StopTypeProfile disables type profile. Disabling releases type profile data
collected so far.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-stopTypeProfile
EXPERIMENTAL.
*/
func (protocol *ProfilerProtocol) StopTypeProfile() error {
	command := NewCommand("Profiler.stopTypeProfile", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
TakePreciseCoverage collects coverage data for the current isolate, and resets
execution counters. Precise code coverage needs to have started.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-takePreciseCoverage
*/
func (protocol *ProfilerProtocol) TakePreciseCoverage() (*profiler.TakePreciseCoverageResult, error) {
	command := NewCommand("Profiler.takePreciseCoverage", nil)
	result := &profiler.TakePreciseCoverageResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
TakeTypeProfile collect type profile.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-takeTypeProfile
EXPERIMENTAL.
*/
func (protocol *ProfilerProtocol) TakeTypeProfile() (*profiler.TakeTypeProfileResult, error) {
	command := NewCommand("Profiler.takeTypeProfile", nil)
	result := &profiler.TakeTypeProfileResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
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

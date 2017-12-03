package chrome

import (
	"app/chrome/protocol"
	"encoding/json"

	log "github.com/Sirupsen/logrus"
)

/*
Profiler - https://chromedevtools.github.io/devtools-protocol/tot/Profiler/
Provides various functionality related to drawing atop the inspected page. EXPERIMENTAL
*/
type Profiler struct{}

/*
Disable disables profiling.
*/
func (Profiler) Disable(socket *Socket) error {
	command := &protocol.Command{
		method: "Profiler.disable",
		params: nil,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Enable enables profiling.
*/
func (Profiler) Enable(socket *Socket) error {
	command := &protocol.Command{
		method: "Profiler.enable",
		params: nil,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetSamplingInterval changes CPU profiler sampling interval. Must be called before CPU profiles
recording started.
*/
func (Profiler) SetSamplingInterval(socket *Socket, params *Profiler.SetSamplingIntervalParams) error {
	command := &protocol.Command{
		method: "Profiler.setSamplingInterval",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Start starts profiling.
*/
func (Profiler) Start(socket *Socket) error {
	command := &protocol.Command{
		method: "Profiler.start",
		params: nil,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Stop stops profiling.
*/
func (Profiler) Stop(socket *Socket) error {
	command := &protocol.Command{
		method: "Profiler.stop",
		params: nil,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
StartPreciseCoverage enable precise code coverage. Coverage data for JavaScript executed before
enabling precise code coverage may be incomplete. Enabling prevents running optimized code and
resets execution counters.
*/
func (Profiler) StartPreciseCoverage(socket *Socket, params *Profiler.StartPreciseCoverageParams) error {
	command := &protocol.Command{
		method: "Profiler.startPreciseCoverage",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
StopPreciseCoverage disable precise code coverage. Disabling releases unnecessary execution count
records and allows executing optimized code.
*/
func (Profiler) StopPreciseCoverage(socket *Socket) error {
	command := &protocol.Command{
		method: "Profiler.stopPreciseCoverage",
		params: nil,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
TakePreciseCoverage collects coverage data for the current isolate, and resets execution counters.
Precise code coverage needs to have started.
*/
func (Profiler) TakePreciseCoverage(socket *Socket) error {
	command := &protocol.Command{
		method: "Profiler.takePreciseCoverage",
		params: nil,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
GetBestEffortCoverage collects coverage data for the current isolate. The coverage data may be
incomplete due to garbage collection.
*/
func (Profiler) GetBestEffortCoverage(socket *Socket) error {
	command := &protocol.Command{
		method: "Profiler.getBestEffortCoverage",
		params: nil,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
StartTypeProfile enables type profile. EXPERIMENTAL
*/
func (Profiler) StartTypeProfile(socket *Socket) error {
	command := &protocol.Command{
		method: "Profiler.startTypeProfile",
		params: nil,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
StopTypeProfile disables type profile. Disabling releases type profile data collected so far.
EXPERIMENTAL
*/
func (Profiler) StopTypeProfile(socket *Socket) error {
	command := &protocol.Command{
		method: "Profiler.stopTypeProfile",
		params: nil,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
TakeTypeProfile collect type profile. EXPERIMENTAL
*/
func (Profiler) TakeTypeProfile(socket *Socket) error {
	command := &protocol.Command{
		method: "Profiler.takeTypeProfile",
		params: nil,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
OnConsoleProfileStarted adds a handler to the Profiler.consoleProfileStarted event.
Profiler.consoleProfileStarted fires when new profile recording is started using console.profile()
call.
*/
func (Profiler) OnConsoleProfileStarted(socket *Socket, callback func(event *profiler.ConsoleProfileStartedEvent)) error {
	handler := protocol.NewEventHandler(
		"Profiler.consoleProfileStarted",
		func(name string, params []byte) {
			event := &profiler.ConsoleProfileStartedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
	return command.Err
}

/*
OnConsoleProfileFinished adds a handler to the Profiler.consoleProfileFinished event.
Profiler.consoleProfileFinished fires when profile recording finishes.
*/
func (Profiler) OnConsoleProfileFinished(socket *Socket, callback func(event *profiler.ConsoleProfileFinishedEvent)) error {
	handler := protocol.NewEventHandler(
		"Profiler.consoleProfileFinished",
		func(name string, params []byte) {
			event := &profiler.ConsoleProfileFinishedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
	return command.Err
}

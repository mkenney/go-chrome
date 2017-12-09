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
func (Profiler) Disable(
	socket *Socket,
) (nil, error) {
	command := &protocol.Command{
		method: "Profiler.disable",
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
Enable enables profiling.
*/
func (Profiler) Enable(
	socket *Socket,
) (nil, error) {
	command := &protocol.Command{
		method: "Profiler.enable",
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
GetBestEffortCoverage collects coverage data for the current isolate. The coverage data may be
incomplete due to garbage collection.
*/
func (Profiler) GetBestEffortCoverage(
	socket *Socket,
) (Profiler.GetBestEffortCoverageResult, error) {
	command := &protocol.Command{
		method: "Profiler.getBestEffortCoverage",
	}
	socket.SendCommand(command)
	return command.Result.(Profiler.GetBestEffortCoverageResult), command.Err
}

/*
SetSamplingInterval changes CPU profiler sampling interval. Must be called before CPU profiles
recording started.
*/
func (Profiler) SetSamplingInterval(
	socket *Socket,
	params *Profiler.SetSamplingIntervalParams,
) (nil, error) {
	command := &protocol.Command{
		method: "Profiler.setSamplingInterval",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
Start starts profiling.
*/
func (Profiler) Start(
	socket *Socket,
) (nil, error) {
	command := &protocol.Command{
		method: "Profiler.start",
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
StartPreciseCoverage enable precise code coverage. Coverage data for JavaScript executed before
enabling precise code coverage may be incomplete. Enabling prevents running optimized code and
resets execution counters.
*/
func (Profiler) StartPreciseCoverage(
	socket *Socket,
	params *Profiler.StartPreciseCoverageParams,
) (nil, error) {
	command := &protocol.Command{
		method: "Profiler.startPreciseCoverage",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
StartTypeProfile enables type profile. EXPERIMENTAL
*/
func (Profiler) StartTypeProfile(
	socket *Socket,
) (nil, error) {
	command := &protocol.Command{
		method: "Profiler.startTypeProfile",
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
Stop stops profiling.
*/
func (Profiler) Stop(
	socket *Socket,
) (Profiler.StopResult, error) {
	command := &protocol.Command{
		method: "Profiler.stop",
	}
	socket.SendCommand(command)
	return command.Result.(Profiler.StopResult), command.Err
}

/*
StopPreciseCoverage disable precise code coverage. Disabling releases unnecessary execution count
records and allows executing optimized code.
*/
func (Profiler) StopPreciseCoverage(
	socket *Socket,
) (nil, error) {
	command := &protocol.Command{
		method: "Profiler.stopPreciseCoverage",
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
StopTypeProfile disables type profile. Disabling releases type profile data collected so far.
EXPERIMENTAL
*/
func (Profiler) StopTypeProfile(
	socket *Socket,
) (nil, error) {
	command := &protocol.Command{
		method: "Profiler.stopTypeProfile",
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
TakePreciseCoverage collects coverage data for the current isolate, and resets execution counters.
Precise code coverage needs to have started.
*/
func (Profiler) TakePreciseCoverage(
	socket *Socket,
) (Profiler.TakePreciseCoverageResult, error) {
	command := &protocol.Command{
		method: "Profiler.takePreciseCoverage",
	}
	socket.SendCommand(command)
	return command.Result.(Profiler.TakePreciseCoverageResult), command.Err
}

/*
TakeTypeProfile collect type profile. EXPERIMENTAL
*/
func (Profiler) TakeTypeProfile(
	socket *Socket,
) (Profiler.TakeTypeProfileResult, error) {
	command := &protocol.Command{
		method: "Profiler.takeTypeProfile",
	}
	socket.SendCommand(command)
	return command.Result.(Profiler.TakeTypeProfileResult), command.Err
}

/*
OnConsoleProfileStarted adds a handler to the Profiler.consoleProfileStarted event.
Profiler.consoleProfileStarted fires when new profile recording is started using console.profile()
call.
*/
func (Profiler) OnConsoleProfileStarted(
	socket *Socket,
	callback func(event *profiler.ConsoleProfileStartedEvent),
) {
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
}

/*
OnConsoleProfileFinished adds a handler to the Profiler.consoleProfileFinished event.
Profiler.consoleProfileFinished fires when profile recording finishes.
*/
func (Profiler) OnConsoleProfileFinished(
	socket *Socket,
	callback func(event *profiler.ConsoleProfileFinishedEvent),
) {
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
}

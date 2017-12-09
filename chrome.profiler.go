package chrome

import (
	profiler "app/chrome/profiler"
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
) error {
	command := &protocol.Command{
		Method: "Profiler.disable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Enable enables profiling.
*/
func (Profiler) Enable(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "Profiler.enable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
GetBestEffortCoverage collects coverage data for the current isolate. The coverage data may be
incomplete due to garbage collection.
*/
func (Profiler) GetBestEffortCoverage(
	socket *Socket,
) (profiler.GetBestEffortCoverageResult, error) {
	command := &protocol.Command{
		Method: "Profiler.getBestEffortCoverage",
	}
	socket.SendCommand(command)
	return command.Result.(profiler.GetBestEffortCoverageResult), command.Err
}

/*
SetSamplingInterval changes CPU profiler sampling interval. Must be called before CPU profiles
recording started.
*/
func (Profiler) SetSamplingInterval(
	socket *Socket,
	params *profiler.SetSamplingIntervalParams,
) error {
	command := &protocol.Command{
		Method: "Profiler.setSamplingInterval",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Start starts profiling.
*/
func (Profiler) Start(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "Profiler.start",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
StartPreciseCoverage enable precise code coverage. Coverage data for JavaScript executed before
enabling precise code coverage may be incomplete. Enabling prevents running optimized code and
resets execution counters.
*/
func (Profiler) StartPreciseCoverage(
	socket *Socket,
	params *profiler.StartPreciseCoverageParams,
) error {
	command := &protocol.Command{
		Method: "Profiler.startPreciseCoverage",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
StartTypeProfile enables type profile. EXPERIMENTAL
*/
func (Profiler) StartTypeProfile(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "Profiler.startTypeProfile",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Stop stops profiling.
*/
func (Profiler) Stop(
	socket *Socket,
) (profiler.StopResult, error) {
	command := &protocol.Command{
		Method: "Profiler.stop",
	}
	socket.SendCommand(command)
	return command.Result.(profiler.StopResult), command.Err
}

/*
StopPreciseCoverage disable precise code coverage. Disabling releases unnecessary execution count
records and allows executing optimized code.
*/
func (Profiler) StopPreciseCoverage(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "Profiler.stopPreciseCoverage",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
StopTypeProfile disables type profile. Disabling releases type profile data collected so far.
EXPERIMENTAL
*/
func (Profiler) StopTypeProfile(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "Profiler.stopTypeProfile",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
TakePreciseCoverage collects coverage data for the current isolate, and resets execution counters.
Precise code coverage needs to have started.
*/
func (Profiler) TakePreciseCoverage(
	socket *Socket,
) (profiler.TakePreciseCoverageResult, error) {
	command := &protocol.Command{
		Method: "Profiler.takePreciseCoverage",
	}
	socket.SendCommand(command)
	return command.Result.(profiler.TakePreciseCoverageResult), command.Err
}

/*
TakeTypeProfile collect type profile. EXPERIMENTAL
*/
func (Profiler) TakeTypeProfile(
	socket *Socket,
) (profiler.TakeTypeProfileResult, error) {
	command := &protocol.Command{
		Method: "Profiler.takeTypeProfile",
	}
	socket.SendCommand(command)
	return command.Result.(profiler.TakeTypeProfileResult), command.Err
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

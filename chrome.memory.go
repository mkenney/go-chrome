package chrome

import (
	memory "app/chrome/memory"
	"app/chrome/protocol"
)

/*
Memory - https://chromedevtools.github.io/devtools-protocol/tot/Memory/
EXPERIMENTAL
*/
type Memory struct{}

/*
GetDOMCounters EXPERIMENTAL
*/
func (Memory) GetDOMCounters(
	socket *Socket,
	params *memory.GetDOMCountersParams,
) error {
	command := &protocol.Command{
		Method: "Memory.getDOMCounters",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
PrepareForLeakDetection EXPERIMENTAL
*/
func (Memory) PrepareForLeakDetection(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "Memory.prepareForLeakDetection",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetPressureNotificationsSuppressed enables/disables suppressing memory pressure notifications in all
processes.
*/
func (Memory) SetPressureNotificationsSuppressed(
	socket *Socket,
	params *memory.SetPressureNotificationsSuppressedParams,
) error {
	command := &protocol.Command{
		Method: "Memory.setPressureNotificationsSuppressed",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SimulatePressureNotification simulates a memory pressure notification in all processes.
*/
func (Memory) SimulatePressureNotification(
	socket *Socket,
	params *memory.SimulatePressureNotificationParams,
) error {
	command := &protocol.Command{
		Method: "Memory.simulatePressureNotification",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

package chrome

import (
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
func (Memory) GetDOMCounters(socket *Socket, params *memory.GetDOMCountersParams) error {
	command := &protocol.Command{
		method: "Memory.getDOMCounters",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
PrepareForLeakDetection EXPERIMENTAL
*/
func (Memory) PrepareForLeakDetection(socket *Socket) error {
	command := &protocol.Command{
		method: "Memory.prepareForLeakDetection",
		params: nil,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetPressureNotificationsSuppressed enables/disables suppressing memory pressure notifications in all
processes.
*/
func (Memory) SetPressureNotificationsSuppressed(socket *Socket, params *memory.SetPressureNotificationsSuppressedParams) error {
	command := &protocol.Command{
		method: "Memory.setPressureNotificationsSuppressed",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SimulatePressureNotification simulates a memory pressure notification in all processes.
*/
func (Memory) SimulatePressureNotification(socket *Socket, params *memory.SimulatePressureNotificationParams) error {
	command := &protocol.Command{
		method: "Memory.simulatePressureNotification",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

package protocol

import (
	memory "github.com/mkenney/go-chrome/protocol/memory"
	sock "github.com/mkenney/go-chrome/socket"
)

/*
Memory - https://chromedevtools.github.io/devtools-protocol/tot/Memory/
EXPERIMENTAL
*/
type Memory struct{}

/*
GetDOMCounters EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Memory/#method-getDOMCounters
*/
func (Memory) GetDOMCounters(
	socket *sock.Socket,
	params *memory.GetDOMCountersParams,
) error {
	command := &sock.Command{
		Method: "Memory.getDOMCounters",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
PrepareForLeakDetection EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Memory/#method-prepareForLeakDetection
*/
func (Memory) PrepareForLeakDetection(
	socket *sock.Socket,
) error {
	command := &sock.Command{
		Method: "Memory.prepareForLeakDetection",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetPressureNotificationsSuppressed enables/disables suppressing memory pressure notifications in all
processes.

https://chromedevtools.github.io/devtools-protocol/tot/Memory/#method-setPressureNotificationsSuppressed
*/
func (Memory) SetPressureNotificationsSuppressed(
	socket *sock.Socket,
	params *memory.SetPressureNotificationsSuppressedParams,
) error {
	command := &sock.Command{
		Method: "Memory.setPressureNotificationsSuppressed",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SimulatePressureNotification simulates a memory pressure notification in all processes.

https://chromedevtools.github.io/devtools-protocol/tot/Memory/#method-simulatePressureNotification
*/
func (Memory) SimulatePressureNotification(
	socket *sock.Socket,
	params *memory.SimulatePressureNotificationParams,
) error {
	command := &sock.Command{
		Method: "Memory.simulatePressureNotification",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

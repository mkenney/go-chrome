package protocol

import (
	memory "github.com/mkenney/go-chrome/protocol/memory"
	sock "github.com/mkenney/go-chrome/socket"
)

/*
Memory provides a namespace for the Chrome Memory protocol methods.

https://chromedevtools.github.io/devtools-protocol/tot/Memory/ EXPERIMENTAL.
*/
var Memory = MemoryProtocol{}

/*
MemoryProtocol defines the Memory protocol methods.
*/
type MemoryProtocol struct{}

/*
GetDOMCounters is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/Memory/#method-getDOMCounters EXPERIMENTAL.
*/
func (MemoryProtocol) GetDOMCounters(
	socket sock.Socketer,
	params *memory.GetDOMCountersParams,
) error {
	command := sock.NewCommand("Memory.getDOMCounters", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
PrepareForLeakDetection experimental

https://chromedevtools.github.io/devtools-protocol/tot/Memory/#method-prepareForLeakDetection
EXPERIMENTAL.
*/
func (MemoryProtocol) PrepareForLeakDetection(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("Memory.prepareForLeakDetection", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetPressureNotificationsSuppressed enables/disables suppressing memory pressure notifications in all
processes.

https://chromedevtools.github.io/devtools-protocol/tot/Memory/#method-setPressureNotificationsSuppressed
*/
func (MemoryProtocol) SetPressureNotificationsSuppressed(
	socket sock.Socketer,
	params *memory.SetPressureNotificationsSuppressedParams,
) error {
	command := sock.NewCommand("Memory.setPressureNotificationsSuppressed", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SimulatePressureNotification simulates a memory pressure notification in all processes.

https://chromedevtools.github.io/devtools-protocol/tot/Memory/#method-simulatePressureNotification
*/
func (MemoryProtocol) SimulatePressureNotification(
	socket sock.Socketer,
	params *memory.SimulatePressureNotificationParams,
) error {
	command := sock.NewCommand("Memory.simulatePressureNotification", params)
	socket.SendCommand(command)
	return command.Error()
}

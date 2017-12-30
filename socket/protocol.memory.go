package socket

import (
	memory "github.com/mkenney/go-chrome/protocol/memory"
)

/*
MemoryProtocol provides a namespace for the Chrome Memory protocol methods.

https://chromedevtools.github.io/devtools-protocol/tot/Memory/ EXPERIMENTAL.
*/
type MemoryProtocol struct {
	Socket Socketer
}

/*
GetDOMCounters is experimental.

https://chromedevtools.github.io/devtools-protocol/tot/Memory/#method-getDOMCounters
EXPERIMENTAL.
*/
func (protocol *MemoryProtocol) GetDOMCounters(
	params *memory.GetDOMCountersParams,
) error {
	command := NewCommand("Memory.getDOMCounters", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
PrepareForLeakDetection experimental

https://chromedevtools.github.io/devtools-protocol/tot/Memory/#method-prepareForLeakDetection
EXPERIMENTAL.
*/
func (protocol *MemoryProtocol) PrepareForLeakDetection() error {
	command := NewCommand("Memory.prepareForLeakDetection", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetPressureNotificationsSuppressed enables/disables suppressing memory pressure
notifications in all processes.

https://chromedevtools.github.io/devtools-protocol/tot/Memory/#method-setPressureNotificationsSuppressed
*/
func (protocol *MemoryProtocol) SetPressureNotificationsSuppressed(
	params *memory.SetPressureNotificationsSuppressedParams,
) error {
	command := NewCommand("Memory.setPressureNotificationsSuppressed", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SimulatePressureNotification simulates a memory pressure notification in all
processes.

https://chromedevtools.github.io/devtools-protocol/tot/Memory/#method-simulatePressureNotification
*/
func (protocol *MemoryProtocol) SimulatePressureNotification(
	params *memory.SimulatePressureNotificationParams,
) error {
	command := NewCommand("Memory.simulatePressureNotification", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

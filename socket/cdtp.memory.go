package socket

import (
	memory "github.com/mkenney/go-chrome/cdtp/memory"
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
) chan *memory.GetDOMCountersResult {
	resultChan := make(chan *memory.GetDOMCountersResult)
	command := NewCommand(protocol.Socket, "Memory.getDOMCounters", params)
	result := &memory.GetDOMCountersResult{}

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
PrepareForLeakDetection experimental

https://chromedevtools.github.io/devtools-protocol/tot/Memory/#method-prepareForLeakDetection
EXPERIMENTAL.
*/
func (protocol *MemoryProtocol) PrepareForLeakDetection() chan *memory.PrepareForLeakDetectionResult {
	resultChan := make(chan *memory.PrepareForLeakDetectionResult)
	command := NewCommand(protocol.Socket, "Memory.prepareForLeakDetection", nil)
	result := &memory.PrepareForLeakDetectionResult{}

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
SetPressureNotificationsSuppressed enables/disables suppressing memory pressure
notifications in all processes.

https://chromedevtools.github.io/devtools-protocol/tot/Memory/#method-setPressureNotificationsSuppressed
*/
func (protocol *MemoryProtocol) SetPressureNotificationsSuppressed(
	params *memory.SetPressureNotificationsSuppressedParams,
) chan *memory.SetPressureNotificationsSuppressedResult {
	resultChan := make(chan *memory.SetPressureNotificationsSuppressedResult)
	command := NewCommand(protocol.Socket, "Memory.setPressureNotificationsSuppressed", params)
	result := &memory.SetPressureNotificationsSuppressedResult{}

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
SimulatePressureNotification simulates a memory pressure notification in all
processes.

https://chromedevtools.github.io/devtools-protocol/tot/Memory/#method-simulatePressureNotification
*/
func (protocol *MemoryProtocol) SimulatePressureNotification(
	params *memory.SimulatePressureNotificationParams,
) chan *memory.SimulatePressureNotificationResult {
	resultChan := make(chan *memory.SimulatePressureNotificationResult)
	command := NewCommand(protocol.Socket, "Memory.simulatePressureNotification", params)
	result := &memory.SimulatePressureNotificationResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

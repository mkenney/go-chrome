package socket

import (
	"encoding/json"

	chromeLog "github.com/mkenney/go-chrome/tot/cdtp/log"
)

/*
LogProtocol provides a namespace for the Chrome Log protocol methods. The Log
protocol provides access to log entries.

https://chromedevtools.github.io/devtools-protocol/tot/Log/
*/
type LogProtocol struct {
	Socket Socketer
}

/*
Clear clears the log.

https://chromedevtools.github.io/devtools-protocol/tot/Log/#method-clear
*/
func (protocol *LogProtocol) Clear() <-chan *chromeLog.ClearResult {
	resultChan := make(chan *chromeLog.ClearResult)
	command := NewCommand(protocol.Socket, "Log.clear", nil)
	result := &chromeLog.ClearResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
Disable disables log domain, prevents further log entries from being reported to
the client.

https://chromedevtools.github.io/devtools-protocol/tot/Log/#method-disable
*/
func (protocol *LogProtocol) Disable() <-chan *chromeLog.DisableResult {
	resultChan := make(chan *chromeLog.DisableResult)
	command := NewCommand(protocol.Socket, "Log.disable", nil)
	result := &chromeLog.DisableResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
Enable enables log domain, sends the entries collected so far to the client by
means of the `entryAdded` notification.

https://chromedevtools.github.io/devtools-protocol/tot/Log/#method-enable
*/
func (protocol *LogProtocol) Enable() <-chan *chromeLog.EnableResult {
	resultChan := make(chan *chromeLog.EnableResult)
	command := NewCommand(protocol.Socket, "Log.enable", nil)
	result := &chromeLog.EnableResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
StartViolationsReport starts violation reporting.

https://chromedevtools.github.io/devtools-protocol/tot/Log/#method-startViolationsReport
*/
func (protocol *LogProtocol) StartViolationsReport(
	params *chromeLog.StartViolationsReportParams,
) <-chan *chromeLog.StartViolationsReportResult {
	resultChan := make(chan *chromeLog.StartViolationsReportResult)
	command := NewCommand(protocol.Socket, "Log.startViolationsReport", params)
	result := &chromeLog.StartViolationsReportResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
StopViolationsReport stops violation reporting.

https://chromedevtools.github.io/devtools-protocol/tot/Log/#method-stopViolationsReport
*/
func (protocol *LogProtocol) StopViolationsReport() <-chan *chromeLog.StopViolationsReportResult {
	resultChan := make(chan *chromeLog.StopViolationsReportResult)
	command := NewCommand(protocol.Socket, "Log.stopViolationsReport", nil)
	result := &chromeLog.StopViolationsReportResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
OnEntryAdded adds a handler to the Log.entryAdded event. Log.entryAdded fires
when a new message is logged.

https://chromedevtools.github.io/devtools-protocol/tot/Log/#event-entryAdded
*/
func (protocol *LogProtocol) OnEntryAdded(
	callback func(event *chromeLog.EntryAddedEvent),
) {
	handler := NewEventHandler(
		"Log.entryAdded",
		func(response *Response) {
			event := &chromeLog.EntryAddedEvent{}
			json.Unmarshal([]byte(response.Result), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

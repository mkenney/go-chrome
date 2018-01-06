package socket

import (
	"encoding/json"

	console "github.com/mkenney/go-chrome/cdtp/console"
	log "github.com/sirupsen/logrus"
)

/*
ConsoleProtocol provides a namespace for the Chrome Console protocol methods.

https://chromedevtools.github.io/devtools-protocol/tot/Console/ DEPRECATED - use
Runtime or Log instead.
*/
type ConsoleProtocol struct {
	Socket Socketer
}

/*
ClearMessages does nothing.

https://chromedevtools.github.io/devtools-protocol/tot/Console/#method-clearMessages
*/
func (protocol *ConsoleProtocol) ClearMessages() chan *console.ClearMessagesResult {
	resultChan := make(chan *console.ClearMessagesResult)
	command := NewCommand(protocol.Socket, "Console.clearMessages", nil)
	result := &console.ClearMessagesResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
Disable disables console domain, prevents further console messages from being
reported to the client.

https://chromedevtools.github.io/devtools-protocol/tot/Console/#method-disable
*/
func (protocol *ConsoleProtocol) Disable() chan *console.DisableResult {
	resultChan := make(chan *console.DisableResult)
	command := NewCommand(protocol.Socket, "Console.disable", nil)
	result := &console.DisableResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
Enable enables console domain, sends the messages collected so far to the client
by means of the messageAdded notification.

https://chromedevtools.github.io/devtools-protocol/tot/Console/#method-enable
*/
func (protocol *ConsoleProtocol) Enable() chan *console.EnableResult {
	resultChan := make(chan *console.EnableResult)
	command := NewCommand(protocol.Socket, "Console.enable", nil)
	result := &console.EnableResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
OnMessageAdded adds a handler to the Console.messageAdded event.
Console.messageAdded fires whenever an active document stylesheet is removed.

https://chromedevtools.github.io/devtools-protocol/tot/Console/#event-messageAdded
*/
func (protocol *ConsoleProtocol) OnMessageAdded(
	callback func(event *console.MessageAddedEvent),
) {
	handler := NewEventHandler(
		"Console.messageAdded",
		func(response *Response) {
			event := &console.MessageAddedEvent{}
			if err := json.Unmarshal([]byte(response.Result), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

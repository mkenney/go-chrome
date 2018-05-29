package socket

import (
	"encoding/json"

	"github.com/mkenney/go-chrome/tot/cdtp/tethering"
)

/*
TetheringProtocol provides a namespace for the Chrome Tethering protocol
methods. The Tethering protocol defines methods and events for browser port
binding.

https://chromedevtools.github.io/devtools-protocol/tot/Tethering/
*/
type TetheringProtocol struct {
	Socket Socketer
}

/*
Bind requests browser port binding.

https://chromedevtools.github.io/devtools-protocol/tot/Tethering/#method-bind
*/
func (protocol *TetheringProtocol) Bind(
	params *tethering.BindParams,
) <-chan *tethering.BindResult {
	resultChan := make(chan *tethering.BindResult)
	command := NewCommand(protocol.Socket, "Tethering.bind", params)
	result := &tethering.BindResult{}

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
Unbind requests browser port unbinding.

https://chromedevtools.github.io/devtools-protocol/tot/Tethering/#method-unbind
*/
func (protocol *TetheringProtocol) Unbind(
	params *tethering.UnbindParams,
) <-chan *tethering.UnbindResult {
	resultChan := make(chan *tethering.UnbindResult)
	command := NewCommand(protocol.Socket, "Tethering.unbind", params)
	result := &tethering.UnbindResult{}

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
OnAccepted adds a handler to the Tethering.accepted event. Tethering.accepted
fires when a port was successfully bound and got a specified connection id.

https://chromedevtools.github.io/devtools-protocol/tot/Tethering/#event-accepted
*/
func (protocol *TetheringProtocol) OnAccepted(
	callback func(event *tethering.AcceptedEvent),
) {
	handler := NewEventHandler(
		"Tethering.accepted",
		func(response *Response) {
			event := &tethering.AcceptedEvent{}
			json.Unmarshal([]byte(response.Params), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

package socket

import (
	"encoding/json"

	tethering "github.com/mkenney/go-chrome/cdtp/tethering"
	log "github.com/sirupsen/logrus"
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
) error {
	command := NewCommand("Tethering.bind", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
Unbind requests browser port unbinding.

https://chromedevtools.github.io/devtools-protocol/tot/Tethering/#method-unbind
*/
func (protocol *TetheringProtocol) Unbind(
	params *tethering.UnbindParams,
) error {
	command := NewCommand("Tethering.unbind", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
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
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

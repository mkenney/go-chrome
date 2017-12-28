package protocol

import (
	"encoding/json"

	tethering "github.com/mkenney/go-chrome/protocol/tethering"
	sock "github.com/mkenney/go-chrome/socket"

	log "github.com/sirupsen/logrus"
)

/*
Tethering is a struct that provides a namespace for the Chrome Tethering protocol methods.

The Tethering protocol defines methods and events for browser port binding.

https://chromedevtools.github.io/devtools-protocol/tot/Tethering/
*/
type Tethering struct{}

/*
Bind requests browser port binding.

https://chromedevtools.github.io/devtools-protocol/tot/Tethering/#method-bind
*/
func (Tethering) Bind(
	socket sock.Socketer,
	params *tethering.BindParams,
) error {
	command := sock.NewCommand("Tethering.bind", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
Unbind requests browser port unbinding.

https://chromedevtools.github.io/devtools-protocol/tot/Tethering/#method-unbind
*/
func (Tethering) Unbind(
	socket sock.Socketer,
	params *tethering.UnbindParams,
) error {
	command := sock.NewCommand("Tethering.unbind", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
OnAccepted adds a handler to the Tethering.accepted event. Tethering.accepted fires when a port was
successfully bound and got a specified connection id.

https://chromedevtools.github.io/devtools-protocol/tot/Tethering/#event-accepted
*/
func (Tethering) OnAccepted(
	socket sock.Socketer,
	callback func(event *tethering.AcceptedEvent),
) {
	handler := sock.NewEventHandler(
		"Tethering.accepted",
		func(response *sock.Response) {
			event := &tethering.AcceptedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

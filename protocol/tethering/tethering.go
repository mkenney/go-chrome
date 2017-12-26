/*
Package tethering provides type definitions for use with the Chrome Tethering protocol

https://chromedevtools.github.io/devtools-protocol/tot/Tethering/
*/
package tethering

/*
BindParams represents Tethering.bind parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Tethering/#method-bind
*/
type BindParams struct {
	// Port number to bind.
	Port int `json:"port"`
}

/*
UnbindParams represents Tethering.unbind parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Tethering/#method-unbind
*/
type UnbindParams struct {
	// Port number to unbind.
	Port int `json:"port"`
}

/*
AcceptedEvent represents Overlay.accepted event data.

https://chromedevtools.github.io/devtools-protocol/tot/Tethering/#event-accepted
*/
type AcceptedEvent struct {
	// Port number that was successfully bound.
	Port int `json:"port"`

	// Connection ID to be used.
	ConnectionID string `json:"connectionId"`
}

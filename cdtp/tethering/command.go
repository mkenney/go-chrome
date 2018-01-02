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

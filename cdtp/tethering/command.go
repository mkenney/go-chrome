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
BindResult represents the result of calls to Tethering.bind.

https://chromedevtools.github.io/devtools-protocol/tot/Tethering/#method-bind
*/
type BindResult struct {
	CDTPError error `json:"-"`
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
UnbindResult represents the result of calls to Tethering.unbind.

https://chromedevtools.github.io/devtools-protocol/tot/Tethering/#method-unbind
*/
type UnbindResult struct {
	CDTPError error `json:"-"`
}

package protocol

import (
	browser "github.com/mkenney/go-chrome/protocol/browser"
	sock "github.com/mkenney/go-chrome/socket"
)

/*
Browser is a struct that provides a namespace for the Chrome Browser protocol methods. The Browser
protocol defines methods and events for browser management.

https://chromedevtools.github.io/devtools-protocol/tot/Browser/
*/
var Browser = _browser{}

type _browser struct{}

/*
Close closes the browser gracefully.

https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-close
*/
func (_browser) Close(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("Browser.close", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
GetVersion returns version information.

https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-getVersion
*/
func (_browser) GetVersion(
	socket sock.Socketer,
) (*browser.GetVersionResult, error) {
	command := sock.NewCommand("Browser.getVersion", nil)
	result := &browser.GetVersionResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
GetWindowBounds sets the position and/or size of the browser window. EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-getWindowBounds
*/
func (_browser) GetWindowBounds(
	socket sock.Socketer,
	params *browser.GetWindowBoundsParams,
) (*browser.GetWindowBoundsResult, error) {
	command := sock.NewCommand("Browser.getWindowBounds", params)
	result := &browser.GetWindowBoundsResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
GetWindowForTarget gets the browser window that contains the devtools target. EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-getWindowForTarget
*/
func (_browser) GetWindowForTarget(
	socket sock.Socketer,
	params *browser.GetWindowForTargetParams,
) (*browser.GetWindowForTargetResult, error) {
	command := sock.NewCommand("Browser.getWindowForTarget", params)
	result := &browser.GetWindowForTargetResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
SetWindowBounds sets the position and/or size of the browser window. EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-setWindowBounds
*/
func (_browser) SetWindowBounds(
	socket sock.Socketer,
	params *browser.SetWindowBoundsParams,
) (*browser.SetWindowBoundsResult, error) {
	command := sock.NewCommand("Browser.setWindowBounds", params)
	result := &browser.SetWindowBoundsResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

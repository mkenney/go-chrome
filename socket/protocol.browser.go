package socket

import (
	browser "github.com/mkenney/go-chrome/protocol/browser"
)

/*
BrowserProtocol provides a namespace for the Chrome Browser protocol methods. The Browser protocol
defines methods and events for browser management.

https://chromedevtools.github.io/devtools-protocol/tot/Browser/
*/
type BrowserProtocol struct {
	Socket Socketer
}

/*
Close closes the browser gracefully.

https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-close
*/
func (protocol *BrowserProtocol) Close() error {
	command := NewCommand("Browser.close", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
GetVersion returns version information.

https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-getVersion
*/
func (protocol *BrowserProtocol) GetVersion() (*browser.GetVersionResult, error) {
	command := NewCommand("Browser.getVersion", nil)
	result := &browser.GetVersionResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
GetWindowBounds sets the position and/or size of the browser window.

https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-getWindowBounds EXPERIMENTAL.
*/
func (protocol *BrowserProtocol) GetWindowBounds(
	params *browser.GetWindowBoundsParams,
) (*browser.GetWindowBoundsResult, error) {
	command := NewCommand("Browser.getWindowBounds", params)
	result := &browser.GetWindowBoundsResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
GetWindowForTarget gets the browser window that contains the devtools target.

https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-getWindowForTarget
EXPERIMENTAL.
*/
func (protocol *BrowserProtocol) GetWindowForTarget(
	params *browser.GetWindowForTargetParams,
) (*browser.GetWindowForTargetResult, error) {
	command := NewCommand("Browser.getWindowForTarget", params)
	result := &browser.GetWindowForTargetResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
SetWindowBounds sets the position and/or size of the browser window.

https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-setWindowBounds EXPERIMENTAL.
*/
func (protocol *BrowserProtocol) SetWindowBounds(
	params *browser.SetWindowBoundsParams,
) (*browser.SetWindowBoundsResult, error) {
	command := NewCommand("Browser.setWindowBounds", params)
	result := &browser.SetWindowBoundsResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

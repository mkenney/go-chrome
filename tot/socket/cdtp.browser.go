package socket

import (
	"encoding/json"

	browser "github.com/mkenney/go-chrome/tot/cdtp/browser"
)

/*
BrowserProtocol provides a namespace for the Chrome Browser protocol methods.
The Browser protocol defines methods and events for browser management.

https://chromedevtools.github.io/devtools-protocol/tot/Browser/
*/
type BrowserProtocol struct {
	Socket Socketer
}

/*
Close closes the browser gracefully.

https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-close
*/
func (protocol *BrowserProtocol) Close() <-chan *browser.CloseResult {
	resultChan := make(chan *browser.CloseResult)
	command := NewCommand(protocol.Socket, "Browser.close", nil)
	result := &browser.CloseResult{}

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
GetVersion returns version information.

https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-getVersion
*/
func (protocol *BrowserProtocol) GetVersion() <-chan *browser.GetVersionResult {
	resultChan := make(chan *browser.GetVersionResult)
	command := NewCommand(protocol.Socket, "Browser.getVersion", nil)
	result := &browser.GetVersionResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
GetWindowBounds sets the position and/or size of the browser window.

https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-getWindowBounds EXPERIMENTAL.
*/
func (protocol *BrowserProtocol) GetWindowBounds(
	params *browser.GetWindowBoundsParams,
) <-chan *browser.GetWindowBoundsResult {
	resultChan := make(chan *browser.GetWindowBoundsResult)
	command := NewCommand(protocol.Socket, "Browser.getWindowBounds", params)
	result := &browser.GetWindowBoundsResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
GetWindowForTarget gets the browser window that contains the devtools target.

https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-getWindowForTarget
EXPERIMENTAL.
*/
func (protocol *BrowserProtocol) GetWindowForTarget(
	params *browser.GetWindowForTargetParams,
) <-chan *browser.GetWindowForTargetResult {
	resultChan := make(chan *browser.GetWindowForTargetResult)
	command := NewCommand(protocol.Socket, "Browser.getWindowForTarget", params)
	result := &browser.GetWindowForTargetResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
SetWindowBounds sets the position and/or size of the browser window.

https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-setWindowBounds EXPERIMENTAL.
*/
func (protocol *BrowserProtocol) SetWindowBounds(
	params *browser.SetWindowBoundsParams,
) <-chan *browser.SetWindowBoundsResult {
	resultChan := make(chan *browser.SetWindowBoundsResult)
	command := NewCommand(protocol.Socket, "Browser.setWindowBounds", params)
	result := &browser.SetWindowBoundsResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

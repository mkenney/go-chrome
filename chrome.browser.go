package chrome

import (
	"encoding/json"

	browser "github.com/mkenney/go-chrome/browser"
	"github.com/mkenney/go-chrome/protocol"
)

/*
Browser - https://chromedevtools.github.io/devtools-protocol/tot/Browser/
Defines methods and events for browser management.
*/
type Browser struct{}

/*
Close closes the browser gracefully.

https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-close
*/
func (Browser) Close(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "Browser.close",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
GetVersion returns version information.

https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-getVersion
*/
func (Browser) GetVersion(
	socket *Socket,
) (browser.GetVersionResult, error) {
	command := &protocol.Command{
		Method: "Browser.getVersion",
	}
	result := browser.GetVersionResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
GetWindowBounds sets the position and/or size of the browser window. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-getWindowBounds
*/
func (Browser) GetWindowBounds(
	socket *Socket,
	params *browser.GetWindowBoundsParams,
) (browser.GetWindowBoundsResult, error) {
	command := &protocol.Command{
		Method: "Browser.getWindowBounds",
		Params: params,
	}
	result := browser.GetWindowBoundsResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
GetWindowForTarget gets the browser window that contains the devtools target. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-getWindowForTarget
*/
func (Browser) GetWindowForTarget(
	socket *Socket,
	params *browser.GetWindowForTargetParams,
) (browser.GetWindowForTargetResult, error) {
	command := &protocol.Command{
		Method: "Browser.getWindowForTarget",
		Params: params,
	}
	result := browser.GetWindowForTargetResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
SetWindowBounds sets the position and/or size of the browser window. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-setWindowBounds
*/
func (Browser) SetWindowBounds(
	socket *Socket,
	params *browser.SetWindowBoundsParams,
) (browser.SetWindowBoundsResult, error) {
	command := &protocol.Command{
		Method: "Browser.setWindowBounds",
		Params: params,
	}
	result := browser.SetWindowBoundsResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

package chrome

import (
	browser "app/chrome/browser"
	"app/chrome/protocol"
)

/*
Browser - https://chromedevtools.github.io/devtools-protocol/tot/Browser/
Defines methods and events for browser management.
*/
type Browser struct{}

/*
Close closes the browser gracefully.
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
*/
func (Browser) GetVersion(
	socket *Socket,
) (browser.GetVersionResult, error) {
	command := &protocol.Command{
		Method: "Browser.getVersion",
	}
	socket.SendCommand(command)
	return command.Result.(browser.GetVersionResult), command.Err
}

/*
GetWindowBounds sets the position and/or size of the browser window. EXPERIMENTAL
*/
func (Browser) GetWindowBounds(
	socket *Socket,
	params *browser.GetWindowBoundsParams,
) (browser.GetWindowBoundsResult, error) {
	command := &protocol.Command{
		Method: "Browser.getWindowBounds",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Result.(browser.GetWindowBoundsResult), command.Err
}

/*
GetWindowForTarget gets the browser window that contains the devtools target. EXPERIMENTAL
*/
func (Browser) GetWindowForTarget(
	socket *Socket,
	params *browser.GetWindowForTargetParams,
) (browser.GetWindowForTargetResult, error) {
	command := &protocol.Command{
		Method: "Browser.getWindowForTarget",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Result.(browser.GetWindowForTargetResult), command.Err
}

/*
SetWindowBounds sets the position and/or size of the browser window. EXPERIMENTAL
*/
func (Browser) SetWindowBounds(
	socket *Socket,
	params *browser.SetWindowBoundsParams,
) (browser.SetWindowBoundsResult, error) {
	command := &protocol.Command{
		Method: "Browser.setWindowBounds",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Result.(browser.SetWindowBoundsResult), command.Err
}

package chrome

import "app/chrome/protocol"

/*
Browser - https://chromedevtools.github.io/devtools-protocol/tot/Browser/
Defines methods and events for browser management.
*/
type Browser struct{}

/*
Close closes the browser gracefully.
*/
func (Browser) Close(socket *Socket) error {
	command := &protocol.Command{
		method: "Browser.close",
		params: nil,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
GetWindowForTarget gets the browser window that contains the devtools target. EXPERIMENTAL
*/
func (Browser) GetWindowForTarget(socket *Socket, params *browser.GetWindowForTargetParams) error {
	command := &protocol.Command{
		method: "Browser.getWindowForTarget",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
GetVersion returns version information.
*/
func (Browser) GetVersion(socket *Socket) error {
	command := &protocol.Command{
		method: "Browser.getVersion",
		params: nil,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetWindowBounds sets the position and/or size of the browser window. EXPERIMENTAL
*/
func (Browser) SetWindowBounds(socket *Socket, params *browser.SetWindowBoundsParams) error {
	command := &protocol.Command{
		method: "Browser.setWindowBounds",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
GetWindowBounds sets the position and/or size of the browser window. EXPERIMENTAL
*/
func (Browser) GetWindowBounds(socket *Socket, params *browser.GetWindowBoundsParams) error {
	command := &protocol.Command{
		method: "Browser.getWindowBounds",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

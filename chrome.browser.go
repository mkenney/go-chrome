package chrome

import (
	"app/chrome/protocol"
	browser "app/chrome/protocol/browser"
)

type Browser struct{}

/*
Close closes the browser gracefully.
*/
func (Browser) Close(socket *Socket) error {
	command := new(protocol.Command)
	command.method = "Browser.close"
	socket.SendCommand(command)
	return command.Err
}

/*
GetWindowForTarget gets the browser window that contains the devtools target. EXPERIMENTAL
*/
func (Browser) GetWindowForTarget(socket *Socket, params *browser.GetWindowForTargetParams) error {
	command := new(protocol.Command)
	command.method = "Browser.getWindowForTarget"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
GetVersion returns version information.
*/
func (Browser) GetVersion(socket *Socket) error {
	command := new(protocol.Command)
	command.method = "Browser.getVersion"
	socket.SendCommand(command)
	return command.Err
}

/*
SetWindowBounds sets the position and/or size of the browser window. EXPERIMENTAL
*/
func (Browser) SetWindowBounds(socket *Socket, params *browser.SetWindowBoundsParams) error {
	command := new(protocol.Command)
	command.method = "Browser.setWindowBounds"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
GetWindowBounds sets the position and/or size of the browser window. EXPERIMENTAL
*/
func (Browser) GetWindowBounds(socket *Socket, params *browser.GetWindowBoundsParams) error {
	command := new(protocol.Command)
	command.method = "Browser.getWindowBounds"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

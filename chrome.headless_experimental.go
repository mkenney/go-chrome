package chrome

import (
	headlessexp "app/chrome/headless_experimental"
	"app/chrome/protocol"
)

/*
HeadlessExperimental provides experimental commands only supported in headless mode. EXPERIMENTAL
*/
type HeadlessExperimental struct{}

/*
Enable enables headless events for the target.
*/
func (HeadlessExperimental) Enable(socket *Socket) error {
	command := new(protocol.Command)
	command.method = "HeadlessExperimental.enable"
	socket.SendCommand(command)
	return command.Err
}

/*
Disable disables headless events for the target.
*/
func (HeadlessExperimental) Disable(socket *Socket) error {
	command := new(protocol.Command)
	command.method = "HeadlessExperimental.disable"
	socket.SendCommand(command)
	return command.Err
}

/*
BeginFrame sends a BeginFrame to the target and returns when the frame was completed. Optionally
captures a screenshot from the resulting frame. Requires that the target was created with enabled
BeginFrameControl.
*/
func (HeadlessExperimental) BeginFrame(socket *Socket, params *headlessexp.BeginFrameParams) error {
	command := new(protocol.Command)
	command.method = "HeadlessExperimental.beginFrame"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
OnNeedsBeginFramesChanged adds a handler to the HeadlessExperimental.needsBeginFramesChanged event.
HeadlessExperimental.needsBeginFramesChanged fires when the target starts or stops needing
BeginFrames.
*/
func (HeadlessExperimental) OnNeedsBeginFramesChanged(socket *Socket, callback func(event *headlessexp.NeedsBeginFramesChangedEvent)) error {
	handler := protocol.NewEventHandler(
		"HeadlessExperimental.needsBeginFramesChanged",
		func(name string, params []byte) {
			event := new(headlessexp.NeedsBeginFramesChangedEvent)
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		}
	)
	socket.AddEventHandler(handler)
	return command.Err
}

/*
OnMainFrameReadyForScreenshots adds a handler to the
HeadlessExperimental.mainFrameReadyForScreenshots event.
HeadlessExperimental.mainFrameReadyForScreenshots fires when the main frame has first submitted a
frame to the browser. May only be fired while a BeginFrame is in flight. Before this event,
screenshotting requests may fail.
*/
func (HeadlessExperimental) OnMainFrameReadyForScreenshots(socket *Socket, callback func(event *headlessexp.MainFrameReadyForScreenshotsEvent)) error {
	handler := protocol.NewEventHandler(
		"HeadlessExperimental.mainFrameReadyForScreenshots",
		func(name string, params []byte) {
			event := new(headlessexp.MainFrameReadyForScreenshotsEvent)
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		}
	)
	socket.AddEventHandler(handler)
	return command.Err
}

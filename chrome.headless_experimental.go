package chrome

import (
	"encoding/json"

	headlessExperimental "github.com/mkenney/go-chrome/headless_experimental"
	"github.com/mkenney/go-chrome/protocol"

	log "github.com/sirupsen/logrus"
)

/*
HeadlessExperimental - https://chromedevtools.github.io/devtools-protocol/tot/HeadlessExperimental/
Provides experimental commands only supported in headless mode. EXPERIMENTAL
*/
type HeadlessExperimental struct{}

/*
BeginFrame sends a BeginFrame to the target and returns when the frame was completed. Optionally
captures a screenshot from the resulting frame. Requires that the target was created with enabled
BeginFrameControl.

https://chromedevtools.github.io/devtools-protocol/tot/HeadlessExperimental/#method-beginFrame
*/
func (HeadlessExperimental) BeginFrame(
	socket *Socket,
	params *headlessExperimental.BeginFrameParams,
) (headlessExperimental.BeginFrameResult, error) {
	command := &protocol.Command{
		Method: "HeadlessExperimental.beginFrame",
		Params: params,
	}
	result := headlessExperimental.BeginFrameResult{}
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
Disable disables headless events for the target.

https://chromedevtools.github.io/devtools-protocol/tot/HeadlessExperimental/#method-disable
*/
func (HeadlessExperimental) Disable(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "HeadlessExperimental.disable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Enable enables headless events for the target.

https://chromedevtools.github.io/devtools-protocol/tot/HeadlessExperimental/#method-enable
*/
func (HeadlessExperimental) Enable(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "HeadlessExperimental.enable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
OnMainFrameReadyForScreenshots adds a handler to the
HeadlessExperimental.mainFrameReadyForScreenshots event.
HeadlessExperimental.mainFrameReadyForScreenshots fires when the main frame has first submitted a
frame to the browser. May only be fired while a BeginFrame is in flight. Before this event,
screenshotting requests may fail.

https://chromedevtools.github.io/devtools-protocol/tot/HeadlessExperimental/#event-mainFrameReadyForScreenshots
*/
func (HeadlessExperimental) OnMainFrameReadyForScreenshots(
	socket *Socket,
	callback func(event *headlessExperimental.MainFrameReadyForScreenshotsEvent),
) {
	handler := protocol.NewEventHandler(
		"HeadlessExperimental.mainFrameReadyForScreenshots",
		func(name string, params []byte) {
			event := &headlessExperimental.MainFrameReadyForScreenshotsEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnNeedsBeginFramesChanged adds a handler to the HeadlessExperimental.needsBeginFramesChanged event.
HeadlessExperimental.needsBeginFramesChanged fires when the target starts or stops needing BeginFrames.

https://chromedevtools.github.io/devtools-protocol/tot/HeadlessExperimental/#event-needsBeginFramesChanged
*/
func (HeadlessExperimental) OnNeedsBeginFramesChanged(
	socket *Socket,
	callback func(event *headlessExperimental.NeedsBeginFramesChangedEvent),
) {
	handler := protocol.NewEventHandler(
		"HeadlessExperimental.needsBeginFramesChanged",
		func(name string, params []byte) {
			event := &headlessExperimental.NeedsBeginFramesChangedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

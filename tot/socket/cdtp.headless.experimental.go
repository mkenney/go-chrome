package socket

import (
	"encoding/json"

	"github.com/mkenney/go-chrome/tot/cdtp/headless/experimental"
)

/*
HeadlessExperimentalProtocol provides a namespace for the Chrome
HeadlessExperimental protocol methods. The HeadlessExperimental protocol
provides experimental commands only supported in headless mode.

https://chromedevtools.github.io/devtools-protocol/tot/HeadlessExperimental/ EXPERIMENTAL.
*/
type HeadlessExperimentalProtocol struct {
	Socket Socketer
}

/*
BeginFrame sends a BeginFrame to the target and returns when the frame was
completed. Optionally captures a screenshot from the resulting frame. Requires
that the target was created with enabled BeginFrameControl.

https://chromedevtools.github.io/devtools-protocol/tot/HeadlessExperimental/#method-beginFrame
*/
func (protocol *HeadlessExperimentalProtocol) BeginFrame(
	params *experimental.BeginFrameParams,
) <-chan *experimental.BeginFrameResult {
	resultChan := make(chan *experimental.BeginFrameResult)
	command := NewCommand(protocol.Socket, "HeadlessExperimental.beginFrame", params)
	result := &experimental.BeginFrameResult{}

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
Disable disables headless events for the target.

https://chromedevtools.github.io/devtools-protocol/tot/HeadlessExperimental/#method-disable
*/
func (protocol *HeadlessExperimentalProtocol) Disable() <-chan *experimental.DisableResult {
	resultChan := make(chan *experimental.DisableResult)
	command := NewCommand(protocol.Socket, "HeadlessExperimental.disable", nil)
	result := &experimental.DisableResult{}

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
Enable enables headless events for the target.

https://chromedevtools.github.io/devtools-protocol/tot/HeadlessExperimental/#method-enable
*/
func (protocol *HeadlessExperimentalProtocol) Enable() <-chan *experimental.EnableResult {
	resultChan := make(chan *experimental.EnableResult)
	command := NewCommand(protocol.Socket, "HeadlessExperimental.enable", nil)
	result := &experimental.EnableResult{}

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
OnMainFrameReadyForScreenshots adds a handler to the HeadlessExperimental.mainFrameReadyForScreenshots
event. HeadlessExperimental.mainFrameReadyForScreenshots fires when the main
frame has first submitted a frame to the browser. May only be fired while a
BeginFrame is in flight. Before this event, screenshotting requests may fail.

https://chromedevtools.github.io/devtools-protocol/tot/HeadlessExperimental/#event-mainFrameReadyForScreenshots
*/
func (protocol *HeadlessExperimentalProtocol) OnMainFrameReadyForScreenshots(
	callback func(event *experimental.MainFrameReadyForScreenshotsEvent),
) {
	handler := NewEventHandler(
		"HeadlessExperimental.mainFrameReadyForScreenshots",
		func(response *Response) {
			event := &experimental.MainFrameReadyForScreenshotsEvent{}
			json.Unmarshal([]byte(response.Params), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnNeedsBeginFramesChanged adds a handler to the HeadlessExperimental.needsBeginFramesChanged
event. HeadlessExperimental.needsBeginFramesChanged fires when the target starts
or stops needing BeginFrames.

https://chromedevtools.github.io/devtools-protocol/tot/HeadlessExperimental/#event-needsBeginFramesChanged
*/
func (protocol *HeadlessExperimentalProtocol) OnNeedsBeginFramesChanged(
	callback func(event *experimental.NeedsBeginFramesChangedEvent),
) {
	handler := NewEventHandler(
		"HeadlessExperimental.needsBeginFramesChanged",
		func(response *Response) {
			event := &experimental.NeedsBeginFramesChangedEvent{}
			json.Unmarshal([]byte(response.Params), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

package chrome

import (
	"app/chrome/protocol"
	"encoding/json"

	log "github.com/Sirupsen/logrus"
)

/*
Emulation - https://chromedevtools.github.io/devtools-protocol/tot/Emulation/
Emulates different environments for the page.
*/
type Emulation struct{}

/*
SetDeviceMetricsOverride overrides the values of device screen dimensions (window.screen.width,
window.screen.height, window.innerWidth, window.innerHeight, and "device-width"/"device-height"
related CSS media query results).
*/
func (Emulation) SetDeviceMetricsOverride(socket *Socket, params *emulation.SetDeviceMetricsOverrideParams) error {
	command := &protocol.Command{
		method: "Emulation.setDeviceMetricsOverride",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
ClearDeviceMetricsOverride clears the overriden device metrics.
*/
func (Emulation) ClearDeviceMetricsOverride(socket *Socket) error {
	command := &protocol.Command{
		method: "Emulation.clearDeviceMetricsOverride",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
ResetPageScaleFactor requests that page scale factor is reset to initial values. EXPERIMENTAL
*/
func (Emulation) ResetPageScaleFactor(socket *Socket) error {
	command := &protocol.Command{
		method: "Emulation.resetPageScaleFactor",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetPageScaleFactor sets a specified page scale factor. EXPERIMENTAL
*/
func (Emulation) SetPageScaleFactor(socket *Socket, params *emulation.SetPageScaleFactorParams) error {
	command := &protocol.Command{
		method: "Emulation.setPageScaleFactor",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetVisibleSize resizes the frame/viewport of the page. Note that this does not affect the frame's
container (e.g. browser window). Can be used to produce screenshots of the specified size. Not
supported on Android. EXPERIMENTAL DEPRECATED
*/
func (Emulation) SetVisibleSize(socket *Socket, params *emulation.SetVisibleSizeParams) error {
	command := &protocol.Command{
		method: "Emulation.setVisibleSize",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetScriptExecutionDisabled switches script execution in the page.
*/
func (Emulation) SetScriptExecutionDisabled(socket *Socket, params *emulation.SetScriptExecutionDisabledParams) error {
	command := &protocol.Command{
		method: "Emulation.setScriptExecutionDisabled",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetGeolocationOverride overrides the Geolocation Position or Error. Omitting any of the parameters
emulates position unavailable.
*/
func (Emulation) SetGeolocationOverride(socket *Socket, params *emulation.SetGeolocationOverrideParams) error {
	command := &protocol.Command{
		method: "Emulation.setGeolocationOverride",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
ClearGeolocationOverride clears the overriden Geolocation Position and Error.
*/
func (Emulation) ClearGeolocationOverride(socket *Socket) error {
	command := &protocol.Command{
		method: "Emulation.clearGeolocationOverride",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetTouchEmulationEnabled enables touch on platforms which do not support it.
*/
func (Emulation) SetTouchEmulationEnabled(socket *Socket, params *emulation.SetTouchEmulationEnabledParams) error {
	command := &protocol.Command{
		method: "Emulation.setTouchEmulationEnabled",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetEmitTouchEventsForMouse enables touch events using a mouse. EXPERIMENTAL
*/
func (Emulation) SetEmitTouchEventsForMouse(socket *Socket, params *emulation.SetEmitTouchEventsForMouseParams) error {
	command := &protocol.Command{
		method: "Emulation.setEmitTouchEventsForMouse",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetEmulatedMedia emulates the given media for CSS media queries.
*/
func (Emulation) SetEmulatedMedia(socket *Socket, params *emulation.SetEmulatedMediaParams) error {
	command := &protocol.Command{
		method: "Emulation.setEmulatedMedia",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetCPUThrottlingRate enables CPU throttling to emulate slow CPUs. EXPERIMENTAL
*/
func (Emulation) SetCPUThrottlingRate(socket *Socket, params *emulation.SetCPUThrottlingRateParams) error {
	command := &protocol.Command{
		method: "Emulation.setCPUThrottlingRate",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
CanEmulate tells whether emulation is supported.
*/
func (Emulation) CanEmulate(socket *Socket) error {
	command := &protocol.Command{
		method: "Emulation.canEmulate",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetVirtualTimePolicy turns on virtual time for all frames (replacing real-time with a synthetic time
source) and sets the current virtual time policy. Note this supersedes any previous time budget.
EXPERIMENTAL
*/
func (Emulation) SetVirtualTimePolicy(socket *Socket, params *emulation.SetVirtualTimePolicyParams) error {
	command := &protocol.Command{
		method: "Emulation.setVirtualTimePolicy",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetNavigatorOverrides overrides value returned by the javascript navigator object. EXPERIMENTAL
*/
func (Emulation) SetNavigatorOverrides(socket *Socket, params *emulation.SetNavigatorOverridesParams) error {
	command := &protocol.Command{
		method: "Emulation.setNavigatorOverrides",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetDefaultBackgroundColorOverride sets or clears an override of the default background color of the
frame. This override is used if the content does not specify one.
*/
func (Emulation) SetDefaultBackgroundColorOverride(socket *Socket, params *emulation.SetDefaultBackgroundColorOverrideParams) error {
	command := &protocol.Command{
		method: "Emulation.setDefaultBackgroundColorOverride",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
OnVirtualTimeBudgetExpired adds a handler to the Emulation.virtualTimeBudgetExpired event.
Emulation.virtualTimeBudgetExpired fires after the virtual time budget for the current
VirtualTimePolicy has run out. EXPERIMENTAL
*/
func (Emulation) OnVirtualTimeBudgetExpired(socket *Socket, callback func(event *emulation.VirtualTimeBudgetExpiredEvent)) {
	handler := protocol.NewEventHandler(
		"Emulation.virtualTimeBudgetExpired",
		func(name string, params []byte) {
			event := &emulation.ResumedEvent{}
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
OnVirtualTimeAdvanced adds a handler to the Emulation.virtualTimeAdvanced event.
Emulation.virtualTimeAdvanced fires after the virtual time has advanced. EXPERIMENTAL
*/
func (Emulation) OnVirtualTimeAdvanced(socket *Socket, callback func(event *emulation.VirtualTimeAdvancedEvent)) {
	handler := protocol.NewEventHandler(
		"Emulation.virtualTimeAdvanced",
		func(name string, params []byte) {
			event := &emulation.VirtualTimeAdvancedEvent{}
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
OnVirtualTimePaused adds a handler to the Emulation.virtualTimePaused event.
Emulation.virtualTimePaused fires after the virtual time has paused. EXPERIMENTAL
*/
func (Emulation) OnVirtualTimePaused(socket *Socket, callback func(event *emulation.VirtualTimePausedEvent)) {
	handler := protocol.NewEventHandler(
		"Emulation.virtualTimePaused",
		func(name string, params []byte) {
			event := &emulation.VirtualTimePausedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

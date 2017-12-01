package chrome

import (
	emulation "app/chrome/emulation"
	"app/chrome/protocol"
)

/*
Emulation emulates different environments for the page.
*/
type Emulation struct{}

/*
SetDeviceMetricsOverride overrides the values of device screen dimensions (window.screen.width,
window.screen.height, window.innerWidth, window.innerHeight, and "device-width"/"device-height"
related CSS media query results).
*/
func (Emulation) SetDeviceMetricsOverride(socket *Socket, params *emulation.SetDeviceMetricsOverrideParams) error {
	command := new(protocol.Command)
	command.method = "Emulation.setDeviceMetricsOverride"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
ClearDeviceMetricsOverride clears the overriden device metrics.
*/
func (Emulation) ClearDeviceMetricsOverride(socket *Socket) error {
	command := new(protocol.Command)
	command.method = "Emulation.clearDeviceMetricsOverride"
	socket.SendCommand(command)
	return command.Err
}

/*
ResetPageScaleFactor requests that page scale factor is reset to initial values. EXPERIMENTAL
*/
func (Emulation) ResetPageScaleFactor(socket *Socket) error {
	command := new(protocol.Command)
	command.method = "Emulation.resetPageScaleFactor"
	socket.SendCommand(command)
	return command.Err
}

/*
SetPageScaleFactor sets a specified page scale factor. EXPERIMENTAL
*/
func (Emulation) SetPageScaleFactor(socket *Socket, params *emulation.SetPageScaleFactorParams) error {
	command := new(protocol.Command)
	command.method = "Emulation.setPageScaleFactor"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
SetVisibleSize resizes the frame/viewport of the page. Note that this does not affect the frame's
container (e.g. browser window). Can be used to produce screenshots of the specified size. Not
supported on Android. EXPERIMENTAL DEPRECATED
*/
func (Emulation) SetVisibleSize(socket *Socket, params *emulation.SetVisibleSizeParams) error {
	command := new(protocol.Command)
	command.method = "Emulation.setVisibleSize"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
SetScriptExecutionDisabled switches script execution in the page.
*/
func (Emulation) SetScriptExecutionDisabled(socket *Socket, params *emulation.SetScriptExecutionDisabledParams) error {
	command := new(protocol.Command)
	command.method = "Emulation.setScriptExecutionDisabled"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
SetGeolocationOverride overrides the Geolocation Position or Error. Omitting any of the parameters
emulates position unavailable.
*/
func (Emulation) SetGeolocationOverride(socket *Socket, params *emulation.SetGeolocationOverrideParams) error {
	command := new(protocol.Command)
	command.method = "Emulation.setGeolocationOverride"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
ClearGeolocationOverride clears the overriden Geolocation Position and Error.
*/
func (Emulation) ClearGeolocationOverride(socket *Socket) error {
	command := new(protocol.Command)
	command.method = "Emulation.clearGeolocationOverride"
	socket.SendCommand(command)
	return command.Err
}

/*
SetTouchEmulationEnabled enables touch on platforms which do not support it.
*/
func (Emulation) SetTouchEmulationEnabled(socket *Socket, params *emulation.SetTouchEmulationEnabledParams) error {
	command := new(protocol.Command)
	command.method = "Emulation.setTouchEmulationEnabled"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
SetEmitTouchEventsForMouse enables touch events using a mouse. EXPERIMENTAL
*/
func (Emulation) SetEmitTouchEventsForMouse(socket *Socket, params *emulation.SetEmitTouchEventsForMouseParams) error {
	command := new(protocol.Command)
	command.method = "Emulation.setEmitTouchEventsForMouse"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
SetEmulatedMedia emulates the given media for CSS media queries.
*/
func (Emulation) SetEmulatedMedia(socket *Socket, params *emulation.SetEmulatedMediaParams) error {
	command := new(protocol.Command)
	command.method = "Emulation.setEmulatedMedia"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
SetCPUThrottlingRate enables CPU throttling to emulate slow CPUs. EXPERIMENTAL
*/
func (Emulation) SetCPUThrottlingRate(socket *Socket, params *emulation.SetCPUThrottlingRateParams) error {
	command := new(protocol.Command)
	command.method = "Emulation.setCPUThrottlingRate"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
CanEmulate tells whether emulation is supported.
*/
func (Emulation) CanEmulate(socket *Socket) error {
	command := new(protocol.Command)
	command.method = "Emulation.canEmulate"
	socket.SendCommand(command)
	return command.Err
}

/*
SetVirtualTimePolicy turns on virtual time for all frames (replacing real-time with a synthetic time
source) and sets the current virtual time policy. Note this supersedes any previous time budget.
EXPERIMENTAL
*/
func (Emulation) SetVirtualTimePolicy(socket *Socket, params *emulation.SetVirtualTimePolicyParams) error {
	command := new(protocol.Command)
	command.method = "Emulation.setVirtualTimePolicy"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
SetNavigatorOverrides overrides value returned by the javascript navigator object. EXPERIMENTAL
*/
func (Emulation) SetNavigatorOverrides(socket *Socket, params *emulation.SetNavigatorOverridesParams) error {
	command := new(protocol.Command)
	command.method = "Emulation.setNavigatorOverrides"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
SetDefaultBackgroundColorOverride sets or clears an override of the default background color of the
frame. This override is used if the content does not specify one.
*/
func (Emulation) SetDefaultBackgroundColorOverride(socket *Socket, params *emulation.SetDefaultBackgroundColorOverrideParams) error {
	command := new(protocol.Command)
	command.method = "Emulation.setDefaultBackgroundColorOverride"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
OnVirtualTimeBudgetExpired adds a handler to the Emulation.virtualTimeBudgetExpired event.
Emulation.virtualTimeBudgetExpired fires after the virtual time budget for the current
VirtualTimePolicy has run out. EXPERIMENTAL
*/
func (Emulation) OnVirtualTimeBudgetExpired(socket *Socket, callback func(event *emulation.VirtualTimeBudgetExpiredEvent)) error {
	handler := protocol.NewEventHandler(
		"Emulation.virtualTimeBudgetExpired",
		func(name string, params []byte) {
			event := new(emulation.ResumedEvent)
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
OnVirtualTimeAdvanced adds a handler to the Emulation.virtualTimeAdvanced event.
Emulation.virtualTimeAdvanced fires after the virtual time has advanced. EXPERIMENTAL
*/
func (Emulation) OnVirtualTimeAdvanced(socket *Socket, callback func(event *emulation.VirtualTimeAdvancedEvent)) error {
	handler := protocol.NewEventHandler(
		"Emulation.virtualTimeAdvanced",
		func(name string, params []byte) {
			event := new(emulation.VirtualTimeAdvancedEvent)
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
OnVirtualTimePaused adds a handler to the Emulation.virtualTimePaused event.
Emulation.virtualTimePaused fires after the virtual time has paused. EXPERIMENTAL
*/
func (Emulation) OnVirtualTimePaused(socket *Socket, callback func(event *emulation.VirtualTimePausedEvent)) error {
	handler := protocol.NewEventHandler(
		"Emulation.virtualTimePaused",
		func(name string, params []byte) {
			event := new(emulation.VirtualTimePausedEvent)
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

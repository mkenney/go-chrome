package chrome

import (
	"encoding/json"

	emulation "github.com/mkenney/go-chrome/emulation"
	"github.com/mkenney/go-chrome/protocol"

	log "github.com/sirupsen/logrus"
)

/*
Emulation - https://chromedevtools.github.io/devtools-protocol/tot/Emulation/
Emulates different environments for the page.
*/
type Emulation struct{}

/*
CanEmulate tells whether emulation is supported.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-canEmulate
*/
func (Emulation) CanEmulate(
	socket *Socket,
) (emulation.CanEmulateResult, error) {
	command := &protocol.Command{
		Method: "Emulation.canEmulate",
	}
	result := emulation.CanEmulateResult{}
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
ClearDeviceMetricsOverride clears the overridden device metrics.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-clearDeviceMetricsOverride
*/
func (Emulation) ClearDeviceMetricsOverride(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "Emulation.clearDeviceMetricsOverride",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
ClearGeolocationOverride clears the overridden Geolocation Position and Error.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-clearGeolocationOverride
*/
func (Emulation) ClearGeolocationOverride(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "Emulation.clearGeolocationOverride",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
ResetPageScaleFactor requests that page scale factor is reset to initial values. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-resetPageScaleFactor
*/
func (Emulation) ResetPageScaleFactor(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "Emulation.resetPageScaleFactor",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetCPUThrottlingRate enables CPU throttling to emulate slow CPUs. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setCPUThrottlingRate
*/
func (Emulation) SetCPUThrottlingRate(
	socket *Socket,
	params *emulation.SetCPUThrottlingRateParams,
) error {
	command := &protocol.Command{
		Method: "Emulation.setCPUThrottlingRate",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetDefaultBackgroundColorOverride sets or clears an override of the default background color of the
frame. This override is used if the content does not specify one.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setDefaultBackgroundColorOverride
*/
func (Emulation) SetDefaultBackgroundColorOverride(
	socket *Socket,
	params *emulation.SetDefaultBackgroundColorOverrideParams,
) error {
	command := &protocol.Command{
		Method: "Emulation.setDefaultBackgroundColorOverride",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetDeviceMetricsOverride overrides the values of device screen dimensions (window.screen.width,
window.screen.height, window.innerWidth, window.innerHeight, and "device-width"/"device-height"
related CSS media query results).

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setDeviceMetricsOverride
*/
func (Emulation) SetDeviceMetricsOverride(
	socket *Socket,
	params *emulation.SetDeviceMetricsOverrideParams,
) error {
	command := &protocol.Command{
		Method: "Emulation.setDeviceMetricsOverride",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetEmitTouchEventsForMouse enables touch events using a mouse. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setEmitTouchEventsForMouse
*/
func (Emulation) SetEmitTouchEventsForMouse(
	socket *Socket,
	params *emulation.SetEmitTouchEventsForMouseParams,
) error {
	command := &protocol.Command{
		Method: "Emulation.setEmitTouchEventsForMouse",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetEmulatedMedia emulates the given media for CSS media queries.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setEmulatedMedia
*/
func (Emulation) SetEmulatedMedia(
	socket *Socket,
	params *emulation.SetEmulatedMediaParams,
) error {
	command := &protocol.Command{
		Method: "Emulation.setEmulatedMedia",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetGeolocationOverride overrides the Geolocation Position or Error. Omitting any of the parameters
emulates position unavailable.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setGeolocationOverride
*/
func (Emulation) SetGeolocationOverride(
	socket *Socket,
	params *emulation.SetGeolocationOverrideParams,
) error {
	command := &protocol.Command{
		Method: "Emulation.setGeolocationOverride",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetNavigatorOverrides overrides value returned by the javascript navigator object. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setNavigatorOverrides
*/
func (Emulation) SetNavigatorOverrides(
	socket *Socket,
	params *emulation.SetNavigatorOverridesParams,
) error {
	command := &protocol.Command{
		Method: "Emulation.setNavigatorOverrides",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetPageScaleFactor sets a specified page scale factor. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setPageScaleFactor
*/
func (Emulation) SetPageScaleFactor(
	socket *Socket,
	params *emulation.SetPageScaleFactorParams,
) error {
	command := &protocol.Command{
		Method: "Emulation.setPageScaleFactor",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetScriptExecutionDisabled switches script execution in the page.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setScriptExecutionDisabled
*/
func (Emulation) SetScriptExecutionDisabled(
	socket *Socket,
	params *emulation.SetScriptExecutionDisabledParams,
) error {
	command := &protocol.Command{
		Method: "Emulation.setScriptExecutionDisabled",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetTouchEmulationEnabled enables touch on platforms which do not support it.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setTouchEmulationEnabled
*/
func (Emulation) SetTouchEmulationEnabled(
	socket *Socket,
	params *emulation.SetTouchEmulationEnabledParams,
) error {
	command := &protocol.Command{
		Method: "Emulation.setTouchEmulationEnabled",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetVirtualTimePolicy turns on virtual time for all frames (replacing real-time with a synthetic time
source) and sets the current virtual time policy. Note this supersedes any previous time budget.
EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setVirtualTimePolicy
*/
func (Emulation) SetVirtualTimePolicy(
	socket *Socket,
	params *emulation.SetVirtualTimePolicyParams,
) error {
	command := &protocol.Command{
		Method: "Emulation.setVirtualTimePolicy",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetVisibleSize resizes the frame/viewport of the page. Note that this does not affect the frame's
container (e.g. browser window). Can be used to produce screenshots of the specified size. Not
supported on Android. EXPERIMENTAL DEPRECATED

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setVisibleSize
*/
func (Emulation) SetVisibleSize(
	socket *Socket,
	params *emulation.SetVisibleSizeParams,
) error {
	command := &protocol.Command{
		Method: "Emulation.setVisibleSize",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
OnVirtualTimeAdvanced adds a handler to the Emulation.virtualTimeAdvanced event.
Emulation.virtualTimeAdvanced fires after the virtual time has advanced. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#event-virtualTimeAdvanced
*/
func (Emulation) OnVirtualTimeAdvanced(
	socket *Socket,
	callback func(event *emulation.VirtualTimeAdvancedEvent),
) {
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
OnVirtualTimeBudgetExpired adds a handler to the Emulation.virtualTimeBudgetExpired event.
Emulation.virtualTimeBudgetExpired fires after the virtual time budget for the current
VirtualTimePolicy has run out. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#event-virtualTimeBudgetExpired
*/
func (Emulation) OnVirtualTimeBudgetExpired(
	socket *Socket,
	callback func(event *emulation.VirtualTimeBudgetExpiredEvent),
) {
	handler := protocol.NewEventHandler(
		"Emulation.virtualTimeBudgetExpired",
		func(name string, params []byte) {
			event := &emulation.VirtualTimeBudgetExpiredEvent{}
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

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#event-virtualTimePaused
*/
func (Emulation) OnVirtualTimePaused(
	socket *Socket,
	callback func(event *emulation.VirtualTimePausedEvent),
) {
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

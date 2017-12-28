package protocol

import (
	"encoding/json"

	emulation "github.com/mkenney/go-chrome/protocol/emulation"
	sock "github.com/mkenney/go-chrome/socket"

	log "github.com/sirupsen/logrus"
)

/*
Emulation is a struct that provides a namespace for the Chrome Emulation protocol methods.

The Emulation protocol emulates different environments for the page.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/
*/
var Emulation = _emulation{}

type _emulation struct{}

/*
CanEmulate tells whether emulation is supported.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-canEmulate
*/
func (_emulation) CanEmulate(
	socket sock.Socketer,
) (emulation.CanEmulateResult, error) {
	command := sock.NewCommand("Emulation.canEmulate", nil)
	result := emulation.CanEmulateResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
ClearDeviceMetricsOverride clears the overridden device metrics.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-clearDeviceMetricsOverride
*/
func (_emulation) ClearDeviceMetricsOverride(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("Emulation.clearDeviceMetricsOverride", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
ClearGeolocationOverride clears the overridden Geolocation Position and Error.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-clearGeolocationOverride
*/
func (_emulation) ClearGeolocationOverride(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("Emulation.clearGeolocationOverride", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
ResetPageScaleFactor requests that page scale factor is reset to initial values. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-resetPageScaleFactor
*/
func (_emulation) ResetPageScaleFactor(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("Emulation.resetPageScaleFactor", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetCPUThrottlingRate enables CPU throttling to emulate slow CPUs. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setCPUThrottlingRate
*/
func (_emulation) SetCPUThrottlingRate(
	socket sock.Socketer,
	params *emulation.SetCPUThrottlingRateParams,
) error {
	command := sock.NewCommand("Emulation.setCPUThrottlingRate", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetDefaultBackgroundColorOverride sets or clears an override of the default background color of the
frame. This override is used if the content does not specify one.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setDefaultBackgroundColorOverride
*/
func (_emulation) SetDefaultBackgroundColorOverride(
	socket sock.Socketer,
	params *emulation.SetDefaultBackgroundColorOverrideParams,
) error {
	command := sock.NewCommand("Emulation.setDefaultBackgroundColorOverride", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetDeviceMetricsOverride overrides the values of device screen dimensions (window.screen.width,
window.screen.height, window.innerWidth, window.innerHeight, and "device-width"/"device-height"
related CSS media query results).

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setDeviceMetricsOverride
*/
func (_emulation) SetDeviceMetricsOverride(
	socket sock.Socketer,
	params *emulation.SetDeviceMetricsOverrideParams,
) error {
	command := sock.NewCommand("Emulation.setDeviceMetricsOverride", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetEmitTouchEventsForMouse enables touch events using a mouse. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setEmitTouchEventsForMouse
*/
func (_emulation) SetEmitTouchEventsForMouse(
	socket sock.Socketer,
	params *emulation.SetEmitTouchEventsForMouseParams,
) error {
	command := sock.NewCommand("Emulation.setEmitTouchEventsForMouse", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetEmulatedMedia emulates the given media for CSS media queries.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setEmulatedMedia
*/
func (_emulation) SetEmulatedMedia(
	socket sock.Socketer,
	params *emulation.SetEmulatedMediaParams,
) error {
	command := sock.NewCommand("Emulation.setEmulatedMedia", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetGeolocationOverride overrides the Geolocation Position or Error. Omitting any of the parameters
emulates position unavailable.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setGeolocationOverride
*/
func (_emulation) SetGeolocationOverride(
	socket sock.Socketer,
	params *emulation.SetGeolocationOverrideParams,
) error {
	command := sock.NewCommand("Emulation.setGeolocationOverride", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetNavigatorOverrides overrides value returned by the javascript navigator object. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setNavigatorOverrides
*/
func (_emulation) SetNavigatorOverrides(
	socket sock.Socketer,
	params *emulation.SetNavigatorOverridesParams,
) error {
	command := sock.NewCommand("Emulation.setNavigatorOverrides", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetPageScaleFactor sets a specified page scale factor. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setPageScaleFactor
*/
func (_emulation) SetPageScaleFactor(
	socket sock.Socketer,
	params *emulation.SetPageScaleFactorParams,
) error {
	command := sock.NewCommand("Emulation.setPageScaleFactor", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetScriptExecutionDisabled switches script execution in the page.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setScriptExecutionDisabled
*/
func (_emulation) SetScriptExecutionDisabled(
	socket sock.Socketer,
	params *emulation.SetScriptExecutionDisabledParams,
) error {
	command := sock.NewCommand("Emulation.setScriptExecutionDisabled", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetTouchEmulationEnabled enables touch on platforms which do not support it.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setTouchEmulationEnabled
*/
func (_emulation) SetTouchEmulationEnabled(
	socket sock.Socketer,
	params *emulation.SetTouchEmulationEnabledParams,
) error {
	command := sock.NewCommand("Emulation.setTouchEmulationEnabled", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetVirtualTimePolicy turns on virtual time for all frames (replacing real-time with a synthetic time
source) and sets the current virtual time policy. Note this supersedes any previous time budget.
EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setVirtualTimePolicy
*/
func (_emulation) SetVirtualTimePolicy(
	socket sock.Socketer,
	params *emulation.SetVirtualTimePolicyParams,
) error {
	command := sock.NewCommand("Emulation.setVirtualTimePolicy", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetVisibleSize resizes the frame/viewport of the page. Note that this does not affect the frame's
container (e.g. browser window). Can be used to produce screenshots of the specified size. Not
supported on Android. EXPERIMENTAL DEPRECATED

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setVisibleSize
*/
func (_emulation) SetVisibleSize(
	socket sock.Socketer,
	params *emulation.SetVisibleSizeParams,
) error {
	command := sock.NewCommand("Emulation.setVisibleSize", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
OnVirtualTimeAdvanced adds a handler to the Emulation.virtualTimeAdvanced event.
Emulation.virtualTimeAdvanced fires after the virtual time has advanced. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#event-virtualTimeAdvanced
*/
func (_emulation) OnVirtualTimeAdvanced(
	socket sock.Socketer,
	callback func(event *emulation.VirtualTimeAdvancedEvent),
) {
	handler := sock.NewEventHandler(
		"Emulation.virtualTimeAdvanced",
		func(response *sock.Response) {
			event := &emulation.VirtualTimeAdvancedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
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
func (_emulation) OnVirtualTimeBudgetExpired(
	socket sock.Socketer,
	callback func(event *emulation.VirtualTimeBudgetExpiredEvent),
) {
	handler := sock.NewEventHandler(
		"Emulation.virtualTimeBudgetExpired",
		func(response *sock.Response) {
			event := &emulation.VirtualTimeBudgetExpiredEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
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
func (_emulation) OnVirtualTimePaused(
	socket sock.Socketer,
	callback func(event *emulation.VirtualTimePausedEvent),
) {
	handler := sock.NewEventHandler(
		"Emulation.virtualTimePaused",
		func(response *sock.Response) {
			event := &emulation.VirtualTimePausedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

package protocol

import (
	"encoding/json"

	emulation "github.com/mkenney/go-chrome/protocol/emulation"
	sock "github.com/mkenney/go-chrome/socket"

	log "github.com/sirupsen/logrus"
)

/*
Emulation provides a namespace for the Chrome Emulation protocol methods. The Emulation protocol
emulates different environments for the page.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/
*/
var Emulation = EmulationProtocol{}

/*
EmulationProtocol defines the Emulation protocol methods.
*/
type EmulationProtocol struct{}

/*
CanEmulate tells whether emulation is supported.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-canEmulate
*/
func (EmulationProtocol) CanEmulate(
	socket sock.Socketer,
) (*emulation.CanEmulateResult, error) {
	command := sock.NewCommand("Emulation.canEmulate", nil)
	result := &emulation.CanEmulateResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
ClearDeviceMetricsOverride clears the overridden device metrics.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-clearDeviceMetricsOverride
*/
func (EmulationProtocol) ClearDeviceMetricsOverride(
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
func (EmulationProtocol) ClearGeolocationOverride(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("Emulation.clearGeolocationOverride", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
ResetPageScaleFactor requests that page scale factor is reset to initial values.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-resetPageScaleFactor
EXPERIMENTAL.
*/
func (EmulationProtocol) ResetPageScaleFactor(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("Emulation.resetPageScaleFactor", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetCPUThrottlingRate enables CPU throttling to emulate slow CPUs.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setCPUThrottlingRate
EXPERIMENTAL.
*/
func (EmulationProtocol) SetCPUThrottlingRate(
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
func (EmulationProtocol) SetDefaultBackgroundColorOverride(
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
func (EmulationProtocol) SetDeviceMetricsOverride(
	socket sock.Socketer,
	params *emulation.SetDeviceMetricsOverrideParams,
) error {
	command := sock.NewCommand("Emulation.setDeviceMetricsOverride", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetEmitTouchEventsForMouse enables touch events using a mouse.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setEmitTouchEventsForMouse
EXPERIMENTAL.
*/
func (EmulationProtocol) SetEmitTouchEventsForMouse(
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
func (EmulationProtocol) SetEmulatedMedia(
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
func (EmulationProtocol) SetGeolocationOverride(
	socket sock.Socketer,
	params *emulation.SetGeolocationOverrideParams,
) error {
	command := sock.NewCommand("Emulation.setGeolocationOverride", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetNavigatorOverrides overrides value returned by the javascript navigator object.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setNavigatorOverrides
EXPERIMENTAL.
*/
func (EmulationProtocol) SetNavigatorOverrides(
	socket sock.Socketer,
	params *emulation.SetNavigatorOverridesParams,
) error {
	command := sock.NewCommand("Emulation.setNavigatorOverrides", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetPageScaleFactor sets a specified page scale factor.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setPageScaleFactor
EXPERIMENTAL.
*/
func (EmulationProtocol) SetPageScaleFactor(
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
func (EmulationProtocol) SetScriptExecutionDisabled(
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
func (EmulationProtocol) SetTouchEmulationEnabled(
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

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setVirtualTimePolicy
EXPERIMENTAL.
*/
func (EmulationProtocol) SetVirtualTimePolicy(
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
supported on Android.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setVisibleSize
EXPERIMENTAL. DEPRECATED.
*/
func (EmulationProtocol) SetVisibleSize(
	socket sock.Socketer,
	params *emulation.SetVisibleSizeParams,
) error {
	command := sock.NewCommand("Emulation.setVisibleSize", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
OnVirtualTimeAdvanced adds a handler to the Emulation.virtualTimeAdvanced event.
Emulation.virtualTimeAdvanced fires after the virtual time has advanced.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#event-virtualTimeAdvanced
EXPERIMENTAL.
*/
func (EmulationProtocol) OnVirtualTimeAdvanced(
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
VirtualTimePolicy has run out.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#event-virtualTimeBudgetExpired
EXPERIMENTAL.
*/
func (EmulationProtocol) OnVirtualTimeBudgetExpired(
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
Emulation.virtualTimePaused fires after the virtual time has paused.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#event-virtualTimePaused
EXPERIMENTAL.
*/
func (EmulationProtocol) OnVirtualTimePaused(
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

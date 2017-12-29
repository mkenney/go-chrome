package socket

import (
	"encoding/json"

	emulation "github.com/mkenney/go-chrome/protocol/emulation"
	log "github.com/sirupsen/logrus"
)

/*
EmulationProtocol provides a namespace for the Chrome Emulation protocol methods. The Emulation
protocol emulates different environments for the page.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/
*/
type EmulationProtocol struct {
	Socket Socketer
}

/*
CanEmulate tells whether emulation is supported.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-canEmulate
*/
func (protocol *EmulationProtocol) CanEmulate() (*emulation.CanEmulateResult, error) {
	command := NewCommand("Emulation.canEmulate", nil)
	result := &emulation.CanEmulateResult{}
	protocol.Socket.SendCommand(command)

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
func (protocol *EmulationProtocol) ClearDeviceMetricsOverride() error {
	command := NewCommand("Emulation.clearDeviceMetricsOverride", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
ClearGeolocationOverride clears the overridden Geolocation Position and Error.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-clearGeolocationOverride
*/
func (protocol *EmulationProtocol) ClearGeolocationOverride() error {
	command := NewCommand("Emulation.clearGeolocationOverride", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
ResetPageScaleFactor requests that page scale factor is reset to initial values.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-resetPageScaleFactor
EXPERIMENTAL.
*/
func (protocol *EmulationProtocol) ResetPageScaleFactor() error {
	command := NewCommand("Emulation.resetPageScaleFactor", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetCPUThrottlingRate enables CPU throttling to emulate slow CPUs.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setCPUThrottlingRate
EXPERIMENTAL.
*/
func (protocol *EmulationProtocol) SetCPUThrottlingRate(
	params *emulation.SetCPUThrottlingRateParams,
) error {
	command := NewCommand("Emulation.setCPUThrottlingRate", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetDefaultBackgroundColorOverride sets or clears an override of the default background color of the
frame. This override is used if the content does not specify one.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setDefaultBackgroundColorOverride
*/
func (protocol *EmulationProtocol) SetDefaultBackgroundColorOverride(
	params *emulation.SetDefaultBackgroundColorOverrideParams,
) error {
	command := NewCommand("Emulation.setDefaultBackgroundColorOverride", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetDeviceMetricsOverride overrides the values of device screen dimensions (window.screen.width,
window.screen.height, window.innerWidth, window.innerHeight, and "device-width"/"device-height"
related CSS media query results).

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setDeviceMetricsOverride
*/
func (protocol *EmulationProtocol) SetDeviceMetricsOverride(
	params *emulation.SetDeviceMetricsOverrideParams,
) error {
	command := NewCommand("Emulation.setDeviceMetricsOverride", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetEmitTouchEventsForMouse enables touch events using a mouse.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setEmitTouchEventsForMouse
EXPERIMENTAL.
*/
func (protocol *EmulationProtocol) SetEmitTouchEventsForMouse(
	params *emulation.SetEmitTouchEventsForMouseParams,
) error {
	command := NewCommand("Emulation.setEmitTouchEventsForMouse", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetEmulatedMedia emulates the given media for CSS media queries.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setEmulatedMedia
*/
func (protocol *EmulationProtocol) SetEmulatedMedia(
	params *emulation.SetEmulatedMediaParams,
) error {
	command := NewCommand("Emulation.setEmulatedMedia", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetGeolocationOverride overrides the Geolocation Position or Error. Omitting any of the parameters
emulates position unavailable.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setGeolocationOverride
*/
func (protocol *EmulationProtocol) SetGeolocationOverride(
	params *emulation.SetGeolocationOverrideParams,
) error {
	command := NewCommand("Emulation.setGeolocationOverride", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetNavigatorOverrides overrides value returned by the javascript navigator object.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setNavigatorOverrides
EXPERIMENTAL.
*/
func (protocol *EmulationProtocol) SetNavigatorOverrides(
	params *emulation.SetNavigatorOverridesParams,
) error {
	command := NewCommand("Emulation.setNavigatorOverrides", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetPageScaleFactor sets a specified page scale factor.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setPageScaleFactor
EXPERIMENTAL.
*/
func (protocol *EmulationProtocol) SetPageScaleFactor(
	params *emulation.SetPageScaleFactorParams,
) error {
	command := NewCommand("Emulation.setPageScaleFactor", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetScriptExecutionDisabled switches script execution in the page.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setScriptExecutionDisabled
*/
func (protocol *EmulationProtocol) SetScriptExecutionDisabled(
	params *emulation.SetScriptExecutionDisabledParams,
) error {
	command := NewCommand("Emulation.setScriptExecutionDisabled", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetTouchEmulationEnabled enables touch on platforms which do not support it.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setTouchEmulationEnabled
*/
func (protocol *EmulationProtocol) SetTouchEmulationEnabled(
	params *emulation.SetTouchEmulationEnabledParams,
) error {
	command := NewCommand("Emulation.setTouchEmulationEnabled", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetVirtualTimePolicy turns on virtual time for all frames (replacing real-time with a synthetic time
source) and sets the current virtual time policy. Note this supersedes any previous time budget.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setVirtualTimePolicy
EXPERIMENTAL.
*/
func (protocol *EmulationProtocol) SetVirtualTimePolicy(
	params *emulation.SetVirtualTimePolicyParams,
) error {
	command := NewCommand("Emulation.setVirtualTimePolicy", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetVisibleSize resizes the frame/viewport of the page. Note that this does not affect the frame's
container (e.g. browser window). Can be used to produce screenshots of the specified size. Not
supported on Android.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setVisibleSize
EXPERIMENTAL. DEPRECATED.
*/
func (protocol *EmulationProtocol) SetVisibleSize(
	params *emulation.SetVisibleSizeParams,
) error {
	command := NewCommand("Emulation.setVisibleSize", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
OnVirtualTimeAdvanced adds a handler to the Emulation.virtualTimeAdvanced event.
Emulation.virtualTimeAdvanced fires after the virtual time has advanced.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#event-virtualTimeAdvanced
EXPERIMENTAL.
*/
func (protocol *EmulationProtocol) OnVirtualTimeAdvanced(
	callback func(event *emulation.VirtualTimeAdvancedEvent),
) {
	handler := NewEventHandler(
		"Emulation.virtualTimeAdvanced",
		func(response *Response) {
			event := &emulation.VirtualTimeAdvancedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnVirtualTimeBudgetExpired adds a handler to the Emulation.virtualTimeBudgetExpired event.
Emulation.virtualTimeBudgetExpired fires after the virtual time budget for the current
VirtualTimePolicy has run out.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#event-virtualTimeBudgetExpired
EXPERIMENTAL.
*/
func (protocol *EmulationProtocol) OnVirtualTimeBudgetExpired(
	callback func(event *emulation.VirtualTimeBudgetExpiredEvent),
) {
	handler := NewEventHandler(
		"Emulation.virtualTimeBudgetExpired",
		func(response *Response) {
			event := &emulation.VirtualTimeBudgetExpiredEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnVirtualTimePaused adds a handler to the Emulation.virtualTimePaused event.
Emulation.virtualTimePaused fires after the virtual time has paused.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#event-virtualTimePaused
EXPERIMENTAL.
*/
func (protocol *EmulationProtocol) OnVirtualTimePaused(
	callback func(event *emulation.VirtualTimePausedEvent),
) {
	handler := NewEventHandler(
		"Emulation.virtualTimePaused",
		func(response *Response) {
			event := &emulation.VirtualTimePausedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

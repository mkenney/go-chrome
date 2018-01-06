package socket

import (
	"encoding/json"

	emulation "github.com/mkenney/go-chrome/cdtp/emulation"
	log "github.com/sirupsen/logrus"
)

/*
EmulationProtocol provides a namespace for the Chrome Emulation protocol
methods. The Emulation protocol emulates different environments for the page.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/
*/
type EmulationProtocol struct {
	Socket Socketer
}

/*
CanEmulate tells whether emulation is supported.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-canEmulate
*/
func (protocol *EmulationProtocol) CanEmulate() chan *emulation.CanEmulateResult {
	resultChan := make(chan *emulation.CanEmulateResult)
	command := NewCommand(protocol.Socket, "Emulation.canEmulate", nil)
	result := &emulation.CanEmulateResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		} else {
			result.CDTPError = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
}

/*
ClearDeviceMetricsOverride clears the overridden device metrics.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-clearDeviceMetricsOverride
*/
func (protocol *EmulationProtocol) ClearDeviceMetricsOverride() chan *emulation.ClearDeviceMetricsOverrideResult {
	resultChan := make(chan *emulation.ClearDeviceMetricsOverrideResult)
	command := NewCommand(protocol.Socket, "Emulation.clearDeviceMetricsOverride", nil)
	result := &emulation.ClearDeviceMetricsOverrideResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
ClearGeolocationOverride clears the overridden Geolocation Position and Error.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-clearGeolocationOverride
*/
func (protocol *EmulationProtocol) ClearGeolocationOverride() chan *emulation.ClearGeolocationOverrideResult {
	resultChan := make(chan *emulation.ClearGeolocationOverrideResult)
	command := NewCommand(protocol.Socket, "Emulation.clearGeolocationOverride", nil)
	result := &emulation.ClearGeolocationOverrideResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
ResetPageScaleFactor requests that page scale factor is reset to initial values.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-resetPageScaleFactor
EXPERIMENTAL.
*/
func (protocol *EmulationProtocol) ResetPageScaleFactor() chan *emulation.ResetPageScaleFactorResult {
	resultChan := make(chan *emulation.ResetPageScaleFactorResult)
	command := NewCommand(protocol.Socket, "Emulation.resetPageScaleFactor", nil)
	result := &emulation.ResetPageScaleFactorResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
SetCPUThrottlingRate enables CPU throttling to emulate slow CPUs.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setCPUThrottlingRate
EXPERIMENTAL.
*/
func (protocol *EmulationProtocol) SetCPUThrottlingRate(
	params *emulation.SetCPUThrottlingRateParams,
) chan *emulation.SetCPUThrottlingRateResult {
	resultChan := make(chan *emulation.SetCPUThrottlingRateResult)
	command := NewCommand(protocol.Socket, "Emulation.setCPUThrottlingRate", params)
	result := &emulation.SetCPUThrottlingRateResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
SetDefaultBackgroundColorOverride sets or clears an override of the default
background color of the frame. This override is used if the content does not
specify one.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setDefaultBackgroundColorOverride
*/
func (protocol *EmulationProtocol) SetDefaultBackgroundColorOverride(
	params *emulation.SetDefaultBackgroundColorOverrideParams,
) chan *emulation.SetDefaultBackgroundColorOverrideResult {
	resultChan := make(chan *emulation.SetDefaultBackgroundColorOverrideResult)
	command := NewCommand(protocol.Socket, "Emulation.setDefaultBackgroundColorOverride", params)
	result := &emulation.SetDefaultBackgroundColorOverrideResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
SetDeviceMetricsOverride overrides the values of device screen dimensions
(window.screen.width, window.screen.height, window.innerWidth,
window.innerHeight, and "device-width"/"device-height" related CSS media query
results).

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setDeviceMetricsOverride
*/
func (protocol *EmulationProtocol) SetDeviceMetricsOverride(
	params *emulation.SetDeviceMetricsOverrideParams,
) chan *emulation.SetDeviceMetricsOverrideResult {
	resultChan := make(chan *emulation.SetDeviceMetricsOverrideResult)
	command := NewCommand(protocol.Socket, "Emulation.setDeviceMetricsOverride", params)
	result := &emulation.SetDeviceMetricsOverrideResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
SetEmitTouchEventsForMouse enables touch events using a mouse.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setEmitTouchEventsForMouse
EXPERIMENTAL.
*/
func (protocol *EmulationProtocol) SetEmitTouchEventsForMouse(
	params *emulation.SetEmitTouchEventsForMouseParams,
) chan *emulation.SetEmitTouchEventsForMouseResult {
	resultChan := make(chan *emulation.SetEmitTouchEventsForMouseResult)
	command := NewCommand(protocol.Socket, "Emulation.setEmitTouchEventsForMouse", params)
	result := &emulation.SetEmitTouchEventsForMouseResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
SetEmulatedMedia emulates the given media for CSS media queries.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setEmulatedMedia
*/
func (protocol *EmulationProtocol) SetEmulatedMedia(
	params *emulation.SetEmulatedMediaParams,
) chan *emulation.SetEmulatedMediaResult {
	resultChan := make(chan *emulation.SetEmulatedMediaResult)
	command := NewCommand(protocol.Socket, "Emulation.setEmulatedMedia", params)
	result := &emulation.SetEmulatedMediaResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
SetGeolocationOverride overrides the Geolocation Position or Error. Omitting any
of the parameters emulates position unavailable.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setGeolocationOverride
*/
func (protocol *EmulationProtocol) SetGeolocationOverride(
	params *emulation.SetGeolocationOverrideParams,
) chan *emulation.SetGeolocationOverrideResult {
	resultChan := make(chan *emulation.SetGeolocationOverrideResult)
	command := NewCommand(protocol.Socket, "Emulation.setGeolocationOverride", params)
	result := &emulation.SetGeolocationOverrideResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
SetNavigatorOverrides overrides value returned by the javascript navigator
object.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setNavigatorOverrides
EXPERIMENTAL.
*/
func (protocol *EmulationProtocol) SetNavigatorOverrides(
	params *emulation.SetNavigatorOverridesParams,
) chan *emulation.SetNavigatorOverridesResult {
	resultChan := make(chan *emulation.SetNavigatorOverridesResult)
	command := NewCommand(protocol.Socket, "Emulation.setNavigatorOverrides", params)
	result := &emulation.SetNavigatorOverridesResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
SetPageScaleFactor sets a specified page scale factor.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setPageScaleFactor
EXPERIMENTAL.
*/
func (protocol *EmulationProtocol) SetPageScaleFactor(
	params *emulation.SetPageScaleFactorParams,
) chan *emulation.SetPageScaleFactorResult {
	resultChan := make(chan *emulation.SetPageScaleFactorResult)
	command := NewCommand(protocol.Socket, "Emulation.setPageScaleFactor", params)
	result := &emulation.SetPageScaleFactorResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
SetScriptExecutionDisabled switches script execution in the page.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setScriptExecutionDisabled
*/
func (protocol *EmulationProtocol) SetScriptExecutionDisabled(
	params *emulation.SetScriptExecutionDisabledParams,
) chan *emulation.SetScriptExecutionDisabledResult {
	resultChan := make(chan *emulation.SetScriptExecutionDisabledResult)
	command := NewCommand(protocol.Socket, "Emulation.setScriptExecutionDisabled", params)
	result := &emulation.SetScriptExecutionDisabledResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
SetTouchEmulationEnabled enables touch on platforms which do not support it.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setTouchEmulationEnabled
*/
func (protocol *EmulationProtocol) SetTouchEmulationEnabled(
	params *emulation.SetTouchEmulationEnabledParams,
) chan *emulation.SetTouchEmulationEnabledResult {
	resultChan := make(chan *emulation.SetTouchEmulationEnabledResult)
	command := NewCommand(protocol.Socket, "Emulation.setTouchEmulationEnabled", params)
	result := &emulation.SetTouchEmulationEnabledResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
SetVirtualTimePolicy turns on virtual time for all frames (replacing real-time
with a synthetic time source) and sets the current virtual time policy. Note
this supersedes any previous time budget.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setVirtualTimePolicy
EXPERIMENTAL.
*/
func (protocol *EmulationProtocol) SetVirtualTimePolicy(
	params *emulation.SetVirtualTimePolicyParams,
) chan *emulation.SetVirtualTimePolicyResult {
	resultChan := make(chan *emulation.SetVirtualTimePolicyResult)
	command := NewCommand(protocol.Socket, "Emulation.SetVirtualTimePolicy", nil)
	result := &emulation.SetVirtualTimePolicyResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		} else {
			result.CDTPError = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
}

/*
SetVisibleSize resizes the frame/viewport of the page. Note that this does not
affect the frame's container (e.g. browser window). Can be used to produce
screenshots of the specified size. Not supported on Android.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setVisibleSize
EXPERIMENTAL. DEPRECATED.
*/
func (protocol *EmulationProtocol) SetVisibleSize(
	params *emulation.SetVisibleSizeParams,
) chan *emulation.SetVisibleSizeResult {
	resultChan := make(chan *emulation.SetVisibleSizeResult)
	command := NewCommand(protocol.Socket, "Emulation.setVisibleSize", params)
	result := &emulation.SetVisibleSizeResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		}
		resultChan <- result
	}()

	return resultChan
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
OnVirtualTimeBudgetExpired adds a handler to the Emulation.virtualTimeBudgetExpired
event. Emulation.virtualTimeBudgetExpired fires after the virtual time budget
for the current VirtualTimePolicy has run out.

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

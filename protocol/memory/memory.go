/*
Package memory provides type definitions for use with the Chrome Memory protocol

https://chromedevtools.github.io/devtools-protocol/tot/Memory/
*/
package memory

/*
GetDOMCountersParams represents Memory.getDOMCounters parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Memory/#method-getDOMCounters
*/
type GetDOMCountersParams struct {
	Documents        int `json:"documents"`
	Nodes            int `json:"nodes"`
	JsEventListeners int `json:"jsEventListeners"`
}

/*
SetPressureNotificationsSuppressedParams represents Memory.setPressureNotificationsSuppressed
parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Memory/#method-setPressureNotificationsSuppressed
*/
type SetPressureNotificationsSuppressedParams struct {
	// If true, memory pressure notifications will be suppressed.
	Suppressed bool `json:"suppressed"`
}

/*
SimulatePressureNotificationParams represents Memory.simulatePressureNotification parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Memory/#method-simulatePressureNotification
*/
type SimulatePressureNotificationParams struct {
	// Memory pressure level of the notification.
	Level PressureLevel `json:"level"`
}

/*
PressureLevel is the memory pressure level.

https://chromedevtools.github.io/devtools-protocol/tot/Memory/#type-PressureLevel
*/
type PressureLevel string

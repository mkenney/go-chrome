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
GetDOMCountersResult represents the result of calls to Memory.getDOMCounters.

https://chromedevtools.github.io/devtools-protocol/tot/Memory/#method-getDOMCounters
*/
type GetDOMCountersResult struct {
	CDTPError error `json:"-"`
}

/*
PrepareForLeakDetectionResult represents the result of calls to Memory.prepareForLeakDetection.

https://chromedevtools.github.io/devtools-protocol/tot/Memory/#method-prepareForLeakDetection
*/
type PrepareForLeakDetectionResult struct {
	CDTPError error `json:"-"`
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
SetPressureNotificationsSuppressedResult represents the result of calls to
Memory.setPressureNotificationsSuppressed.

https://chromedevtools.github.io/devtools-protocol/tot/Memory/#method-setPressureNotificationsSuppressed
*/
type SetPressureNotificationsSuppressedResult struct {
	CDTPError error `json:"-"`
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
SimulatePressureNotificationResult represents the result of calls to
Memory.simulatePressureNotification.

https://chromedevtools.github.io/devtools-protocol/tot/Memory/#method-simulatePressureNotification
*/
type SimulatePressureNotificationResult struct {
	CDTPError error `json:"-"`
}

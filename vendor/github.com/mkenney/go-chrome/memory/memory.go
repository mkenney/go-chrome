package Memory

import (
	"fmt"
)

/*
GetDOMCountersParams represents Memory.getDOMCounters parameters.
*/
type GetDOMCountersParams struct {
	Documents        int `json:"documents"`
	Nodes            int `json:"nodes"`
	JsEventListeners int `json:"jsEventListeners"`
}

/*
SetPressureNotificationsSuppressedParams represents Memory.setPressureNotificationsSuppressed
parameters.
*/
type SetPressureNotificationsSuppressedParams struct {
	// If true, memory pressure notifications will be suppressed.
	Suppressed bool `json:"suppressed"`
}

/*
SimulatePressureNotificationParams represents Memory.simulatePressureNotification parameters.
*/
type SimulatePressureNotificationParams struct {
	// Memory pressure level of the notification.
	Level PressureLevel `json:"level"`
}

/*
PressureLevel is the memory pressure level.
*/
type PressureLevel int

const (
	_moderate PressureLevel = iota
	_critical
)

func (a PressureLevel) String() string {
	if a == 0 {
		return "moderate"
	}
	if a == 1 {
		return "critical"
	}
	panic(fmt.Errorf("Invalid PressureLevel %d", a))
}

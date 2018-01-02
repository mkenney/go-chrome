package emulation

/*
VirtualTimeAdvancedEvent represents Emulation.virtualTimeAdvanced event data.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#event-virtualTimeAdvanced
*/
type VirtualTimeAdvancedEvent struct {
	// The amount of virtual time that has elapsed in milliseconds since virtual
	// time was first enabled.
	VirtualTimeElapsed int `json:"virtualTimeElapsed"`
}

/*
VirtualTimeBudgetExpiredEvent represents Emulation.virtualTimeBudgetExpired event data.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#event-virtualTimeBudgetExpired
*/
type VirtualTimeBudgetExpiredEvent struct{}

/*
VirtualTimePausedEvent represents Emulation.virtualTimePaused event data.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#event-virtualTimePaused
*/
type VirtualTimePausedEvent struct {
	// The amount of virtual time that has elapsed in milliseconds since virtual
	// time was first enabled.
	VirtualTimeElapsed int `json:"virtualTimeElapsed"`
}

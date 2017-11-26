package Emulation

import (
	"fmt"
)

/*
ScreenOrientation represents a screen orientation.
*/
type ScreenOrientation struct {
	// Orientation type. Allowed values: portraitPrimary, portraitSecondary, landscapePrimary,
	// landscapeSecondary.
	Type string `json:"type"`

	// Orientation angle.
	Angle int `json:"angle"`
}

/*
VirtualTimePolicy is the time policy
	- advance: If the scheduler runs out of immediate work, the virtual time base may fast forward
	  to allow the next delayed task (if any) to run
	- pause: The virtual time base may not advance
	- pauseIfNetworkFetchesPending: The virtual time base may not advance if there are any pending
	  resource fetches.
EXPERIMENTAL
*/
type VirtualTimePolicy int

const (
	_advance VirtualTimePolicy = iota
	_pause
	_pauseIfNetworkFetchesPending
)

func (a VirtualTimePolicy) String() string {
	if a == 0 {
		return "advance"
	}
	if a == 1 {
		return "pause"
	}
	if a == 2 {
		return "pauseIfNetworkFetchesPending"
	}
	panic(fmt.Errorf("Invalid VirtualTimePolicy %d", a))
}

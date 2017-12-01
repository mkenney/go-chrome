package Memory

import (
	"fmt"
)

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

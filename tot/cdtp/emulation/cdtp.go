/*
Package emulation provides type definitions for use with the Chrome Emulation protocol

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/
*/
package emulation

/*
ScreenOrientation represents a screen orientation.

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#type-ScreenOrientation
*/
type ScreenOrientation struct {
	// Orientation type. Allowed values:
	//	- OrientationType.PortraitPrimary
	//	- OrientationType.PortraitSecondary
	//	- OrientationType.LandscapePrimary
	//	- OrientationType.LandscapeSecondary
	Type OrientationTypeEnum `json:"type"`

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

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#type-VirtualTimePolicy
*/
type VirtualTimePolicy string

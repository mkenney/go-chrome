package serviceWorker

import (
	"encoding/json"
	"fmt"
)

type versionRunningStatusEnum struct {
	Stopped  VersionRunningStatusEnum
	Starting VersionRunningStatusEnum
	Running  VersionRunningStatusEnum
	Stopping VersionRunningStatusEnum
}

var VersionRunningStatus = versionRunningStatusEnum{
	Stopped:  versionRunningStatusStopped,
	Starting: versionRunningStatusStarting,
	Running:  versionRunningStatusRunning,
	Stopping: versionRunningStatusStopping,
}

/*
VersionRunningStatus is the version running status. Allowed values:
	- VersionRunningStatus.Stopped  "stopped"
	- VersionRunningStatus.Starting "starting"
	- VersionRunningStatus.Running  "running"
	- VersionRunningStatus.Stopping "stopping"

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#type-ServiceWorkerVersionRunningStatus
*/
type VersionRunningStatusEnum int

/*
String implements Stringer
*/
func (enum VersionRunningStatusEnum) String() string {
	return _versionRunningStatusEnums[enum]
}

/*
MarshalJSON implements json.Marshaler
*/
func (enum VersionRunningStatusEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

/*
UnmarshalJSON implements json.Unmarshaler
*/
func (enum *VersionRunningStatusEnum) UnmarshalJSON(bytes []byte) error {
	var err error
	var val string

	err = json.Unmarshal(bytes, &val)
	if nil != err {
		return err
	}

	for k, v := range _versionRunningStatusEnums {
		if v == val {
			*enum = k
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid type value", bytes)
}

const (
	// versionRunningStatusStopped represents the "stopped" value.
	versionRunningStatusStopped VersionRunningStatusEnum = iota + 1
	// versionRunningStatusStarting represents the "starting" value.
	versionRunningStatusStarting
	// versionRunningStatusRunning represents the "running" value.
	versionRunningStatusRunning
	// versionRunningStatusStopping represents the "stopping" value.
	versionRunningStatusStopping
)

var _versionRunningStatusEnums = map[VersionRunningStatusEnum]string{
	versionRunningStatusStopped:  "stopped",
	versionRunningStatusStarting: "starting",
	versionRunningStatusRunning:  "running",
	versionRunningStatusStopping: "stopping",
}

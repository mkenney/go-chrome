package serviceWorker

import (
	"encoding/json"
	"fmt"
)

type versionStatusEnum struct {
	New        VersionStatusEnum
	Installing VersionStatusEnum
	Installed  VersionStatusEnum
	Activating VersionStatusEnum
	Activated  VersionStatusEnum
	Redundant  VersionStatusEnum
}

var VersionStatus = versionStatusEnum{
	New:        versionStatusNew,
	Installing: versionStatusInstalling,
	Installed:  versionStatusInstalled,
	Activating: versionStatusActivating,
	Activated:  versionStatusActivated,
	Redundant:  versionStatusRedundant,
}

/*
VersionStatus is the version status. Allowed values:
	- VersionStatus.New
	- VersionStatus.Installing
	- VersionStatus.Installed
	- VersionStatus.Activating
	- VersionStatus.Activated
	- VersionStatus.Redundant

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#type-ServiceWorkerVersionStatus
*/
type VersionStatusEnum int

/*
String implements Stringer
*/
func (enum VersionStatusEnum) String() string {
	return _versionStatusEnums[enum]
}

/*
MarshalJSON implements json.Marshaler
*/
func (enum VersionStatusEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

/*
UnmarshalJSON implements json.Unmarshaler
*/
func (enum *VersionStatusEnum) UnmarshalJSON(bytes []byte) error {
	var err error
	var val string

	err = json.Unmarshal(bytes, &val)
	if nil != err {
		return err
	}

	for k, v := range _versionStatusEnums {
		if v == val {
			*enum = k
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid type value", bytes)
}

const (
	// versionStatusNew represents the "new" value.
	versionStatusNew VersionStatusEnum = iota + 1
	// versionStatusInstalling represents the "installing" value.
	versionStatusInstalling
	// versionStatusInstalled represents the "installed" value.
	versionStatusInstalled
	// versionStatusActivating represents the "activating" value.
	versionStatusActivating
	// versionStatusActivated represents the "activated" value.
	versionStatusActivated
	// versionStatusRedundant represents the "redundant" value.
	versionStatusRedundant
)

var _versionStatusEnums = map[VersionStatusEnum]string{
	versionStatusNew:        "new",
	versionStatusInstalling: "installing",
	versionStatusInstalled:  "installed",
	versionStatusActivating: "activating",
	versionStatusActivated:  "activated",
	versionStatusRedundant:  "redundant",
}

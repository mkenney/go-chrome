package emulation

import (
	"encoding/json"
	"fmt"
)

type configurationEnum struct {
	Mobile  ConfigurationEnum
	Desktop ConfigurationEnum
}

var Configuration = configurationEnum{
	Mobile:  configurationMobile,
	Desktop: configurationDesktop,
}

/*
Optional. Touch/gesture events configuration. Default: current platform. Allowed \
values:
	- mobile
	- desktop

https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#type-ScreenOrientation
*/
type ConfigurationEnum int

/*
String implements Stringer
*/
func (enum ConfigurationEnum) String() string {
	return _configurationEnums[enum]
}

/*
MarshalJSON implements json.Marshaler
*/
func (enum ConfigurationEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

/*
UnmarshalJSON implements json.Unmarshaler
*/
func (enum *ConfigurationEnum) UnmarshalJSON(bytes []byte) error {
	var err error
	var val string

	err = json.Unmarshal(bytes, &val)
	if nil != err {
		return err
	}

	for k, v := range _configurationEnums {
		if v == val {
			*enum = k
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid configuration value", bytes)
}

const (
	// configurationMobile represents the "mobile" value.
	configurationMobile ConfigurationEnum = iota + 1
	// configurationDesktop represents the "desktop" value.
	configurationDesktop
)

var _configurationEnums = map[ConfigurationEnum]string{
	ConfigurationEnum(0): "",
	configurationMobile:  "mobile",
	configurationDesktop: "desktop",
}

package network

import (
	"encoding/json"
	"fmt"
)

type connectionTypeEnum struct {
	None       ConnectionTypeEnum
	Cellular2g ConnectionTypeEnum
	Cellular3g ConnectionTypeEnum
	Cellular4g ConnectionTypeEnum
	Bluetooth  ConnectionTypeEnum
	Ethernet   ConnectionTypeEnum
	Wifi       ConnectionTypeEnum
	Wimax      ConnectionTypeEnum
	Other      ConnectionTypeEnum
}

/*
ConnectionType provides named acces to the ConnectionTypeEnum values.
*/
var ConnectionType = connectionTypeEnum{
	None:       connectionTypeNone,
	Cellular2g: connectionTypeCellular2g,
	Cellular3g: connectionTypeCellular3g,
	Cellular4g: connectionTypeCellular4g,
	Bluetooth:  connectionTypeBluetooth,
	Ethernet:   connectionTypeEthernet,
	Wifi:       connectionTypeWifi,
	Wimax:      connectionTypeWimax,
	Other:      connectionTypeOther,
}

/*
ConnectionTypeEnum is the underlying connection technology that the browser is
supposedly using. Allowed values:
	- ConnectionType.None       "none"
	- ConnectionType.Cellular2g "cellular2g"
	- ConnectionType.Cellular3g "cellular3g"
	- ConnectionType.Cellular4g "cellular4g"
	- ConnectionType.Bluetooth  "bluetooth"
	- ConnectionType.Ethernet   "ethernet"
	- ConnectionType.Wifi       "wifi"
	- ConnectionType.Wimax      "wimax"
	- ConnectionType.Other      "other"

https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-ConnectionType
*/
type ConnectionTypeEnum int

/*
String implements Stringer
*/
func (enum ConnectionTypeEnum) String() string {
	return _connectionTypeEnums[enum]
}

/*
MarshalJSON implements json.Marshaler
*/
func (enum ConnectionTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

/*
UnmarshalJSON implements json.Unmarshaler
*/
func (enum *ConnectionTypeEnum) UnmarshalJSON(bytes []byte) error {
	var err error
	var val string

	err = json.Unmarshal(bytes, &val)
	if nil != err {
		return err
	}

	for k, v := range _connectionTypeEnums {
		if v == val {
			*enum = k
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid type value", bytes)
}

const (
	// connectionTypeNone represents the "none" value.
	connectionTypeNone ConnectionTypeEnum = iota + 1
	// connectionTypeCellular2g represents the "cellular2g" value.
	connectionTypeCellular2g
	// connectionTypeCellular3g represents the "cellular3g" value.
	connectionTypeCellular3g
	// connectionTypeCellular4g represents the "cellular4g" value.
	connectionTypeCellular4g
	// connectionTypeBluetooth represents the "bluetooth" value.
	connectionTypeBluetooth
	// connectionTypeEthernet represents the "ethernet" value.
	connectionTypeEthernet
	// connectionTypeWifi represents the "wifi" value.
	connectionTypeWifi
	// connectionTypeWimax represents the "wimax" value.
	connectionTypeWimax
	// connectionTypeOther represents the "other" value.
	connectionTypeOther
)

var _connectionTypeEnums = map[ConnectionTypeEnum]string{
	ConnectionTypeEnum(0):    "",
	connectionTypeNone:       "none",
	connectionTypeCellular2g: "cellular2g",
	connectionTypeCellular3g: "cellular3g",
	connectionTypeCellular4g: "cellular4g",
	connectionTypeBluetooth:  "bluetooth",
	connectionTypeEthernet:   "ethernet",
	connectionTypeWifi:       "wifi",
	connectionTypeWimax:      "wimax",
	connectionTypeOther:      "other",
}

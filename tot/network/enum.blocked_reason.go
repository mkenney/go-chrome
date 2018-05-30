package network

import (
	"encoding/json"
	"fmt"
)

type blockedReasonEnum struct {
	Csp               BlockedReasonEnum
	MixedContent      BlockedReasonEnum
	Origin            BlockedReasonEnum
	Inspector         BlockedReasonEnum
	SubresourceFilter BlockedReasonEnum
	Other             BlockedReasonEnum
}

/*
BlockedReason provides named acces to the BlockedReasonEnum values.
*/
var BlockedReason = blockedReasonEnum{
	Csp:               blockedReasonCsp,
	MixedContent:      blockedReasonMixedContent,
	Origin:            blockedReasonOrigin,
	Inspector:         blockedReasonInspector,
	SubresourceFilter: blockedReasonSubresourceFilter,
	Other:             blockedReasonOther,
}

/*
BlockedReasonEnum defines the reason why request was blocked. Allowed values:
	- BlockedReason.Csp               "csp"
	- BlockedReason.MixedContent      "mixed-content"
	- BlockedReason.Origin            "origin"
	- BlockedReason.Inspector         "inspector"
	- BlockedReason.SubresourceFilter "subresource-filter"
	- BlockedReason.Other             "other"

https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-BlockedReason
*/
type BlockedReasonEnum int

/*
String implements Stringer
*/
func (enum BlockedReasonEnum) String() string {
	return _blockedReasonEnums[enum]
}

/*
MarshalJSON implements json.Marshaler
*/
func (enum BlockedReasonEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

/*
UnmarshalJSON implements json.Unmarshaler
*/
func (enum *BlockedReasonEnum) UnmarshalJSON(bytes []byte) error {
	var err error
	var val string

	err = json.Unmarshal(bytes, &val)
	if nil != err {
		return err
	}

	for k, v := range _blockedReasonEnums {
		if v == val {
			*enum = k
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid type value", bytes)
}

const (
	// blockedReasonCsp represents the "csp" value.
	blockedReasonCsp BlockedReasonEnum = iota + 1
	// blockedReasonMixedContent represents the "mixed-content" value.
	blockedReasonMixedContent
	// blockedReasonOrigin represents the "origin" value.
	blockedReasonOrigin
	// blockedReasonInspector represents the "inspector" value.
	blockedReasonInspector
	// blockedReasonSubresourceFilter represents the "subresource-filter" value.
	blockedReasonSubresourceFilter
	// blockedReasonOther represents the "other" value.
	blockedReasonOther
)

var _blockedReasonEnums = map[BlockedReasonEnum]string{
	BlockedReasonEnum(0):           "",
	blockedReasonCsp:               "csp",
	blockedReasonMixedContent:      "mixed-content",
	blockedReasonOrigin:            "origin",
	blockedReasonInspector:         "inspector",
	blockedReasonSubresourceFilter: "subresource-filter",
	blockedReasonOther:             "other",
}

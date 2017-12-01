package Accessibility

import (
	DOM "app/chrome/dom"
	"fmt"
)

/*
PartialAXTreeParams are parameters for GetPartialAXTree.
getPartialAXTree fetches the accessibility node and partial accessibility tree for this DOM node, if
it exists.
*/
type PartialAXTreeParams struct {
	// ID of the node to get the partial accessibility tree for.
	NodeID DOM.NodeID `json:"nodeId"`

	// Optional. Whether to fetch this nodes ancestors, siblings and children. Defaults to true.
	FetchRelatives bool `json:"fetchRelatives,omitempty"`
}

/*
PartialAXTreeResult is the AXNode objects for this DOM node, if it exists, plus its ancestors,
siblings and children, if requested.
*/
type PartialAXTreeResult []*AXNode

/*
AXNodeID is the unique accessibility node identifier.
*/
type AXNodeID string

/*
AXValueType is an enum of possible property types.
*/
type AXValueType int

const (
	_boolean AXValueType = iota
	_tristate
	_booleanOrUndefined
	_idref
	_idrefList
	_integer
	_node
	_nodeList
	_number
	_string
	_computedString
	_token
	_tokenList
	_domRelation
	_role
	_internalRole
	_valueUndefined
)

func (a AXValueType) String() string {
	if a == 0 {
		return "boolean"
	}
	if a == 1 {
		return "tristate"
	}
	if a == 2 {
		return "booleanOrUndefined"
	}
	if a == 3 {
		return "idref"
	}
	if a == 4 {
		return "idrefList"
	}
	if a == 5 {
		return "integer"
	}
	if a == 6 {
		return "node"
	}
	if a == 7 {
		return "nodeList"
	}
	if a == 8 {
		return "number"
	}
	if a == 9 {
		return "string"
	}
	if a == 10 {
		return "computedString"
	}
	if a == 11 {
		return "token"
	}
	if a == 12 {
		return "tokenList"
	}
	if a == 13 {
		return "domRelation"
	}
	if a == 14 {
		return "role"
	}
	if a == 15 {
		return "internalRole"
	}
	if a == 16 {
		return "valueUndefined"
	}
	panic(fmt.Errorf("Invalid AXValueType %d", a))
}

/*
AXValueSourceType is an enum of possible property sources.
*/
type AXValueSourceType int

const (
	_attribute AXValueSourceType = iota
	_implicit
	_style
	_contents
	_placeholder
	_relatedElement
)

func (a AXValueSourceType) String() string {
	if a == 0 {
		return "attribute"
	}
	if a == 1 {
		return "implicit"
	}
	if a == 2 {
		return "style"
	}
	if a == 3 {
		return "contents"
	}
	if a == 4 {
		return "placeholder"
	}
	if a == 5 {
		return "relatedElement"
	}
	panic(fmt.Errorf("Invalid AXValueSourceType %d", a))
}

/*
AXValueNativeSourceType is an enum of possible native property sources (as a subtype of a particular
AXValueSourceType).
*/
type AXValueNativeSourceType int

const (
	_figcaption AXValueNativeSourceType = iota
	_label
	_labelfor
	_labelwrapped
	_legend
	_tablecaption
	_title
	_other
)

func (a AXValueNativeSourceType) String() string {
	if a == 0 {
		return "figcaption"
	}
	if a == 1 {
		return "label"
	}
	if a == 2 {
		return "labelfor"
	}
	if a == 3 {
		return "labelwrapped"
	}
	if a == 4 {
		return "legend"
	}
	if a == 5 {
		return "tablecaption"
	}
	if a == 6 {
		return "title"
	}
	if a == 7 {
		return "other"
	}
	panic(fmt.Errorf("Invalid AXValueNativeSourceType %d", a))
}

/*
AXValueSource is a single source for a computed AX property.
*/
type AXValueSource struct {
	// What type of source this is.
	Type AXValueSourceType `json:"type"`

	// Optional. The value of this property source.
	Value AXValue `json:"value,omitempty"`

	// Optional. The name of the relevant attribute, if any.
	Attribute string `json:"attribute,omitempty"`

	// Optional. The value of the relevant attribute, if any.
	AttributeValue AXValue `json:"attributeValue,omitempty"`

	// Optional. Whether this source is superseded by a higher priority source.
	Superseded bool `json:"superseded,omitempty"`

	// Optional. The native markup source for this value, e.g. a element.
	NativeSource AXValueNativeSourceType `json:"nativeSource,omitempty"`

	// Optional. The value, such as a node or node list, of the native source.
	NativeSourceValue AXValue `json:"nativeSourceValue,omitempty"`

	// Optional. Whether the value for this property is invalid.
	Invalid bool `json:"invalid,omitempty"`

	// Optional. Reason for the value being invalid, if it is.
	InvalidReason string `json:"invalidReason,omitempty"`
}

/*
AXRelatedNode represents a related node
*/
type AXRelatedNode struct {
	// The BackendNodeId of the related DOM node.
	BackendDOMNodeID DOM.BackendNodeID `json:"backendDOMNodeId"`

	// Optional. The IDRef value provided, if any.
	IDRef string `json:"idref,omitempty"`

	// Optional. The text alternative of this node in the current context.
	Text string `json:"text,omitempty"`
}

/*
AXProperty represents a property
*/
type AXProperty struct {
	// The name of this property.
	Name AXPropertyName `json:"name"`

	// The value of this property.
	Value AXValue `json:"value"`
}

/*
AXValue is a single computed AX property.
*/
type AXValue struct {
	// The type of this value.
	Type AXValueType `json:"type"`

	// Optional. The computed value of this property.
	Value interface{} `json:"value,omitempty"`

	// Optional. One or more related nodes, if applicable.
	RelatedNodes []*AXRelatedNode `json:"relatedNodes,omitempty"`

	// Optional. The sources which contributed to the computation of this property.
	Sources []*AXValueSource `json:"sources,omitempty"`
}

/*
AXPropertyName holds
	- values of AXProperty name: from 'busy' to 'roledescription'
	- states which apply to every AX node, from 'live' to 'root'
	- attributes which apply to nodes in live regions, from 'autocomplete' to 'valuetext'
	- attributes which apply to widgets, from 'checked' to 'selected'
	- states which apply to widgets, from 'activedescendant' to 'owns'
	- relationships between elements other than parent/child/sibling.
*/
type AXPropertyName int

func (a AXPropertyName) String() string {
	if a == 0 {
		return "busy"
	}
	if a == 1 {
		return "disabled"
	}
	if a == 2 {
		return "hidden"
	}
	if a == 3 {
		return "hiddenRoot"
	}
	if a == 4 {
		return "invalid"
	}
	if a == 5 {
		return "keyshortcuts"
	}
	if a == 6 {
		return "roledescription"
	}
	if a == 7 {
		return "live"
	}
	if a == 8 {
		return "atomic"
	}
	if a == 9 {
		return "relevant"
	}
	if a == 10 {
		return "root"
	}
	if a == 11 {
		return "autocomplete"
	}
	if a == 12 {
		return "haspopup"
	}
	if a == 13 {
		return "level"
	}
	if a == 14 {
		return "multiselectable"
	}
	if a == 15 {
		return "orientation"
	}
	if a == 16 {
		return "multiline"
	}
	if a == 17 {
		return "readonly"
	}
	if a == 18 {
		return "required"
	}
	if a == 19 {
		return "valuemin"
	}
	if a == 20 {
		return "valuemax"
	}
	if a == 21 {
		return "valuetext"
	}
	if a == 22 {
		return "checked"
	}
	if a == 23 {
		return "expanded"
	}
	if a == 24 {
		return "modal"
	}
	if a == 25 {
		return "pressed"
	}
	if a == 26 {
		return "selected"
	}
	if a == 27 {
		return "activedescendant"
	}
	if a == 28 {
		return "controls"
	}
	if a == 29 {
		return "describedby"
	}
	if a == 30 {
		return "details"
	}
	if a == 31 {
		return "errormessage"
	}
	if a == 32 {
		return "flowto"
	}
	if a == 33 {
		return "labelledby"
	}
	if a == 34 {
		return "owns"
	}
	panic(fmt.Errorf("Invalid AXPropertyName %d", a))
}

const (
	_busy AXPropertyName = iota
	_disabled
	_hidden
	_hiddenRoot
	_invalid
	_keyshortcuts
	_roledescription
	_live
	_atomic
	_relevant
	_root
	_autocomplete
	_haspopup
	_level
	_multiselectable
	_orientation
	_multiline
	_readonly
	_required
	_valuemin
	_valuemax
	_valuetext
	_checked
	_expanded
	_modal
	_pressed
	_selected
	_activedescendant
	_controls
	_describedby
	_details
	_errormessage
	_flowto
	_labelledby
	_owns
)

/*
AXNode is a node in the accessibility tree.
*/
type AXNode struct {
	// Unique identifier for this node.
	NodeID AXNodeID `json:"nodeId"`

	// Whether this node is ignored for accessibility.
	Ignored bool `json:"ignored"`

	// Optional. Collection of reasons why this node is hidden.
	IgnoredReasons []*AXProperty `json:"ignoredReasons,omitempty"`

	// Optional. This Node's role, whether explicit or implicit.
	Role AXValue `json:"role,omitempty"`

	// Optional. The accessible name for this Node.
	Name AXValue `json:"name,omitempty"`

	// Optional. The accessible description for this Node.
	Description AXValue `json:"description,omitempty"`

	// Optional. The value for this Node.
	Value AXValue `json:"value,omitempty"`

	// Optional. All other properties.
	Properties []*AXProperty `json:"properties,omitempty"`

	// Optional. IDs for each of this node's child nodes.
	ChildIds []AXNodeID `json:"childIds,omitempty"`

	// Optional. The backend ID for the associated DOM node, if any.
	BackendDOMNodeID DOM.BackendNodeID `json:"backendDOMNodeId,omitempty"`
}

/*
Package accessibility provides type definitions for use with the Chrome Accessibility protocol

https://chromedevtools.github.io/devtools-protocol/tot/Accessibility/
*/
package accessibility

import (
	"github.com/mkenney/go-chrome/protocol/dom"
)

/*
PartialAXTreeParams represents Accessibility.partialAXTree parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Accessibility/#method-getPartialAXTree
*/
type PartialAXTreeParams struct {
	// ID of the node to get the partial accessibility tree for.
	NodeID dom.NodeID `json:"nodeId"`

	// Optional. Whether to fetch this nodes ancestors, siblings and children.
	// Defaults to true.
	FetchRelatives bool `json:"fetchRelatives,omitempty"`
}

/*
PartialAXTreeResult represents the result of calls to Accessibility.partialAXTree.

https://chromedevtools.github.io/devtools-protocol/tot/Accessibility/#method-getPartialAXTree
*/
type PartialAXTreeResult struct {
	// The `Accessibility.AXNode` for this DOM node, if it exists, plus its
	// ancestors, siblings and children, if requested.
	Nodes []*AXNode `json:"nodes"`
}

/*
AXNode is a node in the accessibility tree.

https://chromedevtools.github.io/devtools-protocol/tot/Accessibility/#type-AXNode
*/
type AXNode struct {
	// Unique identifier for this node.
	NodeID AXNodeID `json:"nodeId"`

	// Whether this node is ignored for accessibility.
	Ignored bool `json:"ignored"`

	// Optional. Collection of reasons why this node is hidden.
	IgnoredReasons []*AXProperty `json:"ignoredReasons,omitempty"`

	// Optional. This Node's role, whether explicit or implicit.
	Role *AXValue `json:"role,omitempty"`

	// Optional. The accessible name for this Node.
	Name *AXValue `json:"name,omitempty"`

	// Optional. The accessible description for this Node.
	Description *AXValue `json:"description,omitempty"`

	// Optional. The value for this Node.
	Value *AXValue `json:"value,omitempty"`

	// Optional. All other properties.
	Properties []*AXProperty `json:"properties,omitempty"`

	// Optional. IDs for each of this node's child nodes.
	ChildIDs []AXNodeID `json:"childIds,omitempty"`

	// Optional. The backend ID for the associated DOM node, if any.
	BackendDOMNodeID dom.BackendNodeID `json:"backendDOMNodeId,omitempty"`
}

/*
AXNodeID is the unique accessibility node identifier.

https://chromedevtools.github.io/devtools-protocol/tot/Accessibility/#type-AXNodeId
*/
type AXNodeID string

/*
AXProperty represents a property

https://chromedevtools.github.io/devtools-protocol/tot/Accessibility/#type-AXProperty
*/
type AXProperty struct {
	// The name of this property.
	Name AXPropertyName `json:"name"`

	// The value of this property.
	Value *AXValue `json:"value"`
}

/*
AXPropertyName holds
	- values of AXProperty name: from 'busy' to 'roledescription'
	- states which apply to every AX node, from 'live' to 'root'
	- attributes which apply to nodes in live regions, from 'autocomplete' to
	'valuetext'
	- attributes which apply to widgets, from 'checked' to 'selected'
	- states which apply to widgets, from 'activedescendant' to 'owns'
	- relationships between elements other than parent/child/sibling.

https://chromedevtools.github.io/devtools-protocol/tot/Accessibility/#type-AXPropertyName
*/
type AXPropertyName string

/*
AXRelatedNode represents a related node

https://chromedevtools.github.io/devtools-protocol/tot/Accessibility/#type-AXRelatedNode
*/
type AXRelatedNode struct {
	// The BackendNodeId of the related DOM node.
	BackendDOMNodeID dom.BackendNodeID `json:"backendDOMNodeId"`

	// Optional. The IDRef value provided, if any.
	IDRef string `json:"idref,omitempty"`

	// Optional. The text alternative of this node in the current context.
	Text string `json:"text,omitempty"`
}

/*
AXValue is a single computed AX property.

https://chromedevtools.github.io/devtools-protocol/tot/Accessibility/#type-AXValue
*/
type AXValue struct {
	// The type of this value.
	Type AXValueType `json:"type"`

	// Optional. The computed value of this property.
	Value interface{} `json:"value,omitempty"`

	// Optional. One or more related nodes, if applicable.
	RelatedNodes []*AXRelatedNode `json:"relatedNodes,omitempty"`

	// Optional. The sources which contributed to the computation of this
	// property.
	Sources []*AXValueSource `json:"sources,omitempty"`
}

/*
AXValueNativeSourceType is an enum of possible native property sources (as a subtype of a particular
AXValueSourceType).

https://chromedevtools.github.io/devtools-protocol/tot/Accessibility/#type-AXValueNativeSourceType
*/
type AXValueNativeSourceType string

/*
AXValueSource is a single source for a computed AX property.
*/
type AXValueSource struct {
	// What type of source this is.
	Type AXValueSourceType `json:"type"`

	// Optional. The value of this property source.
	Value *AXValue `json:"value,omitempty"`

	// Optional. The name of the relevant attribute, if any.
	Attribute string `json:"attribute,omitempty"`

	// Optional. The value of the relevant attribute, if any.
	AttributeValue *AXValue `json:"attributeValue,omitempty"`

	// Optional. Whether this source is superseded by a higher priority source.
	Superseded bool `json:"superseded,omitempty"`

	// Optional. The native markup source for this value, e.g. a element.
	NativeSource AXValueNativeSourceType `json:"nativeSource,omitempty"`

	// Optional. The value, such as a node or node list, of the native source.
	NativeSourceValue *AXValue `json:"nativeSourceValue,omitempty"`

	// Optional. Whether the value for this property is invalid.
	Invalid bool `json:"invalid,omitempty"`

	// Optional. Reason for the value being invalid, if it is.
	InvalidReason string `json:"invalidReason,omitempty"`
}

/*
AXValueSourceType is an enum of possible property sources.

https://chromedevtools.github.io/devtools-protocol/tot/Accessibility/#type-AXValueSourceType
*/
type AXValueSourceType string

/*
AXValueType is an enum of possible property types.

https://chromedevtools.github.io/devtools-protocol/tot/Accessibility/#type-AXValueType
*/
type AXValueType string

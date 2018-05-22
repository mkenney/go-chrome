/*
Package dom provides type definitions for use with the Chrome DOM protocol

https://chromedevtools.github.io/devtools-protocol/tot/DOM/
*/
package dom

import (
	"github.com/mkenney/go-chrome/tot/cdtp/page"
)

/*
NodeID is a unique DOM node identifier.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#type-NodeId
*/
type NodeID int

/*
BackendNodeID is a unique DOM node identifier used to reference a node that may not have been pushed
to the front-end.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#type-BackendNodeId
*/
type BackendNodeID int

/*
BackendNode is a backend node with a friendly name.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#type-BackendNode
*/
type BackendNode struct {
	// Node's nodeType.
	NodeType int `json:"nodeType"`

	// Node's nodeName.
	NodeName string `json:"nodeName"`

	// Node's ID
	BackendNodeID BackendNodeID `json:"backendNodeId"`
}

/*
PseudoType is a pseudo element type.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#type-PseudoType
*/
type PseudoType string

/*
ShadowRootType is a shadow root type.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#type-ShadowRootType
*/
type ShadowRootType string

/*
Node is a base node mirror type. DOM interaction is implemented in terms of mirror objects that
represent the actual DOM nodes.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#type-Node
*/
type Node struct {
	// Node identifier that is passed into the rest of the DOM messages as the
	// nodeId. Backend will only push node with given id once. It is aware of
	// all requested nodes and will only fire DOM events for nodes known to the
	// client.
	NodeID NodeID `json:"nodeId"`

	// Optional. The ID of the parent node if any.
	ParentID NodeID `json:"parentId,omitempty"`

	// The BackendNodeID for this node.
	BackendNodeID BackendNodeID `json:"backendNodeId"`

	// Node's nodeType.
	NodeType int `json:"nodeType"`

	// Node's nodeName.
	NodeName string `json:"nodeName"`

	// Node's localName.
	LocalName string `json:"localName"`

	// Node's nodeValue.
	NodeValue string `json:"nodeValue"`

	// Optional. Child count for Container nodes.
	ChildNodeCount int `json:"childNodeCount,omitempty"`

	// Optional. Child nodes of this node when requested with children.
	Children []*Node `json:"children,omitempty"`

	// Optional. Attributes of the Element node in the form of flat array
	// [name1, value1, name2, value2].
	Attributes []string `json:"attributes,omitempty"`

	// Optional. Document URL that Document or FrameOwner node points to.
	DocumentURL string `json:"documentURL,omitempty"`

	// Optional. Base URL that Document or FrameOwner node uses for URL
	// completion.
	BaseURL string `json:"baseURL,omitempty"`

	// Optional. DocumentType's publicId.
	PublicID string `json:"publicId,omitempty"`

	// Optional. DocumentType's systemId.
	SystemID string `json:"systemId,omitempty"`

	// Optional. DocumentType's internalSubset.
	InternalSubset string `json:"internalSubset,omitempty"`

	// Optional. Document's XML version in case of XML documents.
	XMLVersion string `json:"xmlVersion,omitempty"`

	// Optional. Attr's name.
	Name string `json:"name,omitempty"`

	// Optional. Attr's value.
	Value string `json:"value,omitempty"`

	// Optional. Pseudo element type for this node.
	PseudoType PseudoType `json:"pseudoType,omitempty"`

	// Optional. Shadow root type.
	ShadowRootType ShadowRootType `json:"shadowRootType,omitempty"`

	// Optional. Frame ID for frame owner elements.
	FrameID page.FrameID `json:"frameId,omitempty"`

	// Optional. Content document for frame owner elements.
	ContentDocument *Node `json:"contentDocument,omitempty"`

	// Optional. Shadow root list for given element host.
	ShadowRoots []*Node `json:"shadowRoots,omitempty"`

	// Optional. Content document fragment for template elements.
	TemplateContent *Node `json:"templateContent,omitempty"`

	// Optional. Pseudo elements associated with this node.
	PseudoElements []*Node `json:"pseudoElements,omitempty"`

	// Optional. Import document for the HTMLImport links.
	ImportedDocument *Node `json:"importedDocument,omitempty"`

	// Optional. Distributed nodes for given insertion point.
	DistributedNodes []*BackendNode `json:"distributedNodes,omitempty"`

	// Optional. Whether the node is SVG.
	IsSVG bool `json:"isSVG,omitempty"`
}

/*
RGBA is a structure holding an RGBA color.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#type-RGBA
*/
type RGBA struct {
	// The red component, in the [0-255] range.
	R int `json:"r"`

	// The green component, in the [0-255] range.
	G int `json:"g"`

	// The blue component, in the [0-255] range.
	B int `json:"b"`

	// Optional. The alpha component, in the [0-1] range (default: 1).
	A int `json:"a,omitempty"`
}

/*
Quad is an array of quad vertices, x immediately followed by y for each point, points clock-wise.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#type-Quad
*/
type Quad [2]float64

/*
BoxModel represents the box model.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#type-BoxModel
*/
type BoxModel struct {
	// Content box.
	Content Quad `json:"content"`

	// Padding box.
	Padding Quad `json:"padding"`

	// Border box.
	Border Quad `json:"border"`

	// Margin box.
	Margin Quad `json:"margin"`

	// Node width.
	Width int64 `json:"width"`

	// Node height.
	Height int64 `json:"height"`

	// Optional. Shape outside coordinates.
	ShapeOutside *ShapeOutsideInfo `json:"shapeOutside,omitempty"`
}

/*
ShapeOutsideInfo represents the CSS Shape Outside details.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#type-ShapeOutsideInfo
*/
type ShapeOutsideInfo struct {
	// Shape bounds.
	Bounds Quad `json:"bounds"`

	// Shape coordinate details.
	Shape []interface{} `json:"shape"`

	// Margin shape bounds.
	MarginShape []interface{} `json:"marginShape"`
}

/*
Rect defines a rectangle

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#type-Rect
*/
type Rect struct {
	// X coordinate.
	X int64 `json:"x"`

	// Y coordinate.
	Y int64 `json:"y"`

	// Rectangle width.
	Width int64 `json:"width"`

	// Rectangle height.
	Height int64 `json:"height"`
}

package DOM

import (
	Page "app/chrome/protocol/page"
	"fmt"
)

/*
NodeID is a unique DOM node identifier.
*/
type NodeID int

/*
BackendNodeID is a unique DOM node identifier used to reference a node that may not have been pushed
to the front-end.
*/
type BackendNodeID int

/*
BackendNode is a backend node with a friendly name.
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
*/
type PseudoType int

const (
	_firstLine PseudoType = iota
	_firstLetter
	_before
	_after
	_backdrop
	_selection
	_firstLineInherited
	_scrollbar
	_scrollbarThumb
	_scrollbarButton
	_scrollbarTrack
	_scrollbarTrackPiece
	_scrollbarCorner
	_resizer
	_inputListButton
)

func (a PseudoType) String() string {
	if a == 0 {
		return "first-line"
	}
	if a == 1 {
		return "first-letter"
	}
	if a == 2 {
		return "before"
	}
	if a == 3 {
		return "after"
	}
	if a == 4 {
		return "backdrop"
	}
	if a == 5 {
		return "selection"
	}
	if a == 6 {
		return "first-lineinherited"
	}
	if a == 7 {
		return "scrollbar"
	}
	if a == 8 {
		return "scrollbar-thumb"
	}
	if a == 9 {
		return "scrollbar-button"
	}
	if a == 10 {
		return "scrollbar-track"
	}
	if a == 11 {
		return "scrollbar-track-piece"
	}
	if a == 12 {
		return "scrollbar-corner"
	}
	if a == 13 {
		return "resizer"
	}
	if a == 14 {
		return "inputList-button"
	}
	panic(fmt.Errorf("Invalid PseudoType %d", a))
}

/*
ShadowRootType is a shadow root type.
*/
type ShadowRootType int

const (
	_userAgent ShadowRootType = iota
	_open
	_closed
)

func (a ShadowRootType) String() string {
	if a == 0 {
		return "user-agent"
	}
	if a == 1 {
		return "open"
	}
	if a == 2 {
		return "closed"
	}
	panic(fmt.Errorf("Invalid ShadowRootType %d", a))
}

/*
Node is a base node mirror type. DOM interaction is implemented in terms of mirror objects that
represent the actual DOM nodes.
*/
type Node struct {
	// Node identifier that is passed into the rest of the DOM messages as the nodeId. Backend will
	// only push node with given id once. It is aware of all requested nodes and will only fire DOM
	// events for nodes known to the client.
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
	FrameID Page.FrameID `json:"frameId,omitempty"`

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
*/
type Quad [2]int

/*
BoxModel represents the box model.
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
	Width int `json:"width"`

	// Node height.
	Height int `json:"height"`

	// Optional. Shape outside coordinates.
	ShapeOutside ShapeOutsideInfo `json:"shapeOutside,omitempty"`
}

/*
ShapeOutsideInfo represents the CSS Shape Outside details.
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
*/
type Rect struct {
	// X coordinate.
	X int `json:"x"`

	// Y coordinate.
	Y int `json:"y"`

	// Rectangle width.
	Width int `json:"width"`

	// Rectangle height.
	Height int `json:"height"`
}

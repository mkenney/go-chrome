package DOMSnapshot

import (
	DOM "app/chrome/dom"
	Page "app/chrome/page"
)

/*
GetSnapshotParams represents DOMSnapshot.getSnapshot parameters.
*/
type GetSnapshotParams struct {
	// Whitelist of computed styles to return.
	ComputedStyleWhitelist []string `json:"computedStyleWhitelist"`
}

/*
DOMNode is a Node in the DOM tree.
*/
type DOMNode struct {
	// Node's nodeType.
	NodeType int `json:"nodeType"`

	// Node's nodeName.
	NodeName string `json:"nodeName"`

	// Node's nodeValue.
	NodeValue string `json:"nodeValue"`

	// Optional. Only set for textarea elements, contains the text value.
	TextValue string `json:"textValue,omitempty"`

	// Optional. Only set for input elements, contains the input's associated text value.
	InputValue string `json:"inputValue,omitempty"`

	// Optional. Only set for radio and checkbox input elements, indicates if the element has been
	// checked.
	InputChecked bool `json:"inputChecked,omitempty"`

	// Optional. Only set for option elements, indicates if the element has been selected.
	OptionSelected bool `json:"optionSelected,omitempty"`

	// Optional. Node's ID, corresponds to DOM.Node.backendNodeID.
	BackendNodeID DOM.BackendNodeID `json:"backendNodeId,omitempty"`

	// Optional. The indexes of the node's child nodes in the domNodes array returned by
	// getSnapshot, if any.
	ChildNodeIndexes []int `json:"childNodeIndexes,omitempty"`

	// Optional. Attributes of an Element node.
	Attributes []*NameValue `json:"attributes,omitempty"`

	// Optional. Indexes of pseudo elements associated with this node in the domNodes array returned
	// by getSnapshot, if any.
	PseudoElementIndexes []int `json:"pseudoElementIndexes,omitempty"`

	// Optional. The index of the node's related layout tree node in the layoutTreeNodes array
	// returned by getSnapshot, if any.
	LayoutNodeIndex int `json:"layoutNodeIndex,omitempty"`

	// Optional. Document URL that Document or FrameOwner node points to.
	DocumentURL string `json:"documentURL,omitempty"`

	// Optional. Base URL that Document or FrameOwner node uses for URL completion.
	BaseURL string `json:"baseURL,omitempty"`

	// Optional. Only set for documents, contains the document's content language.
	ContentLanguage string `json:"contentLanguage,omitempty"`

	// Optional. Only set for documents, contains the document's character set encoding.
	DocumentEncoding string `json:"documentEncoding,omitempty"`

	// Optional. DocumentType node's publicId.
	PublicID string `json:"publicId,omitempty"`

	// Optional. DocumentType node's systemId.
	SystemID string `json:"systemId,omitempty"`

	// Optional. Frame ID for frame owner elements and also for the document node.
	FrameID Page.FrameID `json:"frameId,omitempty"`

	// Optional. The index of a frame owner element's content document in the domNodes array
	// returned by getSnapshot, if any.
	ContentDocumentIndex int `json:"contentDocumentIndex,omitempty"`

	// Optional. Index of the imported document's node of a link element in the domNodes array
	// returned by getSnapshot, if any.
	ImportedDocumentIndex int `json:"importedDocumentIndex,omitempty"`

	// Optional. Index of the content node of a template element in the domNodes array returned by
	// getSnapshot.
	TemplateContentIndex int `json:"templateContentIndex,omitempty"`

	// Optional. Type of a pseudo element node.
	PseudoType *DOM.PseudoType `json:"pseudoType,omitempty"`

	// Optional. Whether this DOM node responds to mouse clicks. This includes nodes that have had
	// click event listeners attached via JavaScript as well as anchor tags that naturally navigate
	// when clicked.
	IsClickable bool `json:"isClickable,omitempty"`
}

/*
InlineTextBox contains details of post layout rendered text positions. The exact layout should not
be regarded as stable and may change between versions.
*/
type InlineTextBox struct {
	// The absolute position bounding box.
	BoundingBox *DOM.Rect `json:"boundingBox"`

	// The starting index in characters, for this post layout textbox substring.
	StartCharacterIndex int `json:"startCharacterIndex"`

	// The number of characters in this post layout textbox substring.
	NumCharacters int `json:"numCharacters"`
}

/*
LayoutTreeNode contains details of an element in the DOM tree with a LayoutObject.
*/
type LayoutTreeNode struct {
	// The index of the related DOM node in the domNodes array returned by getSnapshot.
	DomNodeIndex int `json:"domNodeIndex"`

	// The absolute position bounding box.
	BoundingBox *DOM.Rect `json:"boundingBox"`

	// Optional. Contents of the LayoutText, if any.
	LayoutText string `json:"layoutText,omitempty"`

	// Optional. The post-layout inline text nodes, if any.
	InlineTextNodes []*InlineTextBox `json:"inlineTextNodes,omitempty"`

	// Optional. Index into the computedStyles array returned by getSnapshot.
	StyleIndex int `json:"styleIndex,omitempty"`
}

/*
ComputedStyle is a subset of the full ComputedStyle as defined by the request whitelist.
*/
type ComputedStyle struct {
	// Name/value pairs of computed style properties.
	Properties []*NameValue `json:"properties"`
}

/*
NameValue is a name/value pair.
*/
type NameValue struct {
	// Attribute/property name.
	Name string `json:"name"`

	// Attribute/property value.
	Value string `json:"value"`
}

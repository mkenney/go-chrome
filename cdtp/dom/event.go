package dom

/*
AttributeModifiedEvent represents DOM.attributeModified event data.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-attributeModified
*/
type AttributeModifiedEvent struct {
	// ID of the node that has changed.
	NodeID NodeID `json:"nodeId"`

	// Attribute name.
	Name string `json:"name"`

	// Attribute value.
	Value string `json:"value"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
AttributeRemovedEvent represents DOM.attributeRemoved event data.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-attributeRemoved
*/
type AttributeRemovedEvent struct {
	// ID of the node that has changed.
	NodeID NodeID `json:"nodeId"`

	// Attribute name.
	Name string `json:"name"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
CharacterDataModifiedEvent represents DOM.characterDataModified event data.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-characterDataModified
*/
type CharacterDataModifiedEvent struct {
	// ID of the node that has changed.
	NodeID NodeID `json:"nodeId"`

	// New text value.
	CharacterData string `json:"characterData"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
ChildNodeCountUpdatedEvent represents DOM.childNodeCountUpdated event data.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-childNodeCountUpdated
*/
type ChildNodeCountUpdatedEvent struct {
	// ID of the node that has changed.
	ParentNodeID NodeID `json:"parentNodeId"`

	// If of the previous siblint.
	PreviousNodeID NodeID `json:"previousNodeId"`

	// Inserted node data.
	Node *Node `json:"node"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
ChildNodeInsertedEvent represents DOM.childNodeInserted event data.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-childNodeInserted
*/
type ChildNodeInsertedEvent struct {
	// ID of the node that has changed.
	ParentNodeID NodeID `json:"parentNodeId"`

	// If of the previous siblint.
	PreviousNodeID NodeID `json:"previousNodeId"`

	// Inserted node data.
	Node *Node `json:"node"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
ChildNodeRemovedEvent represents DOM.childNodeRemoved event data.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-childNodeRemoved
*/
type ChildNodeRemovedEvent struct {
	// Parent ID.
	ParentNodeID NodeID `json:"parentNodeId"`

	// ID of the node that has been removed.
	NodeID NodeID `json:"nodeId"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
DistributedNodesUpdatedEvent represents DOM.distributedNodesUpdated event data.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-distributedNodesUpdated
*/
type DistributedNodesUpdatedEvent struct {
	// Insertion point where distributed nodes were updated.
	InsertionPointID NodeID `json:"insertionPointId"`

	// Distributed nodes for given insertion point.
	DistributedNodes []*BackendNode `json:"distributedNodes"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
DocumentUpdatedEvent represents DOM.documentUpdated event data.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-documentUpdated
*/
type DocumentUpdatedEvent struct {
	// Error information related to this event
	Err error `json:"-"`
}

/*
InlineStyleInvalidatedEvent represents DOM.inlineStyleInvalidated event data.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-inlineStyleInvalidated
*/
type InlineStyleInvalidatedEvent struct {
	// IDs of the nodes for which the inline styles have been invalidated.
	NodeIDs []NodeID `json:"nodeIds"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
PseudoElementAddedEvent represents DOM.pseudoElementAdded event data.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-pseudoElementAdded
*/
type PseudoElementAddedEvent struct {
	// Pseudo element's parent element ID.
	ParentID NodeID `json:"parentId"`

	// The added pseudo element.
	PseudoElement *Node `json:"pseudoElement"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
PseudoElementRemovedEvent represents DOM.pseudoElementRemoved event data.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-pseudoElementRemoved
*/
type PseudoElementRemovedEvent struct {
	// Pseudo element's parent element ID.
	ParentID NodeID `json:"parentId"`

	// The removed pseudo element ID.
	PseudoElementID NodeID `json:"pseudoElementId"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
SetChildNodesEvent represents DOM.setChildNodes event data.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-setChildNodes
*/
type SetChildNodesEvent struct {
	// Parent node ID to populate with children.
	ParentID NodeID `json:"parentId"`

	// Child nodes array.
	Nodes []*Node `json:"nodes"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
ShadowRootPoppedEvent represents DOM.shadowRootPopped event data.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-shadowRootPopped
*/
type ShadowRootPoppedEvent struct {
	// Host element ID.
	HostID NodeID `json:"hostId"`

	// Shadow root ID.
	RootID NodeID `json:"rootId"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
ShadowRootPushedEvent represents DOM.shadowRootPushed event data.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-shadowRootPushed
*/
type ShadowRootPushedEvent struct {
	// Host element ID.
	HostID NodeID `json:"hostId"`

	// Shadow root.
	Root *Node `json:"root"`

	// Error information related to this event
	Err error `json:"-"`
}

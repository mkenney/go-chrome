package accessibility

import (
	"github.com/mkenney/go-chrome/tot/dom"
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

	// Error information related to executing this method
	Err error `json:"-"`
}

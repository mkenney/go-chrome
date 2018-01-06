package dom

import (
	"github.com/mkenney/go-chrome/cdtp/runtime"
)

/*
CollectClassNamesFromSubtreeParams represents DOM.collectClassNamesFromSubtree parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-collectClassNamesFromSubtree
*/
type CollectClassNamesFromSubtreeParams struct {
	// ID of the node to collect class names.
	NodeID NodeID `json:"nodeId"`
}

/*
CollectClassNamesFromSubtreeResult represents the result of calls to
DOM.collectClassNamesFromSubtree.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-collectClassNamesFromSubtree
*/
type CollectClassNamesFromSubtreeResult struct {
	// Class name list.
	ClassNames []string `json:"classNames"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
CopyToParams represents DOM.copyTo parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-copyTo
*/
type CopyToParams struct {
	// ID of the node to copy.
	NodeID NodeID `json:"nodeId"`

	// ID of the element to drop the copy into.
	TargetNodeID NodeID `json:"targetNodeId"`

	// Optional. Drop the copy before this node (if absent, the copy becomes the
	// last child of targetNodeId).
	InsertBeforeNodeID NodeID `json:"insertBeforeNodeId,omitempty"`
}

/*
CopyToResult represents the result of calls to DOM.copyTo.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-copyTo
*/
type CopyToResult struct {
	// ID of the new cloned node.
	NodeID NodeID `json:"nodeId"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
DescribeNodeParams represents DOM.describeNode parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-describeNode
*/
type DescribeNodeParams struct {
	// Optional. ID of the node.
	NodeID NodeID `json:"nodeId,omitempty"`

	// Optional. ID of the backend node.
	BackendNodeID BackendNodeID `json:"backendNodeId,omitempty"`

	// Optional. JavaScript object ID of the node wrapper.
	ObjectID runtime.RemoteObjectID `json:"objectId,omitempty"`

	// Optional. The maximum depth at which children should be retrieved,
	// defaults to 1. Use -1 for the entire subtree or provide an integer larger
	// than 0.
	Depth int `json:"depth,omitempty"`

	// Optional. Whether or not iframes and shadow roots should be traversed
	// when returning the subtree (default is false).
	Pierce bool `json:"pierce,omitempty"`
}

/*
DescribeNodeResult represents the result of calls to DOM.describeNode.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-describeNode
*/
type DescribeNodeResult struct {
	// ID of the new cloned node.
	NodeID NodeID `json:"nodeId"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
DisableResult represents the result of calls to DOM.disable.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-disable
*/
type DisableResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
DiscardSearchResultsParams represents DOM.discardSearchResults parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-discardSearchResults
*/
type DiscardSearchResultsParams struct {
	// Node description.
	Node *Node `json:"node"`
}

/*
DiscardSearchResultsResult represents the result of calls to DOM.discardSearchResults.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-discardSearchResults
*/
type DiscardSearchResultsResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
EnableResult represents the result of calls to DOM.enable.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-enable
*/
type EnableResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
FocusParams represents DOM.focus parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-focus
*/
type FocusParams struct {
	// Optional. ID of the node.
	NodeID NodeID `json:"nodeId,omitempty"`

	// Optional. ID of the backend node.
	BackendNodeID BackendNodeID `json:"backendNodeId,omitempty"`

	// Optional. JavaScript object ID of the node wrapper.
	ObjectID runtime.RemoteObjectID `json:"objectId,omitempty"`
}

/*
FocusResult represents the result of calls to DOM.focus.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-focus
*/
type FocusResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
GetAttributesParams represents DOM.getAttributes parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getAttributes
*/
type GetAttributesParams struct {
	// ID  of the node to retrieve attibutes for.
	NodeID NodeID `json:"nodeId"`
}

/*
GetAttributesResult represents the result of calls to DOM.getAttributes.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getAttributes
*/
type GetAttributesResult struct {
	// An interleaved array of node attribute names and values.
	Attributes []string `json:"attributes"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
GetBoxModelParams represents DOM.getBoxModel parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getBoxModel
*/
type GetBoxModelParams struct {
	// Optional. ID of the node.
	NodeID NodeID `json:"nodeId,omitempty"`

	// Optional. ID of the backend node.
	BackendNodeID BackendNodeID `json:"backendNodeId,omitempty"`

	// Optional. JavaScript object ID of the node wrapper.
	ObjectID runtime.RemoteObjectID `json:"objectId,omitempty"`
}

/*
GetBoxModelResult represents the result of calls to DOM.getBoxModel.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getBoxModel
*/
type GetBoxModelResult struct {
	// Box model for the node.
	Model *BoxModel `json:"model"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
GetDocumentParams represents DOM.getDocument parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getDocument
*/
type GetDocumentParams struct {
	// Optional. The maximum depth at which children should be retrieved,
	// defaults to 1. Use -1 for the entire subtree or provide an integer larger
	// than 0.
	Depth int `json:"depth,omitempty"`

	// Optional. Whether or not iframes and shadow roots should be traversed
	// when returning the subtree (default is false).
	Pierce bool `json:"pierce,omitempty"`
}

/*
GetDocumentResult represents the result of calls to DOM.getDocument.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getDocument
*/
type GetDocumentResult struct {
	// Resulting node.
	Root *Node `json:"root"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
GetFlattenedDocumentParams represents DOM.getFlattenedDocument parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getFlattenedDocument
*/
type GetFlattenedDocumentParams struct {
	// Optional. The maximum depth at which children should be retrieved,
	// defaults to 1. Use -1 for the entire subtree or provide an integer larger
	// than 0.
	Depth int `json:"depth,omitempty"`

	// Optional. Whether or not iframes and shadow roots should be traversed
	// when returning the subtree (default is false).
	Pierce bool `json:"pierce,omitempty"`
}

/*
GetFlattenedDocumentResult represents the result of calls to DOM.getFlattenedDocument.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getFlattenedDocument
*/
type GetFlattenedDocumentResult struct {
	// Resulting nodes.
	Nodes []*Node `json:"nodes"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
GetNodeForLocationParams represents DOM.getNodeForLocation parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getNodeForLocation
*/
type GetNodeForLocationParams struct {
	// X coordinate.
	X int `json:"x"`

	// Y coordinate.
	Y int `json:"y"`

	// Optional. False to skip to the nearest non-UA shadow root ancestor
	// (default: false).
	IncludeUserAgentShadowDOM bool `json:"includeUserAgentShadowDOM,omitempty"`
}

/*
GetNodeForLocationResult represents the result of calls to DOM.getNodeForLocation.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getNodeForLocation
*/
type GetNodeForLocationResult struct {
	// ID of the node at given coordinates.
	NodeID NodeID `json:"nodeId"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
GetOuterHTMLParams represents DOM.getOuterHTML parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getOuterHTML
*/
type GetOuterHTMLParams struct {
	// Optional. ID of the node.
	NodeID NodeID `json:"nodeId,omitempty"`

	// Optional. ID of the backend node.
	BackendNodeID BackendNodeID `json:"backendNodeId,omitempty"`

	// Optional. JavaScript object ID of the node wrapper.
	ObjectID runtime.RemoteObjectID `json:"objectId,omitempty"`
}

/*
GetOuterHTMLResult represents the result of calls to DOM.getOuterHTML.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getOuterHTML
*/
type GetOuterHTMLResult struct {
	// Outer HTML markup.
	OuterHTML string `json:"outerHTML"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
GetRelayoutBoundaryParams represents DOM.getRelayoutBoundary parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getRelayoutBoundary
*/
type GetRelayoutBoundaryParams struct {
	// ID of the node.
	NodeID NodeID `json:"nodeId"`
}

/*
GetRelayoutBoundaryResult represents the result of calls to DOM.getRelayoutBoundary.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getRelayoutBoundary
*/
type GetRelayoutBoundaryResult struct {
	// Relayout boundary node ID for the given node.
	NodeID NodeID `json:"nodeId"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
GetSearchResultsParams represents DOM.getSearchResults parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getSearchResults
*/
type GetSearchResultsParams struct {
	// Unique search session identifier.
	SearchID string `json:"searchId"`

	// Start index of the search result to be returned.
	FromIndex int `json:"fromIndex"`

	// End index of the search result to be returned.
	ToIndex int `json:"toIndex"`
}

/*
GetSearchResultsResult represents the result of calls to DOM.getSearchResults.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getSearchResults
*/
type GetSearchResultsResult struct {
	// IDs of the search result nodes.
	NodeIDs []NodeID `json:"nodeIds"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
MarkUndoableStateResult represents the result of calls to DOM.markUndoableState.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-markUndoableState
*/
type MarkUndoableStateResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
MoveToParams represents DOM.moveTo parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-moveTo
*/
type MoveToParams struct {
	// ID of the node to move.
	NodeID NodeID `json:"nodeId"`

	// ID of the element to drop the moved node into.
	TargetNodeID NodeID `json:"targetNodeId"`

	// Optional. Drop node before this one (if absent, the moved node becomes
	// the last child of targetNodeId).
	InsertBeforeNodeID NodeID `json:"insertBeforeNodeId,omitempty"`
}

/*
MoveToResult represents the result of calls to DOM.moveTo.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-moveTo
*/
type MoveToResult struct {
	// New ID of the moved node.
	NodeID NodeID `json:"nodeId"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
PerformSearchParams represents DOM.performSearch parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-performSearch
*/
type PerformSearchParams struct {
	// Plain text or query selector or XPath search query.
	Query string `json:"query"`

	// Optional. True to search in user agent shadow DOM.
	IncludeUserAgentShadowDOM bool `json:"includeUserAgentShadowDOM,omitempty"`
}

/*
PerformSearchResult represents the result of calls to DOM.performSearch.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-performSearch
*/
type PerformSearchResult struct {
	// Unique search session identifier.
	SearchID string `json:"searchId"`

	// Number of search results.
	ResultCount int `json:"resultCount"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
PushNodeByPathToFrontendParams represents DOM.pushNodeByPathToFrontend parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-pushNodeByPathToFrontend
*/
type PushNodeByPathToFrontendParams struct {
	// Path to node in the proprietary format.
	Path string `json:"path"`
}

/*
PushNodeByPathToFrontendResult represents the result of calls to DOM.pushNodeByPathToFrontend.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-pushNodeByPathToFrontend
*/
type PushNodeByPathToFrontendResult struct {
	// ID of the node for given path.
	NodeID NodeID `json:"nodeId"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
PushNodesByBackendIDsToFrontendParams represents DOM.pushNodesByBackendIdsToFrontend parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-pushNodesByBackendIdsToFrontend
*/
type PushNodesByBackendIDsToFrontendParams struct {
	// The array of backend node IDs.
	BackendNodeIDs []BackendNodeID `json:"backendNodeIds"`
}

/*
PushNodesByBackendIDsToFrontendResult represents the result of calls to
DOM.pushNodesByBackendIdsToFrontend.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-pushNodesByBackendIdsToFrontend
*/
type PushNodesByBackendIDsToFrontendResult struct {
	// The array of IDs of pushed nodes that correspond to the backend IDs
	// specified in backendNodeIDs.
	NodeIDs []NodeID `json:"nodeIds"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
QuerySelectorParams represents DOM.querySelector parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-querySelector
*/
type QuerySelectorParams struct {
	// ID of the node to query.
	NodeID NodeID `json:"nodeId"`

	// Selector string.
	Selector string `json:"selector"`
}

/*
QuerySelectorResult represents the result of calls to DOM.querySelector.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-querySelector
*/
type QuerySelectorResult struct {
	// Query selector result.
	NodeID NodeID `json:"nodeId"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
QuerySelectorAllParams represents DOM.querySelectorAll parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-querySelectorAll
*/
type QuerySelectorAllParams struct {
	// ID of the node to query upon.
	NodeID NodeID `json:"nodeId"`

	// Selector string.
	Selector string `json:"selector"`
}

/*
QuerySelectorAllResult represents the result of calls to DOM.querySelectorAll.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-querySelectorAll
*/
type QuerySelectorAllResult struct {
	// Query selector result.
	NodeIDs []NodeID `json:"nodeIds"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
RedoResult represents the result of calls to DOM.redo.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-redo
*/
type RedoResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
RemoveAttributeParams represents DOM.removeAttribute parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-removeAttribute
*/
type RemoveAttributeParams struct {
	// ID of the element to remove attribute from.
	NodeID NodeID `json:"nodeId"`

	// Name of the attribute to remove.
	Name string `json:"name"`
}

/*
RemoveAttributeResult represents the result of calls to DOM.removeAttribute.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-removeAttribute
*/
type RemoveAttributeResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
RemoveNodeParams represents DOM.removeNode parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-removeNode
*/
type RemoveNodeParams struct {
	// ID of the node to remove.
	NodeID NodeID `json:"nodeId"`
}

/*
RemoveNodeResult represents the result of calls to DOM.removeNode.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-removeNode
*/
type RemoveNodeResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
RequestChildNodesParams represents DOM.requestChildNodes parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-requestChildNodes
*/
type RequestChildNodesParams struct {
	// ID of the node to get children for.
	NodeID NodeID `json:"nodeId"`

	// Optional. The maximum depth at which children should be retrieved,
	// defaults to 1. Use -1 for the entire subtree or provide an integer larger
	// than 0.
	Depth int `json:"depth,omitempty"`

	// Optional. Whether or not iframes and shadow roots should be traversed
	// when returning the sub-tree (default is false).
	Pierce bool `json:"pierce,omitempty"`
}

/*
RequestChildNodesResult represents the result of calls to DOM.requestChildNodes.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-requestChildNodes
*/
type RequestChildNodesResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
RequestNodeParams represents DOM.requestNode parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-requestNode
*/
type RequestNodeParams struct {
	// JavaScript object ID to convert into node.
	ObjectID runtime.RemoteObjectID `json:"objectId"`
}

/*
RequestNodeResult represents the result of calls to DOM.requestNode.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-requestNode
*/
type RequestNodeResult struct {
	// Node ID for given object.
	NodeID NodeID `json:"nodeId"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
ResolveNodeParams represents DOM.resolveNode parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-resolveNode
*/
type ResolveNodeParams struct {
	// Optional. ID of the node to resolve.
	NodeID NodeID `json:"nodeId,omitempty"`

	// Optional. Backend identifier of the node to resolve.
	BackendNodeID BackendNodeID `json:"backendNodeId,omitempty"`

	// Optional. Symbolic group name that can be used to release multiple
	// objects.
	ObjectGroup string `json:"objectGroup,omitempty"`
}

/*
ResolveNodeResult represents the result of calls to DOM.resolveNode.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-resolveNode
*/
type ResolveNodeResult struct {
	// JavaScript object wrapper for given node.
	Object *runtime.RemoteObject `json:"object"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SetAttributeValueParams represents DOM.setAttributeValue parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setAttributeValue
*/
type SetAttributeValueParams struct {
	// ID of the element to set attribute for.
	NodeID NodeID `json:"nodeId"`

	// Attribute name.
	Name string `json:"name"`

	// Attribute value.
	Value string `json:"value"`
}

/*
SetAttributeValueResult represents the result of calls to DOM.setAttributeValue.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setAttributeValue
*/
type SetAttributeValueResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SetAttributesAsTextParams represents DOM.setAttributesAsText parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setAttributesAsText
*/
type SetAttributesAsTextParams struct {
	// ID of the element to set attributes for.
	NodeID NodeID `json:"nodeId"`

	// Text with a number of attributes. Will parse this text using HTML parser.
	Text string `json:"text"`

	// Optional. Attribute name to replace with new attributes derived from text
	// in case text parsed successfully.
	Name string `json:"name,omitempty"`
}

/*
SetAttributesAsTextResult represents the result of calls to DOM.setAttributesAsText.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setAttributesAsText
*/
type SetAttributesAsTextResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SetFileInputFilesParams represents DOM.setFileInputFiles parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setFileInputFiles
*/
type SetFileInputFilesParams struct {
	// Array of file paths to set.
	Files []string `json:"files"`

	// Optional. ID of the node.
	NodeID NodeID `json:"nodeId,omitempty"`

	// Optional. ID of the backend node.
	BackendNodeID BackendNodeID `json:"backendNodeId,omitempty"`

	// Optional. JavaScript object ID of the node wrapper.
	ObjectID runtime.RemoteObjectID `json:"objectId,omitempty"`
}

/*
SetFileInputFilesResult represents the result of calls to DOM.setFileInputFiles.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setFileInputFiles
*/
type SetFileInputFilesResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SetInspectedNodeParams represents DOM.setInspectedNode parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setInspectedNode
*/
type SetInspectedNodeParams struct {
	// DOM node ID to be accessible by means of $x command line API.
	NodeID NodeID `json:"nodeId"`
}

/*
SetInspectedNodeResult represents the result of calls to DOM.setInspectedNode.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setInspectedNode
*/
type SetInspectedNodeResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SetNodeNameParams represents DOM.setNodeName parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setNodeName
*/
type SetNodeNameParams struct {
	// ID of the node to set name for.
	NodeID NodeID `json:"nodeId"`

	// New node name.
	Name string `json:"name"`
}

/*
SetNodeNameResult represents the result of calls to DOM.setNodeName.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setNodeName
*/
type SetNodeNameResult struct {
	// New node's ID.
	NodeID NodeID `json:"nodeId"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SetNodeValueParams represents DOM.setNodeValue parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setNodeValue
*/
type SetNodeValueParams struct {
	// ID of the node to set value for.
	NodeID NodeID `json:"nodeId"`

	// New node value.
	Value string `json:"value"`
}

/*
SetNodeValueResult represents the result of calls to DOM.setNodeValue.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setNodeValue
*/
type SetNodeValueResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SetOuterHTMLParams represents DOM.setOuterHTML parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setOuterHTML
*/
type SetOuterHTMLParams struct {
	// ID of the node to set markup for.
	NodeID NodeID `json:"nodeId"`

	// Outer HTML markup to set.
	OuterHTML string `json:"outerHTML"`
}

/*
SetOuterHTMLResult represents the result of calls to DOM.setOuterHTML.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setOuterHTML
*/
type SetOuterHTMLResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
UndoResult represents the result of calls to DOM.undo.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-undo
*/
type UndoResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

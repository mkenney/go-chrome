package socket

import (
	"encoding/json"

	dom "github.com/mkenney/go-chrome/cdtp/dom"
	log "github.com/sirupsen/logrus"
)

/*
DOMProtocol provides a namespace for the Chrome DOM protocol methods. The DOM
protocol exposes DOM read/write operations. Each DOM Node is represented with
its mirror object that has an ID. This ID can be used to get additional
information on the Node, resolve it into the JavaScript object wrapper, etc. It
is important that client receives DOM events only for the nodes that are known
to the client. Backend keeps track of the nodes that were sent to the client and
never sends the same node twice. It is client's responsibility to collect
information about the nodes that were sent to the client.

Note that iframe owner elements will return corresponding document elements as their child nodes.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/
*/
type DOMProtocol struct {
	Socket Socketer
}

/*
CollectClassNamesFromSubtree creates a deep copy of the specified node and
places it into the target container before the given anchor.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-collectClassNamesFromSubtree
EXPERIMENTAL.
*/
func (protocol *DOMProtocol) CollectClassNamesFromSubtree(
	params *dom.CollectClassNamesFromSubtreeParams,
) chan *dom.CollectClassNamesFromSubtreeResult {
	resultChan := make(chan *dom.CollectClassNamesFromSubtreeResult)
	command := NewCommand(protocol.Socket, "DOM.collectClassNamesFromSubtree", params)
	result := &dom.CollectClassNamesFromSubtreeResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		} else {
			result.CDTPError = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
}

/*
CopyTo creates a deep copy of the specified node and places it into the target
container before the given anchor.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-copyTo EXPERIMENTAL.
*/
func (protocol *DOMProtocol) CopyTo(
	params *dom.CopyToParams,
) chan *dom.CopyToResult {
	resultChan := make(chan *dom.CopyToResult)
	command := NewCommand(protocol.Socket, "DOM.copyTo", params)
	result := &dom.CopyToResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		} else {
			result.CDTPError = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
}

/*
DescribeNode describes node given its id, does not require domain to be enabled.
Does not start tracking any objects, can be used for automation.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-describeNode
*/
func (protocol *DOMProtocol) DescribeNode(
	params *dom.DescribeNodeParams,
) chan *dom.DescribeNodeResult {
	resultChan := make(chan *dom.DescribeNodeResult)
	command := NewCommand(protocol.Socket, "DOM.describeNode", params)
	result := &dom.DescribeNodeResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		} else {
			result.CDTPError = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
}

/*
Disable disables the DOM agent for the given page.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-disable
*/
func (protocol *DOMProtocol) Disable() chan *dom.DisableResult {
	resultChan := make(chan *dom.DisableResult)
	command := NewCommand(protocol.Socket, "DOM.disable", nil)
	result := &dom.DisableResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
DiscardSearchResults discards search results from the session with the given id.
getSearchResults should no longer be called for that search.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-discardSearchResults
EXPERIMENTAL.
*/
func (protocol *DOMProtocol) DiscardSearchResults(
	params *dom.DiscardSearchResultsParams,
) chan *dom.DiscardSearchResultsResult {
	resultChan := make(chan *dom.DiscardSearchResultsResult)
	command := NewCommand(protocol.Socket, "DOM.discardSearchResults", params)
	result := &dom.DiscardSearchResultsResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
Enable enables the DOM agent for the given page.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-enable
*/
func (protocol *DOMProtocol) Enable() chan *dom.EnableResult {
	resultChan := make(chan *dom.EnableResult)
	command := NewCommand(protocol.Socket, "DOM.enable", nil)
	result := &dom.EnableResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
Focus focuses the given element.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-focus
*/
func (protocol *DOMProtocol) Focus(
	params *dom.FocusParams,
) chan *dom.FocusResult {
	resultChan := make(chan *dom.FocusResult)
	command := NewCommand(protocol.Socket, "DOM.focus", params)
	result := &dom.FocusResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
GetAttributes returns attributes for the specified node.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getAttributes
*/
func (protocol *DOMProtocol) GetAttributes(
	params *dom.GetAttributesParams,
) chan *dom.GetAttributesResult {
	resultChan := make(chan *dom.GetAttributesResult)
	command := NewCommand(protocol.Socket, "DOM.getAttributes", params)
	result := &dom.GetAttributesResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		} else {
			result.CDTPError = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
}

/*
GetBoxModel returns boxes for the given node.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getBoxModel
*/
func (protocol *DOMProtocol) GetBoxModel(
	params *dom.GetBoxModelParams,
) chan *dom.GetBoxModelResult {
	resultChan := make(chan *dom.GetBoxModelResult)
	command := NewCommand(protocol.Socket, "DOM.getBoxModel", params)
	result := &dom.GetBoxModelResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		} else {
			result.CDTPError = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
}

/*
GetDocument returns the root DOM node (and optionally the subtree) to the
caller.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getDocument
*/
func (protocol *DOMProtocol) GetDocument(
	params *dom.GetDocumentParams,
) chan *dom.GetDocumentResult {
	resultChan := make(chan *dom.GetDocumentResult)
	command := NewCommand(protocol.Socket, "DOM.getDocument", params)
	result := &dom.GetDocumentResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		} else {
			result.CDTPError = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
}

/*
GetFlattenedDocument returns the root DOM node (and optionally the subtree) to
the caller.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getFlattenedDocument
*/
func (protocol *DOMProtocol) GetFlattenedDocument(
	params *dom.GetFlattenedDocumentParams,
) chan *dom.GetFlattenedDocumentResult {
	resultChan := make(chan *dom.GetFlattenedDocumentResult)
	command := NewCommand(protocol.Socket, "DOM.getFlattenedDocument", params)
	result := &dom.GetFlattenedDocumentResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		} else {
			result.CDTPError = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
}

/*
GetNodeForLocation returns node id at given location.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getNodeForLocation EXPERIMENTAL.
*/
func (protocol *DOMProtocol) GetNodeForLocation(
	params *dom.GetNodeForLocationParams,
) chan *dom.GetNodeForLocationResult {
	resultChan := make(chan *dom.GetNodeForLocationResult)
	command := NewCommand(protocol.Socket, "DOM.getNodeForLocation", params)
	result := &dom.GetNodeForLocationResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		} else {
			result.CDTPError = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
}

/*
GetOuterHTML returns node's HTML markup.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getOuterHTML
*/
func (protocol *DOMProtocol) GetOuterHTML(
	params *dom.GetOuterHTMLParams,
) chan *dom.GetOuterHTMLResult {
	resultChan := make(chan *dom.GetOuterHTMLResult)
	command := NewCommand(protocol.Socket, "DOM.getOuterHTML", params)
	result := &dom.GetOuterHTMLResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		} else {
			result.CDTPError = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
}

/*
GetRelayoutBoundary returns the id of the nearest ancestor that is a relayout
boundary.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getRelayoutBoundary EXPERIMENTAL.
*/
func (protocol *DOMProtocol) GetRelayoutBoundary(
	params *dom.GetRelayoutBoundaryParams,
) chan *dom.GetRelayoutBoundaryResult {
	resultChan := make(chan *dom.GetRelayoutBoundaryResult)
	command := NewCommand(protocol.Socket, "DOM.getRelayoutBoundary", params)
	result := &dom.GetRelayoutBoundaryResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		} else {
			result.CDTPError = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
}

/*
GetSearchResults returns search results from given fromIndex to given toIndex
from the search with the given identifier.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getSearchResults
EXPERIMENTAL.
*/
func (protocol *DOMProtocol) GetSearchResults(
	params *dom.GetSearchResultsParams,
) chan *dom.GetSearchResultsResult {
	resultChan := make(chan *dom.GetSearchResultsResult)
	command := NewCommand(protocol.Socket, "DOM.getSearchResults", params)
	result := &dom.GetSearchResultsResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		} else {
			result.CDTPError = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
}

/*
MarkUndoableState marks last undoable state.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-markUndoableState
EXPERIMENTAL.
*/
func (protocol *DOMProtocol) MarkUndoableState() chan *dom.MarkUndoableStateResult {
	resultChan := make(chan *dom.MarkUndoableStateResult)
	command := NewCommand(protocol.Socket, "DOM.markUndoableState", nil)
	result := &dom.MarkUndoableStateResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
MoveTo moves node into the new container, places it before the given anchor.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-moveTo
*/
func (protocol *DOMProtocol) MoveTo(
	params *dom.MoveToParams,
) chan *dom.MoveToResult {
	resultChan := make(chan *dom.MoveToResult)
	command := NewCommand(protocol.Socket, "DOM.moveTo", params)
	result := &dom.MoveToResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		} else {
			result.CDTPError = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
}

/*
PerformSearch searches for a given string in the DOM tree. Use getSearchResults
to access search results or cancelSearch to end this search session.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-performSearch
EXPERIMENTAL.
*/
func (protocol *DOMProtocol) PerformSearch(
	params *dom.PerformSearchParams,
) chan *dom.PerformSearchResult {
	resultChan := make(chan *dom.PerformSearchResult)
	command := NewCommand(protocol.Socket, "DOM.performSearch", params)
	result := &dom.PerformSearchResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		} else {
			result.CDTPError = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
}

/*
PushNodeByPathToFrontend requests that the node is sent to the caller given its
path.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-pushNodeByPathToFrontend
EXPERIMENTAL. @TODO, use XPath.
*/
func (protocol *DOMProtocol) PushNodeByPathToFrontend(
	params *dom.PushNodeByPathToFrontendParams,
) chan *dom.PushNodeByPathToFrontendResult {
	resultChan := make(chan *dom.PushNodeByPathToFrontendResult)
	command := NewCommand(protocol.Socket, "DOM.pushNodeByPathToFrontend", params)
	result := &dom.PushNodeByPathToFrontendResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		} else {
			result.CDTPError = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
}

/*
PushNodesByBackendIDsToFrontend requests that a batch of nodes is sent to the
caller given their backend node IDs.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-pushNodesByBackendIdsToFrontend
EXPERIMENTAL.
*/
func (protocol *DOMProtocol) PushNodesByBackendIDsToFrontend(
	params *dom.PushNodesByBackendIDsToFrontendParams,
) chan *dom.PushNodesByBackendIDsToFrontendResult {
	resultChan := make(chan *dom.PushNodesByBackendIDsToFrontendResult)
	command := NewCommand(protocol.Socket, "DOM.pushNodesByBackendIdsToFrontend", params)
	result := &dom.PushNodesByBackendIDsToFrontendResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		} else {
			result.CDTPError = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
}

/*
QuerySelector executes querySelector on a given node.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-querySelector
*/
func (protocol *DOMProtocol) QuerySelector(
	params *dom.QuerySelectorParams,
) chan *dom.QuerySelectorResult {
	resultChan := make(chan *dom.QuerySelectorResult)
	command := NewCommand(protocol.Socket, "DOM.querySelector", params)
	result := &dom.QuerySelectorResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		} else {
			result.CDTPError = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
}

/*
QuerySelectorAll executes querySelectorAll on a given node.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-querySelectorAll
*/
func (protocol *DOMProtocol) QuerySelectorAll(
	params *dom.QuerySelectorAllParams,
) chan *dom.QuerySelectorAllResult {
	resultChan := make(chan *dom.QuerySelectorAllResult)
	command := NewCommand(protocol.Socket, "DOM.querySelectorAll", params)
	result := &dom.QuerySelectorAllResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		} else {
			result.CDTPError = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
}

/*
Redo re-does the last undone action.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-redo EXPERIMENTAL.
*/
func (protocol *DOMProtocol) Redo() chan *dom.RedoResult {
	resultChan := make(chan *dom.RedoResult)
	command := NewCommand(protocol.Socket, "DOM.redo", nil)
	result := &dom.RedoResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
RemoveAttribute removes attribute with given name from an element with given id.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-removeAttribute
*/
func (protocol *DOMProtocol) RemoveAttribute(
	params *dom.RemoveAttributeParams,
) chan *dom.RemoveAttributeResult {
	resultChan := make(chan *dom.RemoveAttributeResult)
	command := NewCommand(protocol.Socket, "DOM.removeAttribute", params)
	result := &dom.RemoveAttributeResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
RemoveNode removes the specified node.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-removeNode
*/
func (protocol *DOMProtocol) RemoveNode(
	params *dom.RemoveNodeParams,
) chan *dom.RemoveNodeResult {
	resultChan := make(chan *dom.RemoveNodeResult)
	command := NewCommand(protocol.Socket, "DOM.removeNode", params)
	result := &dom.RemoveNodeResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
RequestChildNodes requests that children of the node with given id are returned
to the caller in form of setChildNodes events where not only immediate children
are retrieved, but all children down to the specified depth.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-requestChildNodes
*/
func (protocol *DOMProtocol) RequestChildNodes(
	params *dom.RequestChildNodesParams,
) chan *dom.RequestChildNodesResult {
	resultChan := make(chan *dom.RequestChildNodesResult)
	command := NewCommand(protocol.Socket, "DOM.requestChildNodes", params)
	result := &dom.RequestChildNodesResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
RequestNode requests that the node is sent to the caller given the JavaScript
node object reference. All nodes that form the path from the node to the root
are also sent to the client as a series of setChildNodes notifications.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-requestNode
*/
func (protocol *DOMProtocol) RequestNode(
	params *dom.RequestNodeParams,
) chan *dom.RequestNodeResult {
	resultChan := make(chan *dom.RequestNodeResult)
	command := NewCommand(protocol.Socket, "DOM.requestNode", params)
	result := &dom.RequestNodeResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		} else {
			result.CDTPError = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
}

/*
ResolveNode resolves the JavaScript node object for a given NodeID or
BackendNodeID.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-resolveNode
*/
func (protocol *DOMProtocol) ResolveNode(
	params *dom.ResolveNodeParams,
) chan *dom.ResolveNodeResult {
	resultChan := make(chan *dom.ResolveNodeResult)
	command := NewCommand(protocol.Socket, "DOM.resolveNode", params)
	result := &dom.ResolveNodeResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		} else {
			result.CDTPError = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
}

/*
SetAttributeValue sets attribute for an element with given id.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setAttributeValue
*/
func (protocol *DOMProtocol) SetAttributeValue(
	params *dom.SetAttributeValueParams,
) chan *dom.SetAttributeValueResult {
	resultChan := make(chan *dom.SetAttributeValueResult)
	command := NewCommand(protocol.Socket, "DOM.setAttributeValue", params)
	result := &dom.SetAttributeValueResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
SetAttributesAsText sets attributes on element with given id. This method is
useful when user edits some existing attribute value and types in several
attribute name/value pairs.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setAttributesAsText
*/
func (protocol *DOMProtocol) SetAttributesAsText(
	params *dom.SetAttributesAsTextParams,
) chan *dom.SetAttributesAsTextResult {
	resultChan := make(chan *dom.SetAttributesAsTextResult)
	command := NewCommand(protocol.Socket, "DOM.setAttributesAsText", params)
	result := &dom.SetAttributesAsTextResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
SetFileInputFiles sets files for the given file input element.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setFileInputFiles
*/
func (protocol *DOMProtocol) SetFileInputFiles(
	params *dom.SetFileInputFilesParams,
) chan *dom.SetFileInputFilesResult {
	resultChan := make(chan *dom.SetFileInputFilesResult)
	command := NewCommand(protocol.Socket, "DOM.setFileInputFiles", params)
	result := &dom.SetFileInputFilesResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
SetInspectedNode enables console to refer to the node with given id via $x (see
Command Line API for more details $x functions).

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setInspectedNode EXPERIMENTAL.
*/
func (protocol *DOMProtocol) SetInspectedNode(
	params *dom.SetInspectedNodeParams,
) chan *dom.SetInspectedNodeResult {
	resultChan := make(chan *dom.SetInspectedNodeResult)
	command := NewCommand(protocol.Socket, "DOM.setInspectedNode", params)
	result := &dom.SetInspectedNodeResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
SetNodeName sets node name for the specified node.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setNodeName
*/
func (protocol *DOMProtocol) SetNodeName(
	params *dom.SetNodeNameParams,
) chan *dom.SetNodeNameResult {
	resultChan := make(chan *dom.SetNodeNameResult)
	command := NewCommand(protocol.Socket, "DOM.setNodeName", params)
	result := &dom.SetNodeNameResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		} else {
			result.CDTPError = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
}

/*
SetNodeValue sets node value for the specified node.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setNodeValue
*/
func (protocol *DOMProtocol) SetNodeValue(
	params *dom.SetNodeValueParams,
) chan *dom.SetNodeValueResult {
	resultChan := make(chan *dom.SetNodeValueResult)
	command := NewCommand(protocol.Socket, "DOM.setNodeValue", params)
	result := &dom.SetNodeValueResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
SetOuterHTML sets node HTML markup, returns new node id.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setOuterHTML
*/
func (protocol *DOMProtocol) SetOuterHTML(
	params *dom.SetOuterHTMLParams,
) chan *dom.SetOuterHTMLResult {
	resultChan := make(chan *dom.SetOuterHTMLResult)
	command := NewCommand(protocol.Socket, "DOM.setOuterHTML", params)
	result := &dom.SetOuterHTMLResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
Undo undoes the last performed action.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-undo
EXPERIMENTAL.
*/
func (protocol *DOMProtocol) Undo() chan *dom.UndoResult {
	resultChan := make(chan *dom.UndoResult)
	command := NewCommand(protocol.Socket, "DOM.undo", nil)
	result := &dom.UndoResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
OnAttributeModified adds a handler to the DOM.attributeModified event.
DOM.attributeModified fires when Element's attribute is modified.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-attributeModified
*/
func (protocol *DOMProtocol) OnAttributeModified(
	callback func(event *dom.AttributeModifiedEvent),
) {
	handler := NewEventHandler(
		"DOM.attributeModified",
		func(response *Response) {
			event := &dom.AttributeModifiedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnAttributeRemoved adds a handler to the DOM.attributeRemoved event.
DOM.attributeRemoved fires when Element's attribute is modified.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-attributeRemoved
*/
func (protocol *DOMProtocol) OnAttributeRemoved(
	callback func(event *dom.AttributeRemovedEvent),
) {
	handler := NewEventHandler(
		"DOM.attributeRemoved",
		func(response *Response) {
			event := &dom.AttributeRemovedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnCharacterDataModified adds a handler to the DOM.characterDataModified event.
DOM.characterDataModified mirrors the DOMCharacterDataModified event.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-characterDataModified
*/
func (protocol *DOMProtocol) OnCharacterDataModified(
	callback func(event *dom.CharacterDataModifiedEvent),
) {
	handler := NewEventHandler(
		"DOM.characterDataModified",
		func(response *Response) {
			event := &dom.CharacterDataModifiedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnChildNodeCountUpdated adds a handler to the DOM.childNodeCountUpdated event.
DOM.childNodeCountUpdated fires when Container's child node count has changed.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-childNodeCountUpdated
*/
func (protocol *DOMProtocol) OnChildNodeCountUpdated(
	callback func(event *dom.ChildNodeCountUpdatedEvent),
) {
	handler := NewEventHandler(
		"DOM.childNodeCountUpdated",
		func(response *Response) {
			event := &dom.ChildNodeCountUpdatedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnChildNodeInserted adds a handler to the DOM.childNodeInserted event.
DOM.childNodeInserted mirrors the DOMNodeInserted event.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-childNodeInserted
*/
func (protocol *DOMProtocol) OnChildNodeInserted(
	callback func(event *dom.ChildNodeInsertedEvent),
) {
	handler := NewEventHandler(
		"DOM.childNodeInserted",
		func(response *Response) {
			event := &dom.ChildNodeInsertedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnChildNodeRemoved adds a handler to the DOM.childNodeRemoved event.
DOM.childNodeRemoved mirrors the DOMNodeRemoved event.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-childNodeRemoved
*/
func (protocol *DOMProtocol) OnChildNodeRemoved(
	callback func(event *dom.ChildNodeRemovedEvent),
) {
	handler := NewEventHandler(
		"DOM.childNodeRemoved",
		func(response *Response) {
			event := &dom.ChildNodeRemovedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnDistributedNodesUpdated adds a handler to the DOM.distributedNodesUpdated
event. DOM.distributedNodesUpdated fires when distribution is changed.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-distributedNodesUpdated
EXPERIMENTAL.
*/
func (protocol *DOMProtocol) OnDistributedNodesUpdated(
	callback func(event *dom.DistributedNodesUpdatedEvent),
) {
	handler := NewEventHandler(
		"DOM.distributedNodesUpdated",
		func(response *Response) {
			event := &dom.DistributedNodesUpdatedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnDocumentUpdated adds a handler to the DOM.documentUpdated event.
DOM.documentUpdated fires when Document has been totally updated. Node IDs are
no longer valid.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-documentUpdated
*/
func (protocol *DOMProtocol) OnDocumentUpdated(
	callback func(event *dom.DocumentUpdatedEvent),
) {
	handler := NewEventHandler(
		"DOM.documentUpdated",
		func(response *Response) {
			event := &dom.DocumentUpdatedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnInlineStyleInvalidated adds a handler to the DOM.inlineStyleInvalidated event.
DOM.inlineStyleInvalidated fires when Element's attribute is removed.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-inlineStyleInvalidated
*/
func (protocol *DOMProtocol) OnInlineStyleInvalidated(
	callback func(event *dom.InlineStyleInvalidatedEvent),
) {
	handler := NewEventHandler(
		"DOM.inlineStyleInvalidated",
		func(response *Response) {
			event := &dom.InlineStyleInvalidatedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnPseudoElementAdded adds a handler to the DOM.pseudoElementAdded event.
DOM.pseudoElementAdded fires when a pseudo element is added to an element.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-pseudoElementAdded EXPERIMENTAL.
*/
func (protocol *DOMProtocol) OnPseudoElementAdded(
	callback func(event *dom.PseudoElementAddedEvent),
) {
	handler := NewEventHandler(
		"DOM.pseudoElementAdded",
		func(response *Response) {
			event := &dom.PseudoElementAddedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnPseudoElementRemoved adds a handler to the DOM.pseudoElementRemoved event.
DOM.pseudoElementRemoved fires when a pseudo element is removed from an element.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-pseudoElementRemoved EXPERIMENTAL.
*/
func (protocol *DOMProtocol) OnPseudoElementRemoved(
	callback func(event *dom.PseudoElementRemovedEvent),
) {
	handler := NewEventHandler(
		"DOM.pseudoElementRemoved",
		func(response *Response) {
			event := &dom.PseudoElementRemovedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnSetChildNodes adds a handler to the DOM.setChildNodes event. DOM.setChildNodes
fires when backend wants to provide client with the missing DOM structure. This
happens upon most of the calls requesting node IDs.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-setChildNodes
*/
func (protocol *DOMProtocol) OnSetChildNodes(
	callback func(event *dom.SetChildNodesEvent),
) {
	handler := NewEventHandler(
		"DOM.setChildNodes",
		func(response *Response) {
			event := &dom.SetChildNodesEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnShadowRootPopped adds a handler to the DOM.shadowRootPopped event.
DOM.shadowRootPopped fires when shadow root is popped from the element.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-shadowRootPopped EXPERIMENTAL.
*/
func (protocol *DOMProtocol) OnShadowRootPopped(
	callback func(event *dom.ShadowRootPoppedEvent),
) {
	handler := NewEventHandler(
		"DOM.shadowRootPopped",
		func(response *Response) {
			event := &dom.ShadowRootPoppedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnShadowRootPushed adds a handler to the DOM.shadowRootPushed event.
DOM.shadowRootPushed fires when shadow root is pushed into the element.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-shadowRootPushed EXPERIMENTAL.
*/
func (protocol *DOMProtocol) OnShadowRootPushed(
	callback func(event *dom.ShadowRootPushedEvent),
) {
	handler := NewEventHandler(
		"DOM.shadowRootPushed",
		func(response *Response) {
			event := &dom.ShadowRootPushedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

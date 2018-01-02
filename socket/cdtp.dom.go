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
) (*dom.CollectClassNamesFromSubtreeResult, error) {
	command := NewCommand("DOM.collectClassNamesFromSubtree", params)
	result := &dom.CollectClassNamesFromSubtreeResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
CopyTo creates a deep copy of the specified node and places it into the target
container before the given anchor.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-copyTo EXPERIMENTAL.
*/
func (protocol *DOMProtocol) CopyTo(
	params *dom.CopyToParams,
) (*dom.CopyToResult, error) {
	command := NewCommand("DOM.copyTo", params)
	result := &dom.CopyToResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
DescribeNode describes node given its id, does not require domain to be enabled.
Does not start tracking any objects, can be used for automation.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-describeNode
*/
func (protocol *DOMProtocol) DescribeNode(
	params *dom.DescribeNodeParams,
) (*dom.DescribeNodeResult, error) {
	command := NewCommand("DOM.describeNode", params)
	result := &dom.DescribeNodeResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
Disable disables the DOM agent for the given page.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-disable
*/
func (protocol *DOMProtocol) Disable() error {
	command := NewCommand("DOM.disable", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
DiscardSearchResults discards search results from the session with the given id.
getSearchResults should no longer be called for that search.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-discardSearchResults
EXPERIMENTAL.
*/
func (protocol *DOMProtocol) DiscardSearchResults(
	params *dom.DiscardSearchResultsParams,
) error {
	command := NewCommand("DOM.discardSearchResults", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
Enable enables the DOM agent for the given page.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-enable
*/
func (protocol *DOMProtocol) Enable() error {
	command := NewCommand("DOM.enable", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
Focus focuses the given element.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-focus
*/
func (protocol *DOMProtocol) Focus(
	params *dom.FocusParams,
) error {
	command := NewCommand("DOM.focus", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
GetAttributes returns attributes for the specified node.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getAttributes
*/
func (protocol *DOMProtocol) GetAttributes(
	params *dom.GetAttributesParams,
) (*dom.GetAttributesResult, error) {
	command := NewCommand("DOM.getAttributes", params)
	result := &dom.GetAttributesResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
GetBoxModel returns boxes for the given node.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getBoxModel
*/
func (protocol *DOMProtocol) GetBoxModel(
	params *dom.GetBoxModelParams,
) (*dom.GetBoxModelResult, error) {
	command := NewCommand("DOM.getBoxModel", params)
	result := &dom.GetBoxModelResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
GetDocument returns the root DOM node (and optionally the subtree) to the
caller.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getDocument
*/
func (protocol *DOMProtocol) GetDocument(
	params *dom.GetDocumentParams,
) (*dom.GetDocumentResult, error) {
	command := NewCommand("DOM.getDocument", params)
	result := &dom.GetDocumentResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
GetFlattenedDocument returns the root DOM node (and optionally the subtree) to
the caller.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getFlattenedDocument
*/
func (protocol *DOMProtocol) GetFlattenedDocument(
	params *dom.GetFlattenedDocumentParams,
) (*dom.GetFlattenedDocumentResult, error) {
	command := NewCommand("DOM.getFlattenedDocument", params)
	result := &dom.GetFlattenedDocumentResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
GetNodeForLocation returns node id at given location.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getNodeForLocation EXPERIMENTAL.
*/
func (protocol *DOMProtocol) GetNodeForLocation(
	params *dom.GetNodeForLocationParams,
) (*dom.GetNodeForLocationResult, error) {
	command := NewCommand("DOM.getNodeForLocation", params)
	result := &dom.GetNodeForLocationResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
GetOuterHTML returns node's HTML markup.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getOuterHTML
*/
func (protocol *DOMProtocol) GetOuterHTML(
	params *dom.GetOuterHTMLParams,
) (*dom.GetOuterHTMLResult, error) {
	command := NewCommand("DOM.getOuterHTML", params)
	result := &dom.GetOuterHTMLResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
GetRelayoutBoundary returns the id of the nearest ancestor that is a relayout
boundary.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getRelayoutBoundary EXPERIMENTAL.
*/
func (protocol *DOMProtocol) GetRelayoutBoundary(
	params *dom.GetRelayoutBoundaryParams,
) (*dom.GetRelayoutBoundaryResult, error) {
	command := NewCommand("DOM.getRelayoutBoundary", params)
	result := &dom.GetRelayoutBoundaryResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
GetSearchResults returns search results from given fromIndex to given toIndex
from the search with the given identifier.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getSearchResults
EXPERIMENTAL.
*/
func (protocol *DOMProtocol) GetSearchResults(
	params *dom.GetSearchResultsParams,
) (*dom.GetSearchResultsResult, error) {
	command := NewCommand("DOM.getSearchResults", params)
	result := &dom.GetSearchResultsResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
MarkUndoableState marks last undoable state.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-markUndoableState
EXPERIMENTAL.
*/
func (protocol *DOMProtocol) MarkUndoableState() error {
	command := NewCommand("DOM.markUndoableState", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
MoveTo moves node into the new container, places it before the given anchor.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-moveTo
*/
func (protocol *DOMProtocol) MoveTo(
	params *dom.MoveToParams,
) (*dom.MoveToResult, error) {
	command := NewCommand("DOM.moveTo", params)
	result := &dom.MoveToResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
PerformSearch searches for a given string in the DOM tree. Use getSearchResults
to access search results or cancelSearch to end this search session.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-performSearch
EXPERIMENTAL.
*/
func (protocol *DOMProtocol) PerformSearch(
	params *dom.PerformSearchParams,
) (*dom.PerformSearchResult, error) {
	command := NewCommand("DOM.performSearch", params)
	result := &dom.PerformSearchResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
PushNodeByPathToFrontend requests that the node is sent to the caller given its
path.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-pushNodeByPathToFrontend
EXPERIMENTAL. @TODO, use XPath.
*/
func (protocol *DOMProtocol) PushNodeByPathToFrontend(
	params *dom.PushNodeByPathToFrontendParams,
) (*dom.PushNodeByPathToFrontendResult, error) {
	command := NewCommand("DOM.pushNodeByPathToFrontend", params)
	result := &dom.PushNodeByPathToFrontendResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
PushNodesByBackendIDsToFrontend requests that a batch of nodes is sent to the
caller given their backend node IDs.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-pushNodesByBackendIdsToFrontend
EXPERIMENTAL.
*/
func (protocol *DOMProtocol) PushNodesByBackendIDsToFrontend(
	params *dom.PushNodesByBackendIDsToFrontendParams,
) (*dom.PushNodesByBackendIDsToFrontendResult, error) {
	command := NewCommand("DOM.pushNodesByBackendIdsToFrontend", params)
	result := &dom.PushNodesByBackendIDsToFrontendResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
QuerySelector executes querySelector on a given node.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-querySelector
*/
func (protocol *DOMProtocol) QuerySelector(
	params *dom.QuerySelectorParams,
) (*dom.QuerySelectorResult, error) {
	command := NewCommand("DOM.querySelector", params)
	result := &dom.QuerySelectorResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
QuerySelectorAll executes querySelectorAll on a given node.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-querySelectorAll
*/
func (protocol *DOMProtocol) QuerySelectorAll(
	params *dom.QuerySelectorAllParams,
) (*dom.QuerySelectorAllResult, error) {
	command := NewCommand("DOM.querySelectorAll", params)
	result := &dom.QuerySelectorAllResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
Redo re-does the last undone action.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-redo EXPERIMENTAL.
*/
func (protocol *DOMProtocol) Redo() error {
	command := NewCommand("DOM.redo", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
RemoveAttribute removes attribute with given name from an element with given id.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-removeAttribute
*/
func (protocol *DOMProtocol) RemoveAttribute(
	params *dom.RemoveAttributeParams,
) error {
	command := NewCommand("DOM.removeAttribute", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
RemoveNode removes the specified node.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-removeNode
*/
func (protocol *DOMProtocol) RemoveNode(
	params *dom.RemoveNodeParams,
) error {
	command := NewCommand("DOM.removeNode", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
RequestChildNodes requests that children of the node with given id are returned
to the caller in form of setChildNodes events where not only immediate children
are retrieved, but all children down to the specified depth.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-requestChildNodes
*/
func (protocol *DOMProtocol) RequestChildNodes(
	params *dom.RequestChildNodesParams,
) error {
	command := NewCommand("DOM.requestChildNodes", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
RequestNode requests that the node is sent to the caller given the JavaScript
node object reference. All nodes that form the path from the node to the root
are also sent to the client as a series of setChildNodes notifications.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-requestNode
*/
func (protocol *DOMProtocol) RequestNode(
	params *dom.RequestNodeParams,
) (*dom.RequestNodeResult, error) {
	command := NewCommand("DOM.requestNode", params)
	result := &dom.RequestNodeResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
ResolveNode resolves the JavaScript node object for a given NodeID or
BackendNodeID.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-resolveNode
*/
func (protocol *DOMProtocol) ResolveNode(
	params *dom.ResolveNodeParams,
) (*dom.ResolveNodeResult, error) {
	command := NewCommand("DOM.resolveNode", params)
	result := &dom.ResolveNodeResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
SetAttributeValue sets attribute for an element with given id.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setAttributeValue
*/
func (protocol *DOMProtocol) SetAttributeValue(
	params *dom.SetAttributeValueParams,
) error {
	command := NewCommand("DOM.setAttributeValue", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetAttributesAsText sets attributes on element with given id. This method is
useful when user edits some existing attribute value and types in several
attribute name/value pairs.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setAttributesAsText
*/
func (protocol *DOMProtocol) SetAttributesAsText(
	params *dom.SetAttributesAsTextParams,
) error {
	command := NewCommand("DOM.setAttributesAsText", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetFileInputFiles sets files for the given file input element.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setFileInputFiles
*/
func (protocol *DOMProtocol) SetFileInputFiles(
	params *dom.SetFileInputFilesParams,
) error {
	command := NewCommand("DOM.setFileInputFiles", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetInspectedNode enables console to refer to the node with given id via $x (see
Command Line API for more details $x functions).

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setInspectedNode EXPERIMENTAL.
*/
func (protocol *DOMProtocol) SetInspectedNode(
	params *dom.SetInspectedNodeParams,
) error {
	command := NewCommand("DOM.setInspectedNode", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetNodeName sets node name for the specified node.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setNodeName
*/
func (protocol *DOMProtocol) SetNodeName(
	params *dom.SetNodeNameParams,
) (*dom.SetNodeNameResult, error) {
	command := NewCommand("DOM.setNodeName", params)
	result := &dom.SetNodeNameResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
SetNodeValue sets node value for the specified node.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setNodeValue
*/
func (protocol *DOMProtocol) SetNodeValue(
	params *dom.SetNodeValueParams,
) error {
	command := NewCommand("DOM.setNodeValue", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetOuterHTML sets node HTML markup, returns new node id.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setOuterHTML
*/
func (protocol *DOMProtocol) SetOuterHTML(
	params *dom.SetOuterHTMLParams,
) error {
	command := NewCommand("DOM.setOuterHTML", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
Undo undoes the last performed action.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-undo
EXPERIMENTAL.
*/
func (protocol *DOMProtocol) Undo() error {
	command := NewCommand("DOM.undo", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
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

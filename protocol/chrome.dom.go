package protocol

import (
	"encoding/json"

	dom "github.com/mkenney/go-chrome/protocol/dom"
	sock "github.com/mkenney/go-chrome/socket"

	log "github.com/sirupsen/logrus"
)

/*
DOM is a struct that provides a namespace for the Chrome DOM protocol methods.

The DOM protocol exposes DOM read/write operations. Each DOM Node is represented with its mirror object that has an
ID. This ID can be used to get additional information on the Node, resolve it into the JavaScript
object wrapper, etc. It is important that client receives DOM events only for the nodes that are
known to the client. Backend keeps track of the nodes that were sent to the client and never sends
the same node twice. It is client's responsibility to collect information about the nodes that were
sent to the client.

Note that iframe owner elements will return corresponding document elements as their child nodes.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/
*/
type DOM struct{}

/*
CollectClassNamesFromSubtree creates a deep copy of the specified node and places it into the target container before the
given anchor. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-collectClassNamesFromSubtree
*/
func (DOM) CollectClassNamesFromSubtree(
	socket sock.Socketer,
	params *dom.CollectClassNamesFromSubtreeParams,
) (dom.CollectClassNamesFromSubtreeResult, error) {
	command := sock.NewCommand("DOM.collectClassNamesFromSubtree", params)
	result := dom.CollectClassNamesFromSubtreeResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
CopyTo creates a deep copy of the specified node and places it into the target container before the
given anchor. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-copyTo
*/
func (DOM) CopyTo(
	socket sock.Socketer,
	params *dom.CopyToParams,
) (dom.CopyToResult, error) {
	command := sock.NewCommand("DOM.copyTo", params)
	result := dom.CopyToResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
DescribeNode describes node given its id, does not require domain to be enabled. Does not start
tracking any objects, can be used for automation.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-describeNode
*/
func (DOM) DescribeNode(
	socket sock.Socketer,
	params *dom.DescribeNodeParams,
) (dom.DescribeNodeResult, error) {
	command := sock.NewCommand("DOM.describeNode", params)
	result := dom.DescribeNodeResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
Disable disables the DOM agent for the given page.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-disable
*/
func (DOM) Disable(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("DOM.disable", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
DiscardSearchResults discards search results from the session with the given id. getSearchResults
should no longer be called for that search. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-discardSearchResults
*/
func (DOM) DiscardSearchResults(
	socket sock.Socketer,
	params *dom.DiscardSearchResultsParams,
) error {
	command := sock.NewCommand("DOM.discardSearchResults", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
Enable enables the DOM agent for the given page.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-enable
*/
func (DOM) Enable(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("DOM.enable", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
Focus focuses the given element.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-focus
*/
func (DOM) Focus(
	socket sock.Socketer,
	params *dom.FocusParams,
) error {
	command := sock.NewCommand("DOM.focus", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
GetAttributes returns attributes for the specified node.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getAttributes
*/
func (DOM) GetAttributes(
	socket sock.Socketer,
	params *dom.GetAttributesParams,
) (dom.GetAttributesResult, error) {
	command := sock.NewCommand("DOM.getAttributes", params)
	result := dom.GetAttributesResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
GetBoxModel returns boxes for the given node.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getBoxModel
*/
func (DOM) GetBoxModel(
	socket sock.Socketer,
	params *dom.GetBoxModelParams,
) (dom.GetBoxModelResult, error) {
	command := sock.NewCommand("DOM.getBoxModel", params)
	result := dom.GetBoxModelResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
GetDocument returns the root DOM node (and optionally the subtree) to the caller.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getDocument
*/
func (DOM) GetDocument(
	socket sock.Socketer,
	params *dom.GetDocumentParams,
) (dom.GetDocumentResult, error) {
	command := sock.NewCommand("DOM.getDocument", params)
	result := dom.GetDocumentResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
GetFlattenedDocument returns the root DOM node (and optionally the subtree) to the caller.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getFlattenedDocument
*/
func (DOM) GetFlattenedDocument(
	socket sock.Socketer,
	params *dom.GetFlattenedDocumentParams,
) (dom.GetFlattenedDocumentResult, error) {
	command := sock.NewCommand("DOM.getFlattenedDocument", params)
	result := dom.GetFlattenedDocumentResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
GetNodeForLocation returns node id at given location. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getNodeForLocation
*/
func (DOM) GetNodeForLocation(
	socket sock.Socketer,
	params *dom.GetNodeForLocationParams,
) (dom.GetNodeForLocationResult, error) {
	command := sock.NewCommand("DOM.getNodeForLocation", params)
	result := dom.GetNodeForLocationResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
GetOuterHTML returns node's HTML markup.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getOuterHTML
*/
func (DOM) GetOuterHTML(
	socket sock.Socketer,
	params *dom.GetOuterHTMLParams,
) (dom.GetOuterHTMLResult, error) {
	command := sock.NewCommand("DOM.getOuterHTML", params)
	result := dom.GetOuterHTMLResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
GetRelayoutBoundary returns the id of the nearest ancestor that is a relayout boundary. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getRelayoutBoundary
*/
func (DOM) GetRelayoutBoundary(
	socket sock.Socketer,
	params *dom.GetRelayoutBoundaryParams,
) (dom.GetRelayoutBoundaryResult, error) {
	command := sock.NewCommand("DOM.getRelayoutBoundary", params)
	result := dom.GetRelayoutBoundaryResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
GetSearchResults returns search results from given fromIndex to given toIndex from the search with
the given identifier. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getSearchResults
*/
func (DOM) GetSearchResults(
	socket sock.Socketer,
	params *dom.GetSearchResultsParams,
) (dom.GetSearchResultsResult, error) {
	command := sock.NewCommand("DOM.getSearchResults", params)
	result := dom.GetSearchResultsResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
MarkUndoableState marks last undoable state. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-markUndoableState
*/
func (DOM) MarkUndoableState(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("DOM.markUndoableState", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
MoveTo moves node into the new container, places it before the given anchor.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-moveTo
*/
func (DOM) MoveTo(
	socket sock.Socketer,
	params *dom.MoveToParams,
) (dom.MoveToResult, error) {
	command := sock.NewCommand("DOM.moveTo", params)
	result := dom.MoveToResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
PerformSearch searches for a given string in the DOM tree. Use getSearchResults to access search
results or cancelSearch to end this search session. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-performSearch
*/
func (DOM) PerformSearch(
	socket sock.Socketer,
	params *dom.PerformSearchParams,
) (dom.PerformSearchResult, error) {
	command := sock.NewCommand("DOM.performSearch", params)
	result := dom.PerformSearchResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
PushNodeByPathToFrontend requests that the node is sent to the caller given its path. EXPERIMENTAL
@TODO, use XPath.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-pushNodeByPathToFrontend
*/
func (DOM) PushNodeByPathToFrontend(
	socket sock.Socketer,
	params *dom.PushNodeByPathToFrontendParams,
) (dom.PushNodeByPathToFrontendResult, error) {
	command := sock.NewCommand("DOM.pushNodeByPathToFrontend", params)
	result := dom.PushNodeByPathToFrontendResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
PushNodesByBackendIDsToFrontend requests that a batch of nodes is sent to the caller given their
backend node IDs. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-pushNodesByBackendIdsToFrontend
*/
func (DOM) PushNodesByBackendIDsToFrontend(
	socket sock.Socketer,
	params *dom.PushNodesByBackendIDsToFrontendParams,
) (dom.PushNodesByBackendIDsToFrontendResult, error) {
	command := sock.NewCommand("DOM.pushNodesByBackendIdsToFrontend", params)
	result := dom.PushNodesByBackendIDsToFrontendResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
QuerySelector executes querySelector on a given node.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-querySelector
*/
func (DOM) QuerySelector(
	socket sock.Socketer,
	params *dom.QuerySelectorParams,
) (dom.QuerySelectorResult, error) {
	command := sock.NewCommand("DOM.querySelector", params)
	result := dom.QuerySelectorResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
QuerySelectorAll executes querySelectorAll on a given node.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-querySelectorAll
*/
func (DOM) QuerySelectorAll(
	socket sock.Socketer,
	params *dom.QuerySelectorAllParams,
) (dom.QuerySelectorAllResult, error) {
	command := sock.NewCommand("DOM.querySelectorAll", params)
	result := dom.QuerySelectorAllResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
Redo re-does the last undone action. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-redo
*/
func (DOM) Redo(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("DOM.redo", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
RemoveAttribute removes attribute with given name from an element with given id.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-removeAttribute
*/
func (DOM) RemoveAttribute(
	socket sock.Socketer,
	params *dom.RemoveAttributeParams,
) error {
	command := sock.NewCommand("DOM.removeAttribute", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
RemoveNode removes the specified node.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-removeNode
*/
func (DOM) RemoveNode(
	socket sock.Socketer,
	params *dom.RemoveNodeParams,
) error {
	command := sock.NewCommand("DOM.removeNode", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
RequestChildNodes requests that children of the node with given id are returned to the caller in
form of setChildNodes events where not only immediate children are retrieved, but all children down
to the specified depth.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-requestChildNodes
*/
func (DOM) RequestChildNodes(
	socket sock.Socketer,
	params *dom.RequestChildNodesParams,
) error {
	command := sock.NewCommand("DOM.requestChildNodes", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
RequestNode requests that the node is sent to the caller given the JavaScript node object reference.
All nodes that form the path from the node to the root are also sent to the client as a series of
setChildNodes notifications.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-requestNode
*/
func (DOM) RequestNode(
	socket sock.Socketer,
	params *dom.RequestNodeParams,
) (dom.RequestNodeResult, error) {
	command := sock.NewCommand("DOM.requestNode", params)
	result := dom.RequestNodeResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
ResolveNode resolves the JavaScript node object for a given NodeID or BackendNodeID.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-resolveNode
*/
func (DOM) ResolveNode(
	socket sock.Socketer,
	params *dom.ResolveNodeParams,
) (dom.ResolveNodeResult, error) {
	command := sock.NewCommand("DOM.resolveNode", params)
	result := dom.ResolveNodeResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
SetAttributeValue sets attribute for an element with given id.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setAttributeValue
*/
func (DOM) SetAttributeValue(
	socket sock.Socketer,
	params *dom.SetAttributeValueParams,
) error {
	command := sock.NewCommand("DOM.setAttributeValue", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetAttributesAsText sets attributes on element with given id. This method is useful when user edits
some existing attribute value and types in several attribute name/value pairs.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setAttributesAsText
*/
func (DOM) SetAttributesAsText(
	socket sock.Socketer,
	params *dom.SetAttributesAsTextParams,
) error {
	command := sock.NewCommand("DOM.setAttributesAsText", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetFileInputFiles sets files for the given file input element.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setFileInputFiles
*/
func (DOM) SetFileInputFiles(
	socket sock.Socketer,
	params *dom.SetFileInputFilesParams,
) error {
	command := sock.NewCommand("DOM.setFileInputFiles", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetInspectedNode enables console to refer to the node with given id via $x (see Command Line API for
more details $x functions). EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setInspectedNode
*/
func (DOM) SetInspectedNode(
	socket sock.Socketer,
	params *dom.SetInspectedNodeParams,
) error {
	command := sock.NewCommand("DOM.setInspectedNode", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetNodeName sets node name for the specified node.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setNodeName
*/
func (DOM) SetNodeName(
	socket sock.Socketer,
	params *dom.SetNodeNameParams,
) (dom.SetNodeNameResult, error) {
	command := sock.NewCommand("DOM.setNodeName", params)
	result := dom.SetNodeNameResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
SetNodeValue sets node value for the specified node.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setNodeValue
*/
func (DOM) SetNodeValue(
	socket sock.Socketer,
	params *dom.SetNodeValueParams,
) error {
	command := sock.NewCommand("DOM.setNodeValue", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetOuterHTML sets node HTML markup, returns new node id.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setOuterHTML
*/
func (DOM) SetOuterHTML(
	socket sock.Socketer,
	params *dom.SetOuterHTMLParams,
) error {
	command := sock.NewCommand("DOM.setOuterHTML", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
Undo undoes the last performed action. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-undo
*/
func (DOM) Undo(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("DOM.undo", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
OnAttributeModified adds a handler to the DOM.attributeModified event. DOM.attributeModified fires
when Element's attribute is modified.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-attributeModified
*/
func (DOM) OnAttributeModified(
	socket sock.Socketer,
	callback func(event *dom.AttributeModifiedEvent),
) {
	handler := sock.NewEventHandler(
		"DOM.attributeModified",
		func(response *sock.Response) {
			event := &dom.AttributeModifiedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnAttributeRemoved adds a handler to the DOM.attributeRemoved event. DOM.attributeRemoved fires when
Element's attribute is modified.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-attributeRemoved
*/
func (DOM) OnAttributeRemoved(
	socket sock.Socketer,
	callback func(event *dom.AttributeRemovedEvent),
) {
	handler := sock.NewEventHandler(
		"DOM.attributeRemoved",
		func(response *sock.Response) {
			event := &dom.AttributeRemovedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnCharacterDataModified adds a handler to the DOM.characterDataModified event.
DOM.characterDataModified mirrors the DOMCharacterDataModified event.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-characterDataModified
*/
func (DOM) OnCharacterDataModified(
	socket sock.Socketer,
	callback func(event *dom.CharacterDataModifiedEvent),
) {
	handler := sock.NewEventHandler(
		"DOM.characterDataModified",
		func(response *sock.Response) {
			event := &dom.CharacterDataModifiedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnChildNodeCountUpdated adds a handler to the DOM.childNodeCountUpdated event.
DOM.childNodeCountUpdated fires when Container's child node count has changed.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-childNodeCountUpdated
*/
func (DOM) OnChildNodeCountUpdated(
	socket sock.Socketer,
	callback func(event *dom.ChildNodeCountUpdatedEvent),
) {
	handler := sock.NewEventHandler(
		"DOM.childNodeCountUpdated",
		func(response *sock.Response) {
			event := &dom.ChildNodeCountUpdatedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnChildNodeInserted adds a handler to the DOM.childNodeInserted event. DOM.childNodeInserted mirrors
the DOMNodeInserted event.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-childNodeInserted
*/
func (DOM) OnChildNodeInserted(
	socket sock.Socketer,
	callback func(event *dom.ChildNodeInsertedEvent),
) {
	handler := sock.NewEventHandler(
		"DOM.childNodeInserted",
		func(response *sock.Response) {
			event := &dom.ChildNodeInsertedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnChildNodeRemoved adds a handler to the DOM.childNodeRemoved event.DOM.childNodeRemoved mirrors the
DOMNodeRemoved event.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-childNodeRemoved
*/
func (DOM) OnChildNodeRemoved(
	socket sock.Socketer,
	callback func(event *dom.ChildNodeRemovedEvent),
) {
	handler := sock.NewEventHandler(
		"DOM.childNodeRemoved",
		func(response *sock.Response) {
			event := &dom.ChildNodeRemovedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnDistributedNodesUpdated adds a handler to the DOM.distributedNodesUpdated event.
DOM.distributedNodesUpdated fires when distribution is changed. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-distributedNodesUpdated
*/
func (DOM) OnDistributedNodesUpdated(
	socket sock.Socketer,
	callback func(event *dom.DistributedNodesUpdatedEvent),
) {
	handler := sock.NewEventHandler(
		"DOM.distributedNodesUpdated",
		func(response *sock.Response) {
			event := &dom.DistributedNodesUpdatedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnDocumentUpdated adds a handler to the DOM.documentUpdated event. DOM.documentUpdated
fires when Document has been totally updated. Node IDs are no longer valid.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-documentUpdated
*/
func (DOM) OnDocumentUpdated(
	socket sock.Socketer,
	callback func(event *dom.DocumentUpdatedEvent),
) {
	handler := sock.NewEventHandler(
		"DOM.documentUpdated",
		func(response *sock.Response) {
			event := &dom.DocumentUpdatedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnInlineStyleInvalidated adds a handler to the DOM.inlineStyleInvalidated event.
DOM.inlineStyleInvalidated fires when Element's attribute is removed.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-inlineStyleInvalidated
*/
func (DOM) OnInlineStyleInvalidated(
	socket sock.Socketer,
	callback func(event *dom.InlineStyleInvalidatedEvent),
) {
	handler := sock.NewEventHandler(
		"DOM.inlineStyleInvalidated",
		func(response *sock.Response) {
			event := &dom.InlineStyleInvalidatedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnPseudoElementAdded adds a handler to the DOM.pseudoElementAdded event. DOM.pseudoElementAdded
fires when a pseudo element is added to an element. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-pseudoElementAdded
*/
func (DOM) OnPseudoElementAdded(
	socket sock.Socketer,
	callback func(event *dom.PseudoElementAddedEvent),
) {
	handler := sock.NewEventHandler(
		"DOM.pseudoElementAdded",
		func(response *sock.Response) {
			event := &dom.PseudoElementAddedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnPseudoElementRemoved adds a handler to the DOM.pseudoElementRemoved event.
DOM.pseudoElementRemoved fires when a pseudo element is removed from an element. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-pseudoElementRemoved
*/
func (DOM) OnPseudoElementRemoved(
	socket sock.Socketer,
	callback func(event *dom.PseudoElementRemovedEvent),
) {
	handler := sock.NewEventHandler(
		"DOM.pseudoElementRemoved",
		func(response *sock.Response) {
			event := &dom.PseudoElementRemovedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnSetChildNodes adds a handler to the DOM.setChildNodes event. DOM.setChildNodes fires
when backend wants to provide client with the missing DOM structure. This happens upon most of the
calls requesting node IDs.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-setChildNodes
*/
func (DOM) OnSetChildNodes(
	socket sock.Socketer,
	callback func(event *dom.SetChildNodesEvent),
) {
	handler := sock.NewEventHandler(
		"DOM.setChildNodes",
		func(response *sock.Response) {
			event := &dom.SetChildNodesEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnShadowRootPopped adds a handler to the DOM.shadowRootPopped event. DOM.shadowRootPopped fires when
shadow root is popped from the element. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-shadowRootPopped
*/
func (DOM) OnShadowRootPopped(
	socket sock.Socketer,
	callback func(event *dom.ShadowRootPoppedEvent),
) {
	handler := sock.NewEventHandler(
		"DOM.shadowRootPopped",
		func(response *sock.Response) {
			event := &dom.ShadowRootPoppedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnShadowRootPushed adds a handler to the DOM.shadowRootPushed event. DOM.shadowRootPushed fires when
shadow root is pushed into the element. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-shadowRootPushed
*/
func (DOM) OnShadowRootPushed(
	socket sock.Socketer,
	callback func(event *dom.ShadowRootPushedEvent),
) {
	handler := sock.NewEventHandler(
		"DOM.shadowRootPushed",
		func(response *sock.Response) {
			event := &dom.ShadowRootPushedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

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
	socket *sock.Socket,
	params *dom.CollectClassNamesFromSubtreeParams,
) (dom.CollectClassNamesFromSubtreeResult, error) {
	command := &sock.Command{
		Method: "DOM.collectClassNamesFromSubtree",
		Params: params,
	}
	result := dom.CollectClassNamesFromSubtreeResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
CopyTo creates a deep copy of the specified node and places it into the target container before the
given anchor. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-copyTo
*/
func (DOM) CopyTo(
	socket *sock.Socket,
	params *dom.CopyToParams,
) (dom.CopyToResult, error) {
	command := &sock.Command{
		Method: "DOM.copyTo",
		Params: params,
	}
	result := dom.CopyToResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
DescribeNode describes node given its id, does not require domain to be enabled. Does not start
tracking any objects, can be used for automation.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-describeNode
*/
func (DOM) DescribeNode(
	socket *sock.Socket,
	params *dom.DescribeNodeParams,
) (dom.DescribeNodeResult, error) {
	command := &sock.Command{
		Method: "DOM.describeNode",
		Params: params,
	}
	result := dom.DescribeNodeResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
Disable disables the DOM agent for the given page.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-disable
*/
func (DOM) Disable(
	socket *sock.Socket,
) error {
	command := &sock.Command{
		Method: "DOM.disable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
DiscardSearchResults discards search results from the session with the given id. getSearchResults
should no longer be called for that search. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-discardSearchResults
*/
func (DOM) DiscardSearchResults(
	socket *sock.Socket,
	params *dom.DiscardSearchResultsParams,
) error {
	command := &sock.Command{
		Method: "DOM.discardSearchResults",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Enable enables the DOM agent for the given page.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-enable
*/
func (DOM) Enable(
	socket *sock.Socket,
) error {
	command := &sock.Command{
		Method: "DOM.enable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Focus focuses the given element.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-focus
*/
func (DOM) Focus(
	socket *sock.Socket,
	params *dom.FocusParams,
) error {
	command := &sock.Command{
		Method: "DOM.focus",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
GetAttributes returns attributes for the specified node.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getAttributes
*/
func (DOM) GetAttributes(
	socket *sock.Socket,
	params *dom.GetAttributesParams,
) (dom.GetAttributesResult, error) {
	command := &sock.Command{
		Method: "DOM.getAttributes",
		Params: params,
	}
	result := dom.GetAttributesResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
GetBoxModel returns boxes for the given node.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getBoxModel
*/
func (DOM) GetBoxModel(
	socket *sock.Socket,
	params *dom.GetBoxModelParams,
) (dom.GetBoxModelResult, error) {
	command := &sock.Command{
		Method: "DOM.getBoxModel",
		Params: params,
	}
	result := dom.GetBoxModelResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
GetDocument returns the root DOM node (and optionally the subtree) to the caller.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getDocument
*/
func (DOM) GetDocument(
	socket *sock.Socket,
	params *dom.GetDocumentParams,
) (dom.GetDocumentResult, error) {
	command := &sock.Command{
		Method: "DOM.getDocument",
		Params: params,
	}
	result := dom.GetDocumentResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
GetFlattenedDocument returns the root DOM node (and optionally the subtree) to the caller.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getFlattenedDocument
*/
func (DOM) GetFlattenedDocument(
	socket *sock.Socket,
	params *dom.GetFlattenedDocumentParams,
) (dom.GetFlattenedDocumentResult, error) {
	command := &sock.Command{
		Method: "DOM.getFlattenedDocument",
		Params: params,
	}
	result := dom.GetFlattenedDocumentResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
GetNodeForLocation returns node id at given location. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getNodeForLocation
*/
func (DOM) GetNodeForLocation(
	socket *sock.Socket,
	params *dom.GetNodeForLocationParams,
) (dom.GetNodeForLocationResult, error) {
	command := &sock.Command{
		Method: "DOM.getNodeForLocation",
		Params: params,
	}
	result := dom.GetNodeForLocationResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
GetOuterHTML returns node's HTML markup.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getOuterHTML
*/
func (DOM) GetOuterHTML(
	socket *sock.Socket,
	params *dom.GetOuterHTMLParams,
) (dom.GetOuterHTMLResult, error) {
	command := &sock.Command{
		Method: "DOM.getOuterHTML",
		Params: params,
	}
	result := dom.GetOuterHTMLResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
GetRelayoutBoundary returns the id of the nearest ancestor that is a relayout boundary. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getRelayoutBoundary
*/
func (DOM) GetRelayoutBoundary(
	socket *sock.Socket,
	params *dom.GetRelayoutBoundaryParams,
) (dom.GetRelayoutBoundaryResult, error) {
	command := &sock.Command{
		Method: "DOM.getRelayoutBoundary",
		Params: params,
	}
	result := dom.GetRelayoutBoundaryResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
GetSearchResults returns search results from given fromIndex to given toIndex from the search with
the given identifier. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getSearchResults
*/
func (DOM) GetSearchResults(
	socket *sock.Socket,
	params *dom.GetSearchResultsParams,
) (dom.GetSearchResultsResult, error) {
	command := &sock.Command{
		Method: "DOM.getSearchResults",
		Params: params,
	}
	result := dom.GetSearchResultsResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
MarkUndoableState marks last undoable state. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-markUndoableState
*/
func (DOM) MarkUndoableState(
	socket *sock.Socket,
) error {
	command := &sock.Command{
		Method: "DOM.markUndoableState",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
MoveTo moves node into the new container, places it before the given anchor.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-moveTo
*/
func (DOM) MoveTo(
	socket *sock.Socket,
	params *dom.MoveToParams,
) (dom.MoveToResult, error) {
	command := &sock.Command{
		Method: "DOM.moveTo",
		Params: params,
	}
	result := dom.MoveToResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
PerformSearch searches for a given string in the DOM tree. Use getSearchResults to access search
results or cancelSearch to end this search session. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-performSearch
*/
func (DOM) PerformSearch(
	socket *sock.Socket,
	params *dom.PerformSearchParams,
) (dom.PerformSearchResult, error) {
	command := &sock.Command{
		Method: "DOM.performSearch",
		Params: params,
	}
	result := dom.PerformSearchResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
PushNodeByPathToFrontend requests that the node is sent to the caller given its path. EXPERIMENTAL
@TODO, use XPath.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-pushNodeByPathToFrontend
*/
func (DOM) PushNodeByPathToFrontend(
	socket *sock.Socket,
	params *dom.PushNodeByPathToFrontendParams,
) (dom.PushNodeByPathToFrontendResult, error) {
	command := &sock.Command{
		Method: "DOM.pushNodeByPathToFrontend",
		Params: params,
	}
	result := dom.PushNodeByPathToFrontendResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
PushNodesByBackendIDsToFrontend requests that a batch of nodes is sent to the caller given their
backend node IDs. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-pushNodesByBackendIdsToFrontend
*/
func (DOM) PushNodesByBackendIDsToFrontend(
	socket *sock.Socket,
	params *dom.PushNodesByBackendIDsToFrontendParams,
) (dom.PushNodesByBackendIDsToFrontendResult, error) {
	command := &sock.Command{
		Method: "DOM.pushNodesByBackendIdsToFrontend",
		Params: params,
	}
	result := dom.PushNodesByBackendIDsToFrontendResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
QuerySelector executes querySelector on a given node.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-querySelector
*/
func (DOM) QuerySelector(
	socket *sock.Socket,
	params *dom.QuerySelectorParams,
) (dom.QuerySelectorResult, error) {
	command := &sock.Command{
		Method: "DOM.querySelector",
		Params: params,
	}
	result := dom.QuerySelectorResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
QuerySelectorAll executes querySelectorAll on a given node.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-querySelectorAll
*/
func (DOM) QuerySelectorAll(
	socket *sock.Socket,
	params *dom.QuerySelectorAllParams,
) (dom.QuerySelectorAllResult, error) {
	command := &sock.Command{
		Method: "DOM.querySelectorAll",
		Params: params,
	}
	result := dom.QuerySelectorAllResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
Redo re-does the last undone action. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-redo
*/
func (DOM) Redo(
	socket *sock.Socket,
) error {
	command := &sock.Command{
		Method: "DOM.redo",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
RemoveAttribute removes attribute with given name from an element with given id.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-removeAttribute
*/
func (DOM) RemoveAttribute(
	socket *sock.Socket,
	params *dom.RemoveAttributeParams,
) error {
	command := &sock.Command{
		Method: "DOM.removeAttribute",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
RemoveNode removes the specified node.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-removeNode
*/
func (DOM) RemoveNode(
	socket *sock.Socket,
	params *dom.RemoveNodeParams,
) error {
	command := &sock.Command{
		Method: "DOM.removeNode",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
RequestChildNodes requests that children of the node with given id are returned to the caller in
form of setChildNodes events where not only immediate children are retrieved, but all children down
to the specified depth.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-requestChildNodes
*/
func (DOM) RequestChildNodes(
	socket *sock.Socket,
	params *dom.RequestChildNodesParams,
) error {
	command := &sock.Command{
		Method: "DOM.requestChildNodes",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
RequestNode requests that the node is sent to the caller given the JavaScript node object reference.
All nodes that form the path from the node to the root are also sent to the client as a series of
setChildNodes notifications.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-requestNode
*/
func (DOM) RequestNode(
	socket *sock.Socket,
	params *dom.RequestNodeParams,
) (dom.RequestNodeResult, error) {
	command := &sock.Command{
		Method: "DOM.requestNode",
		Params: params,
	}
	result := dom.RequestNodeResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
ResolveNode resolves the JavaScript node object for a given NodeID or BackendNodeID.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-resolveNode
*/
func (DOM) ResolveNode(
	socket *sock.Socket,
	params *dom.ResolveNodeParams,
) (dom.ResolveNodeResult, error) {
	command := &sock.Command{
		Method: "DOM.resolveNode",
		Params: params,
	}
	result := dom.ResolveNodeResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
SetAttributeValue sets attribute for an element with given id.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setAttributeValue
*/
func (DOM) SetAttributeValue(
	socket *sock.Socket,
	params *dom.SetAttributeValueParams,
) error {
	command := &sock.Command{
		Method: "DOM.setAttributeValue",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetAttributesAsText sets attributes on element with given id. This method is useful when user edits
some existing attribute value and types in several attribute name/value pairs.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setAttributesAsText
*/
func (DOM) SetAttributesAsText(
	socket *sock.Socket,
	params *dom.SetAttributesAsTextParams,
) error {
	command := &sock.Command{
		Method: "DOM.setAttributesAsText",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetFileInputFiles sets files for the given file input element.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setFileInputFiles
*/
func (DOM) SetFileInputFiles(
	socket *sock.Socket,
	params *dom.SetFileInputFilesParams,
) error {
	command := &sock.Command{
		Method: "DOM.setFileInputFiles",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetInspectedNode enables console to refer to the node with given id via $x (see Command Line API for
more details $x functions). EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setInspectedNode
*/
func (DOM) SetInspectedNode(
	socket *sock.Socket,
	params *dom.SetInspectedNodeParams,
) error {
	command := &sock.Command{
		Method: "DOM.setInspectedNode",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetNodeName sets node name for the specified node.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setNodeName
*/
func (DOM) SetNodeName(
	socket *sock.Socket,
	params *dom.SetNodeNameParams,
) (dom.SetNodeNameResult, error) {
	command := &sock.Command{
		Method: "DOM.setNodeName",
		Params: params,
	}
	result := dom.SetNodeNameResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
SetNodeValue sets node value for the specified node.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setNodeValue
*/
func (DOM) SetNodeValue(
	socket *sock.Socket,
	params *dom.SetNodeValueParams,
) error {
	command := &sock.Command{
		Method: "DOM.setNodeValue",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetOuterHTML sets node HTML markup, returns new node id.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setOuterHTML
*/
func (DOM) SetOuterHTML(
	socket *sock.Socket,
	params *dom.SetOuterHTMLParams,
) error {
	command := &sock.Command{
		Method: "DOM.setOuterHTML",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Undo undoes the last performed action. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-undo
*/
func (DOM) Undo(
	socket *sock.Socket,
) error {
	command := &sock.Command{
		Method: "DOM.undo",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
OnAttributeModified adds a handler to the DOM.attributeModified event. DOM.attributeModified fires
when Element's attribute is modified.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-attributeModified
*/
func (DOM) OnAttributeModified(
	socket *sock.Socket,
	callback func(event *dom.AttributeModifiedEvent),
) {
	handler := sock.NewEventHandler(
		"DOM.attributeModified",
		func(name string, params []byte) {
			event := &dom.AttributeModifiedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
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
	socket *sock.Socket,
	callback func(event *dom.AttributeRemovedEvent),
) {
	handler := sock.NewEventHandler(
		"DOM.attributeRemoved",
		func(name string, params []byte) {
			event := &dom.AttributeRemovedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
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
	socket *sock.Socket,
	callback func(event *dom.CharacterDataModifiedEvent),
) {
	handler := sock.NewEventHandler(
		"DOM.characterDataModified",
		func(name string, params []byte) {
			event := &dom.CharacterDataModifiedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
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
	socket *sock.Socket,
	callback func(event *dom.ChildNodeCountUpdatedEvent),
) {
	handler := sock.NewEventHandler(
		"DOM.childNodeCountUpdated",
		func(name string, params []byte) {
			event := &dom.ChildNodeCountUpdatedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
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
	socket *sock.Socket,
	callback func(event *dom.ChildNodeInsertedEvent),
) {
	handler := sock.NewEventHandler(
		"DOM.childNodeInserted",
		func(name string, params []byte) {
			event := &dom.ChildNodeInsertedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
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
	socket *sock.Socket,
	callback func(event *dom.ChildNodeRemovedEvent),
) {
	handler := sock.NewEventHandler(
		"DOM.childNodeRemoved",
		func(name string, params []byte) {
			event := &dom.ChildNodeRemovedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
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
	socket *sock.Socket,
	callback func(event *dom.DistributedNodesUpdatedEvent),
) {
	handler := sock.NewEventHandler(
		"DOM.distributedNodesUpdated",
		func(name string, params []byte) {
			event := &dom.DistributedNodesUpdatedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
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
	socket *sock.Socket,
	callback func(event *dom.DocumentUpdatedEvent),
) {
	handler := sock.NewEventHandler(
		"DOM.documentUpdated",
		func(name string, params []byte) {
			event := &dom.DocumentUpdatedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
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
	socket *sock.Socket,
	callback func(event *dom.InlineStyleInvalidatedEvent),
) {
	handler := sock.NewEventHandler(
		"DOM.inlineStyleInvalidated",
		func(name string, params []byte) {
			event := &dom.InlineStyleInvalidatedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
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
	socket *sock.Socket,
	callback func(event *dom.PseudoElementAddedEvent),
) {
	handler := sock.NewEventHandler(
		"DOM.pseudoElementAdded",
		func(name string, params []byte) {
			event := &dom.PseudoElementAddedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
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
	socket *sock.Socket,
	callback func(event *dom.PseudoElementRemovedEvent),
) {
	handler := sock.NewEventHandler(
		"DOM.pseudoElementRemoved",
		func(name string, params []byte) {
			event := &dom.PseudoElementRemovedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
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
	socket *sock.Socket,
	callback func(event *dom.SetChildNodesEvent),
) {
	handler := sock.NewEventHandler(
		"DOM.setChildNodes",
		func(name string, params []byte) {
			event := &dom.SetChildNodesEvent{}
			if err := json.Unmarshal(params, event); err != nil {
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
	socket *sock.Socket,
	callback func(event *dom.ShadowRootPoppedEvent),
) {
	handler := sock.NewEventHandler(
		"DOM.shadowRootPopped",
		func(name string, params []byte) {
			event := &dom.ShadowRootPoppedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
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
	socket *sock.Socket,
	callback func(event *dom.ShadowRootPushedEvent),
) {
	handler := sock.NewEventHandler(
		"DOM.shadowRootPushed",
		func(name string, params []byte) {
			event := &dom.ShadowRootPushedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

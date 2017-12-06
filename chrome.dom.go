package chrome

import (
	"app/chrome/protocol"
	"encoding/json"

	log "github.com/Sirupsen/logrus"
)

/*
DOM - https://chromedevtools.github.io/devtools-protocol/tot/DOM/
Exposes DOM read/write operations. Each DOM Node is represented with its mirror object that has an
ID. This ID can be used to get additional information on the Node, resolve it into the JavaScript
object wrapper, etc. It is important that client receives DOM events only for the nodes that are
known to the client. Backend keeps track of the nodes that were sent to the client and never sends
the same node twice. It is client's responsibility to collect information about the nodes that were
sent to the client.

Note that iframe owner elements will return corresponding document elements as their child nodes.
*/
type DOM struct{}

/*
CollectClassNamesFromSubtree creates a deep copy of the specified node and places it into the target container before the
given anchor. EXPERIMENTAL
*/
func (DOM) CollectClassNamesFromSubtree(
	socket *Socket,
	params *dom.CollectClassNamesFromSubtreeParams,
) (dom.CollectClassNamesFromSubtreeResult, error) {
	command := &protocol.Command{
		method: "DOM.collectClassNamesFromSubtree",
		params: params,
	}
	socket.SendCommand(command)
	return command.Result.(dom.CollectClassNamesFromSubtreeResult), command.Err
}

/*
CopyTo creates a deep copy of the specified node and places it into the target container before the
given anchor. EXPERIMENTAL
*/
func (DOM) CopyTo(
	socket *Socket,
	params *dom.CopyToParams,
) (dom.CopyToResult, error) {
	command := &protocol.Command{
		method: "DOM.copyTo",
		params: params,
	}
	socket.SendCommand(command)
	return command.Result.(dom.CopyToResult), command.Err
}

/*
DescribeNode describes node given its id, does not require domain to be enabled. Does not start
tracking any objects, can be used for automation.
*/
func (DOM) DescribeNode(
	socket *Socket,
	params *dom.DescribeNodeParams,
) (dom.DescribeNodeResult, error) {
	command := &protocol.Command{
		method: "DOM.describeNode",
		params: params,
	}
	socket.SendCommand(command)
	return command.Result.(dom.DescribeNodeResult), command.Err
}

/*
Disable disables the DOM agent for the given page.
*/
func (DOM) Disable(
	socket *Socket,
) (nil, error) {
	command := &protocol.Command{
		method: "DOM.disable",
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
DiscardSearchResults discards search results from the session with the given id. getSearchResults
should no longer be called for that search. EXPERIMENTAL
*/
func (DOM) DiscardSearchResults(
	socket *Socket,
	params *dom.DiscardSearchResultsParams,
) (nil, error) {
	command := &protocol.Command{
		method: "DOM.discardSearchResults",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
Enable enables the DOM agent for the given page.
*/
func (DOM) Enable(
	socket *Socket,
) (nil, error) {
	command := &protocol.Command{
		method: "DOM.enable",
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
Focus focuses the given element.
*/
func (DOM) Focus(
	socket *Socket,
	params *dom.FocusParams,
) (nil, error) {
	command := &protocol.Command{
		method: "DOM.focus",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
GetAttributes returns attributes for the specified node.
*/
func (DOM) GetAttributes(
	socket *Socket,
	params *dom.GetAttributesParams,
) (dom.GetAttributesResult, error) {
	command := &protocol.Command{
		method: "DOM.getAttributes",
		params: params,
	}
	socket.SendCommand(command)
	return command.Result.(dom.GetAttributesResult), command.Err
}

/*
GetBoxModel returns boxes for the given node.
*/
func (DOM) GetBoxModel(
	socket *Socket,
	params *dom.GetBoxModelParams,
) (dom.GetBoxModelResult, error) {
	command := &protocol.Command{
		method: "DOM.getBoxModel",
		params: params,
	}
	socket.SendCommand(command)
	return command.Result.(dom.GetBoxModelResult), command.Err
}

/*
GetDocument returns the root DOM node (and optionally the subtree) to the caller.
*/
func (DOM) GetDocument(
	socket *Socket,
	params *dom.GetDocumentParams,
) (dom.GetDocumentResult, error) {
	command := &protocol.Command{
		method: "DOM.getDocument",
		params: params,
	}
	socket.SendCommand(command)
	return command.Result.(dom.GetDocumentResult), command.Err
}

/*
GetFlattenedDocument returns the root DOM node (and optionally the subtree) to the caller.
*/
func (DOM) GetFlattenedDocument(
	socket *Socket,
	params *dom.GetFlattenedDocumentParams,
) (dom.GetFlattenedDocumentResult, error) {
	command := &protocol.Command{
		method: "DOM.getFlattenedDocument",
		params: params,
	}
	socket.SendCommand(command)
	return command.Result.(dom.GetFlattenedDocumentResult), command.Err
}

/*
GetNodeForLocation returns node id at given location. EXPERIMENTAL
*/
func (DOM) GetNodeForLocation(
	socket *Socket,
	params *dom.GetNodeForLocationParams,
) (dom.GetNodeForLocationResult, error) {
	command := &protocol.Command{
		method: "DOM.getNodeForLocation",
		params: params,
	}
	socket.SendCommand(command)
	return command.Result.(dom.GetNodeForLocationResult), command.Err
}

/*
GetOuterHTML returns node's HTML markup.
*/
func (DOM) GetOuterHTML(
	socket *Socket,
	params *dom.GetOuterHTMLParams,
) (dom.GetOuterHTMLResult, error) {
	command := &protocol.Command{
		method: "DOM.getOuterHTML",
		params: params,
	}
	socket.SendCommand(command)
	return command.Result.(dom.GetOuterHTMLResult), command.Err
}

/*
GetRelayoutBoundary returns the id of the nearest ancestor that is a relayout boundary. EXPERIMENTAL
*/
func (DOM) GetRelayoutBoundary(
	socket *Socket,
	params *dom.GetRelayoutBoundaryParams,
) (dom.GetRelayoutBoundaryResult, error) {
	command := &protocol.Command{
		method: "DOM.getRelayoutBoundary",
		params: params,
	}
	socket.SendCommand(command)
	return command.Result.(dom.GetRelayoutBoundaryResult), command.Err
}

/*
GetSearchResults returns search results from given fromIndex to given toIndex from the search with
the given identifier. EXPERIMENTAL
*/
func (DOM) GetSearchResults(
	socket *Socket,
	params *dom.GetSearchResultsParams,
) (dom.GetSearchResultsResult, error) {
	command := &protocol.Command{
		method: "DOM.getSearchResults",
		params: params,
	}
	socket.SendCommand(command)
	return command.Result.(dom.GetSearchResultsResult), command.Err
}

/*
MarkUndoableState marks last undoable state. EXPERIMENTAL
*/
func (DOM) MarkUndoableState(
	socket *Socket,
) (nil, error) {
	command := &protocol.Command{
		method: "DOM.markUndoableState",
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
MoveTo moves node into the new container, places it before the given anchor.
*/
func (DOM) MoveTo(
	socket *Socket,
	params *dom.MoveToParams,
) (dom.MoveToResult, error) {
	command := &protocol.Command{
		method: "DOM.moveTo",
		params: params,
	}
	socket.SendCommand(command)
	return command.Result.(dom.MoveToResult), command.Err
}

/*
PerformSearch searches for a given string in the DOM tree. Use getSearchResults to access search
results or cancelSearch to end this search session. EXPERIMENTAL
*/
func (DOM) PerformSearch(
	socket *Socket,
	params *dom.PerformSearchParams,
) (dom.PerformSearchResult, error) {
	command := &protocol.Command{
		method: "DOM.performSearch",
		params: params,
	}
	socket.SendCommand(command)
	return command.Result.(dom.PerformSearchResult), command.Err
}

/*
PushNodeByPathToFrontend requests that the node is sent to the caller given its path. EXPERIMENTAL
@TODO, use XPath.
*/
func (DOM) PushNodeByPathToFrontend(
	socket *Socket,
	params *dom.PushNodeByPathToFrontendParams,
) (dom.PushNodeByPathToFrontendResult, error) {
	command := &protocol.Command{
		method: "DOM.pushNodeByPathToFrontend",
		params: params,
	}
	socket.SendCommand(command)
	return command.Result.(dom.PushNodeByPathToFrontendResult), command.Err
}

/*
PushNodesByBackendIDsToFrontend requests that a batch of nodes is sent to the caller given their
backend node IDs. EXPERIMENTAL
*/
func (DOM) PushNodesByBackendIDsToFrontend(
	socket *Socket,
	params *dom.PushNodesByBackendIDsToFrontendParams,
) (dom.PushNodesByBackendIDsToFrontendResult, error) {
	command := &protocol.Command{
		method: "DOM.pushNodesByBackendIdsToFrontend",
		params: params,
	}
	socket.SendCommand(command)
	return command.Result.(dom.PushNodesByBackendIDsToFrontendResult), command.Err
}

/*
QuerySelector executes querySelector on a given node.
*/
func (DOM) QuerySelector(
	socket *Socket,
	params *dom.QuerySelectorParams,
) (dom.QuerySelectorResult, error) {
	command := &protocol.Command{
		method: "DOM.querySelector",
		params: params,
	}
	socket.SendCommand(command)
	return command.Result.(dom.QuerySelectorResult), command.Err
}

/*
QuerySelectorAll executes querySelectorAll on a given node.
*/
func (DOM) QuerySelectorAll(
	socket *Socket,
	params *dom.QuerySelectorAllParams,
) (dom.QuerySelectorAllResult, error) {
	command := &protocol.Command{
		method: "DOM.querySelectorAll",
		params: params,
	}
	socket.SendCommand(command)
	return command.Result.(dom.QuerySelectorAllResult), command.Err
}

/*
Redo re-does the last undone action. EXPERIMENTAL
*/
func (DOM) Redo(
	socket *Socket,
) (nil, error) {
	command := &protocol.Command{
		method: "DOM.redo",
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
RemoveAttribute removes attribute with given name from an element with given id.
*/
func (DOM) RemoveAttribute(
	socket *Socket,
	params *dom.RemoveAttributeParams,
) (nil, error) {
	command := &protocol.Command{
		method: "DOM.removeAttribute",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
RemoveNode removes the specified node.
*/
func (DOM) RemoveNode(
	socket *Socket,
	params *dom.RemoveNodeParams,
) (nil, error) {
	command := &protocol.Command{
		method: "DOM.removeNode",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
RequestChildNodes requests that children of the node with given id are returned to the caller in
form of setChildNodes events where not only immediate children are retrieved, but all children down
to the specified depth.
*/
func (DOM) RequestChildNodes(
	socket *Socket,
	params *dom.RequestChildNodesParams,
) (nil, error) {
	command := &protocol.Command{
		method: "DOM.requestChildNodes",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
RequestNode requests that the node is sent to the caller given the JavaScript node object reference.
All nodes that form the path from the node to the root are also sent to the client as a series of
setChildNodes notifications.
*/
func (DOM) RequestNode(
	socket *Socket,
	params *dom.RequestNodeParams,
) (dom.RequestNodeResult, error) {
	command := &protocol.Command{
		method: "DOM.requestNode",
		params: params,
	}
	socket.SendCommand(command)
	return command.Result.(dom.RequestNodeResult), command.Err
}

/*
ResolveNode resolves the JavaScript node object for a given NodeID or BackendNodeID.
*/
func (DOM) ResolveNode(
	socket *Socket,
	params *dom.ResolveNodeParams,
) (dom.ResolveNodeResult, error) {
	command := &protocol.Command{
		method: "DOM.resolveNode",
		params: params,
	}
	socket.SendCommand(command)
	return command.Result.(dom.ResolveNodeResult), command.Err
}

/*
SetAttributeValue sets attribute for an element with given id.
*/
func (DOM) SetAttributeValue(
	socket *Socket,
	params *dom.SetAttributeValueParams,
) (nil, error) {
	command := &protocol.Command{
		method: "DOM.setAttributeValue",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
SetAttributesAsText sets attributes on element with given id. This method is useful when user edits
some existing attribute value and types in several attribute name/value pairs.
*/
func (DOM) SetAttributesAsText(
	socket *Socket,
	params *dom.SetAttributesAsTextParams,
) (nil, error) {
	command := &protocol.Command{
		method: "DOM.setAttributesAsText",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
SetFileInputFiles sets files for the given file input element.
*/
func (DOM) SetFileInputFiles(
	socket *Socket,
	params *dom.SetFileInputFilesParams,
) (nil, error) {
	command := &protocol.Command{
		method: "DOM.setFileInputFiles",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
SetInspectedNode enables console to refer to the node with given id via $x (see Command Line API for
more details $x functions). EXPERIMENTAL
*/
func (DOM) SetInspectedNode(
	socket *Socket,
	params *dom.SetInspectedNodeParams,
) (nil, error) {
	command := &protocol.Command{
		method: "DOM.setInspectedNode",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
SetNodeName sets node name for the specified node.
*/
func (DOM) SetNodeName(
	socket *Socket,
	params *dom.SetNodeNameParams,
) (dom.SetNodeNameResult, error) {
	command := &protocol.Command{
		method: "DOM.setNodeName",
		params: params,
	}
	socket.SendCommand(command)
	return command.Result.(dom.SetNodeNameResult), command.Err
}

/*
SetNodeValue sets node value for the specified node.
*/
func (DOM) SetNodeValue(
	socket *Socket,
	params *dom.SetNodeValueParams,
) (nil, error) {
	command := &protocol.Command{
		method: "DOM.setNodeValue",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
SetOuterHTML sets node HTML markup, returns new node id.
*/
func (DOM) SetOuterHTML(
	socket *Socket,
	params *dom.SetOuterHTMLParams,
) (nil, error) {
	command := &protocol.Command{
		method: "DOM.setOuterHTML",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
Undo undoes the last performed action. EXPERIMENTAL
*/
func (DOM) Undo(
	socket *Socket,
) (nil, error) {
	command := &protocol.Command{
		method: "DOM.undo",
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
OnAttributeModified adds a handler to the DOM.attributeModified event. DOM.attributeModified fires
when Element's attribute is modified.
*/
func (DOM) OnAttributeModified(
	socket *Socket,
	callback func(event *dom.AttributeModifiedEvent),
) {
	handler := protocol.NewEventHandler(
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
*/
func (DOM) OnAttributeRemoved(
	socket *Socket,
	callback func(event *dom.AttributeRemovedEvent),
) {
	handler := protocol.NewEventHandler(
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
*/
func (DOM) OnCharacterDataModified(
	socket *Socket,
	callback func(event *dom.CharacterDataModifiedEvent),
) {
	handler := protocol.NewEventHandler(
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
*/
func (DOM) OnChildNodeCountUpdated(
	socket *Socket,
	callback func(event *dom.ChildNodeCountUpdatedEvent),
) {
	handler := protocol.NewEventHandler(
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
*/
func (DOM) OnChildNodeInserted(
	socket *Socket,
	callback func(event *dom.ChildNodeInsertedEvent),
) {
	handler := protocol.NewEventHandler(
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
*/
func (DOM) OnChildNodeRemoved(
	socket *Socket,
	callback func(event *dom.ChildNodeRemovedEvent),
) {
	handler := protocol.NewEventHandler(
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
DOM.distributedNodesUpdated fires when distrubution is changed. EXPERIMENTAL
*/
func (DOM) OnDistributedNodesUpdated(
	socket *Socket,
	callback func(event *dom.DistributedNodesUpdatedEvent),
) {
	handler := protocol.NewEventHandler(
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
*/
func (DOM) OnDocumentUpdated(
	socket *Socket,
	callback func(event *dom.DocumentUpdatedEvent),
) {
	handler := protocol.NewEventHandler(
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
*/
func (DOM) OnInlineStyleInvalidated(
	socket *Socket,
	callback func(event *dom.InlineStyleInvalidatedEvent),
) {
	handler := protocol.NewEventHandler(
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
*/
func (DOM) OnPseudoElementAdded(
	socket *Socket,
	callback func(event *dom.PseudoElementAddedEvent),
) {
	handler := protocol.NewEventHandler(
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
*/
func (DOM) OnPseudoElementRemoved(
	socket *Socket,
	callback func(event *dom.PseudoElementRemovedEvent),
) {
	handler := protocol.NewEventHandler(
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
*/
func (DOM) OnSetChildNodes(
	socket *Socket,
	callback func(event *dom.SetChildNodesEvent),
) {
	handler := protocol.NewEventHandler(
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
*/
func (DOM) OnShadowRootPopped(
	socket *Socket,
	callback func(event *dom.ShadowRootPoppedEvent),
) {
	handler := protocol.NewEventHandler(
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
*/
func (DOM) OnShadowRootPushed(
	socket *Socket,
	callback func(event *dom.ShadowRootPushedEvent),
) {
	handler := protocol.NewEventHandler(
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

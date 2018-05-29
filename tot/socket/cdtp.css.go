package socket

import (
	"encoding/json"

	"github.com/mkenney/go-chrome/tot/cdtp/css"
)

/*
CSSProtocol provides a namespace for the Chrome CSS protocol methods. The CSS
protocol exposes CSS read/write operations. All CSS objects (stylesheets, rules,
and styles) have an associated id used in subsequent operations on the related
object. Each object type has a specific ID structure, and those are not
interchangeable between objects of different kinds. CSS objects can be loaded
using the get*ForNode() calls (which accept a DOM node id). A client can also
keep track of stylesheets via the styleSheetAdded/styleSheetRemoved events and
subsequently load the required stylesheet contents using the
getStyleSheet[Text]() methods.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/
*/
type CSSProtocol struct {
	Socket Socketer
}

/*
AddRule inserts a new rule with the given ruleText in a stylesheet with given
styleSheetId, at the position specified by location.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-addRule
*/
func (protocol *CSSProtocol) AddRule(
	params *css.AddRuleParams,
) <-chan *css.AddRuleResult {
	resultChan := make(chan *css.AddRuleResult)
	command := NewCommand(protocol.Socket, "CSS.addRule", params)
	result := &css.AddRuleResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
CollectClassNames returns all class names from specified stylesheet.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-collectClassNames
*/
func (protocol *CSSProtocol) CollectClassNames(
	params *css.CollectClassNamesParams,
) <-chan *css.CollectClassNamesResult {
	resultChan := make(chan *css.CollectClassNamesResult)
	command := NewCommand(protocol.Socket, "CSS.collectClassNames", params)
	result := &css.CollectClassNamesResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
CreateStyleSheet creates a new special "via-inspector" stylesheet in the frame
with given frameId.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-createStyleSheet
*/
func (protocol *CSSProtocol) CreateStyleSheet(
	params *css.CreateStyleSheetParams,
) <-chan *css.CreateStyleSheetResult {
	resultChan := make(chan *css.CreateStyleSheetResult)
	command := NewCommand(protocol.Socket, "CSS.createStyleSheet", params)
	result := &css.CreateStyleSheetResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
Disable disables the CSS agent for the given page.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-disable
*/
func (protocol *CSSProtocol) Disable() <-chan *css.DisableResult {
	resultChan := make(chan *css.DisableResult)
	command := NewCommand(protocol.Socket, "CSS.disable", nil)
	result := &css.DisableResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
Enable enables the CSS agent for the given page. Clients should not assume that
the CSS agent has been enabled until the result of this command is received.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-enable
*/
func (protocol *CSSProtocol) Enable() <-chan *css.EnableResult {
	resultChan := make(chan *css.EnableResult)
	command := NewCommand(protocol.Socket, "CSS.enable", nil)
	result := &css.EnableResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
ForcePseudoState ensures that the given node will have specified pseudo-classes
whenever its style is computed by the browser.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-forcePseudoState
*/
func (protocol *CSSProtocol) ForcePseudoState(
	params *css.ForcePseudoStateParams,
) <-chan *css.ForcePseudoStateResult {
	resultChan := make(chan *css.ForcePseudoStateResult)
	command := NewCommand(protocol.Socket, "CSS.forcePseudoState", params)
	result := &css.ForcePseudoStateResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
GetBackgroundColors gets background colors for a node.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getBackgroundColors
*/
func (protocol *CSSProtocol) GetBackgroundColors(
	params *css.GetBackgroundColorsParams,
) <-chan *css.GetBackgroundColorsResult {
	resultChan := make(chan *css.GetBackgroundColorsResult)
	command := NewCommand(protocol.Socket, "CSS.getBackgroundColors", params)
	result := &css.GetBackgroundColorsResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
GetComputedStyleForNode returns the computed style for a DOM node identified by
nodeId.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getComputedStyleForNode
*/
func (protocol *CSSProtocol) GetComputedStyleForNode(
	params *css.GetComputedStyleForNodeParams,
) <-chan *css.GetComputedStyleForNodeResult {
	resultChan := make(chan *css.GetComputedStyleForNodeResult)
	command := NewCommand(protocol.Socket, "CSS.getComputedStyleForNode", params)
	result := &css.GetComputedStyleForNodeResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
GetInlineStylesForNode returns the styles defined inline (explicitly in the
"style" attribute and  implicitly, using DOM attributes) for a DOM node
identified by nodeId.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getInlineStylesForNode
*/
func (protocol *CSSProtocol) GetInlineStylesForNode(
	params *css.GetInlineStylesForNodeParams,
) <-chan *css.GetInlineStylesForNodeResult {
	resultChan := make(chan *css.GetInlineStylesForNodeResult)
	command := NewCommand(protocol.Socket, "CSS.getInlineStylesForNode", params)
	result := &css.GetInlineStylesForNodeResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
GetMatchedStylesForNode returns requested styles for a DOM node identified by
nodeId.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getMatchedStylesForNode
*/
func (protocol *CSSProtocol) GetMatchedStylesForNode(
	params *css.GetMatchedStylesForNodeParams,
) <-chan *css.GetMatchedStylesForNodeResult {
	resultChan := make(chan *css.GetMatchedStylesForNodeResult)
	command := NewCommand(protocol.Socket, "CSS.getMatchedStylesForNode", params)
	result := &css.GetMatchedStylesForNodeResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
GetMediaQueries returns all media queries parsed by the rendering engine.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getMediaQueries
*/
func (protocol *CSSProtocol) GetMediaQueries() <-chan *css.GetMediaQueriesResult {
	resultChan := make(chan *css.GetMediaQueriesResult)
	command := NewCommand(protocol.Socket, "CSS.getMediaQueries", nil)
	result := &css.GetMediaQueriesResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
GetPlatformFontsForNode requests information about platform fonts which we used
to render child TextNodes in the given node.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getPlatformFontsForNode
*/
func (protocol *CSSProtocol) GetPlatformFontsForNode(
	params *css.GetPlatformFontsForNodeParams,
) <-chan *css.GetPlatformFontsForNodeResult {
	resultChan := make(chan *css.GetPlatformFontsForNodeResult)
	command := NewCommand(protocol.Socket, "CSS.getPlatformFontsForNode", params)
	result := &css.GetPlatformFontsForNodeResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
GetStyleSheetText returns the current textual content and the URL for a
stylesheet.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getStyleSheetText
*/
func (protocol *CSSProtocol) GetStyleSheetText(
	params *css.GetStyleSheetTextParams,
) <-chan *css.GetStyleSheetTextResult {
	resultChan := make(chan *css.GetStyleSheetTextResult)
	command := NewCommand(protocol.Socket, "CSS.getStyleSheetText", params)
	result := &css.GetStyleSheetTextResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
SetEffectivePropertyValueForNode finds a rule with the given active property for
the given node and sets the new value for that property.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setEffectivePropertyValueForNode
*/
func (protocol *CSSProtocol) SetEffectivePropertyValueForNode(
	params *css.SetEffectivePropertyValueForNodeParams,
) <-chan *css.SetEffectivePropertyValueForNodeResult {
	resultChan := make(chan *css.SetEffectivePropertyValueForNodeResult)
	command := NewCommand(protocol.Socket, "CSS.setEffectivePropertyValueForNode", params)
	result := &css.SetEffectivePropertyValueForNodeResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
SetKeyframeKey modifies the keyframe rule key text.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setKeyframeKey
*/
func (protocol *CSSProtocol) SetKeyframeKey(
	params *css.SetKeyframeKeyParams,
) <-chan *css.SetKeyframeKeyResult {
	resultChan := make(chan *css.SetKeyframeKeyResult)
	command := NewCommand(protocol.Socket, "CSS.setKeyframeKey", params)
	result := &css.SetKeyframeKeyResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
SetMediaText modifies the rule selector.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setMediaText
*/
func (protocol *CSSProtocol) SetMediaText(
	params *css.SetMediaTextParams,
) <-chan *css.SetMediaTextResult {
	resultChan := make(chan *css.SetMediaTextResult)
	command := NewCommand(protocol.Socket, "CSS.setMediaText", params)
	result := &css.SetMediaTextResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
SetRuleSelector modifies the rule selector.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setRuleSelector
*/
func (protocol *CSSProtocol) SetRuleSelector(
	params *css.SetRuleSelectorParams,
) <-chan *css.SetRuleSelectorResult {
	resultChan := make(chan *css.SetRuleSelectorResult)
	command := NewCommand(protocol.Socket, "CSS.setRuleSelector", params)
	result := &css.SetRuleSelectorResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
SetStyleSheetText sets the new stylesheet text.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setStyleSheetText
*/
func (protocol *CSSProtocol) SetStyleSheetText(
	params *css.SetStyleSheetTextParams,
) <-chan *css.SetStyleSheetTextResult {
	resultChan := make(chan *css.SetStyleSheetTextResult)
	command := NewCommand(protocol.Socket, "CSS.setStyleSheetText", params)
	result := &css.SetStyleSheetTextResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
SetStyleTexts applies specified style edits one after another in the given order.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setStyleTexts
*/
func (protocol *CSSProtocol) SetStyleTexts(
	params *css.SetStyleTextsParams,
) <-chan *css.SetStyleTextsResult {
	resultChan := make(chan *css.SetStyleTextsResult)
	command := NewCommand(protocol.Socket, "CSS.setStyleTexts", params)
	result := &css.SetStyleTextsResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
StartRuleUsageTracking enables the selector recording.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-startRuleUsageTracking
*/
func (protocol *CSSProtocol) StartRuleUsageTracking() <-chan *css.StartRuleUsageTrackingResult {
	resultChan := make(chan *css.StartRuleUsageTrackingResult)
	command := NewCommand(protocol.Socket, "CSS.startRuleUsageTracking", nil)
	result := &css.StartRuleUsageTrackingResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
StopRuleUsageTracking returns he list of rules with an indication of whether
they were used.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-stopRuleUsageTracking
*/
func (protocol *CSSProtocol) StopRuleUsageTracking() <-chan *css.StopRuleUsageTrackingResult {
	resultChan := make(chan *css.StopRuleUsageTrackingResult)
	command := NewCommand(protocol.Socket, "CSS.stopRuleUsageTracking", nil)
	result := &css.StopRuleUsageTrackingResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
TakeCoverageDelta obtains the list of rules that became used since last call to
this method (or since start of coverage instrumentation).

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-takeCoverageDelta
*/
func (protocol *CSSProtocol) TakeCoverageDelta() <-chan *css.TakeCoverageDeltaResult {
	resultChan := make(chan *css.TakeCoverageDeltaResult)
	command := NewCommand(protocol.Socket, "CSS.takeCoverageDelta", nil)
	result := &css.TakeCoverageDeltaResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
OnFontsUpdated adds a handler to the CSS.fontsUpdated event. CSS.fontsUpdated
fires whenever a web font gets loaded.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#event-fontsUpdated
*/
func (protocol *CSSProtocol) OnFontsUpdated(
	callback func(event *css.FontsUpdatedEvent),
) {
	handler := NewEventHandler(
		"CSS.fontsUpdated",
		func(response *Response) {
			event := &css.FontsUpdatedEvent{}
			json.Unmarshal([]byte(response.Params), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnMediaQueryResultChanged adds a handler to the CSS.mediaQueryResultChanged
event. CSS.mediaQueryResultChanged fires whenever a MediaQuery result changes
(for example, after a browser window has been resized.) The current
implementation considers only viewport-dependent media features.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#event-mediaQueryResultChanged
*/
func (protocol *CSSProtocol) OnMediaQueryResultChanged(
	callback func(event *css.MediaQueryResultChangedEvent),
) {
	handler := NewEventHandler(
		"CSS.mediaQueryResultChanged",
		func(response *Response) {
			event := &css.MediaQueryResultChangedEvent{}
			json.Unmarshal([]byte(response.Params), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnStyleSheetAdded adds a handler to the CSS.styleSheetAdded event.
CSS.styleSheetAdded fires whenever an active document stylesheet is added.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#event-styleSheetAdded
*/
func (protocol *CSSProtocol) OnStyleSheetAdded(
	callback func(event *css.StyleSheetAddedEvent),
) {
	handler := NewEventHandler(
		"CSS.styleSheetAdded",
		func(response *Response) {
			event := &css.StyleSheetAddedEvent{}
			json.Unmarshal([]byte(response.Params), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnStyleSheetChanged adds a handler to the CSS.styleSheetChanged event.
CSS.styleSheetChanged fires whenever a stylesheet is changed as a result of the
client operation.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#event-styleSheetChanged
*/
func (protocol *CSSProtocol) OnStyleSheetChanged(
	callback func(event *css.StyleSheetChangedEvent),
) {
	handler := NewEventHandler(
		"CSS.styleSheetChanged",
		func(response *Response) {
			event := &css.StyleSheetChangedEvent{}
			json.Unmarshal([]byte(response.Params), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnStyleSheetRemoved adds a handler to the CSS.styleSheetRemoved event.
CSS.styleSheetRemoved fires whenever an active document stylesheet is removed.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#event-styleSheetRemoved
*/
func (protocol *CSSProtocol) OnStyleSheetRemoved(
	callback func(event *css.StyleSheetRemovedEvent),
) {
	handler := NewEventHandler(
		"CSS.styleSheetRemoved",
		func(response *Response) {
			event := &css.StyleSheetRemovedEvent{}
			json.Unmarshal([]byte(response.Params), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

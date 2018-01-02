package socket

import (
	"encoding/json"

	css "github.com/mkenney/go-chrome/cdtp/css"
	log "github.com/sirupsen/logrus"
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
) (*css.AddRuleResult, error) {
	command := NewCommand("CSS.addRule", params)
	result := &css.AddRuleResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
CollectClassNames returns all class names from specified stylesheet.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-collectClassNames
*/
func (protocol *CSSProtocol) CollectClassNames(
	params *css.CollectClassNamesParams,
) (*css.CollectClassNamesResult, error) {
	command := NewCommand("CSS.collectClassNames", params)
	result := &css.CollectClassNamesResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
CreateStyleSheet creates a new special "via-inspector" stylesheet in the frame
with given frameId.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-createStyleSheet
*/
func (protocol *CSSProtocol) CreateStyleSheet(
	params *css.CreateStyleSheetParams,
) (*css.CreateStyleSheetResult, error) {
	command := NewCommand("CSS.createStyleSheet", params)
	result := &css.CreateStyleSheetResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
Disable disables the CSS agent for the given page.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-disable
*/
func (protocol *CSSProtocol) Disable() error {
	command := NewCommand("CSS.disable", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
Enable enables the CSS agent for the given page. Clients should not assume that
the CSS agent has been enabled until the result of this command is received.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-enable
*/
func (protocol *CSSProtocol) Enable() error {
	command := NewCommand("CSS.enable", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
ForcePseudoState ensures that the given node will have specified pseudo-classes
whenever its style is computed by the browser.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-forcePseudoState
*/
func (protocol *CSSProtocol) ForcePseudoState(
	params *css.ForcePseudoStateParams,
) error {
	command := NewCommand("CSS.forcePseudoState", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
GetBackgroundColors gets background colors for a node.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getBackgroundColors
*/
func (protocol *CSSProtocol) GetBackgroundColors(
	params *css.GetBackgroundColorsParams,
) (*css.GetBackgroundColorsResult, error) {
	command := NewCommand("CSS.getBackgroundColors", params)
	result := &css.GetBackgroundColorsResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
GetComputedStyleForNode returns the computed style for a DOM node identified by
nodeId.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getComputedStyleForNode
*/
func (protocol *CSSProtocol) GetComputedStyleForNode(
	params *css.GetComputedStyleForNodeParams,
) (*css.GetComputedStyleForNodeResult, error) {
	command := NewCommand("CSS.getComputedStyleForNode", params)
	result := &css.GetComputedStyleForNodeResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
GetInlineStylesForNode returns the styles defined inline (explicitly in the
"style" attribute and  implicitly, using DOM attributes) for a DOM node
identified by nodeId.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getInlineStylesForNode
*/
func (protocol *CSSProtocol) GetInlineStylesForNode(
	params *css.GetInlineStylesForNodeParams,
) (*css.GetInlineStylesForNodeResult, error) {
	command := NewCommand("CSS.getInlineStylesForNode", params)
	result := &css.GetInlineStylesForNodeResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
GetMatchedStylesForNode returns requested styles for a DOM node identified by
nodeId.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getMatchedStylesForNode
*/
func (protocol *CSSProtocol) GetMatchedStylesForNode(
	params *css.GetMatchedStylesForNodeParams,
) (*css.GetMatchedStylesForNodeResult, error) {
	command := NewCommand("CSS.getMatchedStylesForNode", params)
	result := &css.GetMatchedStylesForNodeResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
GetMediaQueries returns all media queries parsed by the rendering engine.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getMediaQueries
*/
func (protocol *CSSProtocol) GetMediaQueries() error {
	command := NewCommand("CSS.getMediaQueries", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
GetPlatformFontsForNode requests information about platform fonts which we used
to render child TextNodes in the given node.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getPlatformFontsForNode
*/
func (protocol *CSSProtocol) GetPlatformFontsForNode(
	params *css.GetPlatformFontsForNodeParams,
) (*css.GetPlatformFontsForNodeResult, error) {
	command := NewCommand("CSS.getPlatformFontsForNode", params)
	result := &css.GetPlatformFontsForNodeResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
GetStyleSheetText returns the current textual content and the URL for a
stylesheet.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getStyleSheetText
*/
func (protocol *CSSProtocol) GetStyleSheetText(
	params *css.GetStyleSheetTextParams,
) (*css.GetStyleSheetTextResult, error) {
	command := NewCommand("CSS.getStyleSheetText", params)
	result := &css.GetStyleSheetTextResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
SetEffectivePropertyValueForNode finds a rule with the given active property for
the given node and sets the new value for that property.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setEffectivePropertyValueForNode
*/
func (protocol *CSSProtocol) SetEffectivePropertyValueForNode(
	params *css.SetEffectivePropertyValueForNodeParams,
) error {
	command := NewCommand("CSS.setEffectivePropertyValueForNode", params)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
SetKeyframeKey modifies the keyframe rule key text.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setKeyframeKey
*/
func (protocol *CSSProtocol) SetKeyframeKey(
	params *css.SetKeyframeKeyParams,
) (*css.SetKeyframeKeyResult, error) {
	command := NewCommand("CSS.setKeyframeKey", params)
	result := &css.SetKeyframeKeyResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
SetMediaText modifies the rule selector.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setMediaText
*/
func (protocol *CSSProtocol) SetMediaText(
	params *css.SetMediaTextParams,
) (*css.SetMediaTextResult, error) {
	command := NewCommand("CSS.setMediaText", params)
	result := &css.SetMediaTextResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
SetRuleSelector modifies the rule selector.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setRuleSelector
*/
func (protocol *CSSProtocol) SetRuleSelector(
	params *css.SetRuleSelectorParams,
) (*css.SetRuleSelectorResult, error) {
	command := NewCommand("CSS.setRuleSelector", params)
	result := &css.SetRuleSelectorResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
SetStyleSheetText sets the new stylesheet text.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setStyleSheetText
*/
func (protocol *CSSProtocol) SetStyleSheetText(
	params *css.SetStyleSheetTextParams,
) (*css.SetStyleSheetTextResult, error) {
	command := NewCommand("CSS.setStyleSheetText", params)
	result := &css.SetStyleSheetTextResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
SetStyleTexts applies specified style edits one after another in the given order.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setStyleTexts
*/
func (protocol *CSSProtocol) SetStyleTexts(
	params *css.SetStyleTextsParams,
) (*css.SetStyleTextsResult, error) {
	command := NewCommand("CSS.setStyleTexts", params)
	result := &css.SetStyleTextsResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
StartRuleUsageTracking enables the selector recording.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-startRuleUsageTracking
*/
func (protocol *CSSProtocol) StartRuleUsageTracking() error {
	command := NewCommand("CSS.startRuleUsageTracking", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
StopRuleUsageTracking returns he list of rules with an indication of whether
they were used.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-stopRuleUsageTracking
*/
func (protocol *CSSProtocol) StopRuleUsageTracking() (*css.StopRuleUsageTrackingResult, error) {
	command := NewCommand("CSS.stopRuleUsageTracking", nil)
	result := &css.StopRuleUsageTrackingResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
TakeCoverageDelta obtains the list of rules that became used since last call to
this method (or since start of coverage instrumentation).

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-takeCoverageDelta
*/
func (protocol *CSSProtocol) TakeCoverageDelta() (*css.TakeCoverageDeltaResult, error) {
	command := NewCommand("CSS.takeCoverageDelta", nil)
	result := &css.TakeCoverageDeltaResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
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
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

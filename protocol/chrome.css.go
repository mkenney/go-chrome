package protocol

import (
	"encoding/json"

	css "github.com/mkenney/go-chrome/protocol/css"
	sock "github.com/mkenney/go-chrome/socket"

	log "github.com/sirupsen/logrus"
)

/*
CSS is a struct that provides a namespace for the Chrome CSS protocol methods.

The CSS protocol exposes CSS read/write operations. All CSS objects (stylesheets, rules, and styles)
have an associated id used in subsequent operations on the related object. Each object type has a
specific ID structure, and those are not interchangeable between objects of different kinds. CSS
objects can be loaded using the get*ForNode() calls (which accept a DOM node id). A client can also
keep track of stylesheets via the styleSheetAdded/styleSheetRemoved events and subsequently load the
required stylesheet contents using the getStyleSheet[Text]() methods.

- https://chromedevtools.github.io/devtools-protocol/tot/CSS/
*/
type CSS struct{}

/*
AddRule inserts a new rule with the given ruleText in a stylesheet with given styleSheetId, at the
position specified by location.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-addRule
*/
func (CSS) AddRule(
	socket *sock.Socket,
	params *css.AddRuleParams,
) (css.AddRuleResult, error) {
	command := &sock.Command{
		Method: "CSS.addRule",
		Params: params,
	}
	result := css.AddRuleResult{}
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
CollectClassNames returns all class names from specified stylesheet.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-collectClassNames
*/
func (CSS) CollectClassNames(
	socket *sock.Socket,
	params *css.CollectClassNamesParams,
) (css.CollectClassNamesResult, error) {
	command := &sock.Command{
		Method: "CSS.collectClassNames",
		Params: params,
	}
	result := css.CollectClassNamesResult{}
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
CreateStyleSheet creates a new special "via-inspector" stylesheet in the frame with given frameId.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-createStyleSheet
*/
func (CSS) CreateStyleSheet(
	socket *sock.Socket,
	params *css.CreateStyleSheetParams,
) (css.CreateStyleSheetResult, error) {
	command := &sock.Command{
		Method: "CSS.createStyleSheet",
		Params: params,
	}
	result := css.CreateStyleSheetResult{}
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
Disable disables the CSS agent for the given page.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-disable
*/
func (CSS) Disable(
	socket *sock.Socket,
) error {
	command := &sock.Command{
		Method: "CSS.disable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Enable enables the CSS agent for the given page. Clients should not assume that the CSS agent has
been enabled until the result of this command is received.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-enable
*/
func (CSS) Enable(
	socket *sock.Socket,
) error {
	command := &sock.Command{
		Method: "CSS.enable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
ForcePseudoState ensures that the given node will have specified pseudo-classes whenever its style
is computed by the browser.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-forcePseudoState
*/
func (CSS) ForcePseudoState(
	socket *sock.Socket,
	params *css.ForcePseudoStateParams,
) error {
	command := &sock.Command{
		Method: "CSS.forcePseudoState",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
GetBackgroundColors gets background colors for a node.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getBackgroundColors
*/
func (CSS) GetBackgroundColors(
	socket *sock.Socket,
	params *css.GetBackgroundColorsParams,
) (css.GetBackgroundColorsResult, error) {
	command := &sock.Command{
		Method: "CSS.getBackgroundColors",
		Params: params,
	}
	result := css.GetBackgroundColorsResult{}
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
GetComputedStyleForNode returns the computed style for a DOM node identified by nodeId.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getComputedStyleForNode
*/
func (CSS) GetComputedStyleForNode(
	socket *sock.Socket,
	params *css.GetComputedStyleForNodeParams,
) (css.GetComputedStyleForNodeResult, error) {
	command := &sock.Command{
		Method: "CSS.getComputedStyleForNode",
		Params: params,
	}
	result := css.GetComputedStyleForNodeResult{}
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
GetInlineStylesForNode returns the styles defined inline (explicitly in the "style" attribute and
implicitly, using DOM attributes) for a DOM node identified by nodeId.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getInlineStylesForNode
*/
func (CSS) GetInlineStylesForNode(
	socket *sock.Socket,
	params *css.GetInlineStylesForNodeParams,
) (css.GetInlineStylesForNodeResult, error) {
	command := &sock.Command{
		Method: "CSS.getInlineStylesForNode",
		Params: params,
	}
	result := css.GetInlineStylesForNodeResult{}
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
GetMatchedStylesForNode returns requested styles for a DOM node identified by nodeId.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getMatchedStylesForNode
*/
func (CSS) GetMatchedStylesForNode(
	socket *sock.Socket,
	params *css.GetMatchedStylesForNodeParams,
) (css.GetMatchedStylesForNodeResult, error) {
	command := &sock.Command{
		Method: "CSS.getMatchedStylesForNode",
		Params: params,
	}
	result := css.GetMatchedStylesForNodeResult{}
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
GetMediaQueries returns all media queries parsed by the rendering engine.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getMediaQueries
*/
func (CSS) GetMediaQueries(
	socket *sock.Socket,
) error {
	command := &sock.Command{
		Method: "CSS.getMediaQueries",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
GetPlatformFontsForNode requests information about platform fonts which we used to render child
TextNodes in the given node.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getPlatformFontsForNode
*/
func (CSS) GetPlatformFontsForNode(
	socket *sock.Socket,
	params *css.GetPlatformFontsForNodeParams,
) (css.GetPlatformFontsForNodeResult, error) {
	command := &sock.Command{
		Method: "CSS.getPlatformFontsForNode",
		Params: params,
	}
	result := css.GetPlatformFontsForNodeResult{}
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
GetStyleSheetText returns the current textual content and the URL for a stylesheet.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getStyleSheetText
*/
func (CSS) GetStyleSheetText(
	socket *sock.Socket,
	params *css.GetStyleSheetTextParams,
) (css.GetStyleSheetTextResult, error) {
	command := &sock.Command{
		Method: "CSS.getStyleSheetText",
		Params: params,
	}
	result := css.GetStyleSheetTextResult{}
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
SetEffectivePropertyValueForNode finds a rule with the given active property for the given node and
sets the new value for that property.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setEffectivePropertyValueForNode
*/
func (CSS) SetEffectivePropertyValueForNode(
	socket *sock.Socket,
	params *css.SetEffectivePropertyValueForNodeParams,
) error {
	command := &sock.Command{
		Method: "CSS.setEffectivePropertyValueForNode",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetKeyframeKey modifies the keyframe rule key text.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setKeyframeKey
*/
func (CSS) SetKeyframeKey(
	socket *sock.Socket,
	params *css.SetKeyframeKeyParams,
) (css.SetKeyframeKeyResult, error) {
	command := &sock.Command{
		Method: "CSS.setKeyframeKey",
		Params: params,
	}
	result := css.SetKeyframeKeyResult{}
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
SetMediaText modifies the rule selector.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setMediaText
*/
func (CSS) SetMediaText(
	socket *sock.Socket,
	params *css.SetMediaTextParams,
) (css.SetMediaTextResult, error) {
	command := &sock.Command{
		Method: "CSS.setMediaText",
		Params: params,
	}
	result := css.SetMediaTextResult{}
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
SetRuleSelector modifies the rule selector.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setRuleSelector
*/
func (CSS) SetRuleSelector(
	socket *sock.Socket,
	params *css.SetRuleSelectorParams,
) (css.SetRuleSelectorResult, error) {
	command := &sock.Command{
		Method: "CSS.setRuleSelector",
		Params: params,
	}
	result := css.SetRuleSelectorResult{}
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
SetStyleSheetText sets the new stylesheet text.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setStyleSheetText
*/
func (CSS) SetStyleSheetText(
	socket *sock.Socket,
	params *css.SetStyleSheetTextParams,
) (css.SetStyleSheetTextResult, error) {
	command := &sock.Command{
		Method: "CSS.setStyleSheetText",
		Params: params,
	}
	result := css.SetStyleSheetTextResult{}
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
SetStyleTexts applies specified style edits one after another in the given order.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setStyleTexts
*/
func (CSS) SetStyleTexts(
	socket *sock.Socket,
	params *css.SetStyleTextsParams,
) (css.SetStyleTextsResult, error) {
	command := &sock.Command{
		Method: "CSS.setStyleTexts",
		Params: params,
	}
	result := css.SetStyleTextsResult{}
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
StartRuleUsageTracking enables the selector recording.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-startRuleUsageTracking
*/
func (CSS) StartRuleUsageTracking(
	socket *sock.Socket,
) error {
	command := &sock.Command{
		Method: "CSS.startRuleUsageTracking",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
StopRuleUsageTracking returns he list of rules with an indication of whether they were used.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-stopRuleUsageTracking
*/
func (CSS) StopRuleUsageTracking(
	socket *sock.Socket,
) (css.StopRuleUsageTrackingResult, error) {
	command := &sock.Command{
		Method: "CSS.stopRuleUsageTracking",
	}
	result := css.StopRuleUsageTrackingResult{}
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
TakeCoverageDelta obtains the list of rules that became used since last call to this method (or
since start of coverage instrumentation).

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-takeCoverageDelta
*/
func (CSS) TakeCoverageDelta(
	socket *sock.Socket,
) (css.TakeCoverageDeltaResult, error) {
	command := &sock.Command{
		Method: "CSS.takeCoverageDelta",
	}
	result := css.TakeCoverageDeltaResult{}
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
OnFontsUpdated adds a handler to the CSS.fontsUpdated event. CSS.fontsUpdated fires whenever a web
font gets loaded.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#event-fontsUpdated
*/
func (CSS) OnFontsUpdated(
	socket *sock.Socket,
	callback func(event *css.FontsUpdatedEvent),
) {
	handler := sock.NewEventHandler(
		"CSS.fontsUpdated",
		func(name string, params []byte) {
			event := &css.FontsUpdatedEvent{}
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
OnMediaQueryResultChanged adds a handler to the CSS.mediaQueryResultChanged event.
CSS.mediaQueryResultChanged fires whenever a MediaQuery result changes (for example, after a browser
window has been resized.) The current implementation considers only viewport-dependent media
features.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#event-mediaQueryResultChanged
*/
func (CSS) OnMediaQueryResultChanged(
	socket *sock.Socket,
	callback func(event *css.MediaQueryResultChangedEvent),
) {
	handler := sock.NewEventHandler(
		"CSS.mediaQueryResultChanged",
		func(name string, params []byte) {
			event := &css.MediaQueryResultChangedEvent{}
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
OnStyleSheetAdded adds a handler to the CSS.styleSheetAdded event. CSS.styleSheetAdded fires
whenever an active document stylesheet is added.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#event-styleSheetAdded
*/
func (CSS) OnStyleSheetAdded(
	socket *sock.Socket,
	callback func(event *css.StyleSheetAddedEvent),
) {
	handler := sock.NewEventHandler(
		"CSS.styleSheetAdded",
		func(name string, params []byte) {
			event := &css.StyleSheetAddedEvent{}
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
OnStyleSheetChanged adds a handler to the CSS.styleSheetChanged event. CSS.styleSheetChanged fires
whenever a stylesheet is changed as a result of the client operation.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#event-styleSheetChanged
*/
func (CSS) OnStyleSheetChanged(
	socket *sock.Socket,
	callback func(event *css.StyleSheetChangedEvent),
) {
	handler := sock.NewEventHandler(
		"CSS.styleSheetChanged",
		func(name string, params []byte) {
			event := &css.StyleSheetChangedEvent{}
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
OnStyleSheetRemoved adds a handler to the CSS.styleSheetRemoved event. CSS.styleSheetRemoved fires
whenever an active document stylesheet is removed.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#event-styleSheetRemoved
*/
func (CSS) OnStyleSheetRemoved(
	socket *sock.Socket,
	callback func(event *css.StyleSheetRemovedEvent),
) {
	handler := sock.NewEventHandler(
		"CSS.styleSheetRemoved",
		func(name string, params []byte) {
			event := &css.StyleSheetRemovedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}
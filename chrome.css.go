package chrome

import (
	css "app/chrome/css"
	"app/chrome/protocol"
	"encoding/json"

	log "github.com/Sirupsen/logrus"
)

/*
CSS - https://chromedevtools.github.io/devtools-protocol/tot/CSS/
Exposes CSS read/write operations. All CSS objects (stylesheets, rules, and styles) have an
associated id used in subsequent operations on the related object. Each object type has a specific
id structure, and those are not interchangeable between objects of different kinds. CSS objects can
be loaded using the get*ForNode() calls (which accept a DOM node id). A client can also keep track
of stylesheets via the styleSheetAdded/styleSheetRemoved events and subsequently load the required
stylesheet contents using the getStyleSheet[Text]() methods.
*/
type CSS struct{}

/*
AddRule inserts a new rule with the given ruleText in a stylesheet with given styleSheetId, at the
position specified by location.
*/
func (CSS) AddRule(
	socket *Socket,
	params *css.AddRuleParams,
) (css.AddRuleResult, error) {
	command := &protocol.Command{
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
*/
func (CSS) CollectClassNames(
	socket *Socket,
	params *css.CollectClassNamesParams,
) (css.CollectClassNamesResult, error) {
	command := &protocol.Command{
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
*/
func (CSS) CreateStyleSheet(
	socket *Socket,
	params *css.CreateStyleSheetParams,
) (css.CreateStyleSheetResult, error) {
	command := &protocol.Command{
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
*/
func (CSS) Disable(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "CSS.disable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Enable enables the CSS agent for the given page. Clients should not assume that the CSS agent has
been enabled until the result of this command is received.
*/
func (CSS) Enable(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "CSS.enable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
ForcePseudoState ensures that the given node will have specified pseudo-classes whenever its style
is computed by the browser.
*/
func (CSS) ForcePseudoState(
	socket *Socket,
	params *css.ForcePseudoStateParams,
) error {
	command := &protocol.Command{
		Method: "CSS.forcePseudoState",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
GetBackgroundColors gets background colors for a node.
*/
func (CSS) GetBackgroundColors(
	socket *Socket,
	params *css.GetBackgroundColorsParams,
) (css.GetBackgroundColorsResult, error) {
	command := &protocol.Command{
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
*/
func (CSS) GetComputedStyleForNode(
	socket *Socket,
	params *css.GetComputedStyleForNodeParams,
) (css.GetComputedStyleForNodeResult, error) {
	command := &protocol.Command{
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
*/
func (CSS) GetInlineStylesForNode(
	socket *Socket,
	params *css.GetInlineStylesForNodeParams,
) (css.GetInlineStylesForNodeResult, error) {
	command := &protocol.Command{
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
*/
func (CSS) GetMatchedStylesForNode(
	socket *Socket,
	params *css.GetMatchedStylesForNodeParams,
) (css.GetMatchedStylesForNodeResult, error) {
	command := &protocol.Command{
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
*/
func (CSS) GetMediaQueries(
	socket *Socket,
	params *css.GetMediaQueriesParams,
) error {
	command := &protocol.Command{
		Method: "CSS.getMediaQueries",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
GetPlatformFontsForNode requests information about platform fonts which we used to render child
TextNodes in the given node.
*/
func (CSS) GetPlatformFontsForNode(
	socket *Socket,
	params *css.GetPlatformFontsForNodeParams,
) (css.GetPlatformFontsForNodeResult, error) {
	command := &protocol.Command{
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
*/
func (CSS) GetStyleSheetText(
	socket *Socket,
	params *css.GetStyleSheetTextParams,
) (css.GetStyleSheetTextResult, error) {
	command := &protocol.Command{
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
*/
func (CSS) SetEffectivePropertyValueForNode(
	socket *Socket,
	params *css.SetEffectivePropertyValueForNodeParams,
) error {
	command := &protocol.Command{
		Method: "CSS.setEffectivePropertyValueForNode",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetKeyframeKey modifies the keyframe rule key text.
*/
func (CSS) SetKeyframeKey(
	socket *Socket,
	params *css.SetKeyframeKeyParams,
) (css.SetKeyframeKeyResult, error) {
	command := &protocol.Command{
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
*/
func (CSS) SetMediaText(
	socket *Socket,
	params *css.SetMediaTextParams,
) (css.SetMediaTextResult, error) {
	command := &protocol.Command{
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
*/
func (CSS) SetRuleSelector(
	socket *Socket,
	params *css.SetRuleSelectorParams,
) (css.SetRuleSelectorResult, error) {
	command := &protocol.Command{
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
*/
func (CSS) SetStyleSheetText(
	socket *Socket,
	params *css.SetStyleSheetTextParams,
) (css.SetStyleSheetTextResult, error) {
	command := &protocol.Command{
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
*/
func (CSS) SetStyleTexts(
	socket *Socket,
	params *css.SetStyleTextsParams,
) (css.SetStyleTextsResult, error) {
	command := &protocol.Command{
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
*/
func (CSS) StartRuleUsageTracking(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "CSS.startRuleUsageTracking",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
StopRuleUsageTracking returns he list of rules with an indication of whether they were used.
*/
func (CSS) StopRuleUsageTracking(
	socket *Socket,
) (css.StopRuleUsageTrackingResult, error) {
	command := &protocol.Command{
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
*/
func (CSS) TakeCoverageDelta(
	socket *Socket,
) (css.TakeCoverageDeltaResult, error) {
	command := &protocol.Command{
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
*/
func (CSS) OnFontsUpdated(
	socket *Socket,
	callback func(event *css.FontsUpdatedEvent),
) {
	handler := protocol.NewEventHandler(
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
*/
func (CSS) OnMediaQueryResultChanged(
	socket *Socket,
	callback func(event *css.MediaQueryResultChangedEvent),
) {
	handler := protocol.NewEventHandler(
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
*/
func (CSS) OnStyleSheetAdded(
	socket *Socket,
	callback func(event *css.StyleSheetAddedEvent),
) {
	handler := protocol.NewEventHandler(
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
*/
func (CSS) OnStyleSheetChanged(
	socket *Socket,
	callback func(event *css.StyleSheetChangedEvent),
) {
	handler := protocol.NewEventHandler(
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
*/
func (CSS) OnStyleSheetRemoved(
	socket *Socket,
	callback func(event *css.StyleSheetRemovedEvent),
) {
	handler := protocol.NewEventHandler(
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

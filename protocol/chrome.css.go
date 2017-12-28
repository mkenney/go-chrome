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
	socket sock.Socketer,
	params *css.AddRuleParams,
) (css.AddRuleResult, error) {
	command := sock.NewCommand("CSS.addRule", params)
	result := css.AddRuleResult{}
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
CollectClassNames returns all class names from specified stylesheet.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-collectClassNames
*/
func (CSS) CollectClassNames(
	socket sock.Socketer,
	params *css.CollectClassNamesParams,
) (css.CollectClassNamesResult, error) {
	command := sock.NewCommand("CSS.collectClassNames", params)
	result := css.CollectClassNamesResult{}
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
CreateStyleSheet creates a new special "via-inspector" stylesheet in the frame with given frameId.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-createStyleSheet
*/
func (CSS) CreateStyleSheet(
	socket sock.Socketer,
	params *css.CreateStyleSheetParams,
) (css.CreateStyleSheetResult, error) {
	command := sock.NewCommand("CSS.createStyleSheet", params)
	result := css.CreateStyleSheetResult{}
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
Disable disables the CSS agent for the given page.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-disable
*/
func (CSS) Disable(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("CSS.disable", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
Enable enables the CSS agent for the given page. Clients should not assume that the CSS agent has
been enabled until the result of this command is received.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-enable
*/
func (CSS) Enable(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("CSS.enable", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
ForcePseudoState ensures that the given node will have specified pseudo-classes whenever its style
is computed by the browser.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-forcePseudoState
*/
func (CSS) ForcePseudoState(
	socket sock.Socketer,
	params *css.ForcePseudoStateParams,
) error {
	command := sock.NewCommand("CSS.forcePseudoState", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
GetBackgroundColors gets background colors for a node.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getBackgroundColors
*/
func (CSS) GetBackgroundColors(
	socket sock.Socketer,
	params *css.GetBackgroundColorsParams,
) (css.GetBackgroundColorsResult, error) {
	command := sock.NewCommand("CSS.getBackgroundColors", params)
	result := css.GetBackgroundColorsResult{}
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
GetComputedStyleForNode returns the computed style for a DOM node identified by nodeId.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getComputedStyleForNode
*/
func (CSS) GetComputedStyleForNode(
	socket sock.Socketer,
	params *css.GetComputedStyleForNodeParams,
) (css.GetComputedStyleForNodeResult, error) {
	command := sock.NewCommand("CSS.getComputedStyleForNode", params)
	result := css.GetComputedStyleForNodeResult{}
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
GetInlineStylesForNode returns the styles defined inline (explicitly in the "style" attribute and
implicitly, using DOM attributes) for a DOM node identified by nodeId.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getInlineStylesForNode
*/
func (CSS) GetInlineStylesForNode(
	socket sock.Socketer,
	params *css.GetInlineStylesForNodeParams,
) (css.GetInlineStylesForNodeResult, error) {
	command := sock.NewCommand("CSS.getInlineStylesForNode", params)
	result := css.GetInlineStylesForNodeResult{}
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
GetMatchedStylesForNode returns requested styles for a DOM node identified by nodeId.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getMatchedStylesForNode
*/
func (CSS) GetMatchedStylesForNode(
	socket sock.Socketer,
	params *css.GetMatchedStylesForNodeParams,
) (css.GetMatchedStylesForNodeResult, error) {
	command := sock.NewCommand("CSS.getMatchedStylesForNode", params)
	result := css.GetMatchedStylesForNodeResult{}
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
GetMediaQueries returns all media queries parsed by the rendering engine.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getMediaQueries
*/
func (CSS) GetMediaQueries(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("CSS.getMediaQueries", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
GetPlatformFontsForNode requests information about platform fonts which we used to render child
TextNodes in the given node.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getPlatformFontsForNode
*/
func (CSS) GetPlatformFontsForNode(
	socket sock.Socketer,
	params *css.GetPlatformFontsForNodeParams,
) (css.GetPlatformFontsForNodeResult, error) {
	command := sock.NewCommand("CSS.getPlatformFontsForNode", params)
	result := css.GetPlatformFontsForNodeResult{}
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
GetStyleSheetText returns the current textual content and the URL for a stylesheet.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getStyleSheetText
*/
func (CSS) GetStyleSheetText(
	socket sock.Socketer,
	params *css.GetStyleSheetTextParams,
) (css.GetStyleSheetTextResult, error) {
	command := sock.NewCommand("CSS.getStyleSheetText", params)
	result := css.GetStyleSheetTextResult{}
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
SetEffectivePropertyValueForNode finds a rule with the given active property for the given node and
sets the new value for that property.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setEffectivePropertyValueForNode
*/
func (CSS) SetEffectivePropertyValueForNode(
	socket sock.Socketer,
	params *css.SetEffectivePropertyValueForNodeParams,
) error {
	command := sock.NewCommand("CSS.setEffectivePropertyValueForNode", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetKeyframeKey modifies the keyframe rule key text.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setKeyframeKey
*/
func (CSS) SetKeyframeKey(
	socket sock.Socketer,
	params *css.SetKeyframeKeyParams,
) (css.SetKeyframeKeyResult, error) {
	command := sock.NewCommand("CSS.setKeyframeKey", params)
	result := css.SetKeyframeKeyResult{}
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
SetMediaText modifies the rule selector.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setMediaText
*/
func (CSS) SetMediaText(
	socket sock.Socketer,
	params *css.SetMediaTextParams,
) (css.SetMediaTextResult, error) {
	command := sock.NewCommand("CSS.setMediaText", params)
	result := css.SetMediaTextResult{}
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
SetRuleSelector modifies the rule selector.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setRuleSelector
*/
func (CSS) SetRuleSelector(
	socket sock.Socketer,
	params *css.SetRuleSelectorParams,
) (css.SetRuleSelectorResult, error) {
	command := sock.NewCommand("CSS.setRuleSelector", params)
	result := css.SetRuleSelectorResult{}
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
SetStyleSheetText sets the new stylesheet text.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setStyleSheetText
*/
func (CSS) SetStyleSheetText(
	socket sock.Socketer,
	params *css.SetStyleSheetTextParams,
) (css.SetStyleSheetTextResult, error) {
	command := sock.NewCommand("CSS.setStyleSheetText", params)
	result := css.SetStyleSheetTextResult{}
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
SetStyleTexts applies specified style edits one after another in the given order.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setStyleTexts
*/
func (CSS) SetStyleTexts(
	socket sock.Socketer,
	params *css.SetStyleTextsParams,
) (css.SetStyleTextsResult, error) {
	command := sock.NewCommand("CSS.setStyleTexts", params)
	result := css.SetStyleTextsResult{}
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
StartRuleUsageTracking enables the selector recording.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-startRuleUsageTracking
*/
func (CSS) StartRuleUsageTracking(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("CSS.startRuleUsageTracking", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
StopRuleUsageTracking returns he list of rules with an indication of whether they were used.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-stopRuleUsageTracking
*/
func (CSS) StopRuleUsageTracking(
	socket sock.Socketer,
) (css.StopRuleUsageTrackingResult, error) {
	command := sock.NewCommand("CSS.stopRuleUsageTracking", nil)
	result := css.StopRuleUsageTrackingResult{}
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
TakeCoverageDelta obtains the list of rules that became used since last call to this method (or
since start of coverage instrumentation).

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-takeCoverageDelta
*/
func (CSS) TakeCoverageDelta(
	socket sock.Socketer,
) (css.TakeCoverageDeltaResult, error) {
	command := sock.NewCommand("CSS.takeCoverageDelta", nil)
	result := css.TakeCoverageDeltaResult{}
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
OnFontsUpdated adds a handler to the CSS.fontsUpdated event. CSS.fontsUpdated fires whenever a web
font gets loaded.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#event-fontsUpdated
*/
func (CSS) OnFontsUpdated(
	socket sock.Socketer,
	callback func(event *css.FontsUpdatedEvent),
) {
	handler := sock.NewEventHandler(
		"CSS.fontsUpdated",
		func(response *sock.Response) {
			event := &css.FontsUpdatedEvent{}
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
OnMediaQueryResultChanged adds a handler to the CSS.mediaQueryResultChanged event.
CSS.mediaQueryResultChanged fires whenever a MediaQuery result changes (for example, after a browser
window has been resized.) The current implementation considers only viewport-dependent media
features.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#event-mediaQueryResultChanged
*/
func (CSS) OnMediaQueryResultChanged(
	socket sock.Socketer,
	callback func(event *css.MediaQueryResultChangedEvent),
) {
	handler := sock.NewEventHandler(
		"CSS.mediaQueryResultChanged",
		func(response *sock.Response) {
			event := &css.MediaQueryResultChangedEvent{}
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
OnStyleSheetAdded adds a handler to the CSS.styleSheetAdded event. CSS.styleSheetAdded fires
whenever an active document stylesheet is added.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#event-styleSheetAdded
*/
func (CSS) OnStyleSheetAdded(
	socket sock.Socketer,
	callback func(event *css.StyleSheetAddedEvent),
) {
	handler := sock.NewEventHandler(
		"CSS.styleSheetAdded",
		func(response *sock.Response) {
			event := &css.StyleSheetAddedEvent{}
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
OnStyleSheetChanged adds a handler to the CSS.styleSheetChanged event. CSS.styleSheetChanged fires
whenever a stylesheet is changed as a result of the client operation.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#event-styleSheetChanged
*/
func (CSS) OnStyleSheetChanged(
	socket sock.Socketer,
	callback func(event *css.StyleSheetChangedEvent),
) {
	handler := sock.NewEventHandler(
		"CSS.styleSheetChanged",
		func(response *sock.Response) {
			event := &css.StyleSheetChangedEvent{}
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
OnStyleSheetRemoved adds a handler to the CSS.styleSheetRemoved event. CSS.styleSheetRemoved fires
whenever an active document stylesheet is removed.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#event-styleSheetRemoved
*/
func (CSS) OnStyleSheetRemoved(
	socket sock.Socketer,
	callback func(event *css.StyleSheetRemovedEvent),
) {
	handler := sock.NewEventHandler(
		"CSS.styleSheetRemoved",
		func(response *sock.Response) {
			event := &css.StyleSheetRemovedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

package css

import (
	"github.com/mkenney/go-chrome/cdtp/dom"
	"github.com/mkenney/go-chrome/cdtp/page"
)

/*
AddRuleParams represents CSS.addRule parameters.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-addRule
*/
type AddRuleParams struct {
	// The css style sheet identifier where a new rule should be inserted.
	StyleSheetID StyleSheetID `json:"styleSheetId"`

	// The text of a new rule.
	RuleText string `json:"ruleText"`

	// Text position of a new rule in the target style sheet.
	Location *SourceRange `json:"location"`
}

/*
AddRuleResult represents the result of calls to CSS.addRule.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-addRule
*/
type AddRuleResult struct {
	// The newly created rule.
	Rule *Rule `json:"rule"`
}

/*
CollectClassNamesParams represents CSS.collectClassNames parameters.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-collectClassNames
*/
type CollectClassNamesParams struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"`
}

/*
CollectClassNamesResult represents the result of calls to CSS.collectClassNames.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-collectClassNames
*/
type CollectClassNamesResult struct {
	// Class name list.
	ClassNames []string `json:"classNames"`
}

/*
CreateStyleSheetParams represents CSS.createStyleSheet parameters.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-createStyleSheet
*/
type CreateStyleSheetParams struct {
	FrameID page.FrameID `json:"frameId"`
}

/*
CreateStyleSheetResult represents the result of calls to CSS.createStyleSheet.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-createStyleSheet
*/
type CreateStyleSheetResult struct {
	// Identifier of the created "via-inspector" stylesheet.
	StyleSheetID StyleSheetID `json:"styleSheetId"`
}

/*
ForcePseudoStateParams represents CSS.forcePseudoState parameters.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-forcePseudoState
*/
type ForcePseudoStateParams struct {
	// The element ID for which to force the pseudo state.
	NodeID dom.NodeID `json:"nodeId"`

	// Element pseudo classes to force when computing the element's style.
	// Allowed values: active, focus, hover, visited.
	ForcedPseudoClasses []string `json:"forcedPseudoClasses"`
}

/*
GetBackgroundColorsParams represents CSS.getBackgroundColors parameters.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getBackgroundColors
*/
type GetBackgroundColorsParams struct {
	// Id of the node to get background colors for.
	NodeID dom.NodeID `json:"nodeId"`
}

/*
GetBackgroundColorsResult represents the result of calls to CSS.getBackgroundColors.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getBackgroundColors
*/
type GetBackgroundColorsResult struct {
	// Optional. The range of background colors behind this element, if it
	// contains any visible text. If no visible text is present, this will be
	// undefined. In the case of a flat background color, this will consist of
	// simply that color. In the case of a gradient, this will consist of each
	// of the color stops. For anything more complicated, this will be an empty
	// array. Images will be ignored (as if the image had failed to load).
	BackgroundColors []string `json:"backgroundColors,omitempty"`

	// Optional. The computed font size for this node, as a CSS computed value
	// string (e.g. '12px').
	ComputedFontSize string `json:"computedFontSize,omitempty"`

	// Optional. The computed font weight for this node, as a CSS computed value
	// string (e.g.
	// 'normal' or '100').
	ComputedFontWeight string `json:"computedFontWeight,omitempty"`

	// Optional. The computed font size for the document body, as a computed CSS
	// value string (e.g. '16px').
	ComputedBodyFontSize string `json:"computedBodyFontSize,omitempty"`
}

/*
GetComputedStyleForNodeParams represents CSS.getComputedStyleForNode parameters.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getComputedStyleForNode
*/
type GetComputedStyleForNodeParams struct {
	NodeID dom.NodeID `json:"nodeId"`
}

/*
GetComputedStyleForNodeResult represents the result of calls to CSS.getComputedStyleForNode.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getComputedStyleForNode
*/
type GetComputedStyleForNodeResult struct {
	// Computed style for the specified DOM node.
	ComputedStyle []*ComputedStyleProperty `json:"computedStyle"`
}

/*
GetInlineStylesForNodeParams represents CSS.getInlineStylesForNode parameters.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getInlineStylesForNode
*/
type GetInlineStylesForNodeParams struct {
	NodeID dom.NodeID `json:"nodeId"`
}

/*
GetInlineStylesForNodeResult represents the result of calls to CSS.getInlineStylesForNode.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getInlineStylesForNode
*/
type GetInlineStylesForNodeResult struct {
	// Optional. Inline style for the specified DOM node.
	InlineStyle *Style `json:"inlineStyle,omitempty"`

	// Optional. Attribute-defined element style (e.g. resulting from "width=20
	// height=100%").
	AttributesStyle *Style `json:"attributesStyle,omitempty"`
}

/*
GetMatchedStylesForNodeParams represents CSS.getMatchedStylesForNode parameters.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getMatchedStylesForNode
*/
type GetMatchedStylesForNodeParams struct {
	NodeID dom.NodeID `json:"nodeId"`
}

/*
GetMatchedStylesForNodeResult represents the result of calls to CSS.getMatchedStylesForNode.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getMatchedStylesForNode
*/
type GetMatchedStylesForNodeResult struct {
	// Inline style for the specified DOM node.
	InlineStyle *Style `json:"inlineStyle,omitempty"`

	// Attribute-defined element style (e.g. resulting from "width=20
	// height=100%").
	AttributesStyle *Style `json:"attributesStyle,omitempty"`

	// CSS rules matching this node, from all applicable stylesheets.
	MatchedRules []*RuleMatch `json:"matchedCSSRules,omitempty"`

	// Pseudo style matches for this node.
	PseudoElements []*PseudoElementMatches `json:"pseudoElements,omitempty"`

	// A chain of inherited styles (from the immediate node parent up to the DOM
	// tree root).
	Inherited []*InheritedStyleEntry `json:"inherited,omitempty"`

	// A list of CSS keyframed animations matching this node.
	KeyframesRules []*KeyframesRule `json:"cssKeyframesRules,omitempty"`
}

/*
GetMediaQueriesResult represents the result of calls to CSS.getMediaQueries.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getMediaQueries
*/
type GetMediaQueriesResult struct {
	Medias []*Media `json:"medias"`
}

/*
GetPlatformFontsForNodeParams represents CSS.getPlatformFontsForNode parameters.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getPlatformFontsForNode
*/
type GetPlatformFontsForNodeParams struct {
	NodeID dom.NodeID `json:"nodeId"`
}

/*
GetPlatformFontsForNodeResult represents the result of calls to CSS.getPlatformFontsForNode.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getPlatformFontsForNode
*/
type GetPlatformFontsForNodeResult struct {
	// Usage statistics for every employed platform font.
	Fonts []*PlatformFontUsage `json:"fonts"`
}

/*
GetStyleSheetTextParams represents CSS.getStyleSheetText parameters.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getStyleSheetText
*/
type GetStyleSheetTextParams struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"`
}

/*
GetStyleSheetTextResult represents the result of calls to CSS.getStyleSheetText.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getStyleSheetText
*/
type GetStyleSheetTextResult struct {
	// The stylesheet text.
	Text string `json:"text"`
}

/*
SetEffectivePropertyValueForNodeParams represents CSS.setEffectivePropertyValueForNode parameters.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setEffectivePropertyValueForNode
*/
type SetEffectivePropertyValueForNodeParams struct {
	// The element id for which to set property.
	NodeID       dom.NodeID `json:"nodeId"`
	PropertyName string     `json:"propertyName"`
	Value        string     `json:"value"`
}

/*
SetKeyframeKeyParams represents CSS.setKeyframeKey parameters.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setKeyframeKey
*/
type SetKeyframeKeyParams struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"`
	Range        *SourceRange `json:"range"`
	Selector     string       `json:"selector"`
}

/*
SetKeyframeKeyResult represents the result of calls to CSS.setKeyframeKey.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setKeyframeKey
*/
type SetKeyframeKeyResult struct {
	// The resulting key text after modification.
	KeyText *Value `json:"keyText"`
}

/*
SetMediaTextParams represents CSS.setMediaText parameters.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setMediaText
*/
type SetMediaTextParams struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"`
	Range        *SourceRange `json:"range"`
	Text         string       `json:"text"`
}

/*
SetMediaTextResult represents the result of calls to CSS.setMediaText.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setMediaText
*/
type SetMediaTextResult struct {
	// The resulting CSS media rule after modification.
	Media Media `json:"media"`
}

/*
SetRuleSelectorParams represents CSS.setRuleSelector parameters.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setRuleSelector
*/
type SetRuleSelectorParams struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"`
	Range        *SourceRange `json:"range"`
	Selector     string       `json:"selector"`
}

/*
SetRuleSelectorResult represents the result of calls to CSS.setRuleSelector.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setRuleSelector
*/
type SetRuleSelectorResult struct {
	// The resulting selector list after modification.
	SelectorList *SelectorList `json:"selectorList"`
}

/*
SetStyleSheetTextParams represents CSS.setStyleSheetText parameters.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setStyleSheetText
*/
type SetStyleSheetTextParams struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"`
	Text         string       `json:"text"`
}

/*
SetStyleSheetTextResult represents the result of calls to CSS.setStyleSheetText.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setStyleSheetText
*/
type SetStyleSheetTextResult struct {
	// URL of source map associated with script (if any).
	SourceMapURL string `json:"sourceMapURL"`
}

/*
SetStyleTextsParams represents CSS.setStyleTexts parameters.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setStyleTexts
*/
type SetStyleTextsParams struct {
	Edits []*StyleDeclarationEdit `json:"edits"`
}

/*
SetStyleTextsResult represents the result of calls to CSS.setStyleTexts.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setStyleTexts
*/
type SetStyleTextsResult struct {
	// The resulting styles after modification.
	Styles []*Style `json:"styles"`
}

/*
StopRuleUsageTrackingResult represents the result of calls to CSS.stopRuleUsageTracking.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-stopRuleUsageTracking
*/
type StopRuleUsageTrackingResult struct {
	RuleUsage []*RuleUsage `json:"ruleUsage"`
}

/*
TakeCoverageDeltaResult represents the result of calls to CSS.takeCoverageDelta.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-takeCoverageDelta
*/
type TakeCoverageDeltaResult struct {
	Coverage []*RuleUsage `json:"coverage"`
}

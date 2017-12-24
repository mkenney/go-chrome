package CSS

import (
	"github.com/mkenney/go-chrome/dom"
	"github.com/mkenney/go-chrome/page"
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
	Location SourceRange `json:"location"`
}

/*
AddRuleResult represents the result of calls to CSS.addRule.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-addRule
*/
type AddRuleResult struct {
	// The newly created rule.
	Rule Rule `json:"rule"`
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
	FrameID Page.FrameID `json:"frameId"`
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
	NodeID DOM.NodeID `json:"nodeId"`

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
	NodeID DOM.NodeID `json:"nodeId"`
}

/*
GetBackgroundColorsResult represents the result of calls to CSS.getBackgroundColors.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getBackgroundColors
*/
type GetBackgroundColorsResult struct {
	// Optional. The range of background colors behind this element, if it contains any visible
	// text. If no visible text is present, this will be undefined. In the case of a flat background
	// color, this will consist of simply that color. In the case of a gradient, this will consist
	// of each of the color stops. For anything more complicated, this will be an empty array.
	// Images will be ignored (as if the image had failed to load).
	BackgroundColors []string `json:"backgroundColors,omitempty"`

	// Optional. The computed font size for this node, as a CSS computed value string (e.g. '12px').
	ComputedFontSize string `json:"computedFontSize,omitempty"`

	// Optional. The computed font weight for this node, as a CSS computed value string (e.g.
	// 'normal' or '100').
	ComputedFontWeight string `json:"computedFontWeight,omitempty"`

	// Optional. The computed font size for the document body, as a computed CSS value string (e.g.
	// '16px').
	ComputedBodyFontSize string `json:"computedBodyFontSize,omitempty"`
}

/*
GetComputedStyleForNodeParams represents CSS.getComputedStyleForNode parameters.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getComputedStyleForNode
*/
type GetComputedStyleForNodeParams struct {
	NodeID DOM.NodeID `json:"nodeId"`
}

/*
GetComputedStyleForNodeResult represents the result of calls to CSS.getComputedStyleForNode.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getComputedStyleForNode
*/
type GetComputedStyleForNodeResult struct {
	// Computed style for the specified DOM node.
	ComputedStyle []ComputedStyleProperty `json:"computedStyle"`
}

/*
GetInlineStylesForNodeParams represents CSS.getInlineStylesForNode parameters.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getInlineStylesForNode
*/
type GetInlineStylesForNodeParams struct {
	NodeID DOM.NodeID `json:"nodeId"`
}

/*
GetInlineStylesForNodeResult represents the result of calls to CSS.getInlineStylesForNode.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getInlineStylesForNode
*/
type GetInlineStylesForNodeResult struct {
	// Optional. Inline style for the specified DOM node.
	InlineStyle Style `json:"inlineStyle,omitempty"`

	// Optional. Attribute-defined element style (e.g. resulting from "width=20 height=100%").
	AttributesStyle Style `json:"attributesStyle,omitempty"`
}

/*
GetMatchedStylesForNodeParams represents CSS.getMatchedStylesForNode parameters.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getMatchedStylesForNode
*/
type GetMatchedStylesForNodeParams struct {
	NodeID DOM.NodeID `json:"nodeId"`
}

/*
GetMatchedStylesForNodeResult represents the result of calls to CSS.getMatchedStylesForNode.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getMatchedStylesForNode
*/
type GetMatchedStylesForNodeResult struct {
	// Inline style for the specified DOM node.
	InlineStyle Style `json:"inlineStyle,omitempty"`

	// Attribute-defined element style (e.g. resulting from "width=20 height=100%").
	AttributesStyle Style `json:"attributesStyle,omitempty"`

	// CSS rules matching this node, from all applicable stylesheets.
	MatchedCSSRules []RuleMatch `json:"matchedCSSRules,omitempty"`

	// Pseudo style matches for this node.
	PseudoElements []PseudoElementMatches `json:"pseudoElements,omitempty"`

	// A chain of inherited styles (from the immediate node parent up to the DOM tree root).
	Inherited []InheritedStyleEntry `json:"inherited,omitempty"`

	// A list of CSS keyframed animations matching this node.
	KeyframesRules []KeyframesRule `json:"cssKeyframesRules,omitempty"`
}

/*
GetMediaQueriesResult represents the result of calls to CSS.getMediaQueries.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getMediaQueries
*/
type GetMediaQueriesResult struct {
	Medias []Media `json:"medias"`
}

/*
GetPlatformFontsForNodeParams represents CSS.getPlatformFontsForNode parameters.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getPlatformFontsForNode
*/
type GetPlatformFontsForNodeParams struct {
	NodeID DOM.NodeID `json:"nodeId"`
}

/*
GetPlatformFontsForNodeResult represents the result of calls to CSS.getPlatformFontsForNode.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getPlatformFontsForNode
*/
type GetPlatformFontsForNodeResult struct {
	// Usage statistics for every employed platform font.
	Fonts []PlatformFontUsage `json:"fonts"`
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
	NodeID       DOM.NodeID `json:"nodeId"`
	PropertyName string     `json:"propertyName"`
	Value        string     `json:"value"`
}

/*
SetKeyframeKeyParams represents CSS.setKeyframeKey parameters.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setKeyframeKey
*/
type SetKeyframeKeyParams struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"`
	Range        SourceRange  `json:"range"`
	Selector     string       `json:"selector"`
}

/*
SetKeyframeKeyResult represents the result of calls to CSS.setKeyframeKey.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setKeyframeKey
*/
type SetKeyframeKeyResult struct {
	// The resulting key text after modification.
	KeyText Value `json:"keyText"`
}

/*
SetMediaTextParams represents CSS.setMediaText parameters.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setMediaText
*/
type SetMediaTextParams struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"`
	Range        SourceRange  `json:"range"`
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
	Range        SourceRange  `json:"range"`
	Selector     string       `json:"selector"`
}

/*
SetRuleSelectorResult represents the result of calls to CSS.setRuleSelector.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setRuleSelector
*/
type SetRuleSelectorResult struct {
	// The resulting selector list after modification.
	SelectorList SelectorList `json:"selectorList"`
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
	Styles []Style `json:"styles"`
}

/*
StopRuleUsageTrackingResult represents the result of calls to CSS.stopRuleUsageTracking.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-stopRuleUsageTracking
*/
type StopRuleUsageTrackingResult struct {
	RuleUsage []RuleUsage `json:"ruleUsage"`
}

/*
TakeCoverageDeltaResult represents the result of calls to CSS.takeCoverageDelta.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-takeCoverageDelta
*/
type TakeCoverageDeltaResult struct {
	Coverage []RuleUsage `json:"coverage"`
}

/*
FontsUpdatedEvent represents CSS.fontsUpdated event data.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#event-fontsUpdated
*/
type FontsUpdatedEvent struct{}

/*
MediaQueryResultChangedEvent represents CSS.mediaQueryResultChanged event data.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#event-mediaQueryResultChanged
*/
type MediaQueryResultChangedEvent struct{}

/*
StyleSheetAddedEvent represents CSS.styleSheetAdded event data.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#event-styleSheetAdded
*/
type StyleSheetAddedEvent struct {
	// Added stylesheet metainfo.
	Header StyleSheetHeader `json:"header"`
}

/*
StyleSheetChangedEvent represents CSS.styleSheetChanged event data.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#event-styleSheetChanged
*/
type StyleSheetChangedEvent struct {
	// Identifier of the changed stylesheet.
	StyleSheetID StyleSheetID `json:"styleSheetId"`
}

/*
StyleSheetRemovedEvent represents CSS.styleSheetRemoved event data.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#event-styleSheetRemoved
*/
type StyleSheetRemovedEvent struct {
	// Identifier of the removed stylesheet.
	StyleSheetID StyleSheetID `json:"styleSheetId"`
}

/*
StyleSheetID is the ID of a stylesheet

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#type-StyleSheetId
*/
type StyleSheetID string

/*
StyleSheetOrigin is an enum defining:
	- "injected" for stylesheets injected via extension
	- "user-agent" for user-agent stylesheets
	- "inspector" for stylesheets created by the inspector (i.e. those holding the "via inspector"
	  rules)
	- "regular" for regular stylesheets.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#type-StyleSheetOrigin
*/
type StyleSheetOrigin string

/*
PseudoElementMatches is a CSS rule collection for a single pseudo style.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#type-PseudoElementMatches
*/
type PseudoElementMatches struct {
	// Pseudo element type.
	PseudoType *DOM.PseudoType `json:"pseudoType"`

	// Matches of CSS rules applicable to the pseudo style.
	Matches []*RuleMatch `json:"matches"`
}

/*
InheritedStyleEntry is an inherited CSS rule collection from ancestor node.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#type-InheritedStyleEntry
*/
type InheritedStyleEntry struct {
	// Optional. The ancestor node's inline style, if any, in the style inheritance chain.
	InlineStyle Style `json:"inlineStyle,omitempty"`

	// Matches of CSS rules matching the ancestor node in the style inheritance chain.
	MatchedCSSRules []*RuleMatch `json:"matchedCSSRules"`
}

/*
RuleMatch is match data for a CSS rule.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#type-RuleMatch
*/
type RuleMatch struct {
	// CSS rule in the match.
	Rule Rule `json:"rule"`

	// Matching selector indices in the rule's selectorList selectors (0-based).
	MatchingSelectors []int `json:"matchingSelectors"`
}

/*
Value represents data for a simple selector (these are delimited by commas in a selector list).

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#type-Value
*/
type Value struct {
	// Value text.
	Text string `json:"text"`

	// Optional. Value range in the underlying resource (if available).
	Range SourceRange `json:"range,omitempty"`
}

/*
SelectorList represents selector list data.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#type-SelectorList
*/
type SelectorList struct {
	// Selectors in the list.
	Selectors []*Value `json:"selectors"`

	// Rule selector text.
	Text string `json:"text"`
}

/*
StyleSheetHeader holds CSS stylesheet metainformation.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#type-CSSStyleSheetHeader
*/
type StyleSheetHeader struct {
	// The stylesheet identifier.
	StyleSheetID StyleSheetID `json:"styleSheetId"`

	// Owner frame identifier.
	FrameID Page.FrameID `json:"frameId"`

	// Stylesheet resource URL.
	SourceURL string `json:"sourceURL"`

	// Optional. URL of source map associated with the stylesheet (if any).
	SourceMapURL string `json:"sourceMapURL,omitempty"`

	// Stylesheet origin.
	Origin StyleSheetOrigin `json:"origin"`

	// Stylesheet title.
	Title string `json:"title"`

	// Optional. The backend ID for the owner node of the stylesheet.
	OwnerNode DOM.BackendNodeID `json:"ownerNode,omitempty"`

	// Denotes whether the stylesheet is disabled.
	Disabled bool `json:"disabled"`

	// Optional. Whether the sourceURL field value comes from the sourceURL comment.
	HasSourceURL bool `json:"hasSourceURL,omitempty"`

	// Whether this stylesheet is created for STYLE tag by parser. This flag is not set for
	// document.written STYLE tags.
	IsInline bool `json:"isInline"`

	// Line offset of the stylesheet within the resource (zero based).
	StartLine float64 `json:"startLine"`

	// Column offset of the stylesheet within the resource (zero based).
	StartColumn float64 `json:"startColumn"`

	// Size of the content (in characters).
	Length float64 `json:"length"`
}

/*
Rule is a CSS rule representation.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#type-CSSRule
*/
type Rule struct {
	// Optional. The css style sheet identifier (absent for user agent stylesheet and user-specified
	// stylesheet rules) this rule came from.
	StyleSheetID StyleSheetID `json:"styleSheetId,omitempty"`

	// Rule selector data.
	SelectorList SelectorList `json:"selectorList"`

	// Parent stylesheet's origin.
	Origin StyleSheetOrigin `json:"origin"`

	// Associated style declaration.
	Style Style `json:"style"`

	// Optional. Media list array (for rules involving media queries). The array enumerates media
	// queries starting with the innermost one, going outwards.
	Media []*Media `json:"media,omitempty"`
}

/*
RuleUsage holds CSS coverage information.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#type-RuleUsage
*/
type RuleUsage struct {
	// The css style sheet identifier (absent for user agent stylesheet and user-specified
	// stylesheet rules) this rule came from.
	StyleSheetID StyleSheetID `json:"styleSheetId"`

	// Offset of the start of the rule (including selector) from the beginning of the stylesheet.
	StartOffset float64 `json:"startOffset"`

	// Offset of the end of the rule body from the beginning of the stylesheet.
	EndOffset float64 `json:"endOffset"`

	// Indicates whether the rule was actually used by some element in the DOM.
	Used bool `json:"used"`
}

/*
SourceRange is a text range within a resource. All numbers are zero-based.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#type-SourceRange
*/
type SourceRange struct {
	// Start line of range.
	StartLine int `json:"startLine"`

	// Start column of range (inclusive).
	StartColumn int `json:"startColumn"`

	// End line of range.
	SndLine int `json:"endLine"`

	// End column of range (exclusive).
	SndColumn int `json:"endColumn"`
}

/*
ShorthandEntry holds a CSS shorthand entry

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#type-ShorthandEntry
*/
type ShorthandEntry struct {
	// Shorthand name.
	Name string `json:"name"`

	// Shorthand value.
	Value string `json:"value"`

	// Optional. Whether the property has "!important" annotation (implies false if absent).
	Important bool `json:"important,omitempty"`
}

/*
ComputedStyleProperty holds a computed style property

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#type-CSSComputedStyleProperty
*/
type ComputedStyleProperty struct {
	// Computed style property name.
	Name string `json:"name"`

	// Computed style property value.
	Value string `json:"value"`
}

/*
Style is a CSS style representation.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#type-CSSStyle
*/
type Style struct {
	// Optional. The css style sheet identifier (absent for user agent stylesheet and user-specified
	// stylesheet rules) this rule came from.
	StyleSheetID StyleSheetID `json:"styleSheetId,omitempty"`

	// CSS properties in the style.
	Properties []*Property `json:"cssProperties"`

	// Computed values for all shorthands found in the style.
	ShorthandEntries []*ShorthandEntry `json:"shorthandEntries"`

	// Style declaration text (if available).
	CSSText string `json:"cssText"`

	// Optional. Style declaration range in the enclosing stylesheet (if available).
	Range SourceRange `json:"range,omitempty"`
}

/*
Property holds CSS property declaration data.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#type-CSSProperty
*/
type Property struct {
	// The property name.
	Name string `json:"name"`

	// The property value.
	Value string `json:"value"`

	// Optional. Whether the property has "!important" annotation (implies false if absent).
	Important bool `json:"important,omitempty"`

	// Optional. Whether the property is implicit (implies false if absent).
	Implicit bool `json:"implicit,omitempty"`

	// Optional. The full property text as specified in the style.
	Text string `json:"text,omitempty"`

	// Optional. Whether the property is understood by the browser (implies true
	// if absent).
	ParsedOk bool `json:"parsedOk,omitempty"`

	// Optional. Whether the property is disabled by the user (present for
	// source-based properties only).
	Disabled bool `json:"disabled,omitempty"`

	// Optional. The entire property range in the enclosing style declaration (if available).
	Range *SourceRange `json:"range,omitempty"`
}

/*
Media holds a CSS media rule descriptor.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#type-CSSMedia
*/
type Media struct {
	// Media query text.
	Text string `json:"text"`

	// Source of the media query:
	//	- "mediaRule" if specified by a @media rule
	//	- "importRule" if specified by an @import rule
	//	- "linkedSheet" if specified by a "media" attribute in a linked stylesheet's LINK tag
	//	- "inlineSheet" if specified by a "media" attribute in an inline stylesheet's STYLE tag.
	//	  Allowed values: mediaRule, importRule, linkedSheet, inlineSheet.
	Source string `json:"source"`

	// Optional. URL of the document containing the media query description.
	SourceURL string `json:"sourceURL,omitempty"`

	// Optional. The associated rule (@media or @import) header range in the enclosing stylesheet
	// (if available).
	Range *SourceRange `json:"range,omitempty"`

	// Optional. Identifier of the stylesheet containing this object (if exists).
	StyleSheetID StyleSheetID `json:"styleSheetId,omitempty"`

	// Optional. Array of media queries.
	MediaList []*MediaQuery `json:"mediaList,omitempty"`
}

/*
MediaQuery is a media query descriptor.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#type-MediaQuery
*/
type MediaQuery struct {
	// Array of media query expressions.
	Expressions []*MediaQueryExpression `json:"expressions"`

	// Whether the media query condition is satisfied.
	Active bool `json:"active"`
}

/*
MediaQueryExpression is a media query expression descriptor.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#type-MediaQueryExpression
*/
type MediaQueryExpression struct {
	// Media query expression value.
	Value float64 `json:"value"`

	// Media query expression units.
	Unit string `json:"unit"`

	// Media query expression feature.
	Feature string `json:"feature"`

	// Optional. The associated range of the value text in the enclosing stylesheet (if available).
	ValueRange *SourceRange `json:"valueRange,omitempty"`

	// Optional. Computed length of media query expression (if applicable).
	ComputedLength float64 `json:"computedLength,omitempty"`
}

/*
PlatformFontUsage holds information about the amount of glyphs that were rendered with given font.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#type-PlatformFontUsage
*/
type PlatformFontUsage struct {
	// Font's family name reported by platform.
	FamilyName string `json:"familyName"`

	// Indicates if the font was downloaded or resolved locally.
	IsCustomFont bool `json:"isCustomFont"`

	// Amount of glyphs that were rendered with this font.
	GlyphCount float64 `json:"glyphCount"`
}

/*
KeyframesRule is a CSS keyframes rule representation.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#type-CSSKeyframesRule
*/
type KeyframesRule struct {
	// Animation name.
	AnimationName *Value `json:"animationName"`

	// List of keyframes.
	Keyframes []*KeyframeRule `json:"keyframes"`
}

/*
KeyframeRule is a CSS keyframe rule representation.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#type-CSSKeyframeRule
*/
type KeyframeRule struct {
	// Optional. The css style sheet identifier (absent for user agent stylesheet and user-specified
	// stylesheet rules) this rule came from.
	StyleSheetID StyleSheetID `json:"styleSheetId,omitempty"`

	// Parent stylesheet's origin.
	Origin StyleSheetOrigin `json:"origin"`

	// Associated key text.
	KeyText Value `json:"keyText"`

	// Associated style declaration.
	Style Style `json:"style"`
}

/*
StyleDeclarationEdit is a descriptor of operation to mutate style declaration text.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#type-StyleDeclarationEdit
*/
type StyleDeclarationEdit struct {
	// The css style sheet identifier.
	StyleSheetID StyleSheetID `json:"styleSheetId"`

	// The range of the style text in the enclosing stylesheet.
	Range *SourceRange `json:"range"`

	// New style text.
	Text string `json:"text"`
}

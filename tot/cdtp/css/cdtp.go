/*
Package css provides type definitions for use with the Chrome CSS protocol

https://chromedevtools.github.io/devtools-protocol/tot/CSS/
*/
package css

import (
	"github.com/mkenney/go-chrome/tot/cdtp/dom"
	"github.com/mkenney/go-chrome/tot/cdtp/page"
)

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
	PseudoType dom.PseudoType `json:"pseudoType"`

	// Matches of CSS rules applicable to the pseudo style.
	Matches []*RuleMatch `json:"matches"`
}

/*
InheritedStyleEntry is an inherited CSS rule collection from ancestor node.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#type-InheritedStyleEntry
*/
type InheritedStyleEntry struct {
	// Optional. The ancestor node's inline style, if any, in the style
	// inheritance chain.
	InlineStyle *Style `json:"inlineStyle,omitempty"`

	// Matches of CSS rules matching the ancestor node in the style
	// inheritance chain.
	MatchedRules []*RuleMatch `json:"matchedCSSRules"`
}

/*
RuleMatch is match data for a CSS rule.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#type-RuleMatch
*/
type RuleMatch struct {
	// CSS rule in the match.
	Rule *Rule `json:"rule"`

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
	Range *SourceRange `json:"range,omitempty"`
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
	FrameID page.FrameID `json:"frameId"`

	// Stylesheet resource URL.
	SourceURL string `json:"sourceURL"`

	// Optional. URL of source map associated with the stylesheet (if any).
	SourceMapURL string `json:"sourceMapURL,omitempty"`

	// Stylesheet origin.
	Origin StyleSheetOrigin `json:"origin"`

	// Stylesheet title.
	Title string `json:"title"`

	// Optional. The backend ID for the owner node of the stylesheet.
	OwnerNode dom.BackendNodeID `json:"ownerNode,omitempty"`

	// Denotes whether the stylesheet is disabled.
	Disabled bool `json:"disabled"`

	// Optional. Whether the sourceURL field value comes from the sourceURL
	// comment.
	HasSourceURL bool `json:"hasSourceURL,omitempty"`

	// Whether this stylesheet is created for STYLE tag by parser. This flag is
	// not set for
	// document.written STYLE tags.
	IsInline bool `json:"isInline"`

	// Line offset of the stylesheet within the resource (zero based).
	StartLine int64 `json:"startLine"`

	// Column offset of the stylesheet within the resource (zero based).
	StartColumn int64 `json:"startColumn"`

	// Size of the content (in characters).
	Length int64 `json:"length"`
}

/*
Rule is a CSS rule representation.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#type-CSSRule
*/
type Rule struct {
	// Optional. The css style sheet identifier (absent for user agent s
	// tylesheet and user-specified stylesheet rules) this rule came from.
	StyleSheetID StyleSheetID `json:"styleSheetId,omitempty"`

	// Rule selector data.
	SelectorList *SelectorList `json:"selectorList"`

	// Parent stylesheet's origin.
	Origin StyleSheetOrigin `json:"origin"`

	// Associated style declaration.
	Style *Style `json:"style"`

	// Optional. Media list array (for rules involving media queries). The array
	// enumerates media queries starting with the innermost one, going outwards.
	Media []*Media `json:"media,omitempty"`
}

/*
RuleUsage holds CSS coverage information.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#type-RuleUsage
*/
type RuleUsage struct {
	// The css style sheet identifier (absent for user agent stylesheet and
	// user-specified stylesheet rules) this rule came from.
	StyleSheetID StyleSheetID `json:"styleSheetId"`

	// Offset of the start of the rule (including selector) from the beginning
	// of the stylesheet.
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
	EndLine int `json:"endLine"`

	// End column of range (exclusive).
	EndColumn int `json:"endColumn"`
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

	// Optional. Whether the property has "!important" annotation (implies false
	// if absent).
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
	// Optional. The css style sheet identifier (absent for user agent
	// stylesheet and user-specified stylesheet rules) this rule came from.
	StyleSheetID StyleSheetID `json:"styleSheetId,omitempty"`

	// CSS properties in the style.
	Properties []*Property `json:"cssProperties"`

	// Computed values for all shorthands found in the style.
	ShorthandEntries []*ShorthandEntry `json:"shorthandEntries"`

	// Style declaration text (if available).
	Text string `json:"cssText"`

	// Optional. Style declaration range in the enclosing stylesheet (if
	// available).
	Range *SourceRange `json:"range,omitempty"`
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

	// Optional. Whether the property has "!important" annotation (implies false
	// if absent).
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

	// Optional. The entire property range in the enclosing style declaration
	// (if available).
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
	//	- "linkedSheet" if specified by a "media" attribute in a linked
	//    stylesheet's LINK tag
	//	- "inlineSheet" if specified by a "media" attribute in an inline
	//    stylesheet's STYLE tag.
	//
	// Allowed values: mediaRule, importRule, linkedSheet, inlineSheet.
	Source string `json:"source"`

	// Optional. URL of the document containing the media query description.
	SourceURL string `json:"sourceURL,omitempty"`

	// Optional. The associated rule (@media or @import) header range in the
	// enclosing stylesheet
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

	// Optional. The associated range of the value text in the enclosing
	// stylesheet (if available).
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
	// Optional. The css style sheet identifier (absent for user agent
	// stylesheet and user-specified stylesheet rules) this rule came from.
	StyleSheetID StyleSheetID `json:"styleSheetId,omitempty"`

	// Parent stylesheet's origin.
	Origin StyleSheetOrigin `json:"origin"`

	// Associated key text.
	KeyText *Value `json:"keyText"`

	// Associated style declaration.
	Style *Style `json:"style"`
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

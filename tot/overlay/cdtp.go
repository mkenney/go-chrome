/*
Package overlay provides type definitions for use with the Chrome Overlay protocol

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/
*/
package overlay

import (
	"github.com/mkenney/go-chrome/tot/dom"
)

/*
HighlightConfig is configuration data for the highlighting of page elements.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#type-HighlightConfig
*/
type HighlightConfig struct {
	// Optional. Whether the node info tooltip should be shown (default: false).
	ShowInfo bool `json:"showInfo,omitempty"`

	// Optional. Whether the rulers should be shown (default: false).
	ShowRulers bool `json:"showRulers,omitempty"`

	// Optional. Whether the extension lines from node to the rulers should be
	// shown (default: false).
	ShowExtensionLines bool `json:"showExtensionLines,omitempty"`

	// Optional. Display as material.
	DisplayAsMaterial bool `json:"displayAsMaterial,omitempty"`

	// Optional. The content box highlight fill color (default: transparent).
	ContentColor *dom.RGBA `json:"contentColor,omitempty"`

	// Optional. The padding highlight fill color (default: transparent).
	PaddingColor *dom.RGBA `json:"paddingColor,omitempty"`

	// Optional. The border highlight fill color (default: transparent).
	BorderColor *dom.RGBA `json:"borderColor,omitempty"`

	// Optional. The margin highlight fill color (default: transparent).
	MarginColor *dom.RGBA `json:"marginColor,omitempty"`

	// Optional. The event target element highlight fill color (default:
	// transparent).
	EventTargetColor *dom.RGBA `json:"eventTargetColor,omitempty"`

	// Optional. The shape outside fill color (default: transparent).
	ShapeColor *dom.RGBA `json:"shapeColor,omitempty"`

	// Optional. The shape margin fill color (default: transparent).
	ShapeMarginColor *dom.RGBA `json:"shapeMarginColor,omitempty"`

	// Optional. Selectors to highlight relevant nodes.
	SelectorList string `json:"selectorList,omitempty"`

	// Optional. The grid layout color (default: transparent).
	CSSGridColor *dom.RGBA `json:"cssGridColor,omitempty"`
}

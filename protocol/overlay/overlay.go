package Overlay

import (
	DOM "app/chrome/protocol/dom"
	"fmt"
)

/*
HighlightConfig is configuration data for the highlighting of page elements.
*/
type HighlightConfig struct {
	// Optional. Whether the node info tooltip should be shown (default: false).
	ShowInfo bool `json:"showInfo,omitempty"`

	// Optional. Whether the rulers should be shown (default: false).
	ShowRulers bool `json:"showRulers,omitempty"`

	// Optional. Whether the extension lines from node to the rulers should be shown (default:
	// false).
	ShowExtensionLines bool `json:"showExtensionLines,omitempty"`

	// Optional. Display as material.
	DisplayAsMaterial bool `json:"displayAsMaterial,omitempty"`

	// Optional. The content box highlight fill color (default: transparent).
	ContentColor *DOM.RGBA `json:"contentColor,omitempty"`

	// Optional. The padding highlight fill color (default: transparent).
	PaddingColor *DOM.RGBA `json:"paddingColor,omitempty"`

	// Optional. The border highlight fill color (default: transparent).
	BorderColor *DOM.RGBA `json:"borderColor,omitempty"`

	// Optional. The margin highlight fill color (default: transparent).
	MarginColor *DOM.RGBA `json:"marginColor,omitempty"`

	// Optional. The event target element highlight fill color (default: transparent).
	EventTargetColor *DOM.RGBA `json:"eventTargetColor,omitempty"`

	// Optional. The shape outside fill color (default: transparent).
	ShapeColor *DOM.RGBA `json:"shapeColor,omitempty"`

	// Optional. The shape margin fill color (default: transparent).
	ShapeMarginColor *DOM.RGBA `json:"shapeMarginColor,omitempty"`

	// Optional. Selectors to highlight relevant nodes.
	SelectorList string `json:"selectorList,omitempty"`

	// Optional. The grid layout color (default: transparent).
	CSSGridColor *DOM.RGBA `json:"cssGridColor,omitempty"`
}

/*
InspectMode is the inspect mode
*/
type InspectMode int

const (
	_searchForNode InspectMode = iota
	_searchForUAShadowDOM
	_none
)

func (a InspectMode) String() string {
	if a == 0 {
		return "searchForNode"
	}
	if a == 1 {
		return "searchForUAShadowDOM"
	}
	if a == 2 {
		return "none"
	}
	panic(fmt.Errorf("Invalid InspectMode %d", a))
}

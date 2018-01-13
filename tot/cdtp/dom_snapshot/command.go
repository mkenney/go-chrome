package domSnapshot

/*
GetParams represents DOMSnapshot.getSnapshot parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOMSnapshot/#method-getSnapshot
*/
type GetParams struct {
	// Whitelist of computed styles to return.
	ComputedStyleWhitelist []string `json:"computedStyleWhitelist"`
}

/*
GetResult represents the result of calls to DOMSnapshot.getSnapshot.

https://chromedevtools.github.io/devtools-protocol/tot/DOMSnapshot/#method-getSnapshot
*/
type GetResult struct {
	// The nodes in the DOM tree. The DOMNode at index 0 corresponds to the root
	// document.
	DOMNodes []*DOMNode `json:"domNodes"`

	// The nodes in the layout tree.
	LayoutTreeNodes []*LayoutTreeNode `json:"layoutTreeNodes"`

	// Whitelisted ComputedStyle properties for each node in the layout tree.
	ComputedStyles []*ComputedStyle `json:"computedStyles"`

	// Error information related to executing this method
	Err error `json:"-"`
}

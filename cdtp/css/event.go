package css

/*
FontsUpdatedEvent represents CSS.fontsUpdated event data.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#event-fontsUpdated
*/
type FontsUpdatedEvent struct {
	// Error information related to this event
	Err error `json:"-"`
}

/*
MediaQueryResultChangedEvent represents CSS.mediaQueryResultChanged event data.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#event-mediaQueryResultChanged
*/
type MediaQueryResultChangedEvent struct {
	// Error information related to this event
	Err error `json:"-"`
}

/*
StyleSheetAddedEvent represents CSS.styleSheetAdded event data.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#event-styleSheetAdded
*/
type StyleSheetAddedEvent struct {
	// Added stylesheet metainfo.
	Header *StyleSheetHeader `json:"header"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
StyleSheetChangedEvent represents CSS.styleSheetChanged event data.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#event-styleSheetChanged
*/
type StyleSheetChangedEvent struct {
	// Identifier of the changed stylesheet.
	StyleSheetID StyleSheetID `json:"styleSheetId"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
StyleSheetRemovedEvent represents CSS.styleSheetRemoved event data.

https://chromedevtools.github.io/devtools-protocol/tot/CSS/#event-styleSheetRemoved
*/
type StyleSheetRemovedEvent struct {
	// Identifier of the removed stylesheet.
	StyleSheetID StyleSheetID `json:"styleSheetId"`

	// Error information related to this event
	Err error `json:"-"`
}

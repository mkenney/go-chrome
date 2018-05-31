package animation

/*
CanceledEvent represents Animation.animationCanceled event data.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#event-animationCanceled
*/
type CanceledEvent struct {
	// ID of the animation that was cancelled.
	ID string `json:"id"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
CreatedEvent represents Animation.animationCreated event data.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#event-animationCreated
*/
type CreatedEvent struct {
	// ID of the animation that was created.
	ID string `json:"id"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
StartedEvent represents Animation.animationStarted event data.

https://chromedevtools.github.io/devtools-protocol/tot/Animation/#event-animationStarted
*/
type StartedEvent struct {
	// Animation that was started.
	Animation *Animation `json:"animation"`

	// Error information related to this event
	Err error `json:"-"`
}

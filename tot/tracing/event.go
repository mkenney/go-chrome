package tracing

import (
	"github.com/mkenney/go-chrome/tot/io"
)

/*
BufferUsageEvent represents Overlay.bufferUsage event data.

https://chromedevtools.github.io/devtools-protocol/tot/Tracing/#event-bufferUsage
*/
type BufferUsageEvent struct {
	// Optional. A number in range [0..1] that indicates the used size of event
	// buffer as a fraction of its total size.
	PercentFull float64 `json:"percentFull,omitempty"`

	// Optional. An approximate number of events in the trace log.
	EventCount int64 `json:"eventCount,omitempty"`

	// Optional. A number in range [0..1] that indicates the used size of event
	// buffer as a fraction of its total size.
	Value float64 `json:"value,omitempty"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
DataCollectedEvent represents Overlay.dataCollected event data.

https://chromedevtools.github.io/devtools-protocol/tot/Tracing/#event-dataCollected
*/
type DataCollectedEvent struct {
	Value []map[string]string `json:"value"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
CompleteEvent represents Overlay.tracingComplete event data.

https://chromedevtools.github.io/devtools-protocol/tot/Tracing/#event-tracingComplete
*/
type CompleteEvent struct {
	// Optional. A handle of the stream that holds resulting trace data.
	Stream io.StreamHandle `json:"stream,omitempty"`

	// Error information related to this event
	Err error `json:"-"`
}

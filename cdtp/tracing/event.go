package tracing

import (
	"github.com/mkenney/go-chrome/cdtp/io"
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
	EventCount float64 `json:"eventCount,omitempty"`

	// Optional. A number in range [0..1] that indicates the used size of event
	// buffer as a fraction of its total size.
	Value float64 `json:"value,omitempty"`
}

/*
DataCollectedEvent represents Overlay.dataCollected event data.

https://chromedevtools.github.io/devtools-protocol/tot/Tracing/#event-dataCollected
*/
type DataCollectedEvent struct {
	Value []map[string]string `json:"value"`
}

/*
CompleteEvent represents Overlay.tracingComplete event data.

https://chromedevtools.github.io/devtools-protocol/tot/Tracing/#event-tracingComplete
*/
type CompleteEvent struct {
	// Optional. A handle of the stream that holds resulting trace data.
	Stream io.StreamHandle `json:"stream,omitempty"`
}

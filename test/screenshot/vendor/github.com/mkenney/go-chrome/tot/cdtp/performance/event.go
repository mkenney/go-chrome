package performance

/*
MetricsEvent represents Performance.metrics event data.

https://chromedevtools.github.io/devtools-protocol/tot/Performance/#event-metrics
*/
type MetricsEvent struct {
	// Current values of the metrics.
	Metrics []*Metric `json:"metrics"`

	// Timestamp title.
	Title string `json:"title"`

	// Error information related to this event
	Err error `json:"-"`
}

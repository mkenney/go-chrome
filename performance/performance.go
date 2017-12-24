package Performance

/*
GetMetricsResult represents the result of calls to Performance.getMetrics.

https://chromedevtools.github.io/devtools-protocol/tot/Performance/#method-getMetrics
*/
type GetMetricsResult struct {
	// Current values for run-time metrics.
	Metrics []Metric `json:"metrics"`
}

/*
MetricsEvent represents Performance.metrics event data.

https://chromedevtools.github.io/devtools-protocol/tot/Performance/#event-metrics
*/
type MetricsEvent struct {
	// Current values of the metrics.
	Metrics []Metric `json:"metrics"`

	// Timestamp title.
	Title string `json:"title"`
}

/*
Metric is a run-time execution metric.

https://chromedevtools.github.io/devtools-protocol/tot/Performance/#type-Metric
*/
type Metric struct {
	// Metric name.
	Name string `json:"name"`

	// Metric value.
	Value int `json:"value"`
}

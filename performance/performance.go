package Performance

/*
Metric is a run-time execution metric.
*/
type Metric struct {
	// Metric name.
	Name string `json:"name"`

	// Metric value.
	Value int `json:"value"`
}

/*
GetMetricsResult represents the result of calls to Performance.getMetrics.
*/
type GetMetricsResult struct {
	// Current values for run-time metrics.
	Metrics []Metric `json:"metrics"`
}

/*
MetricsEvent represents Performance.metrics event data.
*/
type MetricsEvent struct {
	// Current values of the metrics.
	Metrics []Metric `json:"metrics"`

	// Timestamp title.
	Title string `json:"title"`
}

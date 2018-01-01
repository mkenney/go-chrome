package performance

/*
GetMetricsResult represents the result of calls to Performance.getMetrics.

https://chromedevtools.github.io/devtools-protocol/tot/Performance/#method-getMetrics
*/
type GetMetricsResult struct {
	// Current values for run-time metrics.
	Metrics []*Metric `json:"metrics"`
}

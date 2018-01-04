package performance

/*
DisableResult represents the result of calls to Performance.disable.

https://chromedevtools.github.io/devtools-protocol/tot/Performance/#method-disable
*/
type DisableResult struct {
	CDTPError error `json:"-"`
}

/*
EnableResult represents the result of calls to Performance.enable.

https://chromedevtools.github.io/devtools-protocol/tot/Performance/#method-enable
*/
type EnableResult struct {
	CDTPError error `json:"-"`
}

/*
GetMetricsResult represents the result of calls to Performance.getMetrics.

https://chromedevtools.github.io/devtools-protocol/tot/Performance/#method-getMetrics
*/
type GetMetricsResult struct {
	// Current values for run-time metrics.
	Metrics []*Metric `json:"metrics"`

	CDTPError error `json:"-"`
}

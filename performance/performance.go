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

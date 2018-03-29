/*
Package performance provides type definitions for use with the Chrome Performance protocol

https://chromedevtools.github.io/devtools-protocol/tot/Performance/
*/
package performance

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

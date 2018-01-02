package log

/*
StartViolationsReportParams represents LayerTree.startViolationsReport parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Log/#method-startViolationsReport
*/
type StartViolationsReportParams struct {
	// Configuration for violations.
	Config []*ViolationSetting `json:"config"`
}

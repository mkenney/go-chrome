package log

/*
ClearResult represents the result of calls to Log.clear.

https://chromedevtools.github.io/devtools-protocol/tot/Log/#method-clear
*/
type ClearResult struct {
	CDTPError error `json:"-"`
}

/*
DisableResult represents the result of calls to Log.disable.

https://chromedevtools.github.io/devtools-protocol/tot/Log/#method-disable
*/
type DisableResult struct {
	CDTPError error `json:"-"`
}

/*
EnableResult represents the result of calls to Log.enable.

https://chromedevtools.github.io/devtools-protocol/tot/Log/#method-enable
*/
type EnableResult struct {
	CDTPError error `json:"-"`
}

/*
StartViolationsReportParams represents LayerTree.startViolationsReport parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Log/#method-startViolationsReport
*/
type StartViolationsReportParams struct {
	// Configuration for violations.
	Config []*ViolationSetting `json:"config"`
}

/*
StartViolationsReportResult represents the result of calls to Log.startViolationsReport.

https://chromedevtools.github.io/devtools-protocol/tot/Log/#method-startViolationsReport
*/
type StartViolationsReportResult struct {
	CDTPError error `json:"-"`
}

/*
StopViolationsReportResult represents the result of calls to Log.stopViolationsReport.

https://chromedevtools.github.io/devtools-protocol/tot/Log/#method-stopViolationsReport
*/
type StopViolationsReportResult struct {
	CDTPError error `json:"-"`
}

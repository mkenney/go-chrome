package log

/*
ClearResult represents the result of calls to Log.clear.

https://chromedevtools.github.io/devtools-protocol/tot/Log/#method-clear
*/
type ClearResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
DisableResult represents the result of calls to Log.disable.

https://chromedevtools.github.io/devtools-protocol/tot/Log/#method-disable
*/
type DisableResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
EnableResult represents the result of calls to Log.enable.

https://chromedevtools.github.io/devtools-protocol/tot/Log/#method-enable
*/
type EnableResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
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
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
StopViolationsReportResult represents the result of calls to Log.stopViolationsReport.

https://chromedevtools.github.io/devtools-protocol/tot/Log/#method-stopViolationsReport
*/
type StopViolationsReportResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

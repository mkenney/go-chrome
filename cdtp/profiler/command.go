package profiler

/*
DisableResult represents the result of calls to Profiler.disable.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-disable
*/
type DisableResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
EnableResult represents the result of calls to Profiler.enable.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-enable
*/
type EnableResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
GetBestEffortCoverageResult represents the result of calls to Profiler.getBestEffortCoverage.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-getBestEffortCoverage
*/
type GetBestEffortCoverageResult struct {
	// Coverage data for the current isolate.
	Result []*ScriptCoverage `json:"result"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SetSamplingIntervalParams represents Profiler.setSamplingInterval parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-setSamplingInterval
*/
type SetSamplingIntervalParams struct {
	// New sampling interval in microseconds.
	Interval int `json:"interval"`
}

/*
SetSamplingIntervalResult represents the result of calls to Profiler.setSamplingInterval.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-setSamplingInterval
*/
type SetSamplingIntervalResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
StartResult represents the result of calls to Profiler.start.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-start
*/
type StartResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
StartPreciseCoverageParams represents Profiler.startPreciseCoverage parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-startPreciseCoverage
*/
type StartPreciseCoverageParams struct {
	// Collect accurate call counts beyond simple 'covered' or 'not covered'.
	CallCount bool `json:"callCount"`

	// Collect block-based coverage.
	Detailed bool `json:"detailed"`
}

/*
StartPreciseCoverageResult represents the result of calls to Profiler.startPreciseCoverage.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-startPreciseCoverage
*/
type StartPreciseCoverageResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
StartTypeProfileResult represents the result of calls to Profiler.startTypeProfile.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-startTypeProfile
*/
type StartTypeProfileResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
StopResult represents the result of calls to Profiler.stop.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-stop
*/
type StopResult struct {
	// Recorded profile.
	Profile *Profile `json:"profile"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
StopPreciseCoverageResult represents the result of calls to Profiler.stopPreciseCoverage.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-stopPreciseCoverage
*/
type StopPreciseCoverageResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
StopTypeProfileResult represents the result of calls to Profiler.stopTypeProfile.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-stopTypeProfile
*/
type StopTypeProfileResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
TakePreciseCoverageResult represents the result of calls to Profiler.takePreciseCoverage.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-takePreciseCoverage
*/
type TakePreciseCoverageResult struct {
	// Coverage data for the current isolate.
	Result []*ScriptCoverage `json:"result"`

	// Error information related to executing this method
	Err error `json:"-"`
}

/*
TakeTypeProfileResult represents the result of calls to Profiler.takeTypeProfile.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-takeTypeProfile
*/
type TakeTypeProfileResult struct {
	// Type profile for all scripts since startTypeProfile() was turned on.
	Result []*ScriptTypeProfile `json:"result"`

	// Error information related to executing this method
	Err error `json:"-"`
}

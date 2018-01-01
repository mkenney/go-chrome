package profiler

/*
GetBestEffortCoverageResult represents the result of calls to Profiler.getBestEffortCoverage.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-getBestEffortCoverage
*/
type GetBestEffortCoverageResult struct {
	// Coverage data for the current isolate.
	Result []*ScriptCoverage `json:"result"`
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
StopResult represents the result of calls to Profiler.stop.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-stop
*/
type StopResult struct {
	// Recorded profile.
	Profile *Profile `json:"profile"`
}

/*
TakePreciseCoverageResult represents the result of calls to Profiler.takePreciseCoverage.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-takePreciseCoverage
*/
type TakePreciseCoverageResult struct {
	// Coverage data for the current isolate.
	Result []*ScriptCoverage `json:"result"`
}

/*
TakeTypeProfileResult represents the result of calls to Profiler.takeTypeProfile.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-takeTypeProfile
*/
type TakeTypeProfileResult struct {
	// Type profile for all scripts since startTypeProfile() was turned on.
	Result []*ScriptTypeProfile `json:"result"`
}

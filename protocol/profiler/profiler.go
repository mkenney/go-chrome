/*
Package profiler provides type definitions for use with the Chrome Profiler protocol

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/
*/
package profiler

import (
	"github.com/mkenney/go-chrome/protocol/debugger"
	"github.com/mkenney/go-chrome/protocol/runtime"
)

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

/*
ConsoleProfileFinishedEvent represents Overlay.consoleProfileFinished event data.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#event-consoleProfileFinished
*/
type ConsoleProfileFinishedEvent struct {
	// Profile ID.
	ID string `json:"id"`

	// Location of console.profileEnd().
	Location *debugger.Location `json:"location"`

	// Profile data.
	Profile *Profile `json:"profile"`

	// Profile title passed as an argument to console.profile().
	Title string `json:"title"`
}

/*
ConsoleProfileStartedEvent represents Overlay.consoleProfileStarted event data.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#event-consoleProfileStarted
*/
type ConsoleProfileStartedEvent struct {
	// Profile ID.
	ID string `json:"id"`

	// Location of console.profile().
	Location *debugger.Location `json:"location"`

	// Profile title passed as an argument to console.profile().
	Title string `json:"title"`
}

/*
ProfileNode holds callsite information, execution statistics and child nodes.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#type-ProfileNode
*/
type ProfileNode struct {
	// Unique ID of the node.
	ID int `json:"id"`

	// Function location.
	CallFrame *runtime.CallFrame `json:"callFrame"`

	// Optional. Number of samples where this node was on top of the call stack.
	HitCount int `json:"hitCount,omitempty"`

	// Optional. Child node ids.
	Children []int `json:"children,omitempty"`

	// Optional. The reason of being not optimized. The function may be
	// deoptimized or marked as don't optimize.
	DeoptReason string `json:"deoptReason,omitempty"`

	// Optional. An array of source position ticks.
	PositionTicks []*PositionTickInfo `json:"positionTicks,omitempty"`
}

/*
Profile defines a profile

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#type-Profile
*/
type Profile struct {
	// The list of profile nodes. First item is the root node.
	Nodes []*ProfileNode `json:"nodes"`

	// Profiling start timestamp in microseconds.
	StartTime int `json:"startTime"`

	// Profiling end timestamp in microseconds.
	EndTime int `json:"endTime"`

	// Optional. Ids of samples top nodes.
	Samples []int `json:"samples,omitempty"`

	// Optional. Time intervals between adjacent samples in microseconds. The
	// first delta is relative to the profile startTime.
	TimeDeltas []int `json:"timeDeltas,omitempty"`
}

/*
PositionTickInfo specifies a number of samples attributed to a certain source position.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#type-PositionTickInfo
*/
type PositionTickInfo struct {
	// Source line number (1-based).
	Line int `json:"line"`

	// Number of samples attributed to the source line.
	Ticks int `json:"ticks"`
}

/*
CoverageRange defines coverage data for a source range.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#type-CoverageRange
*/
type CoverageRange struct {
	// JavaScript script source offset for the range start.
	StartOffset int `json:"startOffset"`

	// JavaScript script source offset for the range end.
	EndOffset int `json:"endOffset"`

	// Collected execution count of the source range.
	Count int `json:"count"`
}

/*
FunctionCoverage defines coverage data for a JavaScript function.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#type-FunctionCoverage
*/
type FunctionCoverage struct {
	// JavaScript function name.
	FunctionName string `json:"functionName"`

	// Source ranges inside the function with coverage data.
	Ranges []*CoverageRange `json:"ranges"`

	// Whether coverage data for this function has block granularity.
	IsBlockCoverage bool `json:"isBlockCoverage"`
}

/*
ScriptCoverage defines coverage data for a JavaScript script.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#type-ScriptCoverage
*/
type ScriptCoverage struct {
	// JavaScript script ID.
	ScriptID runtime.ScriptID `json:"scriptId"`

	// JavaScript script name or url.
	URL string `json:"url"`

	// Functions contained in the script that has coverage data.
	Functions []*FunctionCoverage `json:"functions"`
}

/*
TypeObject describes a type collected during runtime. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#type-TypeObject
*/
type TypeObject struct {
	// Name of a type collected with type profiling.
	Name string `json:"name"`
}

/*
TypeProfileEntry is the source offset and types for a parameter or return value. EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#type-TypeProfileEntry
*/
type TypeProfileEntry struct {
	// Source offset of the parameter or end of function for return values.
	Offset int `json:"offset"`

	// The types for this parameter or return value.
	Types []*TypeObject `json:"types"`
}

/*
ScriptTypeProfile is type profile data collected during runtime for a JavaScript script.
EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#type-ScriptTypeProfile
*/
type ScriptTypeProfile struct {
	// JavaScript script ID.
	ScriptID runtime.ScriptID `json:"scriptId"`

	// JavaScript script name or url.
	URL string `json:"url"`

	// Type profile entries for parameters and return values of the functions in
	// the script.
	Entries []*TypeProfileEntry `json:"entries"`
}

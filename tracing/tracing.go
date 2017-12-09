package Tracing

import (
	IO "app/chrome/io"
)

/*
GetCategoriesResult represents the result of calls to Tracing.getCategories.
*/
type GetCategoriesResult struct {
	// A list of supported tracing categories.
	Categories []string `json:"categories"`
}

/*
RecordClockSyncMarkerParams represents Tracing.recordClockSyncMarker parameters.
*/
type RecordClockSyncMarkerParams struct {
	// The ID of this clock sync marker
	SyncID string `json:"syncId"`
}

/*
RequestMemoryDumpResult represents the result of calls to Tracing.requestMemoryDump.
*/
type RequestMemoryDumpResult struct {
	// GUID of the resulting global memory dump.
	DumpGUID string `json:"dumpGuid"`

	// True if the global memory dump succeeded.
	Success bool `json:"success"`
}

/*
StartParams represents Tracing.start parameters.
*/
type StartParams struct {
	// Optional. Category/tag filter.
	Categories string `json:"categories,omitempty"`

	// Optional. Tracing options.
	Options string `json:"options,omitempty"`

	// Optional. If set, the agent will issue bufferUsage events at this interval, specified in
	// milliseconds.
	BufferUsageReportingInterval float64 `json:"bufferUsageReportingInterval,omitempty"`

	// Optional. Whether to report trace events as series of dataCollected events or to save trace
	// to a stream (defaults to `ReportEvents`). Allowed values: ReportEvents, ReturnAsStream.
	TransferMode string `json:"transferMode,omitempty"`

	// Optional. Trace config.
	TraceConfig TraceConfig `json:"traceConfig,omitempty"`
}

/*
BufferUsageEvent represents Overlay.bufferUsage event data.
*/
type BufferUsageEvent struct {
	// Optional. A number in range [0..1] that indicates the used size of event buffer as a fraction
	// of its total size.
	PercentFull float64 `json:"percentFull,omitempty"`

	// Optional. An approximate number of events in the trace log.
	EventCount float64 `json:"eventCount,omitempty"`

	// Optional. A number in range [0..1] that indicates the used size of event buffer as a fraction
	// of its total size.
	Value float64 `json:"value,omitempty"`
}

/*
DataCollectedEvent represents Overlay.dataCollected event data.
*/
type DataCollectedEvent struct {
	Value []map[string]string `json:"value"`
}

/*
TracingCompleteEvent represents Overlay.tracingComplete event data.
*/
type TracingCompleteEvent struct {
	// Optional. A handle of the stream that holds resulting trace data.
	Stream IO.StreamHandle `json:"stream,omitempty"`
}

/*
MemoryDumpConfig is the configuration for memory dump. Used only when "memory-infra" category is
enabled.
*/
type MemoryDumpConfig map[string]string

/*
TraceConfig is the trace configuration
*/
type TraceConfig struct {
	// Optional. Controls how the trace buffer stores data. Allowed values: recordUntilFull,
	// recordContinuously, recordAsMuchAsPossible, echoToConsole.
	RecordMode string `json:"recordMode,omitempty"`

	// Optional. Turns on JavaScript stack sampling.
	EnableSampling bool `json:"enableSampling,omitempty"`

	// Optional. Turns on system tracing.
	EnableSystrace bool `json:"enableSystrace,omitempty"`

	// Optional. Turns on argument filter.
	EnableArgumentFilter bool `json:"enableArgumentFilter,omitempty"`

	// Optional. Included category filters.
	IncludedCategories []string `json:"includedCategories,omitempty"`

	// Optional. Excluded category filters.
	ExcludedCategories []string `json:"excludedCategories,omitempty"`

	// Optional. Configuration to synthesize the delays in tracing.
	SyntheticDelays []string `json:"syntheticDelays,omitempty"`

	// Optional. Configuration for memory dump triggers. Used only when "memory-infra" category is
	// enabled.
	MemoryDumpConfig MemoryDumpConfig `json:"memoryDumpConfig,omitempty"`
}

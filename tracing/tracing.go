package Tracing

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

/*
Package tracing provides type definitions for use with the Chrome Tracing protocol

https://chromedevtools.github.io/devtools-protocol/tot/Tracing/
*/
package tracing

/*
MemoryDumpConfig is the configuration for memory dump. Used only when "memory-infra" category is
enabled.

https://chromedevtools.github.io/devtools-protocol/tot/Tracing/#type-MemoryDumpConfig
*/
type MemoryDumpConfig map[string]string

/*
TraceConfig is the trace configuration

https://chromedevtools.github.io/devtools-protocol/tot/Tracing/#type-TraceConfig
*/
type TraceConfig struct {
	// Optional. Controls how the trace buffer stores data.
	//
	// Allowed values:
	//	- recordUntilFull
	//	- recordContinuously
	//	- recordAsMuchAsPossible
	//	- echoToConsole
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

	// Optional. Configuration for memory dump triggers. Used only when
	// "memory-infra" category is enabled.
	MemoryDumpConfig MemoryDumpConfig `json:"memoryDumpConfig,omitempty"`
}

/*
StreamCompression is the ompression type to use for traces returned via streams.

ALLOWED VALUES
	- none
	- gzip

https://chromedevtools.github.io/devtools-protocol/tot/Tracing/#type-StreamCompression
*/
type StreamCompression string

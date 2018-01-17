/*
Package headlessExperimental provides type definitions for use with the Chrome HeadlessExperimental
protocol

https://chromedevtools.github.io/devtools-protocol/tot/HeadlessExperimental/
*/
package headlessExperimental

/*
ScreenshotParams represents encoding options for a screenshot.

https://chromedevtools.github.io/devtools-protocol/tot/HeadlessExperimental/#type-ScreenshotParams
*/
type ScreenshotParams struct {
	// Optional. Image compression format (defaults to png). Allowed values:
	//	- jpeg
	//	- png
	Format FormatEnum `json:"format,omitempty"`

	// Optional. Compression quality from range [0..100] (jpeg only).
	Quality int `json:"quality,omitempty"`
}

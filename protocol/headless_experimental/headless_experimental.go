package HeadlessExperimental

/*
ScreenshotParams contains encoding options for a screenshot.
*/
type ScreenshotParams struct {
	// Optional. Image compression format (defaults to png). Allowed values: jpeg, png.
	Format string `json:"format,omitempty"`

	// Optional. Compression quality from range [0..100] (jpeg only).
	Quality int `json:"quality,omitempty"`
}

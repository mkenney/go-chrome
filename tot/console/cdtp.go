/*
Package console provides type definitions for use with the Chrome Console protocol

https://chromedevtools.github.io/devtools-protocol/tot/Console/
*/
package console

/*
Message represents a console message.

https://chromedevtools.github.io/devtools-protocol/tot/Console/#type-ConsoleMessage
*/
type Message struct {
	// Message source. Allowed values:
	//	- MessageSource.XML         - "xml"
	//	- MessageSource.Javascript  - "javascript"
	//	- MessageSource.Network     - "network"
	//	- MessageSource.ConsoleAPI  - "console-api"
	//	- MessageSource.Storage     - "storage"
	//	- MessageSource.Appcache    - "appcache"
	//	- MessageSource.Rendering   - "rendering"
	//	- MessageSource.Security    - "security"
	//	- MessageSource.Other       - "other"
	//	- MessageSource.Deprecation - "deprecation"
	//	- MessageSource.Worker      - "worker"
	Source MessageSourceEnum `json:"source"`

	// Message severity. Allowed values:
	//	- MessageLevel.Log     - "log"
	//	- MessageLevel.Warning - "warning"
	//	- MessageLevel.Error   - "error"
	//	- MessageLevel.Debug   - "debug"
	//	- MessageLevel.Info    - "info"
	Level MessageLevelEnum `json:"level"`

	// Message text.
	Text string `json:"text"`

	// Optional. URL of the message origin.
	URL string `json:"url,omitempty"`

	// Optional. Line number in the resource that generated this message
	// (1-based).
	Line int `json:"line,omitempty"`

	// Optional. Column number in the resource that generated this message
	// (1-based).
	Column int `json:"column,omitempty"`
}

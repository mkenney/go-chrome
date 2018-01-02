/*
Package log provides type definitions for use with the Chrome Log protocol

https://chromedevtools.github.io/devtools-protocol/tot/Log/
*/
package log

import (
	"github.com/mkenney/go-chrome/cdtp/network"
	"github.com/mkenney/go-chrome/cdtp/runtime"
)

/*
Entry is a log entry.

https://chromedevtools.github.io/devtools-protocol/tot/Log/#type-LogEntry
*/
type Entry struct {
	// Log entry source.
	//
	// Allowed values: xml
	//	- javascript
	//	- network
	//	- storage
	//	- appcache
	//	- rendering
	//	- security
	//	- deprecation
	//	- worker
	//	- violation
	//	- intervention
	//	- recommendation
	//	- other
	Source string `json:"source"`

	// Log entry severity.
	//
	// Allowed values:
	//	- verbose
	//	- info
	//	- warning
	//	- error
	Level string `json:"level"`

	// Logged text.
	Text string `json:"text"`

	// Timestamp when this entry was added.
	Timestamp runtime.Timestamp `json:"timestamp"`

	// Optional. URL of the resource if known.
	URL string `json:"url,omitempty"`

	// Optional. Line number in the resource.
	LineNumber int `json:"lineNumber,omitempty"`

	// Optional. JavaScript stack trace.
	StackTrace *runtime.StackTrace `json:"stackTrace,omitempty"`

	// Optional. Identifier of the network request associated with this entry.
	NetworkRequestID network.RequestID `json:"networkRequestId,omitempty"`

	// Optional. Identifier of the worker associated with this entry.
	WorkerID string `json:"workerId,omitempty"`

	// Optional. Call arguments.
	Args []*runtime.RemoteObject `json:"args,omitempty"`
}

/*
ViolationSetting is a violation configuration setting.

https://chromedevtools.github.io/devtools-protocol/tot/Log/#type-ViolationSetting
*/
type ViolationSetting struct {
	// Violation type.
	//
	// Allowed values:
	//	- longTask
	//	- longLayout
	//	- blockedEvent
	//	- blockedParser
	//	- discouragedAPIUse
	//	- handler
	//	- recurringHandler
	Name string `json:"name"`

	// Time threshold to trigger upon.
	Threshold int `json:"threshold"`
}

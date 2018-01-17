/*
Package log provides type definitions for use with the Chrome Log protocol

https://chromedevtools.github.io/devtools-protocol/tot/Log/
*/
package log

import (
	"github.com/mkenney/go-chrome/tot/cdtp/network"
	"github.com/mkenney/go-chrome/tot/cdtp/runtime"
)

/*
Entry is a log entry.

https://chromedevtools.github.io/devtools-protocol/tot/Log/#type-LogEntry
*/
type Entry struct {
	// Log entry source. Allowed values:
	//	- Source.XML
	//	- Source.Javascript
	//	- Source.Network
	//	- Source.Storage
	//	- Source.Appcache
	//	- Source.Rendering
	//	- Source.Security
	//	- Source.Deprecation
	//	- Source.Worker
	//	- Source.Violation
	//	- Source.Intervention
	//	- Source.Recommendation
	//	- Source.Other
	Source SourceEnum `json:"source"`

	// Log entry severity. Allowed values:
	//	- Level.Verbose
	//	- Level.Info
	//	- Level.Warning
	//	- Level.Error
	Level LevelEnum `json:"level"`

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
	// Violation type. Allowed values:
	//	- Name.LongTask
	//	- Name.LongLayout
	//	- Name.BlockedEvent
	//	- Name.BlockedParser
	//	- Name.DiscouragedAPIUse
	//	- Name.Handler
	//	- Name.RecurringHandler
	Name NameEnum `json:"name"`

	// Time threshold to trigger upon.
	Threshold int `json:"threshold"`
}

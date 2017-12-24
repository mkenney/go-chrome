/*
Package Log provides type definitions for use with the Chrome Log protocol

https://chromedevtools.github.io/devtools-protocol/tot/Log/
*/
package Log

import (
	"github.com/mkenney/go-chrome/network"
	"github.com/mkenney/go-chrome/runtime"
)

/*
StartViolationsReportParams represents LayerTree.startViolationsReport parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Log/#method-startViolationsReport
*/
type StartViolationsReportParams struct {
	// Configuration for violations.
	Config []ViolationSetting `json:"config"`
}

/*
EntryAddedEvent represents LayerTree.entryAdded event data.

https://chromedevtools.github.io/devtools-protocol/tot/Log/#event-entryAdded
*/
type EntryAddedEvent struct {
	// The entry.
	Entry Entry `json:"entry"`
}

/*
Entry is a log entry.

https://chromedevtools.github.io/devtools-protocol/tot/Log/#type-LogEntry
*/
type Entry struct {
	// Log entry source. Allowed values: xml, javascript, network, storage, appcache, rendering,
	// security, deprecation, worker, violation, intervention, recommendation, other.
	Source string `json:"source"`

	// Log entry severity. Allowed values: verbose, info, warning, error.
	Level string `json:"level"`

	// Logged text.
	Text string `json:"text"`

	// Timestamp when this entry was added.
	Timestamp Runtime.Timestamp `json:"timestamp"`

	// Optional. URL of the resource if known.
	URL string `json:"url,omitempty"`

	// Optional. Line number in the resource.
	LineNumber int `json:"lineNumber,omitempty"`

	// Optional. JavaScript stack trace.
	//
	// This is an instance of Runtime.StackTrace, but that doesn't omitempty correctly so it must be
	// added manually.
	StackTrace interface{} `json:"stackTrace,omitempty"`

	// Optional. Identifier of the network request associated with this entry.
	NetworkRequestID Network.RequestID `json:"networkRequestId,omitempty"`

	// Optional. Identifier of the worker associated with this entry.
	WorkerID string `json:"workerId,omitempty"`

	// Optional. Call arguments.
	Args []Runtime.RemoteObject `json:"args,omitempty"`
}

/*
ViolationSetting is a violation configuration setting.

https://chromedevtools.github.io/devtools-protocol/tot/Log/#type-ViolationSetting
*/
type ViolationSetting struct {
	// Violation type. Allowed values: longTask, longLayout, blockedEvent, blockedParser,
	// discouragedAPIUse, handler, recurringHandler.
	Name string `json:"name"`

	// Time threshold to trigger upon.
	Threshold int `json:"threshold"`
}

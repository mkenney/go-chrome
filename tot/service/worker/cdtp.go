/*
Package worker provides type definitions for use with the Chrome ServiceWorker protocol

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/
*/
package worker

import (
	"github.com/mkenney/go-chrome/tot/target"
)

/*
Registration is a ServiceWorker registration.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#type-ServiceWorkerRegistration
*/
type Registration struct {
	// desc.
	RegistrationID string `json:"registrationId"`

	// desc.
	ScopeURL string `json:"scopeURL"`

	// desc.
	IsDeleted bool `json:"isDeleted"`
}

/*
Version is the ServiceWorker version.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#type-ServiceWorkerVersion
*/
type Version struct {
	// versionId.
	VersionID string `json:"versionId"`

	// registrationId.
	RegistrationID string `json:"registrationId"`

	// scriptURL.
	ScriptURL string `json:"scriptURL"`

	// runningStatus.
	RunningStatus VersionRunningStatusEnum `json:"runningStatus"`

	// status.
	Status VersionStatusEnum `json:"status"`

	// Optional. The Last-Modified header value of the main script.
	ScriptLastModified int64 `json:"scriptLastModified,omitempty"`

	// Optional. The time at which the response headers of the main script were
	// received from the server. For cached script it is the last time the cache
	// entry was validated.
	ScriptResponseTime int64 `json:"scriptResponseTime,omitempty"`

	// Optional. controlledClients.
	ControlledClients []target.ID `json:"controlledClients,omitempty"`

	// Optional. targetId.
	TargetID target.ID `json:"targetId,omitempty"`
}

/*
ErrorMessage is a ServiceWorker error message.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#type-ServiceWorkerErrorMessage
*/
type ErrorMessage struct {
	// desc.
	ErrorMessage string `json:"errorMessage"`

	// desc.
	RegistrationID string `json:"registrationId"`

	// desc.
	VersionID string `json:"versionId"`

	// desc.
	SourceURL string `json:"sourceURL"`

	// desc.
	LineNumber int `json:"lineNumber"`

	// desc.
	ColumnNumber int `json:"columnNumber"`
}

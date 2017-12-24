/*
Package ServiceWorker provides type definitions for use with the Chrome ServiceWorker protocol

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/
*/
package ServiceWorker

import (
	"github.com/mkenney/go-chrome/target"
)

/*
DeliverPushMessageParams represents ServiceWorker.deliverPushMessage parameters.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-deliverPushMessage
*/
type DeliverPushMessageParams struct {
	// Origin.
	Origin string `json:"origin"`

	// Registration ID.
	RegistrationID string `json:"registrationId"`

	// Data.
	Data string `json:"data"`
}

/*
DispatchSyncEventParams represents ServiceWorker.dispatchSyncEvent parameters.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-dispatchSyncEvent
*/
type DispatchSyncEventParams struct {
	// Origin.
	Origin string `json:"origin"`

	// Registration ID.
	RegistrationID string `json:"registrationId"`

	// Tag.
	Tag string `json:"tag"`

	// Last chance.
	LastChance bool `json:"lastChance"`
}

/*
InspectWorkerParams represents ServiceWorker.inspectWorker parameters.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-inspectWorker
*/
type InspectWorkerParams struct {
	// Version ID.
	VersionID string `json:"versionId"`
}

/*
SetForceUpdateOnPageLoadParams represents ServiceWorker.setForceUpdateOnPageLoad parameters.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-setForceUpdateOnPageLoad
*/
type SetForceUpdateOnPageLoadParams struct {
	// Force update on page load.
	ForceUpdateOnPageLoad bool `json:"forceUpdateOnPageLoad"`
}

/*
SkipWaitingParams represents ServiceWorker.skipWaiting parameters.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-skipWaiting
*/
type SkipWaitingParams struct {
	// Scope URL.
	ScopeURL string `json:"scopeURL"`
}

/*
StartWorkerParams represents ServiceWorker.startWorker parameters.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-startWorker
*/
type StartWorkerParams struct {
	// Scope URL.
	ScopeURL string `json:"scopeURL"`
}

/*
StopWorkerParams represents ServiceWorker.stopWorker parameters.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-stopWorker
*/
type StopWorkerParams struct {
	// Version ID.
	VersionID string `json:"versionId"`
}

/*
UnregisterParams represents ServiceWorker.unregister parameters.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-unregister
*/
type UnregisterParams struct {
	// Scope URL.
	ScopeURL string `json:"scopeURL"`
}

/*
UpdateRegistrationParams represents ServiceWorker.updateRegistration parameters.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-updateRegistration
*/
type UpdateRegistrationParams struct {
	// Scope URL.
	ScopeURL string `json:"scopeURL"`
}

/*
WorkerErrorReportedEvent represents ServiceWorker.workerErrorReported event data.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#event-workerErrorReported
*/
type WorkerErrorReportedEvent struct {
	// Error message.
	ErrorMessage ErrorMessage `json:"errorMessage"`
}

/*
WorkerRegistrationUpdatedEvent represents ServiceWorker.workerRegistrationUpdated event data.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#event-workerRegistrationUpdated
*/
type WorkerRegistrationUpdatedEvent struct {
	// Registrations.
	Registrations []Registration `json:"registrations"`
}

/*
WorkerVersionUpdatedEvent represents ServiceWorker.workerVersionUpdated event data.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#event-workerVersionUpdated
*/
type WorkerVersionUpdatedEvent struct {
	// Versions.
	Versions []Version `json:"versions"`
}

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
VersionRunningStatus is the version running status

ALLOWED VALUES
	- stopped
	- starting
	- running
	- stopping

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#type-ServiceWorkerVersionRunningStatus
*/
type VersionRunningStatus string

/*
VersionStatus is the version status

ALLOWED VALUES
	- new
	- installing
	- installed
	- activating
	- activated
	- redundant

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#type-ServiceWorkerVersionStatus
*/
type VersionStatus string

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
	RunningStatus VersionRunningStatus `json:"runningStatus"`

	// status.
	Status VersionStatus `json:"status"`

	// Optional. The Last-Modified header value of the main script.
	ScriptLastModified int `json:"scriptLastModified,omitempty"`

	// Optional. The time at which the response headers of the main script were received from the
	// server. For cached script it is the last time the cache entry was validated.
	ScriptResponseTime int `json:"scriptResponseTime,omitempty"`

	// Optional. controlledClients.
	ControlledClients []*Target.ID `json:"controlledClients,omitempty"`

	// Optional. targetId.
	TargetID Target.ID `json:"targetId,omitempty"`
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

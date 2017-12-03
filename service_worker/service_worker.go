package ServiceWorker

import (
	Target "app/chrome/target"
	"fmt"
)

/*
DeliverPushMessageParams represents ServiceWorker.deliverPushMessage parameters.
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
*/
type InspectWorkerParams struct {
	// Version ID.
	VersionID string `json:"versionId"`
}

/*
SetForceUpdateOnPageLoadParams represents ServiceWorker.setForceUpdateOnPageLoad parameters.
*/
type SetForceUpdateOnPageLoadParams struct {
	// Force update on page load.
	ForceUpdateOnPageLoad bool `json:"forceUpdateOnPageLoad"`
}

/*
SkipWaitingParams represents ServiceWorker.skipWaiting parameters.
*/
type SkipWaitingParams struct {
	// Scope URL.
	ScopeURL string `json:"scopeURL"`
}

/*
StartWorkerParams represents ServiceWorker.startWorker parameters.
*/
type StartWorkerParams struct {
	// Scope URL.
	ScopeURL string `json:"scopeURL"`
}

/*
StopWorkerParams represents ServiceWorker.stopWorker parameters.
*/
type StopWorkerParams struct {
	// Version ID.
	VersionID string `json:"versionId"`
}

/*
UnregisterParams represents ServiceWorker.unregister parameters.
*/
type UnregisterParams struct {
	// Scope URL.
	ScopeURL string `json:"scopeURL"`
}

/*
UpdateRegistrationParams represents ServiceWorker.updateRegistration parameters.
*/
type UpdateRegistrationParams struct {
	// Scope URL.
	ScopeURL string `json:"scopeURL"`
}

/*
WorkerErrorReportedEvent represents ServiceWorker.workerErrorReported event data.
*/
type WorkerErrorReportedEvent struct {
	// Error message.
	ErrorMessage ServiceWorkerErrorMessage `json:"errorMessage"`
}

/*
WorkerRegistrationUpdatedEvent represents ServiceWorker.workerRegistrationUpdated event data.
*/
type WorkerRegistrationUpdatedEvent struct {
	// Registrations.
	Registrations []ServiceWorkerRegistration `json:"registrations"`
}

/*
WorkerVersionUpdatedEvent represents ServiceWorker.workerVersionUpdated event data.
*/
type WorkerVersionUpdatedEvent struct {
	// Versions.
	Versions []ServiceWorkerVersion `json:"versions"`
}

////////////////

/*
ServiceWorkerRegistration is a ServiceWorker registration.
*/
type ServiceWorkerRegistration struct {
	// desc.
	RegistrationID string `json:"registrationId"`

	// desc.
	ScopeURL string `json:"scopeURL"`

	// desc.
	IsDeleted bool `json:"isDeleted"`
}

/*
ServiceWorkerVersionRunningStatus - Allowed values: stopped, starting, running, stopping.
*/
type ServiceWorkerVersionRunningStatus int

const (
	_stopped ServiceWorkerVersionRunningStatus = iota
	_starting
	_running
	_stopping
)

func (a ServiceWorkerVersionRunningStatus) String() string {
	if a == 0 {
		return "stopped"
	}
	if a == 1 {
		return "starting"
	}
	if a == 2 {
		return "running"
	}
	if a == 3 {
		return "stopping"
	}
	panic(fmt.Errorf("Invalid ServiceWorkerVersionRunningStatus %d", a))
}

/*
ServiceWorkerVersionStatus - Allowed values: new, installing, installed, activating, activated,
redundant.
*/
type ServiceWorkerVersionStatus int

const (
	_new ServiceWorkerVersionStatus = iota
	_installing
	_installed
	_activating
	_activated
	_redundant
)

func (a ServiceWorkerVersionStatus) String() string {
	if a == 0 {
		return "new"
	}
	if a == 1 {
		return "installing"
	}
	if a == 2 {
		return "installed"
	}
	if a == 3 {
		return "activating"
	}
	if a == 4 {
		return "activated"
	}
	if a == 5 {
		return "redundant"
	}
	panic(fmt.Errorf("Invalid ServiceWorkerVersionStatus %d", a))
}

/*
ServiceWorkerVersion is the ServiceWorker version.
*/
type ServiceWorkerVersion struct {
	// versionId.
	VersionID string `json:"versionId"`

	// registrationId.
	RegistrationID string `json:"registrationId"`

	// scriptURL.
	ScriptURL string `json:"scriptURL"`

	// runningStatus.
	RunningStatus ServiceWorkerVersionRunningStatus `json:"runningStatus"`

	// status.
	Status ServiceWorkerVersionStatus `json:"status"`

	// Optional. The Last-Modified header value of the main script.
	ScriptLastModified int `json:"scriptLastModified,omitempty"`

	// Optional. The time at which the response headers of the main script were received from the
	// server. For cached script it is the last time the cache entry was validated.
	ScriptResponseTime int `json:"scriptResponseTime,omitempty"`

	// Optional. controlledClients.
	ControlledClients []*Target.TargetID `json:"controlledClients,omitempty"`

	// Optional. targetId.
	TargetID Target.TargetID `json:"targetId,omitempty"`
}

/*
ServiceWorkerErrorMessage is a ServiceWorker error message.
*/
type ServiceWorkerErrorMessage struct {
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

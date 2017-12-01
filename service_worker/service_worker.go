package ServiceWorker

import (
	Target "app/chrome/target"
	"fmt"
)

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
ServiceWorkerVersionRunningStatus
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
ServiceWorkerVersionStatus
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

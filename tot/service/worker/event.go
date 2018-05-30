package worker

/*
ErrorReportedEvent represents ServiceWorker.workerErrorReported event data.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#event-workerErrorReported
*/
type ErrorReportedEvent struct {
	// Error message.
	ErrorMessage *ErrorMessage `json:"errorMessage"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
RegistrationUpdatedEvent represents ServiceWorker.workerRegistrationUpdated event data.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#event-workerRegistrationUpdated
*/
type RegistrationUpdatedEvent struct {
	// Registrations.
	Registrations []*Registration `json:"registrations"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
VersionUpdatedEvent represents ServiceWorker.workerVersionUpdated event data.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#event-workerVersionUpdated
*/
type VersionUpdatedEvent struct {
	// Versions.
	Versions []*Version `json:"versions"`

	// Error information related to this event
	Err error `json:"-"`
}

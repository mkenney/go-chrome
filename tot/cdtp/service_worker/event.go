package serviceWorker

/*
WorkerErrorReportedEvent represents ServiceWorker.workerErrorReported event data.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#event-workerErrorReported
*/
type WorkerErrorReportedEvent struct {
	// Error message.
	ErrorMessage *ErrorMessage `json:"errorMessage"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
WorkerRegistrationUpdatedEvent represents ServiceWorker.workerRegistrationUpdated event data.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#event-workerRegistrationUpdated
*/
type WorkerRegistrationUpdatedEvent struct {
	// Registrations.
	Registrations []*Registration `json:"registrations"`

	// Error information related to this event
	Err error `json:"-"`
}

/*
WorkerVersionUpdatedEvent represents ServiceWorker.workerVersionUpdated event data.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#event-workerVersionUpdated
*/
type WorkerVersionUpdatedEvent struct {
	// Versions.
	Versions []*Version `json:"versions"`

	// Error information related to this event
	Err error `json:"-"`
}

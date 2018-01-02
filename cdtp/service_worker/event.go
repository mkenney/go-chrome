package serviceWorker

/*
WorkerErrorReportedEvent represents ServiceWorker.workerErrorReported event data.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#event-workerErrorReported
*/
type WorkerErrorReportedEvent struct {
	// Error message.
	ErrorMessage *ErrorMessage `json:"errorMessage"`
}

/*
WorkerRegistrationUpdatedEvent represents ServiceWorker.workerRegistrationUpdated event data.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#event-workerRegistrationUpdated
*/
type WorkerRegistrationUpdatedEvent struct {
	// Registrations.
	Registrations []*Registration `json:"registrations"`
}

/*
WorkerVersionUpdatedEvent represents ServiceWorker.workerVersionUpdated event data.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#event-workerVersionUpdated
*/
type WorkerVersionUpdatedEvent struct {
	// Versions.
	Versions []*Version `json:"versions"`
}

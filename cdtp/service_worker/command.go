package serviceWorker

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

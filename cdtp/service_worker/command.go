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
DeliverPushMessageResult represents the result of calls to ServiceWorker.deliverPushMessage.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-deliverPushMessage
*/
type DeliverPushMessageResult struct {
	CDTPError error `json:"-"`
}

/*
DisableResult represents the result of calls to ServiceWorker.disable.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-disable
*/
type DisableResult struct {
	CDTPError error `json:"-"`
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
DispatchSyncEventResult represents the result of calls to ServiceWorker.dispatchSyncEvent.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-dispatchSyncEvent
*/
type DispatchSyncEventResult struct {
	CDTPError error `json:"-"`
}

/*
EnableResult represents the result of calls to ServiceWorker.enable.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-enable
*/
type EnableResult struct {
	CDTPError error `json:"-"`
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
InspectWorkerResult represents the result of calls to ServiceWorker.inspectWorker.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-inspectWorker
*/
type InspectWorkerResult struct {
	CDTPError error `json:"-"`
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
SetForceUpdateOnPageLoadResult represents the result of calls to
ServiceWorker.setForceUpdateOnPageLoad.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-setForceUpdateOnPageLoad
*/
type SetForceUpdateOnPageLoadResult struct {
	CDTPError error `json:"-"`
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
SkipWaitingResult represents the result of calls to ServiceWorker.skipWaiting.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-skipWaiting
*/
type SkipWaitingResult struct {
	CDTPError error `json:"-"`
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
StartWorkerResult represents the result of calls to ServiceWorker.startWorker.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-startWorker
*/
type StartWorkerResult struct {
	CDTPError error `json:"-"`
}

/*
StopAllWorkersResult represents the result of calls to ServiceWorker.stopAllWorkers.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-stopAllWorkers
*/
type StopAllWorkersResult struct {
	CDTPError error `json:"-"`
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
StopWorkerResult represents the result of calls to ServiceWorker.stopWorker.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-stopWorker
*/
type StopWorkerResult struct {
	CDTPError error `json:"-"`
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
UnregisterResult represents the result of calls to ServiceWorker.unregister.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-unregister
*/
type UnregisterResult struct {
	CDTPError error `json:"-"`
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
UpdateRegistrationResult represents the result of calls to ServiceWorker.updateRegistration.

https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-updateRegistration
*/
type UpdateRegistrationResult struct {
	CDTPError error `json:"-"`
}

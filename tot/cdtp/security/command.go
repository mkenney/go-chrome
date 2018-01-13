package security

/*
DisableResult represents the result of calls to Security.disable.

https://chromedevtools.github.io/devtools-protocol/tot/Security/#method-disable
*/
type DisableResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
EnableResult represents the result of calls to Security.enable.

https://chromedevtools.github.io/devtools-protocol/tot/Security/#method-enable
*/
type EnableResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
HandleCertificateErrorParams represents Security.handleCertificateError parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Security/#method-handleCertificateError
*/
type HandleCertificateErrorParams struct {
	// The ID of the event.
	EventID int `json:"eventId"`

	// The action to take on the certificate error.
	Action CertificateErrorAction `json:"action"`
}

/*
HandleCertificateErrorResult represents the result of calls to Security.handleCertificateError.

https://chromedevtools.github.io/devtools-protocol/tot/Security/#method-handleCertificateError
*/
type HandleCertificateErrorResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SetIgnoreCertificateErrorsParams represents Security.setIgnoreCertificateErrors parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Security/#method-setIgnoreCertificateErrors
*/
type SetIgnoreCertificateErrorsParams struct {
	// If true, all certificate errors will be ignored.
	Ignore bool `json:"ignore"`
}

/*
SetIgnoreCertificateErrorsResult represents the result of calls to
Security.setIgnoreCertificateErrors.

https://chromedevtools.github.io/devtools-protocol/tot/Security/#method-setIgnoreCertificateErrors
*/
type SetIgnoreCertificateErrorsResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

/*
SetOverrideCertificateErrorsParams represents Security.setOverrideCertificateErrors parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Security/#method-setOverrideCertificateErrors
*/
type SetOverrideCertificateErrorsParams struct {
	// If true, certificate errors will be overridden.
	Override bool `json:"override"`
}

/*
SetOverrideCertificateErrorsResult represents the result of calls to
Security.setOverrideCertificateErrors.

https://chromedevtools.github.io/devtools-protocol/tot/Security/#method-setOverrideCertificateErrors
*/
type SetOverrideCertificateErrorsResult struct {
	// Error information related to executing this method
	Err error `json:"-"`
}

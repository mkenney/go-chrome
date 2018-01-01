package security

/*
SetIgnoreCertificateErrorsParams represents Security.setIgnoreCertificateErrors parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Security/#method-setIgnoreCertificateErrors
*/
type SetIgnoreCertificateErrorsParams struct {
	// If true, all certificate errors will be ignored.
	Ignore bool `json:"ignore"`
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
SetOverrideCertificateErrorsParams represents Security.setOverrideCertificateErrors parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Security/#method-setOverrideCertificateErrors
*/
type SetOverrideCertificateErrorsParams struct {
	// If true, certificate errors will be overridden.
	Override bool `json:"override"`
}

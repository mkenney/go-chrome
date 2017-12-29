/*
Package security provides type definitions for use with the Chrome Security protocol

https://chromedevtools.github.io/devtools-protocol/tot/Security/
*/
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

/*
CertificateErrorEvent represents Security.certificateError event data.

https://chromedevtools.github.io/devtools-protocol/tot/Security/#event-certificateError
*/
type CertificateErrorEvent struct {
	// The ID of the event.
	EventID int `json:"eventId"`

	// The type of the error.
	ErrorType string `json:"errorType"`

	// The url that was requested.
	RequestURL string `json:"requestURL"`
}

/*
StateChangedEvent represents Security.securityStateChanged event data.

https://chromedevtools.github.io/devtools-protocol/tot/Security/#event-securityStateChanged
*/
type StateChangedEvent struct {
	// Security state.
	State State `json:"securityState"`

	// True if the page was loaded over cryptographic transport such as HTTPS.
	SchemeIsCryptographic bool `json:"schemeIsCryptographic"`

	// List of explanations for the security state. If the overall security
	// state is `insecure` or `warning`, at least one corresponding explanation
	// should be included.
	Explanations []*StateExplanation `json:"explanations"`

	// Information about insecure content on the page.
	InsecureContentStatus *InsecureContentStatus `json:"insecureContentStatus"`

	// Optional. Overrides user-visible description of the state.
	Summary string `json:"summary,omitempty"`
}

/*
CertificateID is an internal certificate ID value.

https://chromedevtools.github.io/devtools-protocol/tot/Security/#type-CertificateId
*/
type CertificateID int

/*
MixedContentType is a description of mixed content (HTTP resources on HTTPS pages), as defined by
https://www.w3.org/TR/mixed-content/#categories

ALLOWED VALUES
	- blockable
	- optionally-blockable
	- none

https://chromedevtools.github.io/devtools-protocol/tot/Security/#type-MixedContentType
*/
type MixedContentType string

/*
State is the security level of a page or resource.

ALLOWED VALUES
	- unknown
	- neutral
	- insecure
	- secure
	- info

https://chromedevtools.github.io/devtools-protocol/tot/Security/#type-SecurityState
*/
type State string

/*
StateExplanation is an explanation of a factor contributing to the security state.

https://chromedevtools.github.io/devtools-protocol/tot/Security/#type-SecurityStateExplanation
*/
type StateExplanation struct {
	// Security state representing the severity of the factor being explained.
	State State `json:"securityState"`

	// Short phrase describing the type of factor.
	Summary string `json:"summary"`

	// Full text explanation of the factor.
	Description string `json:"description"`

	// The type of mixed content described by the explanation.
	MixedContentType MixedContentType `json:"mixedContentType"`

	// Page certificate.
	Certificate []string `json:"certificate"`
}

/*
InsecureContentStatus describes information about insecure content on the page.

https://chromedevtools.github.io/devtools-protocol/tot/Security/#type-InsecureContentStatus
*/
type InsecureContentStatus struct {
	// True if the page was loaded over HTTPS and ran mixed (HTTP) content such
	// as scripts.
	RanMixedContent bool `json:"ranMixedContent"`

	// True if the page was loaded over HTTPS and displayed mixed (HTTP) content
	// such as images.
	DisplayedMixedContent bool `json:"displayedMixedContent"`

	// True if the page was loaded over HTTPS and contained a form targeting an
	// insecure url.
	ContainedMixedForm bool `json:"containedMixedForm"`

	// True if the page was loaded over HTTPS without certificate errors, and
	// ran content such as scripts that were loaded with certificate errors.
	RanContentWithCertErrors bool `json:"ranContentWithCertErrors"`

	// True if the page was loaded over HTTPS without certificate errors, and
	// displayed content such as images that were loaded with certificate errors.
	DisplayedContentWithCertErrors bool `json:"displayedContentWithCertErrors"`

	// Security state representing a page that ran insecure content.
	RanInsecureContentStyle State `json:"ranInsecureContentStyle"`

	// Security state representing a page that displayed insecure content.
	DisplayedInsecureContentStyle State `json:"displayedInsecureContentStyle"`
}

/*
CertificateErrorAction describes the action to take when a certificate error occurs. 'continue' will
continue processing the request and 'cancel' will cancel the request.

ALLOWED VALUES
	- continue
	- cancel

https://chromedevtools.github.io/devtools-protocol/tot/Security/#type-CertificateErrorAction
*/
type CertificateErrorAction string

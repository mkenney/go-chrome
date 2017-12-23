package Security

import (
	"fmt"
)

/*
CertificateID is an internal certificate ID value.
*/
type CertificateID int

/*
HandleCertificateErrorParams represents Security.handleCertificateError parameters.
*/
type HandleCertificateErrorParams struct {
	// The ID of the event.
	EventID int `json:"eventId"`

	// The action to take on the certificate error.
	Action CertificateErrorAction `json:"action"`
}

/*
SetOverrideCertificateErrorsParams represents Security.setOverrideCertificateErrors parameters.
*/
type SetOverrideCertificateErrorsParams struct {
	// If true, certificate errors will be overridden.
	Override bool `json:"override"`
}

/*
CertificateErrorEvent represents Security.certificateError event data.
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
SecurityStateChangedEvent represents Security.securityStateChanged event data.
*/
type SecurityStateChangedEvent struct {
	// Security state.
	SecurityState SecurityState `json:"securityState"`

	// True if the page was loaded over cryptographic transport such as HTTPS.
	SchemeIsCryptographic bool `json:"schemeIsCryptographic"`

	// List of explanations for the security state. If the overall security state is `insecure` or
	// `warning`, at least one corresponding explanation should be included.
	Explanations []SecurityStateExplanation `json:"explanations"`

	// Information about insecure content on the page.
	InsecureContentStatus InsecureContentStatus `json:"insecureContentStatus"`

	// Optional. Overrides user-visible description of the state.
	Summary string `json:"summary,omitempty"`
}

/*
MixedContentType is a description of mixed content (HTTP resources on HTTPS pages), as defined by
https://www.w3.org/TR/mixed-content/#categories
*/
type MixedContentType string

func (s MixedContentType) String() string {
	str := string(s)
	if str == "blockable" ||
		str == "optionally-blockable" ||
		str == "none" {
		return str
	}
	panic(fmt.Errorf("Invalid MixedContentType '%s'", str))
}

/*
SecurityState is the security level of a page or resource.
*/
type SecurityState string

func (s SecurityState) String() string {
	str := string(s)
	if str == "unknown" ||
		str == "neutral" ||
		str == "insecure" ||
		str == "secure" ||
		str == "info" {
		return str
	}
	panic(fmt.Errorf("Invalid SecurityState '%s'", str))
}

/*
SecurityStateExplanation is an explanation of a factor contributing to the security state.
*/
type SecurityStateExplanation struct {
	// Security state representing the severity of the factor being explained.
	SecurityState SecurityState `json:"securityState"`

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
*/
type InsecureContentStatus struct {
	// True if the page was loaded over HTTPS and ran mixed (HTTP) content such as scripts.
	RanMixedContent bool `json:"ranMixedContent"`

	// True if the page was loaded over HTTPS and displayed mixed (HTTP) content such as images.
	DisplayedMixedContent bool `json:"displayedMixedContent"`

	// True if the page was loaded over HTTPS and contained a form targeting an insecure url.
	ContainedMixedForm bool `json:"containedMixedForm"`

	// True if the page was loaded over HTTPS without certificate errors, and ran content such as
	// scripts that were loaded with certificate errors.
	RanContentWithCertErrors bool `json:"ranContentWithCertErrors"`

	// True if the page was loaded over HTTPS without certificate errors, and displayed content such
	// as images that were loaded with certificate errors.
	DisplayedContentWithCertErrors bool `json:"displayedContentWithCertErrors"`

	// Security state representing a page that ran insecure content.
	RanInsecureContentStyle SecurityState `json:"ranInsecureContentStyle"`

	// Security state representing a page that displayed insecure content.
	DisplayedInsecureContentStyle SecurityState `json:"displayedInsecureContentStyle"`
}

/*
CertificateErrorAction describes the action to take when a certificate error occurs. 'continue' will
continue processing the request and 'cancel' will cancel the request.
*/
type CertificateErrorAction string

func (s CertificateErrorAction) String() string {
	str := string(s)
	if str == "continue" ||
		str == "cancel" {
		return str
	}
	panic(fmt.Errorf("Invalid CertificateErrorAction '%s'", str))
}

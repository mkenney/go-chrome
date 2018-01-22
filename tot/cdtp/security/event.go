package security

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

	// Error information related to this event
	Err error `json:"-"`
}

/*
StateChangedEvent represents Security.securityStateChanged event data.

https://chromedevtools.github.io/devtools-protocol/tot/Security/#event-securityStateChanged
*/
type StateChangedEvent struct {
	// Security state. Allowed values:
	//	- State.Unknown
	//	- State.Neutral
	//	- State.Insecure
	//	- State.Secure
	//	- State.Info
	State StateEnum `json:"securityState"`

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

	// Error information related to this event
	Err error `json:"-"`
}

/*
Package security provides type definitions for use with the Chrome Security protocol

https://chromedevtools.github.io/devtools-protocol/tot/Security/
*/
package security

/*
CertificateID is an internal certificate ID value.

https://chromedevtools.github.io/devtools-protocol/tot/Security/#type-CertificateId
*/
type CertificateID int

/*
StateExplanation is an explanation of a factor contributing to the security state.

https://chromedevtools.github.io/devtools-protocol/tot/Security/#type-SecurityStateExplanation
*/
type StateExplanation struct {
	// Security state representing the severity of the factor being explained.
	// Allowed values:
	//	- State.Unknown
	//	- State.Neutral
	//	- State.Insecure
	//	- State.Secure
	//	- State.Info
	State StateEnum `json:"securityState"`

	// Short phrase describing the type of factor.
	Summary string `json:"summary"`

	// Full text explanation of the factor.
	Description string `json:"description"`

	// The type of mixed content described by the explanation.
	MixedContentType MixedContentTypeEnum `json:"mixedContentType"`

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

	// Security state representing a page that ran insecure content. Allowed
	// values:
	//	- State.Unknown
	//	- State.Neutral
	//	- State.Insecure
	//	- State.Secure
	//	- State.Info
	RanInsecureContentStyle StateEnum `json:"ranInsecureContentStyle"`

	// Security state representing a page that displayed insecure content.
	// Allowed values:
	//	- State.Unknown
	//	- State.Neutral
	//	- State.Insecure
	//	- State.Secure
	//	- State.Info
	DisplayedInsecureContentStyle StateEnum `json:"displayedInsecureContentStyle"`
}

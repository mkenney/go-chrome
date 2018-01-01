/*
Package target provides type definitions for use with the Chrome Target protocol

https://chromedevtools.github.io/devtools-protocol/tot/Target/
*/
package target

/*
ID is the target ID.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#type-TargetID
*/
type ID string

/*
SessionID is a unique identifier of attached debugging session.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#type-SessionID
*/
type SessionID string

/*
BrowserContextID is EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Target/#type-BrowserContextID
*/
type BrowserContextID string

/*
Info holds the target info.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#type-TargetInfo
*/
type Info struct {
	// desc.
	ID ID `json:"targetId"`

	// desc.
	Type string `json:"type"`

	// desc.
	Title string `json:"title"`

	// desc.
	URL string `json:"url"`

	// Whether the target has an attached client.
	Attached bool `json:"attached"`

	// Optional. Opener target Id.
	OpenerID ID `json:"openerId,omitempty"`
}

/*
RemoteLocation is EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Target/#type-RemoteLocation
*/
type RemoteLocation struct {
	// desc.
	Host string `json:"host"`

	// desc.
	Port int `json:"port"`
}

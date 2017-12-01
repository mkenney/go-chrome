package Target

/*
TargetID is the target ID.
*/
type TargetID string

/*
SessionID is a unique identifier of attached debugging session.
*/
type SessionID string

/*
BrowserContextID is EXPERIMENTAL
*/
type BrowserContextID string

/*
TargetInfo holds the target info
*/
type TargetInfo struct {
	// desc.
	TargetID TargetID `json:"targetId"`

	// desc.
	Type string `json:"type"`

	// desc.
	Title string `json:"title"`

	// desc.
	URL string `json:"url"`

	// Whether the target has an attached client.
	Attached bool `json:"attached"`

	// Optional. Opener target Id.
	OpenerID TargetID `json:"openerId,omitempty"`
}

/*
RemoteLocation is EXPERIMENTAL
*/
type RemoteLocation struct {
	// desc.
	Host string `json:"host"`

	// desc.
	Port int `json:"port"`
}

package Console

/*
MessageAddedEvent represents Console.messageAdded event data.
*/
type MessageAddedEvent struct {
	// Console message that has been added.
	Message ConsoleMessage `json:"message"`
}

/*
ConsoleMessage represents a console message.
*/
type ConsoleMessage struct {
	// Message source. Allowed values: xml, javascript, network, console-api, storage, appcache,
	// rendering, security, other, deprecation, worker.
	Source string `json:"source"`

	// Message severity. Allowed values: log, warning, error, debug, info.
	Level string `json:"level"`

	// Message text.
	Text string `json:"text"`

	// Optional. URL of the message origin.
	URL string `json:"url,omitempty"`

	// Optional. Line number in the resource that generated this message (1-based).
	Line int `json:"line,omitempty"`

	// Optional. Column number in the resource that generated this message (1-based).
	Column int `json:"column,omitempty"`
}

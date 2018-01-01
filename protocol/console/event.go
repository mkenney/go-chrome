package console

/*
MessageAddedEvent represents Console.messageAdded event data.

https://chromedevtools.github.io/devtools-protocol/tot/Console/#event-messageAdded
*/
type MessageAddedEvent struct {
	// Console message that has been added.
	Message *Message `json:"message"`
}

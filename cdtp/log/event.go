package log

/*
EntryAddedEvent represents LayerTree.entryAdded event data.

https://chromedevtools.github.io/devtools-protocol/tot/Log/#event-entryAdded
*/
type EntryAddedEvent struct {
	// The entry.
	Entry *Entry `json:"entry"`
}

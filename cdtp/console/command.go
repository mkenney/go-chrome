/*
Package console provides type definitions for use with the Chrome Console protocol

https://chromedevtools.github.io/devtools-protocol/tot/Console/
*/
package console

/*
ClearMessagesResult represents the result of calls to Console.ClearMessages.

https://chromedevtools.github.io/devtools-protocol/tot/Console/#method-clearMessages
*/
type ClearMessagesResult struct {
	CDTPError error `json:"-"`
}

/*
DisableResult represents the result of calls to Console.disable.

https://chromedevtools.github.io/devtools-protocol/tot/Console/#method-disable
*/
type DisableResult struct {
	CDTPError error `json:"-"`
}

/*
EnableResult represents the result of calls to Console.enable.

https://chromedevtools.github.io/devtools-protocol/tot/Console/#method-enable
*/
type EnableResult struct {
	CDTPError error `json:"-"`
}

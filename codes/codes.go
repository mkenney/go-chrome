package codes

import (
	errs "github.com/bdlm/errors"
	std "github.com/bdlm/std/error"
)

const (
	// Unspecified - 1000: The error code was unspecified
	Unspecified std.Code = iota + 1000
	// Unknown - 1001: An unspecified error occurred.
	Unknown
	// MockErr - 1002: An error occurred in a test mock.
	MockErr
)

////////////////////////////////////////////////////////////////////////////
// Chrome errors
////////////////////////////////////////////////////////////////////////////
const (
	// ChromeCannotOpenStderr - 2000: Cannot open error output file.
	ChromeCannotOpenStderr std.Code = iota + 2000
	// ChromeCannotOpenStdout - 2001: Cannot open standard output file.
	ChromeCannotOpenStdout
	// ChromeExitTimeout - 2002: Error waiting for process exit.
	ChromeExitTimeout
	// ChromeInvalidWorkdir - 2003: Cannot create working directory.
	ChromeInvalidWorkdir
	// ChromeQueryFailed - 2004: Chrome.Query() returned an error.
	ChromeQueryFailed
	// ChromeSigintFailed - 2005: Chromium process interrupt failed.
	ChromeSigintFailed
	// ChromeStartTimeout - 2006: Chromium took too long to start.
	ChromeStartTimeout
	// ChromeTabNotFound - 2007: Chromium tab not found.
	ChromeTabNotFound
	// ChromeVersionQueryFailed - 2008: Chromium version query failed.
	ChromeVersionQueryFailed
)

////////////////////////////////////////////////////////////////////////////
// Flag errors
////////////////////////////////////////////////////////////////////////////
const (
	// FlagDoesNotExist - 3000: The specified argument does not exist.
	FlagDoesNotExist std.Code = iota + 3000
	// FlagTypeInvalid - 3001: Invalid data type for the specified argument.
	FlagTypeInvalid
)

////////////////////////////////////////////////////////////////////////////
// Tab errors
////////////////////////////////////////////////////////////////////////////
const (

	// TabQueryFailed - 4000: The new tab query failed.
	TabQueryFailed std.Code = iota + 4000
	// TabURLInvalid - 4001: Invalid URL passed to NewTab.
	TabURLInvalid
	// TabWebsocketURLInvalid - 4002: Invalid websocket URL.
	TabWebsocketURLInvalid
)

////////////////////////////////////////////////////////////////////////////
// Socket errors
////////////////////////////////////////////////////////////////////////////
const (
	// SocketCloseFailed - 5000: A failure occurred while closing a websocket.
	SocketCloseFailed std.Code = iota + 5000
	// SocketReadFailed - 5001: A failure occurred while reading from a websocket.
	SocketReadFailed
	// SocketCmdHandlerNotFound - 5002: Command handler not found.
	SocketCmdHandlerNotFound
	// SocketReadFailed - 5002: Attempted to add a duplicate handler for event.
	SocketDuplicateEventHandler
	// SocketEventHandlerNotFound - 5003: No event listeners found for an event.
	SocketEventHandlerNotFound
	// SocketConnectFailed - 5003: Connect() failed while creating socket.
	SocketConnectFailed
	// SocketNotConnected - 5003: Not connected.
	SocketNotConnected
	// SocketWriteFailed - 5003: Socket write failed.
	SocketWriteFailed
	// SocketPanic - 5003: A panic occurred while reading from a websocket.
	SocketPanic
)

////////////////////////////////////////////////////////////////////////////
// Webocket errors
////////////////////////////////////////////////////////////////////////////
const (
	// WebsocketConnectFailed - 6000: Websocket connection failed.
	WebsocketConnectFailed std.Code = iota + 6000
	// WebsocketNotConnected - 6001: Websocket not connected.
	WebsocketNotConnected
	// WebsocketPanic - 5002: A panic occurred while reading from a websocket.
	WebsocketPanic
)

func init() {
	errs.Codes[Unspecified] = errs.ErrCode{Int: "The error code was unspecified", Ext: "An unknown error occurred", HTTP: 500}
	errs.Codes[Unknown] = errs.ErrCode{Int: "An unspecified error occurred", Ext: "An unknown error occurred", HTTP: 500}
	errs.Codes[MockErr] = errs.ErrCode{Int: "An error occurred in a test mock", Ext: "An unknown error occurred", HTTP: 500}

	errs.Codes[ChromeCannotOpenStderr] = errs.ErrCode{Int: "Cannot open error output file", Ext: "An unknown error occurred", HTTP: 500}
	errs.Codes[ChromeCannotOpenStdout] = errs.ErrCode{Int: "Cannot open standard output file", Ext: "An unknown error occurred", HTTP: 500}
	errs.Codes[ChromeExitTimeout] = errs.ErrCode{Int: "Error waiting for process exit", Ext: "An unknown error occurred", HTTP: 500}
	errs.Codes[ChromeInvalidWorkdir] = errs.ErrCode{Int: "Cannot create working directory", Ext: "An unknown error occurred", HTTP: 500}
	errs.Codes[ChromeQueryFailed] = errs.ErrCode{Int: "Chrome.Query() returned an error", Ext: "An unknown error occurred", HTTP: 500}
	errs.Codes[ChromeSigintFailed] = errs.ErrCode{Int: "Chromium process interrupt failed", Ext: "An unknown error occurred", HTTP: 500}
	errs.Codes[ChromeStartTimeout] = errs.ErrCode{Int: "Chromium took too long to start", Ext: "An unknown error occurred", HTTP: 500}
	errs.Codes[ChromeTabNotFound] = errs.ErrCode{Int: "Chromium tab not found", Ext: "An unknown error occurred", HTTP: 500}
	errs.Codes[ChromeVersionQueryFailed] = errs.ErrCode{Int: "Chromium version query failed", Ext: "An unknown error occurred", HTTP: 500}

	errs.Codes[FlagDoesNotExist] = errs.ErrCode{Int: "The specified argument does not exist", Ext: "An unknown error occurred", HTTP: 500}
	errs.Codes[FlagTypeInvalid] = errs.ErrCode{Int: "Invalid data type for the specified argument", Ext: "An unknown error occurred", HTTP: 500}

	errs.Codes[TabQueryFailed] = errs.ErrCode{Int: "The new tab query failed", Ext: "An unknown error occurred", HTTP: 500}
	errs.Codes[TabURLInvalid] = errs.ErrCode{Int: "Invalid URL passed to NewTab", Ext: "An unknown error occurred", HTTP: 500}
	errs.Codes[TabWebsocketURLInvalid] = errs.ErrCode{Int: "Invalid websocket URL", Ext: "An unknown error occurred", HTTP: 500}

	errs.Codes[SocketCloseFailed] = errs.ErrCode{Int: "A failure occurred while closing a websocket", Ext: "An unknown error occurred", HTTP: 500}
	errs.Codes[SocketReadFailed] = errs.ErrCode{Int: "A failure occurred while reading from a websocket", Ext: "An unknown error occurred", HTTP: 500}
	errs.Codes[SocketPanic] = errs.ErrCode{Int: "A panic occurred while reading from a websocket", Ext: "An unknown error occurred", HTTP: 500}

	errs.Codes[WebsocketConnectFailed] = errs.ErrCode{Int: "Websocket connection failed", Ext: "An unknown error occurred", HTTP: 500}
	errs.Codes[WebsocketNotConnected] = errs.ErrCode{Int: "Websocket not connected", Ext: "An unknown error occurred", HTTP: 500}
	errs.Codes[WebsocketPanic] = errs.ErrCode{Int: "A panic occurred while reading from a websocket", Ext: "An unknown error occurred", HTTP: 500}
}

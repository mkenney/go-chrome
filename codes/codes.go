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
)
const (
	////////////////////////////////////////////////////////////////////////////
	// Socket errors
	////////////////////////////////////////////////////////////////////////////

	// SocketPanic - 2000: A panic occurred while reading from a websocket
	SocketPanic std.Code = iota + 2000
	// SocketReadFailed - 2001: A failure occurred while reading from a websocket
	SocketReadFailed
	// SocketCloseFailed - 2002: A failure occurred while closing a websocket
	SocketCloseFailed
)

func init() {
	errs.Codes[Unspecified] = errs.ErrCode{Ext: "An unknown error occurred", Int: "An unknown error occurred", HTTP: 500}
	errs.Codes[Unknown] = errs.ErrCode{Ext: "An unknown error occurred", Int: "An unspecified error occurred", HTTP: 500}
	errs.Codes[SocketPanic] = errs.ErrCode{Ext: "An unknown error occurred", Int: "A panic occurred while reading from a websocket", HTTP: 500}
	errs.Codes[SocketReadFailed] = errs.ErrCode{Ext: "An unknown error occurred", Int: "A failure occurred while reading from a websocket", HTTP: 500}
	errs.Codes[SocketCloseFailed] = errs.ErrCode{Ext: "An unknown error occurred", Int: "A failure occurred while closing a websocket", HTTP: 500}
}

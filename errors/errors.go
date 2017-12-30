package errors

import (
	"fmt"
	"path"
	"runtime"
	"strconv"
)

/*
SocketReadFailed is used when an attempt to read from a socket connection fails.
*/
type SocketReadFailed struct{ Type }

/*
SocketWriteFailed is used when an attempt to write to a socket connection fails.
*/
type SocketWriteFailed struct{ Type }

/*
SocketErrorResponse is used when a socket connection responds with error data.
*/
type SocketErrorResponse struct{ Type }

/*
Type defines the data type for errors.
*/
type Type struct {
	Caller map[string]string
	Err    error
	Msg    string
}

/*
Error implements errors.
*/
func (err Type) Error() string {
	msg := fmt.Sprintf(
		"[%s] %s:%s - %s",
		err.Caller["pc"],
		err.Caller["file"],
		err.Caller["line"],
		err.Msg,
	)
	if nil != err.Err {
		msg = fmt.Sprintf("%s; %s", msg, err.Err.Error())
	}
	return msg
}

/*
GetCaller returns caller information.
*/
func GetCaller() map[string]string {
	caller := make(map[string]string)
	a := 1
	for {
		if pc, file, line, ok := runtime.Caller(a); ok {
			caller["file"] = path.Base(file)
			caller["line"] = strconv.Itoa(line)
			caller["pc"] = runtime.FuncForPC(pc).Name()
			break
		} else {
			a++
		}

		if a > 100 {
			break
		}
	}
	return caller
}

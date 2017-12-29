package error

import (
	"fmt"

	"github.com/prometheus/common/log"
)

/*
Error codes
*/
const (
	CodeUnspecifiedError        = 0
	CodeInvalidLevel            = 1
	CodeErrorOpeningOutputFile  = 2
	CodeCannotCreateWorkDir     = 3
	CodeOutOfRange              = 4
	CodeCouldNotConnectToSocket = 5
	CodeJSONError               = 6

	LevelDebug = 1
	LevelInfo  = 2
	LevelWarn  = 3
	LevelError = 4
	LevelFatal = 5
)

/*
New returns a new Error
*/
func New(msg string, level, code int, errs ...error) *Error {
	return &Error{
		Code:  code,
		Errs:  errs,
		Level: level,
		Msg:   msg,
	}
}

/*
NewFromErr returns a new Error based on an error interface
*/
func NewFromErr(err error) *Error {
	return &Error{
		Code:  0,
		Level: LevelError,
		Msg:   err.Error(),
	}
}

/*
Error extends the error type
*/
type Error struct {
	Code  int
	Errs  []error
	Level int
	Msg   string
}

func (e Error) Error() string {
	output := fmt.Sprintf("#%d", e.Code)

	if "" != e.Msg {
		output = fmt.Sprintf("%s: %s", output, e.Msg)
	}

	if len(e.Errs) > 0 {
		output = fmt.Sprintf("%s -", output)
		for _, err := range e.Errs {
			output = fmt.Sprintf(`%s msg="%s"; `, output, err.Error())
		}
	}
	return output
}

/*
Log logs the error
*/
func (e *Error) Log() (err Error) {
	switch e.Level {
	default:
		return Error{
			Code:  CodeInvalidLevel,
			Level: LevelWarn,
			Msg:   "Invalid log level '0'",
		}
	case 1:
		log.Debug(e.Error())
	case 2:
		log.Info(e.Error())
	case 3:
		log.Warn(e.Error())
	case 4:
		log.Error(e.Error())
	case 5:
		log.Fatal(e.Error())
	}
	return
}

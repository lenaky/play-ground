package qerr

import "fmt"

type Error struct {
	Err error
}

func (e *Error) Error() string {
	return e.Err.Error()
}

var ThisError = Error{Err: fmt.Errorf("this error")}

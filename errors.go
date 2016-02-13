package errors

import (
	"fmt"
)

func NewType(name string) *ErrorType {
	return &ErrorType{name: name, parent: nil}
}

func SubType(name string, parent *ErrorType) *ErrorType {
	return &ErrorType{name: name, parent: parent}
}

type ErrorType struct {
	name   string
	parent *ErrorType
}

func (tp *ErrorType) Error(msg string) error {
	return &Error{msg: msg, tp: tp}
}

func (tp *ErrorType) Errorf(msg string, args ...interface{}) error {
	return &Error{msg: fmt.Sprintf(msg, args...), tp: tp}
}

func (tp *ErrorType) IsTypeOf(err error) bool {
	e, ok := err.(*Error)
	if !ok {
		return false
	}
	current := e.tp
	for {
		if current == nil {
			return false
		}
		if tp == current {
			return true
		}
		current = current.parent
	}
}

type Error struct {
	msg string
	tp  *ErrorType
}

func (er *Error) Error() string {
	return er.msg
}

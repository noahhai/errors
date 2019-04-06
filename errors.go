package errors

import "fmt"

type formatFunc func(e *Error) string

var defaultFormatFunc formatFunc = func(e *Error) string {
	errString := e.msg
	curr := e
	for curr.child != nil {
		if curr.child.msg != "" {
			errString += "\n"
			errString += curr.child.msg
		}
		curr = curr.child
	}
	return errString
}

type Error struct {
	child  *Error
	msg    string
	format formatFunc
}

func New(msg string) *Error {
	return &Error{msg: msg}
}

func NewF(format string, a ...interface{}) *Error {
	return New(fmt.Sprintf(format, a...))
}

func From(e error) *Error {
	if asErr, ok := e.(*Error); ok {
		return asErr
	}
	return New(e.Error())
}

func FromTuple(o interface{}, e error) (interface{}, *Error) {
	return o, From(e)
}

func GetFromTupleAdd(msg string) func(o interface{}, e error) (interface{}, *Error) {
	return func(o interface{}, e error) (interface{}, *Error) {
		return o, From(e).Add(msg)
	}
}

func (e *Error) Add(msg string) *Error {
	n := Error{msg: msg}
	n.child = e
	return &n
}

func (e *Error) AddF(format string, a ...interface{}) *Error {
	return e.Add(fmt.Sprintf(format, a...))
}

func Or(err1, err2 error) *Error {
	if err1 != nil {
		return From(err1)
	} else if err2 != nil {
		return From(err2)
	}
	return nil
}

func (e *Error) Or(err *Error) *Error {
	if e == nil {
		return err
	}
	return e
}

func (e *Error) SetFormatter(f formatFunc) {
	e.format = f
}

func (e *Error) Error() string {
	if e.format != nil {
		return e.format(e)
	}
	return defaultFormatFunc(e)
}

func (e *Error) Symptom() string {
	return e.msg
}

func Symptom(e error) string {
	if asErr, ok := e.(*Error); ok {
		return asErr.Symptom()
	}
	return e.Error()
}

func (e *Error) Cause() string {
	last := e
	for last.child != nil {
		last = last.child
	}
	return last.msg
}

func Cause(e error) string {
	if asErr, ok := e.(*Error); ok {
		return asErr.Cause()
	}
	return e.Error()
}
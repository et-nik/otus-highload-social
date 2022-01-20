package web

import "net/http"

type Error struct {
	err        error
	httpStatus int
	message    string
}

func (e *Error) HTTPStatus() int {
	return e.httpStatus
}

func (e *Error) Message() string {
	return e.message
}

func (e *Error) Error() string {
	return e.err.Error()
}

func (e *Error) Unwrap() error {
	return e.err
}

func NewServerInternalError(err error, message string) *Error {
	return &Error{err: err, httpStatus: http.StatusInternalServerError, message: message}
}

func NewError(err error, status int, message string) *Error {
	return &Error{err: err, httpStatus: status, message: message}
}

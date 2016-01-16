package core

import (
	"log"
	"net/http"
)

type ErrorHandler interface {
	ServeError(http.ResponseWriter, *http.Request, error)
}

type ErrorHandlerFunc func(http.ResponseWriter, *http.Request, error)

func (fn ErrorHandlerFunc) ServeError(w http.ResponseWriter, r *http.Request, err error) {
	fn(w, r, err)
}

var DefaultErrorHandler ErrorHandler = ErrorHandlerFunc(defaultErrorHandlerFunc)

func defaultErrorHandlerFunc(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("[Error]: %q\r\n", err.Error())
}

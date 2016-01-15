package core

import (
	"log"
	"net/http"
)

type ErrorHandler interface {
	ServeError(http.ResponseWriter, *http.Request, error)
}

var DefaultErrorHandler ErrorHandler = ErrorHandlerFunc(func(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("[Error]: %q\r\n", err.Error())
})

type ErrorHandlerFunc func(http.ResponseWriter, *http.Request, error)

func (fn ErrorHandlerFunc) ServeError(w http.ResponseWriter, r *http.Request, err error) {
	fn(w, r, err)
}

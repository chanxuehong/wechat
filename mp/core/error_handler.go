package core

import (
	"log"
	"net/http"
	"os"
)

type ErrorHandler interface {
	ServeError(http.ResponseWriter, *http.Request, error)
}

var DefaultErrorHandler ErrorHandler = ErrorHandlerFunc(defaultErrorHandlerFunc)

type ErrorHandlerFunc func(http.ResponseWriter, *http.Request, error)

func (fn ErrorHandlerFunc) ServeError(w http.ResponseWriter, r *http.Request, err error) {
	fn(w, r, err)
}

var errorLogger = log.New(os.Stderr, "[WECHAT_ERROR] ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile)

func defaultErrorHandlerFunc(w http.ResponseWriter, r *http.Request, err error) {
	errorLogger.Output(3, err.Error())
}

package util

import (
	"net/http"
	"time"
)

var DefaultHttpClient *http.Client

func init() {
	client := *http.DefaultClient
	client.Timeout = time.Second * 5
	DefaultHttpClient = &client
}

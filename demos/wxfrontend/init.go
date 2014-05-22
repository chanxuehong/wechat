package wxfrontend

import (
	"github.com/chanxuehong/util/pool"
	"net/http"
)

func init() {
	var err error

	requestMsgPool, err = pool.New(newRequestMsg, 10)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", requestHandler)
}

// +build !wechatdebug

package callback

import (
	"encoding/xml"
	"io"
	"net/http"
)

func DebugPrintRequest(r *http.Request) {}

func XmlHttpRequestBodyUnmarshal(r io.Reader, v interface{}) error {
	return xml.NewDecoder(r).Decode(v)
}

func DebugPrintPlainMessage(msg []byte) {}

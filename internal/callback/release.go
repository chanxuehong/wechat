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

func DebugPrintPlainRequestMessage(msg []byte) {}

func XmlRawResponse(w io.Writer, msg interface{}) (err error) {
	return xml.NewEncoder(w).Encode(msg)
}

func DebugPrintPlainResponseMessage(msg []byte) {}

func DebugPrintCipherResponseMessage(msg, msgSignature, timestamp, nonce string) {}

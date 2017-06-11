// +build !wechat_debug

package callback

import (
	"encoding/xml"
	"io"
	"net/http"
)

func DebugPrintRequest(r *http.Request) {}

func DebugPrintPlainRequestMessage(msg []byte) {}

func XmlMarshalResponseMessage(msg interface{}) ([]byte, error) {
	return xml.Marshal(msg)
}

func XmlEncodeResponseMessage(w io.Writer, msg interface{}) error {
	return xml.NewEncoder(w).Encode(msg)
}

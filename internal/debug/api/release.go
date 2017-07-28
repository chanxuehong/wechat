// +build !wechat_debug

package api

import (
	"encoding/json"
	"io"
)

func DebugPrintGetRequest(url string) {}

func DebugPrintPostJSONRequest(url string, body []byte) {}

func DebugPrintPostMultipartRequest(url string, body []byte) {}

func DecodeJSONHttpResponse(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

// +build !wechatdebug

package api

import (
	"io"

	"github.com/chanxuehong/wechat.v2/json"
)

func DebugPrintGetRequest(url string) {}

func DebugPrintPostJSONRequest(url string, body []byte) {}

func DebugPrintPostMultipartRequest(url string, body []byte) {}

func DecodeJSONHttpResponse(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

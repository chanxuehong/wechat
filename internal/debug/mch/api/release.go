// +build !wechat_debug

package api

import (
	"io"

	"github.com/chanxuehong/util"
)

func DebugPrintGetRequest(url string) {}

func DebugPrintPostXMLRequest(url string, body []byte) {}

func DecodeXMLHttpResponse(r io.Reader) (map[string]string, error) {
	return util.DecodeXMLToMap(r)
}

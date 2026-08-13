//go:build !wechat_debug
// +build !wechat_debug

package api

import (
	"io"

	"gopkg.in/chanxuehong/wechat.v2/internal/util"
)

func DebugPrintGetRequest(url string) {}

func DebugPrintPostXMLRequest(url string, body []byte) {}

func DecodeXMLHttpResponse(r io.Reader) (map[string]string, error) {
	return util.DecodeXMLToMap(r)
}

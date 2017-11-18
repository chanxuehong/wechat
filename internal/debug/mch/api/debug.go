// +build wechat_debug

package api

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"

	"github.com/chanxuehong/util"
)

func DebugPrintGetRequest(url string) {
	log.Println("[WECHAT_DEBUG] [MCH] [API] GET", url)
}

func DebugPrintPostXMLRequest(url string, body []byte) {
	const format = "[WECHAT_DEBUG] [MCH] [API] XML POST %s\n" +
		"http request body:\n%s\n"
	log.Printf(format, url, body)
}

func DecodeXMLHttpResponse(r io.Reader) (map[string]string, error) {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	log.Printf("[WECHAT_DEBUG] [MCH] [API] http response body:\n%s\n", body)

	return util.DecodeXMLToMap(bytes.NewReader(body))
}

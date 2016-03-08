// +build wechatdebug

package api

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
)

func DebugPrintGetRequest(url string) {
	log.Println("[WECHAT_DEBUG] [API] GET", url)
}

func DebugPrintPostJSONRequest(url string, body []byte) {
	log.Println("[WECHAT_DEBUG] [API] POST", url)
	log.Printf("[WECHAT_DEBUG] [API] request body:\n%s\n", body)
}

func DebugPrintPostMultipartRequest(url string, body []byte) {
	log.Println("[WECHAT_DEBUG] [API] multipart/form-data POST", url)
}

func JsonHttpResponseUnmarshal(r io.Reader, v interface{}) error {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	log.Printf("[WECHAT_DEBUG] [API] response body:\n%s\n", body)

	return json.Unmarshal(body, v)
}

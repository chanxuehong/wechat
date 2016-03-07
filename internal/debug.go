// +build wechatdebug

package internal

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
)

func DebugPrintGetRequest(url string) {
	log.Println("[WECHAT_DEBUG] GET:", url)
}

func DebugPrintPostJSONRequest(url string, body []byte) {
	log.Println("[WECHAT_DEBUG] POST:", url)
	log.Println("[WECHAT_DEBUG] request body:", string(body))
}

func DebugPrintPostMultipartRequest(url string, body []byte) {
	log.Println("[WECHAT_DEBUG] POST Multipart:", url)
}

func JsonHttpResponseUnmarshal(r io.Reader, v interface{}) error {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	log.Println("[WECHAT_DEBUG] response body:", string(body))

	return json.Unmarshal(body, v)
}

// RETRY ===============================================================================================================

// access_token 过期重试之前打印相应信息
func DebugPrintRetryError(errcode int64, errmsg string, token string) {
	log.Printf("[WECHAT_RETRY] errcode: %d, errmsg: %s\n", errcode, errmsg)
	log.Println("[WECHAT_RETRY] current token:", token)
}

// access_token 过期重试过程中打印新的 access_token
func DebugPrintRetryNewToken(token string) {
	log.Println("[WECHAT_RETRY] new token:", token)
}

// access_token 过期重试失败打印对应的 access_token
func DebugPrintRetryFallthrough(token string) {
	log.Println("[WECHAT_RETRY] fallthrough, current token:", token)
}

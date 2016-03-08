// +build !wechatdebug

package internal

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"net/http"
)

func DebugPrintGetRequest(url string) {}

func DebugPrintPostJSONRequest(url string, body []byte) {}

func DebugPrintPostMultipartRequest(url string, body []byte) {}

func JsonHttpResponseUnmarshal(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

// RETRY ===============================================================================================================

// access_token 过期重试之前打印相应信息
func DebugPrintRetryError(errcode int64, errmsg string, token string) {}

// access_token 过期重试过程中打印新的 access_token
func DebugPrintRetryNewToken(token string) {}

// access_token 过期重试失败打印对应的 access_token
func DebugPrintRetryFallthrough(token string) {}

// callback ============================================================================================================

func DebugPrintCallbackRequest(r *http.Request) {}

func CallbackAesXmlRequestBodyUnmarshal(r io.Reader, v interface{}) error {
	return xml.NewDecoder(r).Decode(v)
}

func DebugPrintCallbackPlainMessage(msg []byte) {}

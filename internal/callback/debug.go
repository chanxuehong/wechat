// +build wechatdebug

package callback

import (
	"encoding/xml"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func DebugPrintRequest(r *http.Request) {
	log.Println("[WECHAT_DEBUG] [CALLBACK]", r.Method, r.RequestURI)
}

func AesXmlRequestBodyUnmarshal(r io.Reader, v interface{}) error {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	log.Printf("[WECHAT_DEBUG] [CALLBACK] http request body:\n%s\n", body)

	return xml.Unmarshal(body, v)
}

func DebugPrintPlainMessage(msg []byte) {
	log.Printf("[WECHAT_DEBUG] [CALLBACK] plain message:\n%s\n", msg)
}

// +build wechat_debug

package callback

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"github.com/chanxuehong/util"
)

func DebugPrintRequest(r *http.Request) {
	log.Println("[WECHAT_DEBUG] [MCH] [CALLBACK]", r.Method, r.RequestURI)
}

func DebugPrintRequestMessage(msg []byte) {
	log.Printf("[WECHAT_DEBUG] [MCH] [CALLBACK] http request body:\n%s\n", msg)
}

func EncodeXMLResponseMessage(w io.Writer, msg map[string]string) (err error) {
	var buf bytes.Buffer
	if err = util.EncodeXMLFromMap(&buf, msg, "xml"); err != nil {
		return
	}
	log.Printf("[WECHAT_DEBUG] [MCH] [CALLBACK] http response body:\n%s\n", buf.Bytes())

	_, err = buf.WriteTo(w)
	return
}

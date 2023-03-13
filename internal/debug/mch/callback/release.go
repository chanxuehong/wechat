//go:build !wechat_debug
// +build !wechat_debug

package callback

import (
	"io"
	"net/http"

	"github.com/bububa/wechat/util"
)

func DebugPrintRequest(r *http.Request) {}

func DebugPrintRequestMessage(msg []byte) {}

func EncodeXMLResponseMessage(w io.Writer, msg map[string]string) (err error) {
	return util.EncodeXMLFromMap(w, msg, "xml")
}

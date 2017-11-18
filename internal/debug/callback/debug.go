// +build wechat_debug

package callback

import (
	"encoding/xml"
	"io"
	"log"
	"net/http"
)

func DebugPrintRequest(r *http.Request) {
	log.Println("[WECHAT_DEBUG] [CALLBACK]", r.Method, r.RequestURI)
}

func DebugPrintPlainRequestMessage(msg []byte) {
	log.Printf("[WECHAT_DEBUG] [CALLBACK] plain request message:\n%s\n", msg)
}

func XmlMarshalResponseMessage(msg interface{}) ([]byte, error) {
	bs, err := xml.MarshalIndent(msg, "", "    ")
	if err != nil {
		return nil, err
	}
	log.Printf("[WECHAT_DEBUG] [CALLBACK] plain response message:\n%s\n", bs)
	return bs, nil
}

func XmlEncodeResponseMessage(w io.Writer, msg interface{}) error {
	bs, err := xml.MarshalIndent(msg, "", "    ")
	if err != nil {
		return err
	}
	log.Printf("[WECHAT_DEBUG] [CALLBACK] plain response message:\n%s\n", bs)

	_, err = w.Write(bs)
	return err
}

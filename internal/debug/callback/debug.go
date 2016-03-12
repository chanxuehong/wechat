// +build wechatdebug

package callback

import (
	"bytes"
	"encoding/xml"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func DebugPrintRequest(r *http.Request) {
	log.Println("[WECHAT_DEBUG] [CALLBACK]", r.Method, r.RequestURI)
}

func XmlHttpRequestBodyUnmarshal(r io.Reader, v interface{}) error {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	log.Printf("[WECHAT_DEBUG] [CALLBACK] cipher request message:\n%s\n", body)

	return xml.Unmarshal(body, v)
}

func DebugPrintPlainRequestMessage(msg []byte) {
	log.Printf("[WECHAT_DEBUG] [CALLBACK] plain request message:\n%s\n", msg)
}

func XmlRawResponse(w io.Writer, msg interface{}) (err error) {
	body, err := xml.Marshal(msg)
	if err != nil {
		return
	}
	log.Printf("[WECHAT_DEBUG] [CALLBACK] plain response message:\n%s\n", body)

	_, err = w.Write(body)
	return
}

func DebugPrintPlainResponseMessage(msg []byte) {
	log.Printf("[WECHAT_DEBUG] [CALLBACK] plain response message:\n%s\n", msg)
}

func DebugPrintCipherResponseMessage(msg, msgSignature, timestamp, nonce string) {
	var buf bytes.Buffer
	xml.EscapeText(&buf, []byte(nonce))
	nonce = buf.String()

	ciphertext := "<xml><Encrypt>" + msg + "</Encrypt><MsgSignature>" + msgSignature + "</MsgSignature><TimeStamp>" +
		timestamp + "</TimeStamp><Nonce>" + nonce + "</Nonce></xml>"
	log.Printf("[WECHAT_DEBUG] [CALLBACK] cipher response message:\n%s\n", ciphertext)
}

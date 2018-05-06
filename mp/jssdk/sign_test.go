package jssdk

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"strings"
	"testing"
)

func TestWXConfigSign(t *testing.T) {
	jsapiTicket := "sM4AOVdWfPE4DxkXGEs8VMCPGGVi4C3VM0P37wVUCFvkVAy_90u5h9nbSlYy3-Sl-HhTdfl2fzFy1AOcHKP7qg"
	nonceStr := "Wm3WZYTPz0wzccnW"
	timestamp := "1414587457"
	url := "http://mp.weixin.qq.com?params=value#xxxx"

	wantSignature := "0f9de62fce790f9a083d5c99e95740ceb90c27ed"

	haveSignature := WXConfigSign(jsapiTicket, nonceStr, timestamp, url)
	if haveSignature != wantSignature {
		t.Errorf("test WXConfigSign failed,\nhave: %s\nwant: %s\n", haveSignature, wantSignature)
		return
	}

	haveSignature2 := WXConfigSign2(jsapiTicket, nonceStr, timestamp, url)
	if haveSignature2 != wantSignature {
		t.Errorf("test WXConfigSign2 failed,\nhave: %s\nwant: %s\n", haveSignature2, wantSignature)
		return
	}

	haveSignature3 := WXConfigSign3(jsapiTicket, nonceStr, timestamp, url)
	if haveSignature3 != wantSignature {
		t.Errorf("test WXConfigSign3 failed,\nhave: %s\nwant: %s\n", haveSignature3, wantSignature)
		return
	}
}

func TestCardSign(t *testing.T) {
	api_ticket := "aaaa"
	timestamp := "bbbb"
	nonce_str := "cccc"
	card_id := "dddd"
	code := "eeee"
	open_id := "ffff"

	wantSignature := "89a0e60888a9471f75dc5eb0ee86431ddbec1fd9"

	haveSignature := CardSign([]string{open_id, code, timestamp, card_id, api_ticket, nonce_str})
	if haveSignature != wantSignature {
		t.Errorf("tests CardSign failed,\nhave: %s\nwant: %s\n", haveSignature, wantSignature)
		return
	}
}

func BenchmarkWXConfigSign(b *testing.B) {
	jsapiTicket := "sM4AOVdWfPE4DxkXGEs8VMCPGGVi4C3VM0P37wVUCFvkVAy_90u5h9nbSlYy3-Sl-HhTdfl2fzFy1AOcHKP7qg"
	nonceStr := "Wm3WZYTPz0wzccnW"
	timestamp := "1414587457"
	url := "http://mp.weixin.qq.com?params=value"

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		WXConfigSign(jsapiTicket, nonceStr, timestamp, url)
	}
}

func BenchmarkWXConfigSign2(b *testing.B) {
	jsapiTicket := "sM4AOVdWfPE4DxkXGEs8VMCPGGVi4C3VM0P37wVUCFvkVAy_90u5h9nbSlYy3-Sl-HhTdfl2fzFy1AOcHKP7qg"
	nonceStr := "Wm3WZYTPz0wzccnW"
	timestamp := "1414587457"
	url := "http://mp.weixin.qq.com?params=value"

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		WXConfigSign2(jsapiTicket, nonceStr, timestamp, url)
	}
}

func BenchmarkWXConfigSign3(b *testing.B) {
	jsapiTicket := "sM4AOVdWfPE4DxkXGEs8VMCPGGVi4C3VM0P37wVUCFvkVAy_90u5h9nbSlYy3-Sl-HhTdfl2fzFy1AOcHKP7qg"
	nonceStr := "Wm3WZYTPz0wzccnW"
	timestamp := "1414587457"
	url := "http://mp.weixin.qq.com?params=value"

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		WXConfigSign3(jsapiTicket, nonceStr, timestamp, url)
	}
}

func WXConfigSign2(jsapiTicket, nonceStr, timestamp, url string) (signature string) {
	if i := strings.IndexByte(url, '#'); i >= 0 {
		url = url[:i]
	}
	h := sha1.New()

	bufw := bufio.NewWriterSize(h, 128)
	bufw.WriteString("jsapi_ticket=")
	bufw.WriteString(jsapiTicket)
	bufw.WriteString("&noncestr=")
	bufw.WriteString(nonceStr)
	bufw.WriteString("&timestamp=")
	bufw.WriteString(timestamp)
	bufw.WriteString("&url=")
	bufw.WriteString(url)
	bufw.Flush()

	hashsum := h.Sum(nil)
	return hex.EncodeToString(hashsum)
}

var (
	jsapiTicketKey = []byte("jsapi_ticket=")
	nonceStrKey    = []byte("&noncestr=")
	timestampKey   = []byte("&timestamp=")
	urlKey         = []byte("&url=")
)

func WXConfigSign3(jsapiTicket, nonceStr, timestamp, url string) (signature string) {
	if i := strings.IndexByte(url, '#'); i >= 0 {
		url = url[:i]
	}
	h := sha1.New()

	h.Write(jsapiTicketKey)
	h.Write([]byte(jsapiTicket))
	h.Write(nonceStrKey)
	h.Write([]byte(nonceStr))
	h.Write(timestampKey)
	h.Write([]byte(timestamp))
	h.Write(urlKey)
	h.Write([]byte(url))

	hashsum := h.Sum(nil)
	return hex.EncodeToString(hashsum)
}

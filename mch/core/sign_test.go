package core

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"hash"
	"sort"
	"testing"
)

func TestSign(t *testing.T) {
	params := map[string]string{
		"appid":       "wxd930ea5d5a258f4f",
		"mch_id":      "10000100",
		"device_info": "1000",
		"body":        "test",
		"nonce_str":   "ibuaiVcKdpRxkhJA",
	}
	apiKey := "192006250b4c09247ec02edce69f6a2d"

	haveSignature := Sign(params, apiKey, nil)
	wantSignature := "9A0A8659F005D6984697E2CA0A9CF3B7"
	if haveSignature != wantSignature {
		t.Errorf("signature mismatch, have %s, want %s", haveSignature, wantSignature)
		return
	}
}

func TestJsapiSign(t *testing.T) {
	appId := "appId_value"
	timeStamp := "123456789"
	nonceStr := "nonceStr_value"
	packageStr := "prepay_id=asdfasdfasdf"
	signType := SignType_MD5

	params := map[string]string{
		"appId":     appId,
		"timeStamp": timeStamp,
		"nonceStr":  nonceStr,
		"package":   packageStr,
		"signType":  signType,
	}
	apiKey := "192006250b4c09247ec02edce69f6a2d"

	haveSignature := JsapiSign(appId, timeStamp, nonceStr, packageStr, signType, apiKey)
	wantSignature := Sign(params, apiKey, nil)
	if haveSignature != wantSignature {
		t.Errorf("signature mismatch, have %s, want %s", haveSignature, wantSignature)
		return
	}
}

func BenchmarkSign(b *testing.B) {
	b.StopTimer()
	var (
		params = map[string]string{
			"asjdfsadfsadfasdfasd1":  "sdfasdfasdfsadfsadfasdfasdfasd",
			"asjdfsadfsadfasdfasd2":  "sdfasdfasdfsadfsadfasdfasdfasd",
			"asjdfsadfsadfasdfasd3":  "sdfasdfasdfsadfsadfasdfasdfasd",
			"asjdfsadfsadfasdfasd4":  "sdfasdfasdfsadfsadfasdfasdfasd",
			"asjdfsadfsadfasdfasd5":  "sdfasdfasdfsadfsadfasdfasdfasd",
			"asjdfsadfsadfasdfasd6":  "sdfasdfasdfsadfsadfasdfasdfasd",
			"asjdfsadfsadfasdfasd7":  "sdfasdfasdfsadfsadfasdfasdfasd",
			"asjdfsadfsadfasdfasd8":  "sdfasdfasdfsadfsadfasdfasdfasd",
			"asjdfsadfsadfasdfasd9":  "sdfasdfasdfsadfsadfasdfasdfasd",
			"asjdfsadfsadfasdfasd11": "sdfasdfasdfsadfsadfasdfasdfasd",
			"asjdfsadfsadfasdfasd12": "sdfasdfasdfsadfsadfasdfasdfasd",
			"asjdfsadfsadfasdfasd13": "sdfasdfasdfsadfsadfasdfasdfasd",
			"asjdfsadfsadfasdfasd14": "sdfasdfasdfsadfsadfasdfasdfasd",
			"asjdfsadfsadfasdfasd15": "sdfasdfasdfsadfsadfasdfasdfasd",
		}
		apiKey = "afadskfjaskldjflkasdjflkashdljkfhalsdjkfhl"
	)
	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Sign(params, apiKey, nil)
	}
}

func BenchmarkClassicalSign(b *testing.B) {
	b.StopTimer()
	var (
		params = map[string]string{
			"asjdfsadfsadfasdfasd1":  "sdfasdfasdfsadfsadfasdfasdfasd",
			"asjdfsadfsadfasdfasd2":  "sdfasdfasdfsadfsadfasdfasdfasd",
			"asjdfsadfsadfasdfasd3":  "sdfasdfasdfsadfsadfasdfasdfasd",
			"asjdfsadfsadfasdfasd4":  "sdfasdfasdfsadfsadfasdfasdfasd",
			"asjdfsadfsadfasdfasd5":  "sdfasdfasdfsadfsadfasdfasdfasd",
			"asjdfsadfsadfasdfasd6":  "sdfasdfasdfsadfsadfasdfasdfasd",
			"asjdfsadfsadfasdfasd7":  "sdfasdfasdfsadfsadfasdfasdfasd",
			"asjdfsadfsadfasdfasd8":  "sdfasdfasdfsadfsadfasdfasdfasd",
			"asjdfsadfsadfasdfasd9":  "sdfasdfasdfsadfsadfasdfasdfasd",
			"asjdfsadfsadfasdfasd11": "sdfasdfasdfsadfsadfasdfasdfasd",
			"asjdfsadfsadfasdfasd12": "sdfasdfasdfsadfsadfasdfasdfasd",
			"asjdfsadfsadfasdfasd13": "sdfasdfasdfsadfsadfasdfasdfasd",
			"asjdfsadfsadfasdfasd14": "sdfasdfasdfsadfsadfasdfasdfasd",
			"asjdfsadfsadfasdfasd15": "sdfasdfasdfsadfsadfasdfasdfasd",
		}
		apiKey = "afadskfjaskldjflkasdjflkashdljkfhalsdjkfhl"
	)
	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ClassicalSign(params, apiKey, nil)
	}
}

// 传统的签名代码, Sign 是优化后的代码, 要提高 30% 的速度
func ClassicalSign(params map[string]string, apiKey string, fn func() hash.Hash) string {
	if fn == nil {
		fn = md5.New
	}
	h := fn()

	keys := make([]string, 0, len(params))
	for k := range params {
		if k == "sign" {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		v := params[k]
		if v == "" {
			continue
		}
		h.Write([]byte(k))
		h.Write([]byte{'='})
		h.Write([]byte(v))
		h.Write([]byte{'&'})
	}
	h.Write([]byte("key="))
	h.Write([]byte(apiKey))

	signature := make([]byte, hex.EncodedLen(h.Size()))
	hex.Encode(signature, h.Sum(nil))
	return string(bytes.ToUpper(signature))
}

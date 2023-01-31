package api

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
)

func DebugPrintGetRequest(url string, debug bool) {
	if !debug {
		return
	}
	log.Println("[WECHAT_DEBUG] [API] GET", url)
}

func DebugPrintPostJSONRequest(url string, body []byte, debug bool) {
	if !debug {
		return
	}
	const format = "[WECHAT_DEBUG] [API] JSON POST %s\n" +
		"http request body:\n%s\n"

	buf := bytes.NewBuffer(make([]byte, 0, len(body)+1024))
	if err := json.Indent(buf, body, "", "    "); err == nil {
		body = buf.Bytes()
	}
	log.Printf(format, url, body)
}

func DebugPrintPostMultipartRequest(url string, body []byte, debug bool) {
	if !debug {
		return
	}
	log.Println("[WECHAT_DEBUG] [API] multipart/form-data POST", url)
}

func DecodeJSONHttpResponse(r io.Reader, v interface{}, debug bool) error {
	if !debug {
		return json.NewDecoder(r).Decode(v)
	}
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	body2 := body
	buf := bytes.NewBuffer(make([]byte, 0, len(body2)+1024))
	if err := json.Indent(buf, body2, "", "    "); err == nil {
		body2 = buf.Bytes()
	}
	log.Printf("[WECHAT_DEBUG] [API] http response body:\n%s\n", body2)

	return json.Unmarshal(body, v)
}

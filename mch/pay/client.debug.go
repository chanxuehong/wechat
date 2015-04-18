// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

// +build wechatdebug

package pay

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"

	"github.com/chanxuehong/util"
)

type Client struct {
	apiKey     string
	httpClient *http.Client
}

// 创建一个新的 Client.
//  如果 httpClient == nil 则默认用 http.DefaultClient.
func NewClient(apiKey string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &Client{
		apiKey:     apiKey,
		httpClient: httpClient,
	}
}

// 微信支付通用请求方法.
//  注意: err == nil 表示协议状态都为 SUCCESS.
func (clt *Client) PostXML(url string, req map[string]string) (resp map[string]string, err error) {
	bodyBuf := textBufferPool.Get().(*bytes.Buffer)
	bodyBuf.Reset()
	defer textBufferPool.Put(bodyBuf)

	if err = util.FormatMapToXML(bodyBuf, req); err != nil {
		return
	}

	debugPrefix := "pay.Client.PostXML"
	if _, file, line, ok := runtime.Caller(1); ok {
		debugPrefix += fmt.Sprintf("(called at %s:%d)", file, line)
	}
	log.Println(debugPrefix, "request url:", url)
	log.Println(debugPrefix, "request xml:", bodyBuf.String())

	httpResp, err := clt.httpClient.Post(url, "text/xml; charset=utf-8", bodyBuf)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		err = fmt.Errorf("http.Status: %s", httpResp.Status)
		return
	}

	respBody, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return
	}
	log.Println(debugPrefix, "response xml:", string(respBody))

	if resp, err = util.ParseXMLToMap(bytes.NewReader(respBody)); err != nil {
		return
	}

	// 判断协议状态
	ReturnCode, ok := resp["return_code"]
	if !ok {
		err = errors.New("no return_code parameter")
		return
	}
	if ReturnCode != ReturnCodeSuccess {
		err = &Error{
			ReturnCode: ReturnCode,
			ReturnMsg:  resp["return_msg"],
		}
		return
	}

	// 认证签名
	signature1, ok := resp["sign"]
	if !ok {
		err = errors.New("no sign parameter")
		return
	}
	signature2 := Sign(resp, clt.apiKey, nil)
	if signature1 != signature2 {
		err = fmt.Errorf("check signature failed, \r\ninput: %q, \r\nlocal: %q", signature1, signature2)
		return
	}
	return
}

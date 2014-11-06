// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay2

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/chanxuehong/wechat/mp/pay"
)

// 封装 退款及对账 功能
type TenpayClient struct {
	partnerId, partnerKey string
	httpClient            *http.Client
}

// 创建一个新的 TenpayClient.
//  如果 httpClient == nil 则默认用 http.DefaultClient
func NewTenpayClient(partnerId, partnerKey string, httpClient *http.Client) *TenpayClient {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &TenpayClient{
		partnerId:  partnerId,
		partnerKey: partnerKey,
		httpClient: httpClient,
	}
}

func (c *TenpayClient) postXML(url_ string, request map[string]string, response map[string]string) (err error) {
	buf := textBufferPool.Get().(*bytes.Buffer) // io.ReadWriter
	buf.Reset()                                 // important
	defer textBufferPool.Put(buf)               // important

	if err = pay.FormatMapToXML(buf, request); err != nil {
		return
	}

	resp, err := c.httpClient.Post(url_, "text/xml; charset=utf-8", buf)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http.Status: %s", resp.Status)
	}

	if err = pay.ParseXMLToMap(resp.Body, response); err != nil {
		return
	}

	return
}

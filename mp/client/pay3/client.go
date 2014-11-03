// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay3

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/chanxuehong/wechat/mp/pay"
	"github.com/chanxuehong/wechat/mp/pay/pay3"
	"net/http"
)

type Client struct {
	appKey     string
	httpClient *http.Client
}

// 创建一个新的 Client.
//  如果 httpClient == nil 则默认用 http.DefaultClient
func NewClient(appKey string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &Client{
		appKey:     appKey,
		httpClient: httpClient,
	}
}

// 统一支付接口
func (c *Client) UnifiedOrder(req pay3.UnifiedOrderRequest) (resp pay3.UnifiedOrderResponse, err error) {
	if req == nil {
		err = errors.New("req == nil")
		return
	}

	url_ := "https://api.mch.weixin.qq.com/pay/unifiedorder"
	result := make(map[string]string)

	if err = c.postMapXML(url_, req, result); err != nil {
		return
	}

	resp = result
	return
}

// 订单查询接口
func (c *Client) OrderQuery(req pay3.OrderQueryRequest) (resp pay3.OrderQueryResponse, err error) {
	if req == nil {
		err = errors.New("req == nil")
		return
	}

	url_ := "https://api.mch.weixin.qq.com/pay/orderquery"
	result := make(map[string]string)

	if err = c.postMapXML(url_, req, result); err != nil {
		return
	}

	resp = result
	return
}

// 关闭订单接口
func (c *Client) OrderClose(req pay3.OrderCloseRequest) (resp pay3.OrderCloseResponse, err error) {
	if req == nil {
		err = errors.New("req == nil")
		return
	}

	url_ := "https://api.mch.weixin.qq.com/pay/closeorder"
	result := make(map[string]string)

	if err = c.postMapXML(url_, req, result); err != nil {
		return
	}

	resp = result
	return
}

// 退款申请接口
func (c *Client) Refund(req pay3.RefundRequest) (resp pay3.RefundResponse, err error) {
	if req == nil {
		err = errors.New("req == nil")
		return
	}

	url_ := "https://api.mch.weixin.qq.com/secapi/pay/refund"
	result := make(map[string]string)

	if err = c.postMapXML(url_, req, result); err != nil {
		return
	}

	resp = result
	return
}

// 退款查询接口
func (c *Client) RefundQuery(req pay3.RefundQueryRequest) (resp pay3.RefundQueryResponse, err error) {
	if req == nil {
		err = errors.New("req == nil")
		return
	}

	url_ := "https://api.mch.weixin.qq.com/pay/refundquery"
	result := make(map[string]string)

	if err = c.postMapXML(url_, req, result); err != nil {
		return
	}

	resp = result
	return
}

// 短链接转换接口
func (c *Client) ShortURL(req pay3.ShortURLRequest) (resp pay3.ShortURLResponse, err error) {
	if req == nil {
		err = errors.New("req == nil")
		return
	}

	url_ := "https://api.mch.weixin.qq.com/tools/shorturl"
	result := make(map[string]string)

	if err = c.postMapXML(url_, req, result); err != nil {
		return
	}

	resp = result
	return
}

// 用于微信支付
func (c *Client) postMapXML(url_ string, request map[string]string, response map[string]string) (err error) {
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

	if RetCode := response["return_code"]; RetCode != pay3.RET_CODE_SUCCESS {
		err = &Error{
			ErrCode: RetCode,
			ErrMsg:  response["return_msg"],
		}
		return
	}

	if err = pay3.CheckSignature(response, c.appKey); err != nil {
		return
	}

	return
}

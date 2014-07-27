// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"errors"
	"github.com/chanxuehong/wechat/pay"
)

// 微信支付发货通知
func (c *Client) PayDeliverNotify(data *pay.DeliverNotifyData) (err error) {
	if data == nil {
		return errors.New("data == nil")
	}

	var result Error

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	_url := payDeliverNotifyURL(token)
	if err = c.postJSON(_url, data, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		return

	case errCodeTimeout:
		if !hasRetry {
			hasRetry = true
			timeoutRetryWait()
			goto RETRY
		}
		fallthrough

	default:
		err = result
		return
	}
}

// 微信支付订单查询
func (c *Client) PayOrderQuery(req *pay.OrderQueryRequest) (resp *pay.OrderQueryResponse, err error) {
	if req == nil {
		err = errors.New("req == nil")
		return
	}

	var result struct {
		Error
		OrderInfo pay.OrderQueryResponse `json:"order_info"`
	}

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	_url := payOrderQueryURL(token)
	if err = c.postJSON(_url, req, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		resp = &result.OrderInfo
		return

	case errCodeTimeout:
		if !hasRetry {
			hasRetry = true
			timeoutRetryWait()
			goto RETRY
		}
		fallthrough

	default:
		err = result.Error
		return
	}
}

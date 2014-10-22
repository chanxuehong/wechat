// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"errors"
	"github.com/chanxuehong/wechat/mp/pay2"
)

// 微信支付发货通知
func (c *Client) Pay2DeliverNotify(data *pay2.DeliverNotifyData) (err error) {
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
	url_ := pay2DeliverNotifyURL(token)

	if err = c.postJSON(url_, data, &result); err != nil {
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
		err = &result
		return
	}
}

// 微信支付订单查询
func (c *Client) Pay2OrderQuery(req *pay2.OrderQueryRequest) (resp *pay2.OrderQueryResponse, err error) {
	if req == nil {
		err = errors.New("req == nil")
		return
	}

	var result struct {
		Error
		OrderInfo pay2.OrderQueryResponse `json:"order_info"`
	}

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	url_ := pay2OrderQueryURL(token)

	if err = c.postJSON(url_, req, &result); err != nil {
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
		err = &result.Error
		return
	}
}

// 标记客户的投诉处理状态
func (c *Client) Pay2FeedbackUpdate(openid string, feedbackid int64) (err error) {
	var result Error

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	url_ := pay2FeedbackUpdateURL(token, openid, feedbackid)

	if err = c.getJSON(url_, &result); err != nil {
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
		err = &result
		return
	}
}

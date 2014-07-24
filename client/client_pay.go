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

	token, err := c.Token()
	if err != nil {
		return
	}
	_url := payDeliverNotifyURL(token)

	var result Error
	if err = c.postJSON(_url, data, &result); err != nil {
		return
	}

	if result.ErrCode != 0 {
		return &result
	}
	return
}

// 微信支付订单查询
func (c *Client) PayOrderQuery(req *pay.OrderQueryRequest) (resp *pay.OrderQueryResponse, err error) {
	if req == nil {
		err = errors.New("req == nil")
		return
	}

	token, err := c.Token()
	if err != nil {
		return
	}
	_url := payOrderQueryURL(token)

	var result struct {
		Error
		OrderInfo pay.OrderQueryResponse `json:"order_info"`
	}
	if err = c.postJSON(_url, req, &result); err != nil {
		return
	}

	if result.ErrCode != 0 {
		err = &result.Error
		return
	}

	resp = &result.OrderInfo
	return
}

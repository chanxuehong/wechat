// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"errors"

	"github.com/chanxuehong/wechat/mp/message/active/massbygroup"
)

// 根据分组群发文本消息.
func (c *Client) MsgMassSendTextByGroup(msg *massbygroup.Text) (msgid int64, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.msgMassSendByGroup(msg)
}

// 根据分组群发图片消息.
func (c *Client) MsgMassSendImageByGroup(msg *massbygroup.Image) (msgid int64, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.msgMassSendByGroup(msg)
}

// 根据分组群发语音消息.
func (c *Client) MsgMassSendVoiceByGroup(msg *massbygroup.Voice) (msgid int64, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.msgMassSendByGroup(msg)
}

// 根据分组群发视频消息.
func (c *Client) MsgMassSendVideoByGroup(msg *massbygroup.Video) (msgid int64, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.msgMassSendByGroup(msg)
}

// 根据分组群发图文消息.
func (c *Client) MsgMassSendNewsByGroup(msg *massbygroup.News) (msgid int64, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.msgMassSendByGroup(msg)
}

func (c *Client) msgMassSendByGroup(msg interface{}) (msgid int64, err error) {
	var result struct {
		Error
		MsgId int64 `json:"msg_id"`
	}

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	url_ := messageMassSendByGroupURL(token)

	if err = c.postJSON(url_, msg, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		msgid = result.MsgId
		return

	case errCodeInvalidCredential, errCodeTimeout:
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

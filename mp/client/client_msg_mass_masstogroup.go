// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"errors"

	"github.com/chanxuehong/wechat/mp/message/active/mass/masstogroup"
)

// 根据分组群发文本消息.
func (c *Client) MsgMassSendTextToGroup(msg *masstogroup.Text) (msgid int64, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.msgMassSendToGroup(msg)
}

// 根据分组群发图片消息.
func (c *Client) MsgMassSendImageToGroup(msg *masstogroup.Image) (msgid int64, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.msgMassSendToGroup(msg)
}

// 根据分组群发语音消息.
func (c *Client) MsgMassSendVoiceToGroup(msg *masstogroup.Voice) (msgid int64, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.msgMassSendToGroup(msg)
}

// 根据分组群发视频消息.
func (c *Client) MsgMassSendVideoToGroup(msg *masstogroup.Video) (msgid int64, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.msgMassSendToGroup(msg)
}

// 根据分组群发图文消息.
func (c *Client) MsgMassSendNewsToGroup(msg *masstogroup.News) (msgid int64, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.msgMassSendToGroup(msg)
}

func (c *Client) msgMassSendToGroup(msg interface{}) (msgid int64, err error) {
	var result struct {
		Error
		MsgId int64 `json:"msg_id"`
	}

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url_ := messageMassSendToGroupURL(token)

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

			if token, err = getNewToken(c.tokenService, token); err != nil {
				return
			}
			goto RETRY
		}
		fallthrough
	default:
		err = &result.Error
		return
	}
}

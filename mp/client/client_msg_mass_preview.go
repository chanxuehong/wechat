// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"errors"

	"github.com/chanxuehong/wechat/mp/message/active/mass/preview"
)

func (c *Client) MsgMassPreviewText(msg *preview.Text) (msgid int64, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.msgMassPreview(msg)
}

func (c *Client) MsgMassPreviewImage(msg *preview.Image) (msgid int64, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.msgMassPreview(msg)
}

func (c *Client) MsgMassPreviewVoice(msg *preview.Voice) (msgid int64, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.msgMassPreview(msg)
}

func (c *Client) MsgMassPreviewVideo(msg *preview.Video) (msgid int64, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.msgMassPreview(msg)
}

func (c *Client) MsgMassPreviewNews(msg *preview.News) (msgid int64, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.msgMassPreview(msg)
}

func (c *Client) msgMassPreview(msg interface{}) (msgid int64, err error) {
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
	url_ := messageMassPreviewURL(token)

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

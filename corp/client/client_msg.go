// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"errors"

	"github.com/chanxuehong/wechat/corp/message/active/common"
)

func (c *Client) MsgSendText(msg *common.Text) (result *common.Result, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.msgSend(msg)
}

func (c *Client) MsgSendImage(msg *common.Image) (result *common.Result, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.msgSend(msg)
}

func (c *Client) MsgSendVoice(msg *common.Voice) (result *common.Result, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.msgSend(msg)
}

func (c *Client) MsgSendVideo(msg *common.Video) (result *common.Result, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.msgSend(msg)
}

func (c *Client) MsgSendFile(msg *common.File) (result *common.Result, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.msgSend(msg)
}

func (c *Client) MsgSendNews(msg *common.News) (result *common.Result, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	if err = msg.CheckValid(); err != nil {
		return
	}
	return c.msgSend(msg)
}

func (c *Client) MsgSendMPNews(msg *common.MPNews) (result *common.Result, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	if err = msg.CheckValid(); err != nil {
		return
	}
	return c.msgSend(msg)
}

func (c *Client) msgSend(msg interface{}) (result *common.Result, err error) {
	var resultx struct {
		Error
		common.Result
	}

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	if err = c.postJSON(_MsgSendURL(token), msg, &resultx); err != nil {
		return
	}

	switch resultx.ErrCode {
	case errCodeOK:
		result = &resultx.Result
		return
	case errCodeTimeout, errCodeInvalidCredential:
		if !hasRetry {
			hasRetry = true

			if token, err = c.TokenRefresh(); err != nil {
				return
			}
			goto RETRY
		}
		fallthrough
	default:
		err = &resultx.Error
		return
	}
}

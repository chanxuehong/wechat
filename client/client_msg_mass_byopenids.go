// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"errors"
	"github.com/chanxuehong/wechat/message/massbyopenids"
)

// 根据 OpenId列表 群发消息, 之所以不暴露这个接口是因为怕接收到不合法的参数.
func (c *Client) msgMassSendByOpenIds(msg interface{}) (msgid int64, err error) {
	token, err := c.Token()
	if err != nil {
		return
	}
	_url := messageMassSendByOpenIdsURL(token)

	var result struct {
		Error
		MsgId int64 `json:"msg_id"`
	}
	if err = c.postJSON(_url, msg, &result); err != nil {
		return
	}

	if result.ErrCode != 0 {
		err = &result.Error
		return
	}
	msgid = result.MsgId
	return
}

// 根据用户列表群发文本消息.
func (c *Client) MsgMassSendTextByOpenIds(msg *massbyopenids.Text) (msgid int64, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.msgMassSendByOpenIds(msg)
}

// 根据用户列表群发图片消息.
func (c *Client) MsgMassSendImageByOpenIds(msg *massbyopenids.Image) (msgid int64, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.msgMassSendByOpenIds(msg)
}

// 根据用户列表群发语音消息.
func (c *Client) MsgMassSendVoiceByOpenIds(msg *massbyopenids.Voice) (msgid int64, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.msgMassSendByOpenIds(msg)
}

// 根据用户列表群发视频消息.
func (c *Client) MsgMassSendVideoByOpenIds(msg *massbyopenids.Video) (msgid int64, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.msgMassSendByOpenIds(msg)
}

// 根据用户列表群发图文消息.
func (c *Client) MsgMassSendNewsByOpenIds(msg *massbyopenids.News) (msgid int64, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.msgMassSendByOpenIds(msg)
}

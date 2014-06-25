package client

import (
	"errors"
	"github.com/chanxuehong/wechat/message/massbygroup"
)

// 根据分组群发消息, 之所以不暴露这个接口是因为怕接收到不合法的参数.
func (c *Client) msgMassSendByGroup(msg interface{}) (msgid int, err error) {
	token, err := c.Token()
	if err != nil {
		return
	}
	_url := messageMassSendByGroupURL(token)

	var result struct {
		Error
		MsgId int `json:"msg_id"`
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

// 根据分组群发文本消息.
func (c *Client) MsgMassSendTextByGroup(msg *massbygroup.Text) (msgid int, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.msgMassSendByGroup(msg)
}

// 根据分组群发图片消息.
func (c *Client) MsgMassSendImageByGroup(msg *massbygroup.Image) (msgid int, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.msgMassSendByGroup(msg)
}

// 根据分组群发语音消息.
func (c *Client) MsgMassSendVoiceByGroup(msg *massbygroup.Voice) (msgid int, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.msgMassSendByGroup(msg)
}

// 根据分组群发视频消息.
func (c *Client) MsgMassSendVideoByGroup(msg *massbygroup.Video) (msgid int, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.msgMassSendByGroup(msg)
}

// 根据分组群发图文消息.
func (c *Client) MsgMassSendNewsByGroup(msg *massbygroup.News) (msgid int, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.msgMassSendByGroup(msg)
}

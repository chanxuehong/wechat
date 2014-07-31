// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"fmt"
	"github.com/chanxuehong/wechat/message/massbyopenid"
)

// 根据用户列表群发文本消息.
func (c *Client) MsgMassSendTextByOpenId(toUser []string, content string) (msgid int64, err error) {
	if len(toUser) > massbyopenid.ToUserCountLimit {
		err = fmt.Errorf("用户列表的长度不能超过 %d, 现在为 %d", massbyopenid.ToUserCountLimit, len(toUser))
		return
	}

	var text massbyopenid.Text
	text.ToUser = toUser
	text.MsgType = massbyopenid.MSG_TYPE_TEXT
	text.Text.Content = content

	return c.msgMassSendByOpenId(&text)
}

// 根据用户列表群发图片消息.
func (c *Client) MsgMassSendImageByOpenId(toUser []string, mediaId string) (msgid int64, err error) {
	if len(toUser) > massbyopenid.ToUserCountLimit {
		err = fmt.Errorf("用户列表的长度不能超过 %d, 现在为 %d", massbyopenid.ToUserCountLimit, len(toUser))
		return
	}

	var image massbyopenid.Image
	image.ToUser = toUser
	image.MsgType = massbyopenid.MSG_TYPE_IMAGE
	image.Image.MediaId = mediaId

	return c.msgMassSendByOpenId(&image)
}

// 根据用户列表群发语音消息.
func (c *Client) MsgMassSendVoiceByOpenId(toUser []string, mediaId string) (msgid int64, err error) {
	if len(toUser) > massbyopenid.ToUserCountLimit {
		err = fmt.Errorf("用户列表的长度不能超过 %d, 现在为 %d", massbyopenid.ToUserCountLimit, len(toUser))
		return
	}

	var voice massbyopenid.Voice
	voice.ToUser = toUser
	voice.MsgType = massbyopenid.MSG_TYPE_VOICE
	voice.Voice.MediaId = mediaId

	return c.msgMassSendByOpenId(&voice)
}

// 根据用户列表群发视频消息.
//  title, description 可以为 ""
//  NOTE: mediaId 应该通过 Client.MediaCreateVideo 得到
func (c *Client) MsgMassSendVideoByOpenId(toUser []string, mediaId string,
	title, description string) (msgid int64, err error) {

	if len(toUser) > massbyopenid.ToUserCountLimit {
		err = fmt.Errorf("用户列表的长度不能超过 %d, 现在为 %d", massbyopenid.ToUserCountLimit, len(toUser))
		return
	}

	var video massbyopenid.Video
	video.ToUser = toUser
	video.MsgType = massbyopenid.MSG_TYPE_VIDEO
	video.Video.MediaId = mediaId
	video.Video.Title = title
	video.Video.Description = description

	return c.msgMassSendByOpenId(&video)
}

// 根据用户列表群发图文消息.
//  NOTE: mediaId 应该通过 Client.MediaCreateNews 得到
func (c *Client) MsgMassSendNewsByOpenId(toUser []string, mediaId string) (msgid int64, err error) {
	if len(toUser) > massbyopenid.ToUserCountLimit {
		err = fmt.Errorf("用户列表的长度不能超过 %d, 现在为 %d", massbyopenid.ToUserCountLimit, len(toUser))
		return
	}

	var news massbyopenid.News
	news.ToUser = toUser
	news.MsgType = massbyopenid.MSG_TYPE_NEWS
	news.News.MediaId = mediaId

	return c.msgMassSendByOpenId(&news)
}

// 根据 OpenId列表 群发消息, 之所以不暴露这个接口是因为怕接收到不合法的参数.
func (c *Client) msgMassSendByOpenId(msg interface{}) (msgid int64, err error) {
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
	_url := messageMassSendByOpenIdURL(token)
	if err = c.postJSON(_url, msg, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		msgid = result.MsgId
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

// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"github.com/chanxuehong/wechat/message/massbygroup"
)

// 根据分组群发文本消息.
func (c *Client) MsgMassSendTextByGroup(groupId int64, content string) (msgid int64, err error) {
	var text massbygroup.Text
	text.Filter.GroupId = groupId
	text.MsgType = massbygroup.MSG_TYPE_TEXT
	text.Text.Content = content

	return c.msgMassSendByGroup(&text)
}

// 根据分组群发图片消息.
func (c *Client) MsgMassSendImageByGroup(groupId int64, mediaId string) (msgid int64, err error) {
	var image massbygroup.Image
	image.Filter.GroupId = groupId
	image.MsgType = massbygroup.MSG_TYPE_IMAGE
	image.Image.MediaId = mediaId

	return c.msgMassSendByGroup(&image)
}

// 根据分组群发语音消息.
func (c *Client) MsgMassSendVoiceByGroup(groupId int64, mediaId string) (msgid int64, err error) {
	var voice massbygroup.Voice
	voice.Filter.GroupId = groupId
	voice.MsgType = massbygroup.MSG_TYPE_VOICE
	voice.Voice.MediaId = mediaId

	return c.msgMassSendByGroup(&voice)
}

// 根据分组群发视频消息.
//  NOTE: mediaId 应该通过 Client.MediaCreateVideo 得到
func (c *Client) MsgMassSendVideoByGroup(groupId int64, mediaId string) (msgid int64, err error) {
	var video massbygroup.Video
	video.Filter.GroupId = groupId
	video.MsgType = massbygroup.MSG_TYPE_VIDEO
	video.Video.MediaId = mediaId

	return c.msgMassSendByGroup(&video)
}

// 根据分组群发图文消息.
//  NOTE: mediaId 应该通过 Client.MediaCreateNews 得到
func (c *Client) MsgMassSendNewsByGroup(groupId int64, mediaId string) (msgid int64, err error) {
	var news massbygroup.News
	news.Filter.GroupId = groupId
	news.MsgType = massbygroup.MSG_TYPE_NEWS
	news.News.MediaId = mediaId

	return c.msgMassSendByGroup(&news)
}

// 根据分组群发消息, 之所以不暴露这个接口是因为怕接收到不合法的参数.
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
	_url := messageMassSendByGroupURL(token)
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

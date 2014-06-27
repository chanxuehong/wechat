// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong@gmail.com

package client

// 删除群发 =====================================================================
//  NOTE: 只有已经发送成功的消息才能删除删除消息只是将消息的图文详情页失效，已经收到的用户，
//  还是能在其本地看到消息卡片。 另外，删除群发消息只能删除图文消息和视频消息，
//  其他类型的消息一经发送，无法删除。
func (c *Client) MsgMassDelete(msgid int) (err error) {
	token, err := c.Token()
	if err != nil {
		return
	}
	_url := messageMassDeleteURL(token)

	var request struct {
		MsgId int `json:"msgid"`
	}
	request.MsgId = msgid

	var result Error
	if err = c.postJSON(_url, request, &result); err != nil {
		return
	}
	if result.ErrCode != 0 {
		return &result
	}
	return
}

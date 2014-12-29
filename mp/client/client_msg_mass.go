// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"github.com/chanxuehong/wechat/mp/message/active/mass"
)

// 删除群发.
//  NOTE: 只有已经发送成功的消息才能删除删除消息只是将消息的图文详情页失效，已经收到的用户，
//  还是能在其本地看到消息卡片。 另外，删除群发消息只能删除图文消息和视频消息，
//  其他类型的消息一经发送，无法删除。
func (c *Client) MsgMassDelete(msgid int64) (err error) {
	var request = struct {
		MsgId int64 `json:"msg_id"`
	}{
		MsgId: msgid,
	}

	var result Error

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url_ := messageMassDeleteURL(token)

	if err = c.postJSON(url_, request, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
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
		err = &result
		return
	}
}

// 查询群发消息发送状态
func (c *Client) MsgMassGetStatus(msgid int64) (status mass.MassStatus, err error) {
	var request = struct {
		MsgId int64 `json:"msg_id"`
	}{
		MsgId: msgid,
	}

	var result struct {
		Error
		mass.MassStatus
	}

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url_ := messageMassGetStatusURL(token)

	if err = c.postJSON(url_, request, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		status = result.MassStatus
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

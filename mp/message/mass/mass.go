// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package mass

import (
	"errors"

	"github.com/chanxuehong/wechat/mp"
)

// 群发消息给所有用户
func (clt *Client) MassToAll(msg interface{}) (msgid int64, err error) {
	if msg == nil {
		err = errors.New("nil msg")
		return
	}

	var result struct {
		mp.Error
		MsgId int64 `json:"msg_id"`
	}

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/message/mass/sendall?access_token="
	if err = ((*mp.Client)(clt)).PostJSON(incompleteURL, msg, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	msgid = result.MsgId
	return
}

// 群发消息给指定分组
func (clt *Client) MassToGroup(msg interface{}) (msgid int64, err error) {
	if msg == nil {
		err = errors.New("nil msg")
		return
	}

	var result struct {
		mp.Error
		MsgId int64 `json:"msg_id"`
	}

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/message/mass/sendall?access_token="
	if err = ((*mp.Client)(clt)).PostJSON(incompleteURL, msg, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	msgid = result.MsgId
	return
}

// 群发消息给指定用户列表
func (clt *Client) MassToUsers(msg interface{}) (msgid int64, err error) {
	if msg == nil {
		err = errors.New("nil msg")
		return
	}

	var result struct {
		mp.Error
		MsgId int64 `json:"msg_id"`
	}

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/message/mass/send?access_token="
	if err = ((*mp.Client)(clt)).PostJSON(incompleteURL, msg, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	msgid = result.MsgId
	return
}

// 预览消息
func (clt *Client) Preview(msg interface{}) (msgid int64, err error) {
	if msg == nil {
		err = errors.New("nil msg")
		return
	}

	var result struct {
		mp.Error
		MsgId int64 `json:"msg_id"`
	}

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/message/mass/preview?access_token="
	if err = ((*mp.Client)(clt)).PostJSON(incompleteURL, msg, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	msgid = result.MsgId
	return
}

// 删除群发.
func (clt *Client) DeleteMass(msgid int64) (err error) {
	var request = struct {
		MsgId int64 `json:"msg_id"`
	}{
		MsgId: msgid,
	}

	var result mp.Error

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/message/mass/delete?access_token="
	if err = ((*mp.Client)(clt)).PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result
		return
	}
	return
}

type MassStatus struct {
	MsgId  int64  `json:"msg_id"`
	Status string `json:"msg_status"` // 消息发送后的状态, SEND_SUCCESS表示发送成功
}

// 查询群发消息发送状态
func (clt *Client) GetMassStatus(msgid int64) (status *MassStatus, err error) {
	var request = struct {
		MsgId int64 `json:"msg_id"`
	}{
		MsgId: msgid,
	}

	var result struct {
		mp.Error
		MassStatus
	}

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/message/mass/get?access_token="
	if err = ((*mp.Client)(clt)).PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	status = &result.MassStatus
	return
}

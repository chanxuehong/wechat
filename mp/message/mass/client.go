// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package mass

import (
	"net/http"

	"github.com/chanxuehong/wechat/mp"
)

type Client struct {
	*mp.WechatClient
}

// 兼容保留, 建議實際項目全局維護一個 *mp.WechatClient
func NewClient(AccessTokenServer mp.AccessTokenServer, httpClient *http.Client) Client {
	return Client{
		WechatClient: mp.NewWechatClient(AccessTokenServer, httpClient),
	}
}

// 删除群发.
//  请注意:
//  只有已经发送成功的消息才能删除删除消息只是将消息的图文详情页失效，已经收到的用户，
//  还是能在其本地看到消息卡片。 另外，删除群发消息只能删除图文消息和视频消息，
//  其他类型的消息一经发送，无法删除。
func (clt Client) DeleteMass(msgid int64) (err error) {
	var request = struct {
		MsgId int64 `json:"msg_id"`
	}{
		MsgId: msgid,
	}

	var result mp.Error

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/message/mass/delete?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
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
	Status string `json:"msg_status"` // 消息发送后的状态，SEND_SUCCESS表示发送成功
}

// 查询群发消息发送状态
func (clt Client) GetMassStatus(msgid int64) (status *MassStatus, err error) {
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
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	status = &result.MassStatus
	return
}

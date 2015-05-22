// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package mass

type MassStatus struct {
	MsgId  int64  `json:"msg_id"`
	Status string `json:"msg_status"` // 消息发送后的状态，SEND_SUCCESS表示发送成功
}

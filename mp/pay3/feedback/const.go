// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package feedback

const (
	MSG_TYPE_COMPLAIN = "request" // 用户提交投诉消息
	MSG_TYPE_CONFIRM  = "confirm" // 用户确认消除投诉
	MSG_TYPE_REJECT   = "reject"  // 用户拒绝消除投诉
)

// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package server

func msgDispatch(agent Agent, msg *mixedRequest, para *RequestParameters) {
	switch msg.InfoType {
	case MsgTypeSuiteTicket:
		agent.ServeSuiteTicketMsg(msg.GetSuiteTicket(), para)

	case MsgTypeChangeAuth:
		agent.ServeChangeAuthMsg(msg.GetChangeAuth(), para)

	case MsgTypeCancelAuth:
		agent.ServeCancelAuthMsg(msg.GetCancelAuth(), para)

	default: // unknown message type
		agent.ServeUnknownMsg(para)
	}
}

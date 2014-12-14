// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package server

func msgDispatch(inputPara *InputParameters, msg *mixedRequest, agent Agent) {
	switch msg.InfoType {
	case MsgTypeSuiteTicket:
		agent.ServeSuiteTicketMsg(inputPara, msg.GetSuiteTicket())

	case MsgTypeChangeAuth:
		agent.ServeChangeAuthMsg(inputPara, msg.GetChangeAuth())

	case MsgTypeCancelAuth:
		agent.ServeCancelAuthMsg(inputPara, msg.GetCancelAuth())

	default: // unknown message type
		agent.ServeUnknownMsg(inputPara)
	}
}

// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"strconv"
)

// https://api.weixin.qq.com/pay/delivernotify?access_token=xxxxxx
func pay2DeliverNotifyURL(accesstoken string) string {
	return "https://api.weixin.qq.com/pay/delivernotify?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/pay/orderquery?access_token=xxxxxx
func pay2OrderQueryURL(accesstoken string) string {
	return "https://api.weixin.qq.com/pay/orderquery?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/payfeedback/update?access_token=xxxxx&openid=XXXX&feedbackid=xxxx
func pay2FeedbackUpdateURL(accesstoken, openid string, feedbackid int64) string {
	feedbackidStr := strconv.FormatInt(feedbackid, 10)
	return "https://api.weixin.qq.com/payfeedback/update?access_token=" +
		accesstoken + "&openid=" + openid + "&feedbackid=" + feedbackidStr
}

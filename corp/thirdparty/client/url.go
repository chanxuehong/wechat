// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

// https://qyapi.weixin.qq.com/cgi-bin/service/get_pre_auth_code?suite_access_token=xxx
func _GetPreAuthCodeURL(suiteAccessToken string) string {
	return "https://qyapi.weixin.qq.com/cgi-bin/service/get_pre_auth_code?suite_access_token=" +
		suiteAccessToken
}

// https://qyapi.weixin.qq.com/cgi-bin/service/get_permanent_code?suite_access_token=xxxx
func _GetPermanentCodeURL(suiteAccessToken string) string {
	return "https://qyapi.weixin.qq.com/cgi-bin/service/get_permanent_code?suite_access_token=" +
		suiteAccessToken
}

// https://qyapi.weixin.qq.com/cgi-bin/service/get_auth_info?suite_access_token=xxxx
func _GetAuthInfoURL(suiteAccessToken string) string {
	return "https://qyapi.weixin.qq.com/cgi-bin/service/get_auth_info?suite_access_token=" +
		suiteAccessToken
}

// https://qyapi.weixin.qq.com/cgi-bin/service/get_agent?suite_access_token=xxxx
func _GetAgentURL(suiteAccessToken string) string {
	return "https://qyapi.weixin.qq.com/cgi-bin/service/get_agent?suite_access_token=" +
		suiteAccessToken
}

// https://qyapi.weixin.qq.com/cgi-bin/service/set_agent?suite_access_token=xxxx
func _SetAgentURL(suiteAccessToken string) string {
	return "https://qyapi.weixin.qq.com/cgi-bin/service/set_agent?suite_access_token=" +
		suiteAccessToken
}

// https://qyapi.weixin.qq.com/cgi-bin/service/get_corp_token?suite_access_token=xxxx
func _GetCorpAccessToken(suiteAccessToken string) string {
	return "https://qyapi.weixin.qq.com/cgi-bin/service/get_corp_token?suite_access_token=" +
		suiteAccessToken
}

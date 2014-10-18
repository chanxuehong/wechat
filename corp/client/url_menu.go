// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"strconv"
)

// https://qyapi.weixin.qq.com/cgi-bin/menu/create?access_token=ACCESS_TOKEN&agentid=001
func _MenuCreateURL(accesstoken string, agentId int64) string {
	return "https://qyapi.weixin.qq.com/cgi-bin/menu/create?access_token=" + accesstoken +
		"&agentid=" + strconv.FormatInt(agentId, 10)
}

// https://qyapi.weixin.qq.com/cgi-bin/menu/delete?access_token=ACCESS_TOKEN&agentid=001
func _MenuDeleteURL(accesstoken string, agentId int64) string {
	return "https://qyapi.weixin.qq.com/cgi-bin/menu/delete?access_token=" + accesstoken +
		"&agentid=" + strconv.FormatInt(agentId, 10)
}

// https://qyapi.weixin.qq.com/cgi-bin/menu/get?access_token=ACCESS_TOKEN&agentid=001
func _MenuGetURL(accesstoken string, agentId int64) string {
	return "https://qyapi.weixin.qq.com/cgi-bin/menu/get?access_token=" + accesstoken +
		"&agentid=" + strconv.FormatInt(agentId, 10)
}

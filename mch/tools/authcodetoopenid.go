// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package tools

import (
	"github.com/chanxuehong/wechat/mch"
)

// 授权码查询OPENID接口.
func AuthCodeToOpenId(pxy *mch.Proxy, req map[string]string) (resp map[string]string, err error) {
	return pxy.PostXML("https://api.mch.weixin.qq.com/tools/authcodetoopenid", req)
}

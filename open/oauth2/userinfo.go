package oauth2

import (
	"net/http"

	mpoauth2 "github.com/chanxuehong/wechat/mp/oauth2"
)

const (
	LanguageZhCN = mpoauth2.LanguageZhCN
	LanguageZhTW = mpoauth2.LanguageZhTW
	LanguageEN   = mpoauth2.LanguageEN
)

const (
	SexUnknown = mpoauth2.SexUnknown
	SexMale    = mpoauth2.SexMale
	SexFemale  = mpoauth2.SexFemale
)

type UserInfo mpoauth2.UserInfo

// GetUserInfo 获取用户信息.
//  accessToken: 网页授权接口调用凭证
//  openId:      用户的唯一标识
//  lang:        返回国家地区语言版本，zh_CN 简体，zh_TW 繁体，en 英语, 如果留空 "" 则默认为 zh_CN
//  httpClient:  如果不指定则默认为 util.DefaultHttpClient
func GetUserInfo(accessToken, openId, lang string, httpClient *http.Client) (info *UserInfo, err error) {
	infox, err := mpoauth2.GetUserInfo(accessToken, openId, lang, httpClient)
	if err != nil {
		return
	}
	info = (*UserInfo)(infox)
	return
}

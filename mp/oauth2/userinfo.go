package oauth2

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/chanxuehong/wechat/internal/debug/api"
	"github.com/chanxuehong/wechat/oauth2"
	"github.com/chanxuehong/wechat/util"
)

const (
	LanguageZhCN = "zh_CN" // 简体中文
	LanguageZhTW = "zh_TW" // 繁体中文
	LanguageEN   = "en"    // 英文
)

const (
	SexUnknown = 0 // 未知
	SexMale    = 1 // 男性
	SexFemale  = 2 // 女性
)

type UserInfo struct {
	OpenId   string `json:"openid"`   // 用户的唯一标识
	Nickname string `json:"nickname"` // 用户昵称
	Sex      int    `json:"sex"`      // 用户的性别, 值为1时是男性, 值为2时是女性, 值为0时是未知
	City     string `json:"city"`     // 普通用户个人资料填写的城市
	Province string `json:"province"` // 用户个人资料填写的省份
	Country  string `json:"country"`  // 国家, 如中国为CN

	// 用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），
	// 用户没有头像时该项为空。若用户更换头像，原有头像URL将失效。
	HeadImageURL string `json:"headimgurl,omitempty"`

	Privilege []string `json:"privilege,omitempty"` // 用户特权信息，json 数组，如微信沃卡用户为（chinaunicom）
	UnionId   string   `json:"unionid,omitempty"`   // 只有在用户将公众号绑定到微信开放平台帐号后，才会出现该字段。
}

// GetUserInfo 获取用户信息.
//  accessToken: 网页授权接口调用凭证
//  openId:      用户的唯一标识
//  lang:        返回国家地区语言版本，zh_CN 简体，zh_TW 繁体，en 英语, 如果留空 "" 则默认为 zh_CN
//  httpClient:  如果不指定则默认为 util.DefaultHttpClient
func GetUserInfo(accessToken, openId, lang string, httpClient *http.Client) (info *UserInfo, err error) {
	switch lang {
	case "":
		lang = LanguageZhCN
	case LanguageZhCN, LanguageZhTW, LanguageEN:
	default:
		lang = LanguageZhCN
	}

	if httpClient == nil {
		httpClient = util.DefaultHttpClient
	}

	_url := "https://api.weixin.qq.com/sns/userinfo?access_token=" + url.QueryEscape(accessToken) +
		"&openid=" + url.QueryEscape(openId) +
		"&lang=" + lang
	api.DebugPrintGetRequest(_url)
	httpResp, err := httpClient.Get(_url)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		err = fmt.Errorf("http.Status: %s", httpResp.Status)
		return
	}

	var result struct {
		oauth2.Error
		UserInfo
	}
	if err = api.DecodeJSONHttpResponse(httpResp.Body, &result); err != nil {
		return
	}
	if result.ErrCode != oauth2.ErrCodeOK {
		err = &result.Error
		return
	}
	info = &result.UserInfo
	return
}

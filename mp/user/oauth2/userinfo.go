// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package oauth2

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/chanxuehong/wechat/mp"
)

const (
	Language_zh_CN = "zh_CN" // 简体中文
	Language_zh_TW = "zh_TW" // 繁体中文
	Language_en    = "en"    // 英文
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

	// 用户头像, 最后一个数值代表正方形头像大小(有0, 46, 64, 96, 132数值可选, 0代表640*640正方形头像),
	// 用户没有头像时该项为空
	HeadImageURL string `json:"headimgurl"`

	// 用户特权信息, json 数组, 如微信沃卡用户为(chinaunicom)
	Privilege []string `json:"privilege"`

	// 用户统一标识. 针对一个微信开放平台帐号下的应用, 同一用户的unionid是唯一的.
	UnionId string `json:"unionid"`
}

var ErrNoHeadImage = errors.New("没有头像")

// 获取用户图像的大小, 如果用户没有图像则返回 ErrNoHeadImage 错误.
func (info *UserInfo) HeadImageSize() (size int, err error) {
	HeadImageURL := info.HeadImageURL
	if HeadImageURL == "" {
		err = ErrNoHeadImage
		return
	}

	lastSlashIndex := strings.LastIndex(HeadImageURL, "/")
	if lastSlashIndex == -1 {
		err = fmt.Errorf("invalid HeadImageURL: %s", HeadImageURL)
		return
	}
	HeadImageIndex := lastSlashIndex + 1
	if HeadImageIndex == len(HeadImageURL) {
		err = fmt.Errorf("invalid HeadImageURL: %s", HeadImageURL)
		return
	}

	sizeStr := HeadImageURL[HeadImageIndex:]

	size, err = strconv.Atoi(sizeStr)
	if err != nil {
		err = fmt.Errorf("invalid HeadImageURL: %s", HeadImageURL)
		return
	}

	if size == 0 {
		size = 640
	}
	return
}

// 获取用户信息(需scope为 snsapi_userinfo).
//  lang 可能的取值是 zh_CN, zh_TW, en, 如果留空 "" 则默认为 zh_CN.
func (clt *Client) UserInfo(lang string) (info *UserInfo, err error) {
	switch lang {
	case "":
		lang = Language_zh_CN
	case Language_zh_CN, Language_zh_TW, Language_en:
	default:
		lang = Language_zh_CN
	}

	if clt.Config == nil {
		err = errors.New("nil Config")
		return
	}

	tk, err := clt.getToken()
	if err != nil {
		return
	}

	// 过期自动刷新 Token
	if tk.AccessTokenExpired() {
		if tk, err = clt.tokenRefresh(tk); err != nil {
			return
		}
	}

	var result struct {
		mp.Error
		UserInfo
	}
	if err = clt.getJSON(clt.Config.UserInfoURL(tk.AccessToken, tk.OpenId, lang), &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	info = &result.UserInfo
	return
}

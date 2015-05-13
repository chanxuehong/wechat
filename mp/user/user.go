// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package user

import (
	"errors"
	"fmt"
	"net/url"
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
	OpenId   string `json:"openid"`   // 用户的标识，对当前公众号唯一
	Nickname string `json:"nickname"` // 用户的昵称
	Sex      int    `json:"sex"`      // 用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
	Language string `json:"language"` // 用户的语言，zh_CN，zh_TW，en
	City     string `json:"city"`     // 用户所在城市
	Province string `json:"province"` // 用户所在省份
	Country  string `json:"country"`  // 用户所在国家

	// 用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），
	// 用户没有头像时该项为空
	HeadImageURL string `json:"headimgurl,omitempty"`

	// 用户关注时间，为时间戳。如果用户曾多次关注，则取最后关注时间
	SubscribeTime int64 `json:"subscribe_time"`

	// 只有在用户将公众号绑定到微信开放平台帐号后，才会出现该字段。
	UnionId string `json:"unionid,omitempty"`

	// 备注名
	Remark string `json:"remark,omitempty"`
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

	size64, err := strconv.ParseUint(sizeStr, 10, 64)
	if err != nil {
		err = fmt.Errorf("invalid HeadImageURL: %s", HeadImageURL)
		return
	}

	if size64 == 0 {
		size64 = 640
	}
	size = int(size64)
	return
}

var ErrUserNotSubscriber = errors.New("用户没有订阅公众号")

// 获取用户基本信息, 如果用户没有订阅公众号, 返回 ErrUserNotSubscriber 错误.
//  lang 可以是 zh_CN, zh_TW, en, 如果留空 "" 则默认为 zh_CN.
func (clt Client) UserInfo(openId string, lang string) (userinfo *UserInfo, err error) {
	if openId == "" {
		err = errors.New("empty openId")
		return
	}

	switch lang {
	case "":
		lang = Language_zh_CN
	case Language_zh_CN, Language_zh_TW, Language_en:
	default:
		err = errors.New("invalid lang: " + lang)
		return
	}

	var result struct {
		mp.Error
		Subscribed int `json:"subscribe"` // 用户是否订阅该公众号标识，值为0时，代表此用户没有关注该公众号，拉取不到其余信息。
		UserInfo
	}

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/user/info?openid=" + url.QueryEscape(openId) +
		"&lang=" + url.QueryEscape(lang) + "&access_token="
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	if result.Subscribed == 0 {
		err = ErrUserNotSubscriber
		return
	}
	userinfo = &result.UserInfo
	return
}

// 开发者可以通过该接口对指定用户设置备注名.
//  NOTE: 该接口暂时开放给微信认证的服务号.
func (clt Client) UserUpdateRemark(openId, remark string) (err error) {
	var request = struct {
		OpenId string `json:"openid"`
		Remark string `json:"remark"`
	}{
		OpenId: openId,
		Remark: remark,
	}

	var result mp.Error

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/user/info/updateremark?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result
		return
	}
	return
}

// 获取关注者列表返回的数据结构
type UserListResult struct {
	TotalCount int `json:"total"` // 关注该公众账号的总用户数
	GotCount   int `json:"count"` // 拉取的OPENID个数，最大值为10000

	Data struct {
		OpenId []string `json:"openid,omitempty"`
	} `json:"data"` // 列表数据，OPENID的列表

	// 拉取列表的后一个用户的OPENID, 如果 next_openid == "" 则表示没有了用户数据
	NextOpenId string `json:"next_openid"`
}

// 获取关注者列表, 每次最多能获取 10000 个用户, 如果 beginOpenId == "" 则表示从头获取
func (clt Client) UserList(beginOpenId string) (data *UserListResult, err error) {
	var result struct {
		mp.Error
		UserListResult
	}

	var incompleteURL string
	if beginOpenId == "" {
		incompleteURL = "https://api.weixin.qq.com/cgi-bin/user/get?access_token="
	} else {
		incompleteURL = "https://api.weixin.qq.com/cgi-bin/user/get?next_openid=" +
			url.QueryEscape(beginOpenId) + "&access_token="
	}

	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	data = &result.UserListResult
	return
}

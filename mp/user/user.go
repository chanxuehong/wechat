package user

import (
	"net/url"

	"gopkg.in/chanxuehong/wechat.v2/mp/core"
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
	IsSubscriber int    `json:"subscribe"` // 用户是否订阅该公众号标识, 值为0时, 代表此用户没有关注该公众号, 拉取不到其余信息
	OpenId       string `json:"openid"`    // 用户的标识, 对当前公众号唯一
	Nickname     string `json:"nickname"`  // 用户的昵称
	Sex          int    `json:"sex"`       // 用户的性别, 值为1时是男性, 值为2时是女性, 值为0时是未知
	Language     string `json:"language"`  // 用户的语言, zh_CN, zh_TW, en
	City         string `json:"city"`      // 用户所在城市
	Province     string `json:"province"`  // 用户所在省份
	Country      string `json:"country"`   // 用户所在国家

	// 用户头像, 最后一个数值代表正方形头像大小(有0, 46, 64, 96, 132数值可选, 0代表640*640正方形头像), 用户没有头像时该项为空
	HeadImageURL string `json:"headimgurl"`

	SubscribeTime int64  `json:"subscribe_time"`    // 用户关注时间, 为时间戳. 如果用户曾多次关注, 则取最后关注时间
	UnionId       string `json:"unionid,omitempty"` // 只有在用户将公众号绑定到微信开放平台帐号后, 才会出现该字段.
	Remark        string `json:"remark"`            // 公众号运营者对粉丝的备注, 公众号运营者可在微信公众平台用户管理界面对粉丝添加备注
	GroupId       int64  `json:"groupid"`           // 用户所在的分组ID

	TagIdList      []int  `json:"tagid_list"`      // Tag List
	SubscribeScene string `json:"subscribe_scene"` // 返回用户关注的渠道来源
	QrScene        int    `json:"qr_scene"`        // 二维码扫码场景（开发者自定义）场景值ID，临时二维码时为32位非0整型，永久二维码时最大值为100000（目前参数只支持1--100000）
	QrSceneStr     string `json:"qr_scene_str"`    // 二维码扫码场景描述（开发者自定义）场景值ID（字符串形式的ID），字符串类型，长度限制为1到64
}

// Get 获取用户基本信息.
//  注意:
//  1. 需要判断返回的 UserInfo.IsSubscriber 是等于 1 还是 0
//  2. lang 指定返回国家地区语言版本，zh_CN 简体，zh_TW 繁体，en 英语, 默认为 zh_CN
func Get(clt *core.Client, openId string, lang string) (info *UserInfo, err error) {
	switch lang {
	case "":
		lang = LanguageZhCN
	case LanguageZhCN, LanguageZhTW, LanguageEN:
	default:
		lang = LanguageZhCN
	}

	var incompleteURL = "https://api.weixin.qq.com/cgi-bin/user/info?openid=" + url.QueryEscape(openId) +
		"&lang=" + lang + "&access_token="
	var result struct {
		core.Error
		UserInfo
	}
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	info = &result.UserInfo
	return
}

type batchGetRequestItem struct {
	OpenId   string `json:"openid"`
	Language string `json:"lang,omitempty"`
}

// 批量获取用户基本信息
//  注意: 需要对返回的 UserInfoList 的每个 UserInfo.IsSubscriber 做判断
func BatchGet(clt *core.Client, openIdList []string, lang string) (list []UserInfo, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/user/info/batchget?access_token="

	if len(openIdList) <= 0 {
		return
	}

	switch lang {
	case "", LanguageZhCN, LanguageZhTW, LanguageEN:
	default:
		lang = ""
	}

	var request struct {
		UserList []batchGetRequestItem `json:"user_list,omitempty"`
	}
	request.UserList = make([]batchGetRequestItem, len(openIdList))
	for i := 0; i < len(openIdList); i++ {
		request.UserList[i].OpenId = openIdList[i]
		request.UserList[i].Language = lang
	}

	var result struct {
		core.Error
		UserInfoList []UserInfo `json:"user_info_list"`
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	list = result.UserInfoList
	return
}

// UpdateRemark 设置用户备注名.
func UpdateRemark(clt *core.Client, openId, remark string) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/user/info/updateremark?access_token="

	var request = struct {
		OpenId string `json:"openid"`
		Remark string `json:"remark"`
	}{
		OpenId: openId,
		Remark: remark,
	}
	var result core.Error
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	return
}

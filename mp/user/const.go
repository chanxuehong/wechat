// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package user

const (
	GroupCountLimit   = 500   // 每个公众号分组个数不能超过 500
	UserPageSizeLimit = 10000 // 每次拉取的OPENID个数最大值为10000
)

const (
	Language_zh_CN = "zh_CN" // 简体中文
	Language_zh_TW = "zh_TW" // 繁体中文
	Language_en    = "en"    // 英文
)

const (
	SEX_UNKNOWN = 0 // 未知
	SEX_MALE    = 1 // 男性
	SEX_FEMALE  = 2 // 女性
)

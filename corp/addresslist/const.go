// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package addresslist

const (
	USER_GENDER_MALE   = 0 // 男性
	USER_GENDER_FEMALE = 1 // 女性

	USER_STATUS_SUBSCRIBED   = 1 // 已关注
	USER_STATUS_BLOCKED      = 2 // 已冻结
	USER_STATUS_NOSUBSCRIBED = 4 // 未关注

	// 获取部门成员 status 的参数取值, 可以是下面的任何一个或者后面三个的叠加
	USERSIMPLELIST_STATUS_ALL          = 0 // 全部员工
	USERSIMPLELIST_STATUS_SUBSCRIBED   = 1 // 已关注成员
	USERSIMPLELIST_STATUS_DISABLED     = 2 // 禁用成员
	USERSIMPLELIST_STATUS_NOSUBSCRIBED = 4 // 未关注成员
)

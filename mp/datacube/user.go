// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     magicshui(shuiyuzhe@gmail.com)

package datacube

type UserSummaryData struct {
	RefDate    string `json:"ref_date"`    // 数据日期
	UserSource int    `json:"user_source"` // 用户渠道，0代表其他 30代表扫二维码 17代表名片分享 35代表搜号码（即微信添加朋友页的搜索） 39代表查询微信公众帐号 43代表图文页右上角菜单
	NewUser    int    `json:"new_user"`    // 新增用户
	CancelUser int    `json:"cancel_user"` // 取消关注的用户
}

type UserCumulateData struct {
	RefDate    string `json:"ref_date"`    // 数据日期
	UserSource int    `json:"user_source"` // 用户渠道，0代表其他 30代表扫二维码 17代表名片分享 35代表搜号码（即微信添加朋友页的搜索） 39代表查询微信公众帐号 43代表图文页右上角菜单
	Cumulate   int    `json:"cumulate_user"`
}

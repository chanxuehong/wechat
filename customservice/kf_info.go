// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package customservice

// 客服基本信息
type KFInfo struct {
	Id       string `json:"kf_id"`      // 客服工号
	Account  string `json:"kf_account"` // 客服账号@微信别名; 微信别名如有修改，旧账号返回旧的微信别名，新增的账号返回新的微信别名
	Nickname string `json:"kf_nick"`    // 客服昵称
}

// 在线客服接待信息
type OnlineKFInfo struct {
	Id                  string `json:"kf_id"`         // 客服工号
	Account             string `json:"kf_account"`    // 客服账号@微信别名; 微信别名如有修改，旧账号返回旧的微信别名，新增的账号返回新的微信别名
	Status              int    `json:"status"`        // 客服在线状态 1：pc在线，2：手机在线, 若pc和手机同时在线则为 1+2=3
	AutoAcceptThreshold int    `json:"auto_accept"`   // 客服设置的最大自动接入数
	AcceptingNumber     int    `json:"accepted_case"` // 客服当前正在接待的会话数
}

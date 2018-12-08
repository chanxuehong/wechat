// 多客服会话控制
package session

import (
	"net/url"

	"github.com/chanxuehong/wechat/mp/core"
)

// Create 创建会话.
//  openId:    必须, 客户openid
//  kfAccount: 必须, 完整客服账号，格式为：账号前缀@公众号微信号
//  text:      可选, 附加信息，文本会展示在客服人员的多客服客户端
func Create(clt *core.Client, openId, kfAccount, text string) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/customservice/kfsession/create?access_token="

	request := struct {
		KfAccount string `json:"kf_account"`
		OpenId    string `json:"openid"`
		Text      string `json:"text,omitempty"`
	}{
		KfAccount: kfAccount,
		OpenId:    openId,
		Text:      text,
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

// Close 关闭会话.
//  openId:    必须, 客户openid
//  kfAccount: 必须, 完整客服账号，格式为：账号前缀@公众号微信号
//  text:      可选, 附加信息，文本会展示在客服人员的多客服客户端
func Close(clt *core.Client, openId, kfAccount, text string) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/customservice/kfsession/close?access_token="

	request := struct {
		KfAccount string `json:"kf_account"`
		OpenId    string `json:"openid"`
		Text      string `json:"text,omitempty"`
	}{
		KfAccount: kfAccount,
		OpenId:    openId,
		Text:      text,
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

type Session struct {
	OpenId     string `json:"openid"`     // 客户openid
	KfAccount  string `json:"kf_account"` // 正在接待的客服，为空表示没有人在接待
	CreateTime int64  `json:"createtime"` // 会话接入的时间
}

// Get 获取客户的会话
func Get(clt *core.Client, openId string) (ss *Session, err error) {
	incompleteURL := "https://api.weixin.qq.com/customservice/kfsession/getsession?openid=" +
		url.QueryEscape(openId) + "&access_token="

	var result struct {
		core.Error
		Session
	}
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	result.Session.OpenId = openId
	ss = &result.Session
	return
}

// List 获取客服的会话列表, 开发者可以通过本接口获取某个客服正在接待的会话列表.
func List(clt *core.Client, kfAccount string) (list []Session, err error) {
	// TODO
	//	incompleteURL := "https://api.weixin.qq.com/customservice/kfsession/getsessionlist?kf_account=" +
	//		url.QueryEscape(kfAccount) + "&access_token="
	incompleteURL := "https://api.weixin.qq.com/customservice/kfsession/getsessionlist?kf_account=" +
		kfAccount + "&access_token="

	var result struct {
		core.Error
		SessionList []Session `json:"sessionlist"`
	}
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	for i := 0; i < len(result.SessionList); i++ {
		result.SessionList[i].KfAccount = kfAccount
	}
	list = result.SessionList
	return
}

type WaitCaseListResult struct {
	TotalCount int       `json:"count"`                  // 未接入会话数量
	ItemCount  int       `json:"item_count"`             // 本次返回的未接入会话列表数量
	Items      []Session `json:"waitcaselist,omitempty"` // 本次返回的未接入会话列表
}

// WaitCaseList 获取未接入会话列表.
func WaitCaseList(clt *core.Client) (rslt *WaitCaseListResult, err error) {
	const incompleteURL = "https://api.weixin.qq.com/customservice/kfsession/getwaitcase?access_token="

	var result struct {
		core.Error
		WaitCaseListResult
	}
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	result.WaitCaseListResult.ItemCount = len(result.WaitCaseListResult.Items)
	rslt = &result.WaitCaseListResult
	return
}

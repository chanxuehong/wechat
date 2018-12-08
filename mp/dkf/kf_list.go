package dkf

import (
	"encoding/json"

	"github.com/chanxuehong/wechat/mp/core"
)

// 客服基本信息
type KfInfo struct {
	Id           json.Number `json:"kf_id"`         // 客服工号
	Account      string      `json:"kf_account"`    // 完整客服账号，格式为：账号前缀@公众号微信号
	Nickname     string      `json:"kf_nick"`       // 客服昵称
	HeadImageURL string      `json:"kf_headimgurl"` // 客服头像
}

// KfList 获取客服基本信息.
func KfList(clt *core.Client) (list []KfInfo, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/customservice/getkflist?access_token="

	var result struct {
		core.Error
		KfList []KfInfo `json:"kf_list"`
	}
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	list = result.KfList
	return
}

const (
	OnlineKfInfoStatusPC          = 1
	OnlineKfInfoStatusMobile      = 2
	OnlineKfInfoStatusPCAndMobile = 3
)

// 在线客服接待信息
type OnlineKfInfo struct {
	Id               json.Number `json:"kf_id"`         // 客服工号
	Account          string      `json:"kf_account"`    // 完整客服账号，格式为：账号前缀@公众号微信号
	Status           int         `json:"status"`        // 客服在线状态 1：pc在线，2：手机在线。若pc和手机同时在线则为 1+2=3
	AutoAcceptNumber int         `json:"auto_accept"`   // 客服设置的最大自动接入数
	AcceptingNumber  int         `json:"accepted_case"` // 客服当前正在接待的会话数
}

// OnlineKfList 获取在线客服接待信息.
func OnlineKfList(clt *core.Client) (list []OnlineKfInfo, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/customservice/getonlinekflist?access_token="

	var result struct {
		core.Error
		OnlineKfInfoList []OnlineKfInfo `json:"kf_online_list"`
	}
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	list = result.OnlineKfInfoList
	return
}

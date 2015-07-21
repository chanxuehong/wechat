// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package provider

import (
	"github.com/chanxuehong/wechat/corp"
)

type CorpInfo struct {
	CorpId            string `json:"corpid"`
	CorpName          string `json:"corp_name"`
	CorpType          string `json:"corp_type"`
	CorpRoundLogoURL  string `json:"corp_round_logo_url"`
	CorpSquareLogoURL string `json:"corp_square_logo_url"`
	CorpUserMax       int64  `json:"corp_user_max"`
	CorpAgentMax      int64  `json:"corp_agent_max"`
}

type AuthInfo struct {
	DepartmentList []AuthInfoDepartment `json:"department,omitempty"`
}

type Agent struct {
	AgentId  int64   `json:"agentid"`
	AuthType int64   `json:"auth_type"`
}

type AuthInfoDepartment struct {
	Id       int64  `json:"id"`
	Writable bool   `json:"writable"`
}


type UserInfo struct {
	Email  string       `json:"email"`
	Userid string       `json:"userid"`
	Name   string       `json:"name"`
	Avatar string       `json:"avatar"`
	Mobile string       `json:"mobile"`
}


type LoginInfo struct {
	IsInner   bool       `json:"is_inner"`
	IsSys     bool       `json:"is_sys"`
	UserInfo  UserInfo `json:"user_info"`
	CorpInfo  CorpInfo `json:"corp_info"`
	AgentList []Agent  `json:"agent,omitempty"`
	AuthInfo  AuthInfo     `json:"auth_info"`
}


// 获取企业号管理员登录信息
//  authCode:  oauth2.0授权企业号管理员登录产生的code .
func (clt *Client) GetLoginInfo(authCode string) (info *LoginInfo, err error) {
	request := struct {
		AuthCode string `json:"auth_code"`
	}{
		AuthCode: authCode,
	}

	var result struct {
		corp.Error
		LoginInfo
	}

	incompleteURL := "https://qyapi.weixin.qq.com/cgi-bin/service/get_login_info?provider_access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != corp.ErrCodeOK {
		err = &result.Error
		return
	}
	info = &result.LoginInfo
	return
}

// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package thirdparty

import (
	"github.com/chanxuehong/wechat/corp"
)

// 获取企业号的授权信息
//  AuthCorpId:    授权方corpid
//  PermanentCode: 永久授权码，通过get_permanent_code获取
func (clt *SuiteClient) GetAuthInfo(AuthCorpId, PermanentCode string) (info *AuthInfo, err error) {
	request := struct {
		SuiteId       string `json:"suite_id"`
		AuthCorpId    string `json:"auth_corpid"`
		PermanentCode string `json:"permanent_code"`
	}{
		SuiteId:       clt.SuiteId,
		AuthCorpId:    AuthCorpId,
		PermanentCode: PermanentCode,
	}

	var result struct {
		corp.Error
		AuthInfo
	}

	incompleteURL := "https://qyapi.weixin.qq.com/cgi-bin/service/get_agent?suite_access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != corp.ErrCodeOK {
		err = &result.Error
		return
	}
	info = &result.AuthInfo
	return
}

type AuthInfo struct {
	AuthCorpInfo struct {
		CorpId            string `json:"corpid"`
		CorpName          string `json:"corp_name"`
		CorpType          string `json:"corp_type"`
		CorpRoundLogoURL  string `json:"corp_round_logo_url"`
		CorpSquareLogoURL string `json:"corp_square_logo_url"`
		CorpUserMax       int64  `json:"corp_user_max"`
		CorpAgentMax      int64  `json:"corp_agent_max"`
		CorpWxQrCode      string `json:"corp_wxqrcode"`
	} `json:"auth_corp_info"`

	AuthInfo struct {
		Agent      []AuthInfoAgent      `json:"agent,omitempty"`
		Department []AuthInfoDepartment `json:"department,omitempty"`
	} `json:"auth_info"`
}

type AuthInfoAgent struct {
	AgentId       int64    `json:"agentid"`
	Name          string   `json:"name"`
	RoundLogoURL  string   `json:"round_logo_url"`
	SquareLogoURL string   `json:"square_logo_url"`
	AppId         int64    `json:"app_id"`
	APIGroup      []string `json:"api_group,omitempty"`
}

type AuthInfoDepartment struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	ParentId int64  `json:"parentid"`
	Writable bool   `json:"writable,omitempty"`
}

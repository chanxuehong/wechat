// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"encoding/json"
)

type AgentInfo struct {
	AgentId        json.Number `json:"agentid"`
	Name           string      `json:"name"`
	SquareLogoURL  string      `json:"square_logo_url"`
	RoundLogoURL   string      `json:"round_logo_url"`
	Description    string      `json:"description"`
	AllowUserInfos struct {
		Users []struct {
			Userid string      `json:"userid"`
			Status json.Number `json:"status"`
		} `json:"user,omitempty"`
	} `json:"allow_userinfos"`
	AllowParties struct {
		PartyIds []int64 `json:"partyid,omitempty"`
	} `json:"allow_partys"`
	AllowTags struct {
		TagIds []int64 `json:"tagid,omitempty"`
	} `json:"allow_tags"`
	Closed             json.Number `json:"close"`
	RedirectDomain     string      `json:"redirect_domain"`
	ReportLocationFlag json.Number `json:"report_location_flag"`
	IsReportUser       json.Number `json:"isreportuser"`
}

// 获取企业号应用
func (c *SuiteClient) GetAgent(AuthCorpId, PermanentCode string, AgentId int64) (info *AgentInfo, err error) {
	request := struct {
		SuiteId       string `json:"suite_id"`
		AuthCorpId    string `json:"auth_corpid"`
		PermanentCode string `json:"permanent_code"`
		AgentId       int64  `json:"agentid,string"`
	}{
		SuiteId:       c.suiteId,
		AuthCorpId:    AuthCorpId,
		PermanentCode: PermanentCode,
		AgentId:       AgentId,
	}

	var result struct {
		Error
		AgentInfo
	}

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url_ := _GetAgentURL(token)

	if err = c.postJSON(url_, &request, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		info = &result.AgentInfo
		return
	case errCodeInvalidCredential, errCodeTimeout:
		if !hasRetry {
			hasRetry = true

			if token, err = getNewToken(c.tokenService, token); err != nil {
				return
			}
			goto RETRY
		}
		fallthrough
	default:
		err = &result.Error
		return
	}
}

// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"errors"
)

type AgentParameters struct {
	AgentId            int64   `json:"agentid,string"`
	ReportLocationFlag *int    `json:"report_location_flag,string,omitempty"`
	LogoMediaId        *string `json:"logo_mediaid,omitempty"`
	Name               *string `json:"name,omitempty"`
	Description        *string `json:"description,omitempty"`
	RedirectDomain     *string `json:"redirect_domain,omitempty"`
	IsReportUser       *int    `json:"isreportuser,omitempty"`
}

func Int(v int) *int {
	return &v
}

func String(v string) *string {
	return &v
}

func (c *SuiteClient) SetAgent(AuthCorpId, PermanentCode string, para *AgentParameters) (err error) {
	if para == nil {
		return errors.New("para == nil")
	}

	request := struct {
		SuiteId       string           `json:"suite_id"`
		AuthCorpId    string           `json:"auth_corpid"`
		PermanentCode string           `json:"permanent_code"`
		Agent         *AgentParameters `json:"agent,omitempty"`
	}{
		SuiteId:       c.suiteId,
		AuthCorpId:    AuthCorpId,
		PermanentCode: PermanentCode,
		Agent:         para,
	}

	var result Error

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url_ := _SetAgentURL(token)

	if err = c.postJSON(url_, &request, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		return
	case errCodeExpired:
		if !hasRetry {
			hasRetry = true

			if token, err = getNewToken(c.tokenService, token); err != nil {
				return
			}
			goto RETRY
		}
		fallthrough
	default:
		err = &result
		return
	}
}

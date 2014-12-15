// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// 获取企业号的授权信息
//  AuthCorpId:    授权方corpid
//  PermanentCode: 永久授权码，通过get_permanent_code获取
func (c *SuiteClient) GetAuthInfo(AuthCorpId, PermanentCode string) (info *AuthInfo, err error) {
	request := struct {
		SuiteId       string `json:"suite_id"`
		AuthCorpId    string `json:"auth_corpid"`
		PermanentCode string `json:"permanent_code"`
	}{
		SuiteId:       c.suiteId,
		AuthCorpId:    AuthCorpId,
		PermanentCode: PermanentCode,
	}

	var result struct {
		Error
		AuthInfo
	}

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url_ := _GetAuthInfoURL(token)

	if err = c.postJSON(url_, &request, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		info = &result.AuthInfo
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
		err = &result.Error
		return
	}
}

type AuthInfo struct {
	AuthCorpInfo struct {
		CorpId            string      `json:"corpid"`
		CorpName          string      `json:"corp_name"`
		CorpType          string      `json:"corp_type"`
		CorpRoundLogoURL  string      `json:"corp_round_logo_url"`
		CorpSquareLogoURL string      `json:"corp_square_logo_url"`
		CorpUserMax       json.Number `json:"corp_user_max"`
		CorpAgentMax      json.Number `json:"corp_agent_max"`
	} `json:"auth_corp_info"`

	AuthInfo struct {
		Agent      []AuthInfoAgent      `json:"agent,omitempty"`
		Department []AuthInfoDepartment `json:"department,omitempty"`
	} `json:"auth_info"`
}

type AuthInfoAgent struct {
	AgentId       json.Number `json:"agentid"`
	Name          string      `json:"name"`
	RoundLogoURL  string      `json:"round_logo_url"`
	SquareLogoURL string      `json:"square_logo_url"`
	AppId         json.Number `json:"app_id"`
	APIGroup      []string    `json:"api_group,omitempty"`
}

type AuthInfoDepartment struct {
	Id            json.Number     `json:"id"`
	Name          string          `json:"name"`
	ParentId      json.Number     `json:"parentid"`
	WritableBytes json.RawMessage `json:"writable,omitempty"`
}

var (
	json_false_bytes        = []byte(`false`)
	json_false_bytes_quoted = []byte(`"false"`)
	json_true_bytes         = []byte(`true`)
	json_true_bytes_quoted  = []byte(`"true"`)
)

func (this *AuthInfoDepartment) Writable() (b bool, err error) {
	src := this.WritableBytes

	if bytes.Equal(src, json_true_bytes_quoted) ||
		bytes.Equal(src, json_true_bytes) {
		b = true
		return
	}
	if bytes.Equal(src, json_false_bytes_quoted) ||
		bytes.Equal(src, json_false_bytes) {
		return
	}
	err = fmt.Errorf("invalid writable: %q", src)
	return
}

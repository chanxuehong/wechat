// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"github.com/chanxuehong/wechat/corp/client"
)

type CorpAccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

var _ client.TokenGetter = new(CorpAccessTokenGetter)

type CorpAccessTokenGetter struct {
	SuiteClient   *SuiteClient
	AuthCorpId    string
	PermanentCode string
}

func (getter *CorpAccessTokenGetter) GetToken() (tk string, err error) {
	request := struct {
		SuiteId       string `json:"suite_id"`
		AuthCorpId    string `json:"auth_corpid"`
		PermanentCode string `json:"permanent_code"`
	}{
		SuiteId:       getter.SuiteClient.suiteId,
		AuthCorpId:    getter.AuthCorpId,
		PermanentCode: getter.PermanentCode,
	}

	var result struct {
		Error
		CorpAccessToken
	}

	token, err := getter.SuiteClient.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url_ := _GetCorpAccessToken(token)

	if err = getter.SuiteClient.postJSON(url_, &request, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		tk = result.CorpAccessToken.AccessToken
		return
	case errCodeInvalidCredential, errCodeTimeout:
		if !hasRetry {
			hasRetry = true

			if token, err = getNewToken(getter.SuiteClient.tokenService, token); err != nil {
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

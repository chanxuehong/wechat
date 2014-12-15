// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

type GetPermanentCodeResponse struct {
	PermanentCode string `json:"permanent_code"`
	CorpAccessToken
	AuthInfo
}

// 获取企业号的永久授权码
//  AuthCode: 临时授权码会在授权成功时附加在redirect_uri中跳转回应用提供商网站。
func (c *SuiteClient) GetPermanentCode(AuthCode string) (resp *GetPermanentCodeResponse, err error) {
	request := struct {
		SuiteId  string `json:"suite_id"`
		AuthCode string `json:"auth_code"`
	}{
		SuiteId:  c.suiteId,
		AuthCode: AuthCode,
	}

	var result struct {
		Error
		GetPermanentCodeResponse
	}

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url_ := _GetPermanentCodeURL(token)

	if err = c.postJSON(url_, &request, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		resp = &result.GetPermanentCodeResponse
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

// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

type PreAuthCode struct {
	PreAuthCode string `json:"pre_auth_code"`
	ExpiresIn   int64  `json:"expires_in"`
}

// 获取预授权码
//  AppId: 表示用户能对本套件内的哪些应用授权, AppId == nil 时默认用户有全部授权权限
func (c *SuiteClient) GetPreAuthCode(AppId []int64) (code *PreAuthCode, err error) {
	request := struct {
		SuiteId string  `json:"suite_id"`
		AppId   []int64 `json:"appid,omitempty"`
	}{
		SuiteId: c.suiteId,
		AppId:   AppId,
	}

	var result struct {
		Error
		PreAuthCode
	}

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url_ := _GetPreAuthCodeURL(token)

	if err = c.postJSON(url_, &request, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		code = &result.PreAuthCode
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

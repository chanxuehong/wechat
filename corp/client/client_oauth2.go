// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

// 根据code获取成员信息
//  agentid: 跳转链接时所在的企业应用ID
//  code:    通过员工授权获取到的code，每次员工授权带上的code将不一样，code只能使用一次，5分钟未被使用自动过期
func (c *Client) OAuth2GetUserInfo(agentid, code string) (userid string, err error) {
	var result struct {
		Error
		UserId string `json:"USERID"`
	}

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	if err = c.getJSON(_OAuth2GetUserInfoURL(token, code, agentid), &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		userid = result.UserId
		return

	case errCodeTimeout:
		if !hasRetry {
			hasRetry = true

			if token, err = c.TokenRefresh(); err != nil {
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

// 企业在开启二次验证时，必须填写企业二次验证页面的url。当员工绑定通讯录中的帐号后，会收到一条图文消息，
// 引导员工到企业的验证页面验证身份。在跳转到企业的验证页面时，会带上如下参数：code=CODE&state=STATE，
// 企业可以调用oauth2接口，根据code获取员工的userid。
//
// 企业在员工验证成功后，调用这个方法即可让员工关注成功。
func (c *Client) OAuth2UserAuthSuccessfully(userid string) (err error) {
	var result Error

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	if err = c.getJSON(_OAuth2UserAuthSuccessfullyURL(token, userid), &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		return

	case errCodeTimeout:
		if !hasRetry {
			hasRetry = true

			if token, err = c.TokenRefresh(); err != nil {
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

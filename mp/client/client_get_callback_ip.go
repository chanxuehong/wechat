// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

// 如果公众号基于安全等考虑，需要获知微信服务器的IP地址列表，以便进行相关限制，
// 可以通过该接口获得微信服务器IP地址列表。
func (c *Client) GetCallbackIP() (ipList []string, err error) {
	var result struct {
		Error
		IPList []string `json:"ip_list"`
	}

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	url_ := getCallbackIPURL(token)

	if err = c.getJSON(url_, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		ipList = result.IPList
		return

	case errCodeInvalidCredential, errCodeTimeout:
		if !hasRetry {
			hasRetry = true
			timeoutRetryWait()
			goto RETRY
		}
		fallthrough

	default:
		err = &result.Error
		return
	}
}

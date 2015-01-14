// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     magicshui(shuiyuzhe@gmail.com)

package client

func (c *Client) TicketGetTicket() (ticket string, err error) {
	var result struct {
		Error
		Ticket    string `json:"ticket"`
		ExpiresIn int    `json:"expires_in"`
	}
	token, err := c.Token()
	if err != nil {
		return
	}
	hasRetry := false
RETRY:
	url := ticketGetTicketURL(token)
	if err = c.getJSON(url, &result); err != nil {
		return
	}
	switch result.ErrCode {
	case errCodeOK:
		ticket = result.Ticket
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

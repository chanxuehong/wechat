package wechat

import (
	"errors"
	"time"
)

// 获取当前的 access token.
// 如果过期了自动从服务器拉取新的 access token, 如果拉取失败则返回空串, 并返回错误信息.
func (c *Client) Token() (string, error) {
	token := c.accessToken.Token()
	if token != "" {
		return token, nil
	}

	// 当前的 access token 过期了, 则重新拉取
	resp, err := c.getNewToken()
	if err != nil {
		return "", err
	}

	c.accessToken.TokenValue = resp.AccessToken
	c.accessToken.Expires = time.Now().Unix() + resp.ExpiresIn
	return c.accessToken.TokenValue, nil
}

// 从微信服务器获取新的 access_token
func (c *Client) getNewToken() (*accessTokenResponse, error) {
	return nil, errors.New("没有实现")
}

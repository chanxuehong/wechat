package wechat

import (
	"errors"
	"time"
)

func (c *Client) Token() (string, error) {
	token := c.accessToken.Token()
	if token != "" {
		return token, nil
	}

	resp, err := c.getNewToken()
	if err != nil {
		return "", err
	}

	c.accessToken.TokenValue = resp.AccessToken
	c.accessToken.Expires = time.Now().Unix() + resp.ExpiresIn
	return c.accessToken.TokenValue, nil
}

// 从服务器获取 acces_token 成功时返回的消息格式
type accessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

// 从微信服务器获取新的 access_token
func (c *Client) getNewToken() (*accessTokenResponse, error) {
	return nil, errors.New("没有实现")
}

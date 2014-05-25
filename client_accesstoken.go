package wechat

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

	c.accessToken.Update(resp.AccessToken, resp.ExpiresIn)
	return resp.AccessToken, nil
}

// 从服务器获取 acces_token 成功时返回的消息格式
type accessTokenResponse struct {
	AccessToken string `json:"access_token"` // 获取到的凭证
	ExpiresIn   int64  `json:"expires_in"`   // 凭证有效时间，单位：秒
}

// 从微信服务器获取新的 access_token
func (c *Client) getNewToken() (*accessTokenResponse, error) {
	url := fmt.Sprintf(getAccessTokenUrlFormat, c.appid, c.appsecret)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		accessTokenResponse
		Error
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, &result.Error
	}
	return &result.accessTokenResponse, nil
}

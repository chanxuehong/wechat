package wechat

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// 获取 access token, 出错返回 "", 并返回错误信息; 否则返回 nil 错误.
func (c *Client) Token() (token string, err error) {
	return c.accessToken.Token()
}

// 强制刷新 access token, 正常情况下不要调用.
// 请使用 Client.Token() 获取 access token
func (c *Client) RefreshToken() (token string, err error) {
	resp, err := c.getNewToken()
	switch {
	case err != nil:
		c.accessToken.Update("", err)

	case resp.ExpiresIn <= 0: // 正常情况下不会出现
		err = fmt.Errorf("access token 过期时间是负数: %d", resp.ExpiresIn)
		c.accessToken.Update("", err)

	case resp.ExpiresIn <= 10: // 正常情况下不会出现
		c.accessToken.Update(resp.AccessToken, nil)
		token = resp.AccessToken
		c.resetTickChan <- time.Duration(resp.ExpiresIn) * time.Second

	default: // resp.ExpiresIn > 10
		c.accessToken.Update(resp.AccessToken, nil)
		token = resp.AccessToken
		// 设置新的 currentTickDuration, 考虑到网络延时, 提前 10 秒过期
		c.resetTickChan <- time.Duration(resp.ExpiresIn-10) * time.Second
	}
	return
}

// 负责更新 access token.
// 使用这种复杂的实现是减少 time.Now() 的调用.
func (c *Client) accessTokenService() {
	const defaultTickDuration = time.Minute
	// 获取新的 access token 时间间隔, 设置 44 秒以上就不会超过限制
	currentTickDuration := defaultTickDuration

OuterLoop: // 改变 currentTickDuration 重新开始
	for {
		tk := time.NewTicker(currentTickDuration)
		for {
			select {
			case currentTickDuration = <-c.resetTickChan:
				tk.Stop()
				break OuterLoop
			case <-tk.C:
				resp, err := c.getNewToken()
				switch {
				case err != nil:
					c.accessToken.Update("", err)
					// 出错则重置到 defaultTickDuration
					if currentTickDuration != defaultTickDuration {
						tk.Stop()
						currentTickDuration = defaultTickDuration
						break OuterLoop
					}
				case resp.ExpiresIn <= 0: // 正常情况下不会出现
					c.accessToken.Update("", fmt.Errorf("access token 过期时间是负数: %d", resp.ExpiresIn))
					// 出错则重置到 defaultTickDuration
					if currentTickDuration != defaultTickDuration {
						tk.Stop()
						currentTickDuration = defaultTickDuration
						break OuterLoop
					}
				case resp.ExpiresIn <= 10: // 正常情况下不会出现
					c.accessToken.Update(resp.AccessToken, nil)
					nextTickDuration := time.Duration(resp.ExpiresIn) * time.Second
					if currentTickDuration != nextTickDuration {
						currentTickDuration = nextTickDuration
						tk.Stop()
						break OuterLoop
					}
				default: // resp.ExpiresIn > 10
					c.accessToken.Update(resp.AccessToken, nil)
					// 设置新的 currentTickDuration, 考虑到网络延时, 提前 10 秒过期
					nextTickDuration := time.Duration(resp.ExpiresIn-10) * time.Second
					if currentTickDuration != nextTickDuration {
						currentTickDuration = nextTickDuration
						tk.Stop()
						break OuterLoop
					}
				}
			}
		}
	}
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

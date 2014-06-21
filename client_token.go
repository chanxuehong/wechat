package wechat

import (
	"fmt"
	"time"
)

// 获取 access token.
func (c *Client) Token() (token string, err error) {
	c.currentToken.rwmutex.RLock()
	token = c.currentToken.token
	err = c.currentToken.err
	c.currentToken.rwmutex.RUnlock()
	return
}

// see Client.TokenRefresh() and Client.tokenService()
func (c *Client) update(token string, err error) {
	c.currentToken.rwmutex.Lock()
	c.currentToken.token = token
	c.currentToken.err = err
	c.currentToken.rwmutex.Unlock()
}

// 从微信服务器获取新的 access token, 并保存到本地.
//  NOTE: 正常情况下无需调用该函数, 请使用 Client.Token() 获取 access token.
func (c *Client) TokenRefresh() (token string, err error) {
	resp, err := c.getNewToken()
	switch {
	case err != nil:
		c.update("", err)
		return
	case resp.ExpiresIn > 10: // 正常情况
		c.update(resp.Token, nil)
		token = resp.Token
		// 通知 goroutine tokenService() 重置定时器
		// 考虑到网络延时, 提前 10 秒过期
		c.resetRefreshTokenTickChan <- time.Duration(resp.ExpiresIn-10) * time.Second
		return
	case resp.ExpiresIn > 0: // (0, 10], 正常情况下不会出现
		c.update(resp.Token, nil)
		token = resp.Token
		// 通知 goroutine tokenService() 重置定时器
		c.resetRefreshTokenTickChan <- time.Duration(resp.ExpiresIn) * time.Second
		return
	default: // resp.ExpiresIn <= 0, 正常情况下不会出现
		err = fmt.Errorf("access token 过期时间应该是正整数, 现在 ==%d", resp.ExpiresIn)
		c.update("", err)
		return
	}
}

// 负责定时更新 access token.
//  NOTE: 使用这种复杂的实现是减少 time.Now() 的调用, 不然每次都要比较 time.Now().
func (c *Client) tokenService() {
	const defaultTickDuration = time.Minute // 设置 44 秒以上就不会超过限制(2000次/日 的限制)

	// 当前定时器的时间间隔, 正常情况下等于当前的 access token 的过期时间减去 10 秒;
	// 异常情况下就要尝试不断的获取, 时间间隔就是 defaultTickDuration.
	currentTickDuration := defaultTickDuration
	var tk *time.Ticker

NewTickDuration:
	for {
		tk = time.NewTicker(currentTickDuration)
		for {
			select {
			case currentTickDuration = <-c.resetRefreshTokenTickChan: // 在别的地方成功获取了 access token, 重置定时器.
				tk.Stop()
				break NewTickDuration

			case <-tk.C:
				resp, err := c.getNewToken()
				switch {
				case err != nil:
					c.update("", err)
					// 出错则重置到 defaultTickDuration
					if currentTickDuration != defaultTickDuration { // 这个判断的目的是避免重置定时器开销
						tk.Stop()
						currentTickDuration = defaultTickDuration
						break NewTickDuration
					}
				case resp.ExpiresIn > 10: // 正常情况
					c.update(resp.Token, nil)
					// 根据返回的过期时间来重新设置定时器
					// 设置新的 currentTickDuration, 考虑到网络延时, 提前 10 秒过期
					nextTickDuration := time.Duration(resp.ExpiresIn-10) * time.Second
					if currentTickDuration != nextTickDuration { // 这个判断的目的是避免重置定时器开销
						tk.Stop()
						currentTickDuration = nextTickDuration
						break NewTickDuration
					}
				case resp.ExpiresIn > 0: // (0, 10], 正常情况下不会出现
					c.update(resp.Token, nil)
					// 根据返回的过期时间来重新设置定时器
					nextTickDuration := time.Duration(resp.ExpiresIn) * time.Second
					if currentTickDuration != nextTickDuration { // 这个判断的目的是避免重置定时器开销
						tk.Stop()
						currentTickDuration = nextTickDuration
						break NewTickDuration
					}
				default: // resp.ExpiresIn <= 0, 正常情况下不会出现
					c.update("", fmt.Errorf("access token 过期时间应该是正整数, 现在 ==%d", resp.ExpiresIn))
					// 出错则重置到 defaultTickDuration
					if currentTickDuration != defaultTickDuration { // 这个判断的目的是避免重置定时器开销
						tk.Stop()
						currentTickDuration = defaultTickDuration
						break NewTickDuration
					}
				}
			}
		}
	}
}

// 从服务器获取 acces_token 成功时返回的消息格式
type clientTokenResponse struct {
	Token     string `json:"access_token"` // 获取到的凭证
	ExpiresIn int64  `json:"expires_in"`   // 凭证有效时间，单位：秒
}

// 从微信服务器获取新的 access_token
func (c *Client) getNewToken() (*clientTokenResponse, error) {
	_url := clientTokenGetURL(c.appid, c.appsecret)
	var result struct {
		clientTokenResponse
		Error
	}
	if err := c.getJSON(_url, &result); err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, &result.Error
	}
	return &result.clientTokenResponse, nil
}

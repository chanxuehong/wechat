package wechat

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

// 从本地缓存获取 access token.
func (c *Client) Token() (token string, err error) {
	c._token.rwmutex.RLock()
	token = c._token.token
	err = c._token.err
	c._token.rwmutex.RUnlock()
	return
}

// see Client.TokenRefresh() and Client.tokenService()
func (c *Client) update(token string, err error) {
	c._token.rwmutex.Lock()
	c._token.token = token
	c._token.err = err
	c._token.rwmutex.Unlock()
}

// 从微信服务器获取 access token, 并更新本地缓存.
//  NOTE: 正常情况下无需调用该函数, 请使用 Client.Token() 获取 access token.
func (c *Client) TokenRefresh() (token string, err error) {
	resp, err := c.getNewToken()
	switch {
	case err != nil:
		c.update("", err)
		return
	case resp.ExpiresIn <= 0: // 正常情况下不会出现
		err = fmt.Errorf("access token 过期时间是负数: %d", resp.ExpiresIn)
		c.update("", err)
		return
	case resp.ExpiresIn <= 10: // 正常情况下不会出现
		c.update(resp.Token, nil)
		token = resp.Token
		// 通知 goroutine tokenService() 重置定时器
		c.resetTickChan <- time.Duration(resp.ExpiresIn) * time.Second
		return
	default: // resp.ExpiresIn > 10
		c.update(resp.Token, nil)
		token = resp.Token
		// 通知 goroutine tokenService() 重置定时器
		// 考虑到网络延时, 提前 10 秒过期
		c.resetTickChan <- time.Duration(resp.ExpiresIn-10) * time.Second
		return
	}
}

// 负责定时更新本地缓存的 access token.
// 使用这种复杂的实现是减少 time.Now() 的调用, 不然每次都要比较 time.Now().
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
			case currentTickDuration = <-c.resetTickChan: // 在别的地方成功获取了 access token, 重置定时器.
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
				case resp.ExpiresIn <= 0: // 正常情况下不会出现
					c.update("", fmt.Errorf("access token 过期时间是负数: %d", resp.ExpiresIn))
					// 出错则重置到 defaultTickDuration
					if currentTickDuration != defaultTickDuration { // 这个判断的目的是避免重置定时器开销
						tk.Stop()
						currentTickDuration = defaultTickDuration
						break NewTickDuration
					}
				case resp.ExpiresIn <= 10: // 正常情况下不会出现
					c.update(resp.Token, nil)
					// 根据返回的过期时间来重新设置定时器
					nextTickDuration := time.Duration(resp.ExpiresIn) * time.Second
					if currentTickDuration != nextTickDuration { // 这个判断的目的是避免重置定时器开销
						tk.Stop()
						currentTickDuration = nextTickDuration
						break NewTickDuration
					}
				default: // resp.ExpiresIn > 10
					c.update(resp.Token, nil)
					// 根据返回的过期时间来重新设置定时器
					// 设置新的 currentTickDuration, 考虑到网络延时, 提前 10 秒过期
					nextTickDuration := time.Duration(resp.ExpiresIn-10) * time.Second
					if currentTickDuration != nextTickDuration { // 这个判断的目的是避免重置定时器开销
						tk.Stop()
						currentTickDuration = nextTickDuration
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
	resp, err := c.httpClient.Get(_url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		clientTokenResponse
		Error
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, &result.Error
	}
	return &result.clientTokenResponse, nil
}

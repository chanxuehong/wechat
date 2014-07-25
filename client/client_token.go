// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"fmt"
	"time"
)

// 获取 access token.
func (c *Client) Token() (token string, err error) {
	if c.tokenCache != nil {
		var tk Token
		tk, err = c.tokenCache.Token()
		if err != nil {
			if err != ErrTokenCacheMiss { // 从缓存中获取 access token 失败
				return
			}

			// 缓存中没有 access token, 去微信服务器获取
			return c.xaTokenRefresh()
		}

		if time.Now().Unix() > tk.ExpiresAt { // 从缓存中获取的 access token 过期了
			return c.xaTokenRefresh()
		}

		token = tk.Value
		return

	} else {
		c.currentToken.rwmutex.RLock()
		token = c.currentToken.token
		err = c.currentToken.err
		c.currentToken.rwmutex.RUnlock()
		return
	}
}

// see Client.TokenRefresh() and Client.tokenAutoUpdate()
func (c *Client) updateCurrentToken(token string, err error) {
	c.currentToken.rwmutex.Lock()
	c.currentToken.token = token
	c.currentToken.err = err
	c.currentToken.rwmutex.Unlock()
}

// 从微信服务器获取新的 access token.
//  NOTE: 正常情况下无需调用该函数, 请使用 Client.Token() 获取 access token.
func (c *Client) TokenRefresh() (token string, err error) {
	if c.tokenCache != nil {
		return c.xaTokenRefresh()

	} else {
		var resp tokenResponse
		resp, err = c.getNewToken()
		if err != nil {
			c.updateCurrentToken("", err)
			c.resetRefreshTokenTickChan <- time.Minute // 一分钟后尝试
			return
		}

		c.updateCurrentToken(resp.Token, nil)
		c.resetRefreshTokenTickChan <- time.Duration(resp.ExpiresIn) * time.Second
		token = resp.Token
		return
	}
}

// 负责定时更新 access token.
//  NOTE: 使用这种实现可以减少 time.Now().Unix() 的调用, 不然每次都要比较 time.Now().Unix().
func (c *Client) tokenAutoUpdate(tickDuration time.Duration) {
	const defaultTickDuration = time.Minute // 设置 44 秒以上就不会超过限制(2000次/日)

NEW_TICK_DURATION:
	ticker := time.NewTicker(tickDuration)
	for {
		select {
		case tickDuration = <-c.resetRefreshTokenTickChan:
			ticker.Stop()
			goto NEW_TICK_DURATION

		case <-ticker.C:
			resp, err := c.getNewToken()
			if err != nil {
				c.updateCurrentToken("", err)
				if tickDuration != defaultTickDuration { // 出错则重置到 defaultTickDuration
					ticker.Stop()
					tickDuration = defaultTickDuration
					goto NEW_TICK_DURATION
				}
			} else {
				c.updateCurrentToken(resp.Token, nil)
				newTickDuration := time.Duration(resp.ExpiresIn) * time.Second
				if tickDuration != newTickDuration {
					ticker.Stop()
					tickDuration = newTickDuration
					goto NEW_TICK_DURATION
				}
			}
		}
	}
}

// 从服务器获取 acces_token 成功时返回的消息格式
type tokenResponse struct {
	Token     string `json:"access_token"` // 获取到的凭证
	ExpiresIn int64  `json:"expires_in"`   // 凭证有效时间，单位：秒
}

// 从微信服务器获取新的 access_token
func (c *Client) getNewToken() (resp tokenResponse, err error) {
	_url := tokenGetURL(c.appid, c.appsecret)
	var result struct {
		tokenResponse
		Error
	}
	if err = c.getJSON(_url, &result); err != nil {
		return
	}

	if result.ErrCode != 0 {
		err = result.Error
		return
	}

	// 由于网络的延时 以及 分布式服务器之间的时间可能不是绝对同步, access token 过期时间留了一个缓冲区;
	// 正常情况下微信服务器会返回 7200, 则缓冲区的大小为 20 分钟, 这样分布式服务器之间的时间差
	// 在 20 分钟内基本不会出现问题!
	switch {
	case result.ExpiresIn > 60*60: // 返回的过期时间大于 1 个小时, 缓冲区为 20 分钟
		result.ExpiresIn -= 60 * 20
		resp = result.tokenResponse

	case result.ExpiresIn > 60*30: // 返回的过期时间大于 30 分钟, 缓冲区为 10 分钟
		result.ExpiresIn -= 60 * 10
		resp = result.tokenResponse

	case result.ExpiresIn > 60*15: // 返回的过期时间大于 15 分钟, 缓冲区为 5 分钟
		result.ExpiresIn -= 60 * 5
		resp = result.tokenResponse

	case result.ExpiresIn > 60*5: // 返回的过期时间大于 5 分钟, 缓冲区为 1 分钟
		result.ExpiresIn -= 60
		resp = result.tokenResponse

	case result.ExpiresIn > 60: // 返回的过期时间大于 1 分钟, 缓冲区为 20 秒
		result.ExpiresIn -= 20
		resp = result.tokenResponse

	case result.ExpiresIn > 0: // 没有办法了, 死马当做活马医了
		resp = result.tokenResponse

	default:
		err = fmt.Errorf("expires_in 应该是正整数, 现在为: %d", result.ExpiresIn)
	}
	return
}

// 分布式情况下, 从微信服务器获取新的 access token, 并更新到 c.tokenCache.
func (c *Client) xaTokenRefresh() (token string, err error) {
	resp, err := c.getNewToken()
	if err != nil {
		return
	}

	var tk Token
	tk.Value = resp.Token
	tk.ExpiresAt = time.Now().Unix() + resp.ExpiresIn
	if err = c.tokenCache.SetToken(tk); err != nil {
		return
	}

	token = tk.Value
	return
}

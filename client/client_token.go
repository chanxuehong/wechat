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
	if c.tokenService != nil {
		return c.tokenService.Token()
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
	if c.tokenService != nil {
		return c.tokenService.TokenRefresh()
	} else {
		var resp *tokenResponse
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
func (c *Client) getNewToken() (resp *tokenResponse, err error) {
	_url := tokenGetURL(c.appid, c.appsecret)
	var result struct {
		tokenResponse
		Error
	}
	if err = c.getJSON(_url, &result); err != nil {
		return
	}

	if result.ErrCode != 0 {
		err = &result.Error
		return
	}

	switch {
	case result.ExpiresIn > 10:
		result.ExpiresIn -= 10 // 考虑到网络延时, 提前 10 秒过期
		resp = &result.tokenResponse
		return

	case result.ExpiresIn > 0: // (0, 10], 正常情况下不会出现
		resp = &result.tokenResponse
		return

	default: // result.ExpiresIn <= 0, 正常情况下不会出现
		err = fmt.Errorf("expires_in 应该是正整数, 现在 ==%d", result.ExpiresIn)
		return
	}
}

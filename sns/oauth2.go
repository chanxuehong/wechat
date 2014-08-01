// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package sns

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// 用户相关的 oauth2 token 信息
//  NOTE: 每个用户对应一个这样的结构, 应该缓存起来, 一般缓存在 session 中.
type OAuth2Token struct {
	AccessToken  string
	RefreshToken string
	ExpiresAt    int64 // unixtime

	OpenId string
	Scopes []string // 用户授权的作用域
}

// 判断授权的 access token 是否过期
func (token *OAuth2Token) accessTokenExpired() bool {
	return time.Now().Unix() > token.ExpiresAt
}

// OAuth2Token 缓存接口
type TokenCache interface {
	Token() (*OAuth2Token, error)
	PutToken(*OAuth2Token) error
}

type Client struct {
	*OAuth2Config

	// 下面两个字段必须指定其中一个(通过code换取网页授权access_token 时除外)
	// 如果指定 OAuth2Token 则表示手动管理 OAuth2Token, 需要手动更新 OAuth2Token
	// 如果指定 TokenCache 则表示自动管理 OAuth2Token, 程序会自动更新 OAuth2Token
	*OAuth2Token
	TokenCache TokenCache

	//  如果 httpClient == nil 则默认用 http.DefaultClient,
	//  see ../CommonHttpClient 和 ../MediaHttpClient.
	HttpClient *http.Client
}

func (c *Client) httpClient() *http.Client {
	if c.HttpClient != nil {
		return c.HttpClient
	}
	return http.DefaultClient
}

// 通过code换取网页授权 access_token
//  NOTE: 如果成功 c.OAuth2Token 字段也会得到更新,
//        如果指定了 TokenCach, TokenCach.PutToken 也会被调用
func (c *Client) Exchange(code string) (token *OAuth2Token, err error) {
	if c.OAuth2Config == nil {
		err = errors.New("没有提供 OAuth2Config")
		return
	}

	tok := c.OAuth2Token
	if tok == nil && c.TokenCache != nil {
		tok, _ = c.TokenCache.Token()
	}
	if tok == nil {
		tok = new(OAuth2Token)
	}

	if err = c.updateToken(tok, oauth2ExchangeTokenURL(c.AppId, c.AppSecret, code)); err != nil {
		return
	}

	c.OAuth2Token = tok
	if c.TokenCache != nil {
		if err = c.TokenCache.PutToken(tok); err != nil {
			return
		}
	}
	token = tok
	return
}

// 刷新access_token（如果需要）
//  NOTE: 如果成功 c.OAuth2Token 字段也会得到更新,
//        如果指定了 TokenCach, TokenCach.PutToken 也会被调用
func (c *Client) TokenRefresh() (token *OAuth2Token, err error) {
	if c.OAuth2Config == nil {
		err = errors.New("没有提供 OAuth2Config")
		return
	}

	tok := c.OAuth2Token
	if tok == nil {
		if c.TokenCache == nil {
			err = errors.New("OAuth2Token 和 TokenCache 都没有提供")
			return
		}
		if tok, err = c.TokenCache.Token(); err != nil {
			return
		}
	}
	if tok == nil {
		err = errors.New("没有有效的 OAuth2Token")
		return
	}
	if len(tok.RefreshToken) == 0 {
		err = errors.New("没有有效的 RefreshToken")
		return
	}

	if err = c.updateToken(tok, oauth2RefreshTokenURL(c.AppId, tok.RefreshToken)); err != nil {
		return
	}

	c.OAuth2Token = tok
	if c.TokenCache != nil {
		if err = c.TokenCache.PutToken(tok); err != nil {
			return
		}
	}
	token = tok
	return
}

// 检查 access_token 是否有效
func (c *Client) CheckAccessTokenValid() (valid bool, err error) {
	if c.OAuth2Config == nil {
		err = errors.New("没有提供 OAuth2Config")
		return
	}

	tok := c.OAuth2Token
	if tok == nil {
		if c.TokenCache == nil {
			err = errors.New("OAuth2Token 和 TokenCache 都没有提供")
			return
		}
		if tok, err = c.TokenCache.Token(); err != nil {
			return
		}
	}
	if tok == nil {
		err = errors.New("没有有效的 OAuth2Token")
		return
	}
	if len(tok.AccessToken) == 0 {
		err = errors.New("没有有效的 AccessToken")
		return
	}
	if len(tok.OpenId) == 0 {
		err = errors.New("没有有效的 OpenId")
		return
	}

	c.OAuth2Token = tok // tok 有可能是从缓存读取的

	resp, err := c.httpClient().Get(checkAccessTokenValidURL(c.AccessToken, c.OpenId))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("http.Status: %s", resp.Status)
		return
	}

	var result Error
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return
	}

	switch result.ErrCode {
	case 0:
		valid = true
		return
	case 40001:
		return
	default:
		err = &result
		return
	}
}

// 从服务器获取新的 token 更新 tok
func (c *Client) updateToken(tok *OAuth2Token, _url string) (err error) {
	resp, err := c.httpClient().Get(_url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http.Status: %s", resp.Status)
	}

	var result struct {
		Error
		AccessToken  string `json:"access_token"`  // 网页授权接口调用凭证,注意：此access_token与基础支持的access_token不同
		RefreshToken string `json:"refresh_token"` // 用户刷新access_token
		ExpiresIn    int64  `json:"expires_in"`    // access_token接口调用凭证超时时间，单位（秒）
		OpenId       string `json:"openid"`        // 用户唯一标识，请注意，在未关注公众号时，用户访问公众号的网页，也会产生一个用户和公众号唯一的OpenID
		Scope        string `json:"scope"`         // 用户授权的作用域，使用逗号（,）分隔
	}

	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return
	}

	if result.ErrCode != 0 {
		return &result.Error
	}

	// 由于网络的延时 以及 分布式服务器之间的时间可能不是绝对同步, access token 过期时间留了一个缓冲区;
	// 正常情况下微信服务器会返回 7200, 则缓冲区的大小为 20 分钟, 这样分布式服务器之间的时间差
	// 在 20 分钟内基本不会出现问题!
	switch {
	case result.ExpiresIn > 60*60: // 返回的过期时间大于 1 个小时, 缓冲区为 20 分钟
		result.ExpiresIn -= 60 * 20

	case result.ExpiresIn > 60*30: // 返回的过期时间大于 30 分钟, 缓冲区为 10 分钟
		result.ExpiresIn -= 60 * 10

	case result.ExpiresIn > 60*15: // 返回的过期时间大于 15 分钟, 缓冲区为 5 分钟
		result.ExpiresIn -= 60 * 5

	case result.ExpiresIn > 60*5: // 返回的过期时间大于 5 分钟, 缓冲区为 1 分钟
		result.ExpiresIn -= 60

	case result.ExpiresIn > 60: // 返回的过期时间大于 1 分钟, 缓冲区为 20 秒
		result.ExpiresIn -= 20

	case result.ExpiresIn > 0: // 没有办法了, 死马当做活马医了

	default:
		err = fmt.Errorf("expires_in 应该是正整数, 现在为: %d", result.ExpiresIn)
		return
	}

	tok.AccessToken = result.AccessToken
	if len(result.RefreshToken) > 0 {
		tok.RefreshToken = result.RefreshToken
	}
	tok.ExpiresAt = time.Now().Unix() + result.ExpiresIn

	tok.OpenId = result.OpenId
	tok.Scopes = strings.Split(result.Scope, ",")

	return
}

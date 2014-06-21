// refer code.google.com/p/goauth2/oauth

package wechat

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/chanxuehong/wechat/user"
	"github.com/chanxuehong/wechat/user/sns"
	"net/http"
	"strings"
	"time"
)

// OAuth2Config is the configuration of an OAuth consumer.
type OAuth2Config struct {
	// ClientId is the OAuth client identifier used when communicating with
	// the configured OAuth provider.
	ClientId string

	// ClientSecret is the OAuth client secret used when communicating with
	// the configured OAuth provider.
	ClientSecret string

	// Scope identifies the level of access being requested. Multiple scope
	// values should be provided as a space-delimited string.
	Scope string

	// RedirectURL is the URL to which the user will be returned after
	// granting (or denying) access.
	RedirectURL string

	// Optional, "online" (default) or "offline", no refresh token if "online"
	AccessType string

	// ApprovalPrompt indicates whether the user should be
	// re-prompted for consent. If set to "auto" (default) the
	// user will be prompted only if they haven't previously
	// granted consent and the code can only be exchanged for an
	// access token.
	// If set to "force" the user will always be prompted, and the
	// code can be exchanged for a refresh token.
	ApprovalPrompt string
}

func NewOAuth2Config(appid, appsecret, redirectURL string, scope ...string) *OAuth2Config {
	return &OAuth2Config{
		ClientId:     appid,
		ClientSecret: appsecret,
		Scope:        strings.Join(scope, " "),
		RedirectURL:  redirectURL,
	}
}

// AuthCodeURL returns a URL that the end-user should be redirected to,
// so that they may obtain an authorization code.
func (c *OAuth2Config) AuthCodeURL(state string) string {
	return snsOAuth2AuthURL(c.ClientId, c.RedirectURL, c.Scope, state)
}

// OAuth2Token contains an end-user's tokens.
// This is the data you must store to persist authentication.
type OAuth2Token struct {
	AccessToken  string
	RefreshToken string
	Expiry       int64 // unixtime; If zero the token has no (known) expiry time.

	// wechat extra
	OpenId string
	Scope  string
}

func (t *OAuth2Token) Expired() bool {
	if t.Expiry == 0 {
		return false
	}
	return time.Now().Unix() > t.Expiry
}

type SNSClient struct {
	*OAuth2Config
	*OAuth2Token

	// It will default to http.DefaultClient if nil.
	HttpClient *http.Client
}

func (c *SNSClient) httpClient() *http.Client {
	if c.HttpClient != nil {
		return c.HttpClient
	}
	return http.DefaultClient
}

// 通过code换取网页授权access_token
func (c *SNSClient) Exchange(code string) (*OAuth2Token, error) {
	if len(code) == 0 {
		return nil, errors.New(`code == ""`)
	}
	if c.OAuth2Config == nil {
		return nil, errors.New("no OAuth2Config supplied")
	}

	// If the SNSClient has a token, preserve existing refresh token.
	tok := c.OAuth2Token
	if tok == nil {
		tok = new(OAuth2Token)
	}

	_url := snsOAuth2TokenURL(c.ClientId, c.ClientSecret, code)
	resp, err := c.httpClient().Get(_url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http.Status: %s", resp.Status)
	}

	type tokenResponse struct {
		AccessToken  string `json:"access_token"`  // 网页授权接口调用凭证,注意：此access_token与基础支持的access_token不同
		RefreshToken string `json:"refresh_token"` // 用户刷新access_token
		ExpiresIn    int64  `json:"expires_in"`    // access_token接口调用凭证超时时间，单位（秒）
		OpenId       string `json:"openid"`        // 用户唯一标识，请注意，在未关注公众号时，用户访问公众号的网页，也会产生一个用户和公众号唯一的OpenID
		Scope        string `json:"scope"`         // 用户授权的作用域，使用逗号（,）分隔
	}
	var result struct {
		tokenResponse
		Error
	}
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	if result.ErrCode != 0 {
		return nil, &result.Error
	}

	tok.AccessToken = result.AccessToken
	// Don't overwrite `RefreshToken` with an empty value
	if len(result.RefreshToken) > 0 {
		tok.RefreshToken = result.RefreshToken
	}
	switch {
	case result.ExpiresIn > 10: // 正常情况下远大于 10
		// 考虑到网络延时，提前 10 秒过期
		tok.Expiry = time.Now().Unix() + result.ExpiresIn - 10
	case result.ExpiresIn > 0:
		tok.Expiry = time.Now().Unix() + result.ExpiresIn
	case result.ExpiresIn == 0:
		tok.Expiry = 0
	default:
		return nil, fmt.Errorf("token ExpiresIn: %d < 0", result.ExpiresIn)
	}
	tok.OpenId = result.OpenId
	tok.Scope = result.Scope

	c.OAuth2Token = tok
	return tok, nil
}

// 刷新access_token（如果需要）
func (c *SNSClient) Refresh() error {
	if c.OAuth2Config == nil {
		return errors.New("no OAuth2Config supplied")
	}
	if c.OAuth2Token == nil {
		return errors.New("no OAuth2Token supplied")
	}
	if c.RefreshToken == "" {
		return errors.New("no Refresh Token")
	}

	_url := snsOAuth2RefreshTokenURL(c.ClientId, c.RefreshToken)
	resp, err := c.httpClient().Get(_url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http.Status: %s", resp.Status)
	}

	type tokenResponse struct {
		AccessToken  string `json:"access_token"`  // 网页授权接口调用凭证,注意：此access_token与基础支持的access_token不同
		RefreshToken string `json:"refresh_token"` // 用户刷新access_token
		ExpiresIn    int64  `json:"expires_in"`    // access_token接口调用凭证超时时间，单位（秒）
		OpenId       string `json:"openid"`        // 用户唯一标识，请注意，在未关注公众号时，用户访问公众号的网页，也会产生一个用户和公众号唯一的OpenID
		Scope        string `json:"scope"`         // 用户授权的作用域，使用逗号（,）分隔
	}
	var result struct {
		tokenResponse
		Error
	}
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}
	if result.ErrCode != 0 {
		return &result.Error
	}

	c.OAuth2Token.AccessToken = result.AccessToken
	// Don't overwrite `RefreshToken` with an empty value
	if len(result.RefreshToken) > 0 {
		c.OAuth2Token.RefreshToken = result.RefreshToken
	}
	switch {
	case result.ExpiresIn > 10: // 正常情况下远大于 10
		// 考虑到网络延时，提前 10 秒过期
		c.OAuth2Token.Expiry = time.Now().Unix() + result.ExpiresIn - 10
	case result.ExpiresIn > 0:
		c.OAuth2Token.Expiry = time.Now().Unix() + result.ExpiresIn
	case result.ExpiresIn == 0:
		c.OAuth2Token.Expiry = 0
	default:
		return fmt.Errorf("token ExpiresIn: %d < 0", result.ExpiresIn)
	}
	c.OAuth2Token.OpenId = result.OpenId
	c.OAuth2Token.Scope = result.Scope
	return nil
}

// 拉取用户信息(需scope为 snsapi_userinfo).
//  lang 可能的取值是 zh_CN, zh_TW, en; 如果留空 "" 则默认为 zh_CN.
func (c *SNSClient) UserInfo(openid, lang string) (*sns.UserInfo, error) {
	if len(openid) == 0 {
		return nil, errors.New(`openid == ""`)
	}
	switch lang {
	case "":
		lang = user.Language_zh_CN
	case user.Language_zh_CN, user.Language_zh_TW, user.Language_en:
	default:
		return nil, errors.New(`lang 必须是 "", zh_CN, zh_TW, en 之一`)
	}

	if c.OAuth2Token == nil {
		return nil, errors.New("no OAuth2Token supplied")
	}

	// Refresh the OAuth2Token if it has expired.
	if c.Expired() {
		if err := c.Refresh(); err != nil {
			return nil, err
		}
	}

	_url := snsUserInfoURL(c.AccessToken, openid, lang)
	resp, err := c.httpClient().Get(_url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http.Status: %s", resp.Status)
	}

	var result struct {
		sns.UserInfo
		Error
	}
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, &result.Error
	}
	return &result.UserInfo, nil
}

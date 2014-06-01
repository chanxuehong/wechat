// refer code.google.com/p/goauth2/oauth

package wechat

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/chanxuehong/wechat/user"
	"github.com/chanxuehong/wechat/user/sns"
	"net/http"
	"net/url"
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

	// AuthURL is the URL the user will be directed to in order to grant
	// access.
	AuthURL string

	// TokenURL is the URL used to Exchange OAuth tokens.
	TokenURL string

	// RefreshTokenURL is the URL used to Refresh OAuth tokens.
	RefreshTokenURL string

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
		ClientId:        appid,
		ClientSecret:    appsecret,
		Scope:           strings.Join(scope, " "),
		AuthURL:         snsOAuth2AuthURL,
		TokenURL:        snsOAuth2TokenURL,
		RefreshTokenURL: snsOAuth2RefreshTokenURL,
		RedirectURL:     redirectURL,
	}
}

// AuthCodeURL returns a URL that the end-user should be redirected to,
// so that they may obtain an authorization code.
func (c *OAuth2Config) AuthCodeURL(state string) string {
	_url, err := url.Parse(c.AuthURL)
	if err != nil {
		panic("AuthURL malformed: " + err.Error())
	}
	q := url.Values{
		"appid":         {c.ClientId},
		"redirect_uri":  {c.RedirectURL},
		"response_type": {"code"},
		"scope":         {c.Scope},
		"state":         {state},
	}.Encode()

	if _url.RawQuery == "" {
		_url.RawQuery = q
	} else {
		_url.RawQuery += "&" + q
	}
	return _url.String() + "#wechat_redirect"
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
	// 考虑到网络延时, 提前 10 秒过期; 如果返回的过期时间小于 10 秒则不能正常工作!
	return time.Now().Unix()+10 > t.Expiry
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
	if c.OAuth2Config == nil {
		return nil, errors.New("Exchange: no OAuth2Config supplied")
	}

	// If the SNSClient has a token, preserve existing refresh token.
	tok := c.OAuth2Token
	if tok == nil {
		tok = new(OAuth2Token)
	}

	_url, err := url.Parse(c.TokenURL)
	if err != nil {
		panic("TokenURL malformed: " + err.Error())
	}

	q := url.Values{
		"appid":      {c.ClientId},
		"secret":     {c.ClientSecret},
		"grant_type": {"authorization_code"},
		"code":       {code},
	}.Encode()

	if _url.RawQuery == "" {
		_url.RawQuery = q
	} else {
		_url.RawQuery += "&" + q
	}

	client := c.httpClient()
	req, err := http.NewRequest("GET", _url.String(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Exchange: %s", resp.Status)
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
	if result.ExpiresIn == 0 {
		tok.Expiry = 0
	} else {
		tok.Expiry = time.Now().Unix() + result.ExpiresIn
	}
	tok.OpenId = result.OpenId
	tok.Scope = result.Scope

	c.OAuth2Token = tok
	return tok, nil
}

// 刷新access_token（如果需要）
func (c *SNSClient) Refresh() error {
	if c.OAuth2Config == nil {
		return errors.New("Refresh: no OAuth2Config supplied")
	}
	if c.OAuth2Token == nil {
		return errors.New("Refresh: no OAuth2Token supplied")
	}
	if c.RefreshToken == "" {
		return errors.New("Refresh: no Refresh Token")
	}

	_url, err := url.Parse(c.RefreshTokenURL)
	if err != nil {
		panic("RefreshTokenURL malformed: " + err.Error())
	}

	q := url.Values{
		"appid":         {c.ClientId},
		"grant_type":    {"refresh_token"},
		"refresh_token": {c.RefreshToken},
	}.Encode()

	if _url.RawQuery == "" {
		_url.RawQuery = q
	} else {
		_url.RawQuery += "&" + q
	}

	client := c.httpClient()
	req, err := http.NewRequest("GET", _url.String(), nil)
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Refresh: %s", resp.Status)
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
	if result.ExpiresIn == 0 {
		c.OAuth2Token.Expiry = 0
	} else {
		c.OAuth2Token.Expiry = time.Now().Unix() + result.ExpiresIn
	}
	c.OAuth2Token.OpenId = result.OpenId
	c.OAuth2Token.Scope = result.Scope
	return nil
}

// 拉取用户信息(需scope为 snsapi_userinfo).
//  lang 可能的取值是 zh_CN, zh_TW, en; 如果留空 "" 则默认为 zh_CN.
func (c *SNSClient) UserInfo(openid, lang string) (*sns.UserInfo, error) {
	switch lang {
	case "":
		lang = user.Language_zh_CN
	case user.Language_zh_CN, user.Language_zh_TW, user.Language_en:
	default:
		return nil, errors.New(`lang 必须是 "", zh_CN, zh_TW, en 之一`)
	}

	if c.OAuth2Token == nil {
		return nil, errors.New("UserInfo: no OAuth2Token supplied")
	}

	// Refresh the OAuth2Token if it has expired.
	if c.Expired() {
		if err := c.Refresh(); err != nil {
			return nil, err
		}
	}

	_url := fmt.Sprintf(snsUserInfoURLFormat, c.AccessToken, openid, lang)
	client := c.httpClient()
	req, err := http.NewRequest("GET", _url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("UserInfo: %s", resp.Status)
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

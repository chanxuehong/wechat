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

// oauth2 相关配置 ( 一般全局只用保存一个变量 )
type OAuth2Config struct {
	ClientId     string // appid
	ClientSecret string // appsecret

	// 应用授权作用域，拥有多个作用域用逗号（,）分隔;
	// 网页应用目前仅填写snsapi_login即可.
	Scope string

	// 用户授权后跳转的目的地址
	// 用户授权后跳转到 RedirectURL?code=CODE&state=STATE
	// 用户禁止授权跳转到 RedirectURL?state=STATE
	RedirectURL string
}

func NewOAuth2Config(appid, appsecret, redirectURL string, scope ...string) *OAuth2Config {
	return &OAuth2Config{
		ClientId:     appid,
		ClientSecret: appsecret,
		Scope:        strings.Join(scope, ","),
		RedirectURL:  redirectURL,
	}
}

// 请求用户授权时跳转的地址.
//  这里的格式为:
//  https://open.weixin.qq.com/connect/oauth2/authorize?appid=APPID&redirect_uri=REDIRECT_URI&response_type=code&scope=SCOPE&state=STATE#wechat_redirect
func (cfg *OAuth2Config) AuthCodeURL(state string) string {
	return oauth2AuthURL(cfg.ClientId, cfg.RedirectURL, cfg.Scope, state)
}

// 通过code换取网页授权access_token 和 刷新 access_token 返回的数据结构.
//  NOTE: 每个用户对应一个这样的结构, 应该缓存起来, 一般缓存在 session 中.
type OAuth2Token struct {
	AccessToken  string
	RefreshToken string
	ExpiresAt    int64 // unixtime; 如果等于 0 则表示不过期或者不知道具体过期时间.

	// 用户唯一标识，请注意，在未关注公众号时，用户访问公众号的网页，
	// 也会产生一个用户和公众号唯一的OpenID
	OpenId string
	Scopes []string // 用户授权的作用域
}

// 判断授权的 access token 是否过期
func (token *OAuth2Token) AccessTokenExpired() bool {
	if token.ExpiresAt == 0 {
		return false
	}
	return time.Now().Unix() > token.ExpiresAt
}

// OAuth2Token 字段正常情况下要从缓存中获取，因为是用户相关，一般都是存储在 session 中。
// 每次请求结束都要从这个字段获取最新的 OAuth2Token 再存入 session 中。
type Client struct {
	*OAuth2Config
	*OAuth2Token

	// It will default to http.DefaultClient if nil.
	// see ../CommonHttpClient and ../MediaHttpClient
	HttpClient *http.Client
}

func (c *Client) httpClient() *http.Client {
	if c.HttpClient != nil {
		return c.HttpClient
	}
	return http.DefaultClient
}

// 交换 token 和 刷新 token 成功时返回的收据结构
type tokenResponse struct {
	AccessToken  string `json:"access_token"`  // 网页授权接口调用凭证,注意：此access_token与基础支持的access_token不同
	RefreshToken string `json:"refresh_token"` // 用户刷新access_token
	ExpiresIn    int64  `json:"expires_in"`    // access_token接口调用凭证超时时间，单位（秒）
	OpenId       string `json:"openid"`        // 用户唯一标识，请注意，在未关注公众号时，用户访问公众号的网页，也会产生一个用户和公众号唯一的OpenID
	Scope        string `json:"scope"`         // 用户授权的作用域，使用逗号（,）分隔
}

// 通过code换取网页授权access_token
//  NOTE: 如果成功 c.OAuth2Token 字段也会得到更新, 后续操作不需要重新赋值.
func (c *Client) Exchange(code string) (token *OAuth2Token, err error) {
	if c.OAuth2Config == nil {
		err = errors.New("no OAuth2Config supplied")
		return
	}

	// If the Client has a token, preserve existing refresh token.
	tok := c.OAuth2Token
	if tok == nil {
		tok = new(OAuth2Token)
	}

	_url := oauth2ExchangeTokenURL(c.ClientId, c.ClientSecret, code)
	resp, err := c.httpClient().Get(_url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("http.Status: %s", resp.Status)
		return
	}

	var result struct {
		Error
		tokenResponse
	}
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = result.Error
		return
	}

	switch {
	case result.ExpiresIn > 10: // 正常情况下远大于 10
		tok.ExpiresAt = time.Now().Unix() + result.ExpiresIn - 10 // 考虑到网络延时，提前 10 秒过期

	case result.ExpiresIn > 0:
		tok.ExpiresAt = time.Now().Unix() + result.ExpiresIn

	case result.ExpiresIn == 0:
		tok.ExpiresAt = 0

	default:
		err = fmt.Errorf("token ExpiresIn(==%d) < 0", result.ExpiresIn)
		return
	}

	tok.AccessToken = result.AccessToken
	// Don't overwrite `RefreshToken` with an empty value
	if len(result.RefreshToken) > 0 {
		tok.RefreshToken = result.RefreshToken
	}
	tok.OpenId = result.OpenId
	tok.Scopes = strings.Split(result.Scope, ",")

	c.OAuth2Token = tok
	token = tok
	return
}

// 刷新access_token（如果需要）
//  NOTE: 如果成功则会更新 c.OAuth2Token 字段.
func (c *Client) TokenRefresh() (err error) {
	if c.OAuth2Config == nil {
		return errors.New("no OAuth2Config supplied")
	}
	if c.OAuth2Token == nil {
		return errors.New("no OAuth2Token supplied")
	}
	if len(c.RefreshToken) == 0 {
		return errors.New("no Refresh Token")
	}

	_url := oauth2RefreshTokenURL(c.ClientId, c.RefreshToken)
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
		tokenResponse
	}
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return
	}

	if result.ErrCode != 0 {
		return result.Error
	}

	switch {
	case result.ExpiresIn > 10: // 正常情况下远大于 10
		c.ExpiresAt = time.Now().Unix() + result.ExpiresIn - 10 // 考虑到网络延时，提前 10 秒过期

	case result.ExpiresIn > 0:
		c.ExpiresAt = time.Now().Unix() + result.ExpiresIn

	case result.ExpiresIn == 0:
		c.ExpiresAt = 0

	default:
		return fmt.Errorf("token ExpiresIn(==%d) < 0", result.ExpiresIn)
	}

	c.AccessToken = result.AccessToken
	// Don't overwrite `RefreshToken` with an empty value
	if len(result.RefreshToken) > 0 {
		c.RefreshToken = result.RefreshToken
	}
	c.OpenId = result.OpenId
	c.Scopes = strings.Split(result.Scope, ",")

	return
}

// 检查 access_token 是否有效
func (c *Client) CheckAccessTokenValid() (valid bool, err error) {
	if c.OAuth2Token == nil {
		err = errors.New("no OAuth2Token supplied")
		return
	}
	if len(c.AccessToken) == 0 {
		err = errors.New("no Access Token")
		return
	}
	if len(c.OpenId) == 0 {
		err = errors.New(`no OpenId`)
		return
	}

	_url := checkAccessTokenValidURL(c.AccessToken, c.OpenId)
	resp, err := c.httpClient().Get(_url)
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

	// 出错则表示无效
	if result.ErrCode != 0 {
		return
	}

	valid = true
	return
}

// 获取用户信息(需scope为 snsapi_userinfo).
//  lang 可能的取值是 zh_CN, zh_TW, en; 如果留空 "" 则默认为 zh_CN.
func (c *Client) UserInfo(lang string) (info *UserInfo, err error) {
	switch lang {
	case "":
		lang = Language_zh_CN
	case Language_zh_CN, Language_zh_TW, Language_en:
	default:
		err = fmt.Errorf("lang 必须是 \"\",%s,%s,%s 其中之一",
			Language_zh_CN, Language_zh_TW, Language_en)
		return
	}

	if c.OAuth2Token == nil {
		err = errors.New("no OAuth2Token supplied")
		return
	}
	if len(c.AccessToken) == 0 {
		err = errors.New("no Access Token")
		return
	}
	if len(c.OpenId) == 0 {
		err = errors.New(`no OpenId`)
		return
	}

	// Refresh the OAuth2Token if it has expired.
	if c.AccessTokenExpired() {
		if err = c.TokenRefresh(); err != nil {
			return
		}
	}

	_url := userInfoURL(c.AccessToken, c.OpenId, lang)
	resp, err := c.httpClient().Get(_url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("http.Status: %s", resp.Status)
		return
	}

	var result struct {
		UserInfo
		Error
	}
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return
	}

	if result.ErrCode != 0 {
		err = result.Error
		return
	}

	info = &result.UserInfo
	return
}

// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

// +build wechatdebug

package oauth2

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/chanxuehong/wechat/mp"
)

// 构造请求用户授权获取code的地址.
//  appId:       公众号的唯一标识
//  redirectURL: 授权后重定向的回调链接地址
//               如果用户同意授权，页面将跳转至 redirect_uri/?code=CODE&state=STATE。
//               若用户禁止授权，则重定向后不会带上code参数，仅会带上state参数redirect_uri?state=STATE
//  scope:       应用授权作用域，
//               snsapi_base （不弹出授权页面，直接跳转，只能获取用户openid），
//               snsapi_userinfo （弹出授权页面，可通过openid拿到昵称、性别、所在地。
//               并且，即使在未关注的情况下，只要用户授权，也能获取其信息）
//  state:       重定向后会带上state参数，开发者可以填写a-zA-Z0-9的参数值，最多128字节
func AuthCodeURL(appId, redirectURL, scope, state string) string {
	return "https://open.weixin.qq.com/connect/oauth2/authorize" +
		"?appid=" + url.QueryEscape(appId) +
		"&redirect_uri=" + url.QueryEscape(redirectURL) +
		"&response_type=code&scope=" + url.QueryEscape(scope) +
		"&state=" + url.QueryEscape(state) +
		"#wechat_redirect"
}

// 用户相关的 oauth2 token 信息
type OAuth2Token struct {
	AccessToken  string
	RefreshToken string
	ExpiresAt    int64 // 过期时间, unixtime, 分布式系统要求时间同步, 建议使用 NTP

	OpenId string
	Scopes []string // 用户授权的作用域
}

// 判断授权的 OAuth2Token.AccessToken 是否过期, 过期返回 true, 否则返回 false
func (token *OAuth2Token) accessTokenExpired() bool {
	return time.Now().Unix() >= token.ExpiresAt
}

type Client struct {
	*OAuth2Config
	*OAuth2Token // 程序会自动更新最新的 OAuth2Token 到这个字段, 如有必要该字段可以保存起来

	HttpClient *http.Client // 如果 httpClient == nil 则默认用 http.DefaultClient
}

func (clt *Client) httpClient() *http.Client {
	if clt.HttpClient != nil {
		return clt.HttpClient
	}
	return http.DefaultClient
}

// 通过code换取网页授权access_token.
//  NOTE:
//  1. Client 需要指定 OAuth2Config
//  2. 如果指定了 OAuth2Token, 则会更新这个 OAuth2Token, 同时返回的也是指定的 OAuth2Token;
//     否则会重新分配一个 OAuth2Token.
func (clt *Client) Exchange(code string) (token *OAuth2Token, err error) {
	if clt.OAuth2Config == nil {
		err = errors.New("没有提供 OAuth2Config")
		return
	}

	tk := clt.OAuth2Token
	if tk == nil {
		tk = new(OAuth2Token)
	}

	_url := "https://api.weixin.qq.com/sns/oauth2/access_token" +
		"?appid=" + url.QueryEscape(clt.AppId) +
		"&secret=" + url.QueryEscape(clt.AppSecret) +
		"&code=" + url.QueryEscape(code) +
		"&grant_type=authorization_code"
	if err = clt.updateToken(tk, _url); err != nil {
		return
	}

	clt.OAuth2Token = tk
	token = tk
	return
}

// 刷新access_token（如果需要）.
//  NOTE: Client 需要指定 OAuth2Config, OAuth2Token
func (clt *Client) TokenRefresh() (token *OAuth2Token, err error) {
	if clt.OAuth2Config == nil {
		err = errors.New("没有提供 OAuth2Config")
		return
	}
	if clt.OAuth2Token == nil {
		err = errors.New("没有提供 OAuth2Token")
		return
	}
	if clt.RefreshToken == "" {
		err = errors.New("没有有效的 RefreshToken")
		return
	}

	_url := "https://api.weixin.qq.com/sns/oauth2/refresh_token" +
		"?appid=" + url.QueryEscape(clt.AppId) +
		"&grant_type=refresh_token&refresh_token=" + url.QueryEscape(clt.RefreshToken)
	if err = clt.updateToken(clt.OAuth2Token, _url); err != nil {
		return
	}

	token = clt.OAuth2Token
	return
}

// 检验授权凭证（access_token）是否有效.
//  NOTE:
//  1. Client 需要指定 OAuth2Token
//  2. 先判断 err 然后再判断 valid
func (clt *Client) CheckAccessTokenValid() (valid bool, err error) {
	if clt.OAuth2Token == nil {
		err = errors.New("没有提供 OAuth2Token")
		return
	}
	if clt.AccessToken == "" {
		err = errors.New("没有有效的 AccessToken")
		return
	}
	if clt.OpenId == "" {
		err = errors.New("没有有效的 OpenId")
		return
	}

	_url := "https://api.weixin.qq.com/sns/auth?access_token=" + url.QueryEscape(clt.AccessToken) +
		"&openid=" + url.QueryEscape(clt.OpenId)
	httpResp, err := clt.httpClient().Get(_url)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		err = fmt.Errorf("http.Status: %s", httpResp.Status)
		return
	}

	var result mp.Error

	respBody, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return
	}

	log.Println("request url:", _url)
	log.Println("response json:", string(respBody))

	if err = json.Unmarshal(respBody, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case mp.ErrCodeOK:
		valid = true
		return
	case 40001:
		return
	default:
		err = &result
		return
	}
}

// 从服务器获取新的 token 更新 tk
func (clt *Client) updateToken(tk *OAuth2Token, url string) (err error) {
	if tk == nil {
		return errors.New("nil OAuth2Token")
	}

	httpResp, err := clt.httpClient().Get(url)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return fmt.Errorf("http.Status: %s", httpResp.Status)
	}

	var result struct {
		mp.Error
		AccessToken  string `json:"access_token"`  // 网页授权接口调用凭证,注意：此access_token与基础支持的access_token不同
		RefreshToken string `json:"refresh_token"` // 用户刷新access_token
		ExpiresIn    int64  `json:"expires_in"`    // access_token接口调用凭证超时时间，单位（秒）
		OpenId       string `json:"openid"`        // 用户唯一标识，请注意，在未关注公众号时，用户访问公众号的网页，也会产生一个用户和公众号唯一的OpenID
		Scope        string `json:"scope"`         // 用户授权的作用域，使用逗号（,）分隔
	}

	respBody, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return
	}

	log.Println("request url:", url)
	log.Println("response json:", string(respBody))

	if err = json.Unmarshal(respBody, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		return &result.Error
	}

	// 由于网络的延时, 分布式服务器之间的时间可能不是绝对同步, access_token 过期时间留了一个缓冲区;
	switch {
	case result.ExpiresIn > 60*60:
		result.ExpiresIn -= 60 * 20
	case result.ExpiresIn > 60*30:
		result.ExpiresIn -= 60 * 10
	case result.ExpiresIn > 60*15:
		result.ExpiresIn -= 60 * 5
	case result.ExpiresIn > 60*5:
		result.ExpiresIn -= 60
	case result.ExpiresIn > 60:
		result.ExpiresIn -= 20
	case result.ExpiresIn > 0:
	default:
		err = fmt.Errorf("invalid expires_in: %d", result.ExpiresIn)
		return
	}

	tk.AccessToken = result.AccessToken
	if result.RefreshToken != "" {
		tk.RefreshToken = result.RefreshToken
	}
	tk.ExpiresAt = time.Now().Unix() + result.ExpiresIn

	tk.OpenId = result.OpenId

	strs := strings.Split(result.Scope, ",")
	tk.Scopes = make([]string, 0, len(strs))
	for _, str := range strs {
		str = strings.TrimSpace(str)
		if str == "" {
			continue
		}
		tk.Scopes = append(tk.Scopes, str)
	}
	return
}

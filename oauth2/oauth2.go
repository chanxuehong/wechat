<<<<<<< HEAD
// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package oauth2

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

// 通过code换取网页授权access_token.
//  返回的 token == clt.Token
func (clt *Client) Exchange(code string) (token *Token, err error) {
	if clt.Config == nil {
		err = errors.New("nil Config")
		return
	}

	var tk *Token
	if clt.TokenStorage != nil {
		if tk, _ = clt.TokenStorage.Get(); tk == nil {
			tk = clt.Token
		} else {
			clt.Token = tk // update local
		}
	} else {
		tk = clt.Token
	}
	if tk == nil {
		tk = new(Token)
	}

	if err = clt.updateToken(tk, clt.Config.ExchangeTokenURL(code)); err != nil {
		return
	}

	if err = clt.putToken(tk); err != nil {
		return
	}
	token = tk
	return
}

// 刷新access_token(如果需要).
//  返回的 token == clt.Token
func (clt *Client) TokenRefresh() (token *Token, err error) {
	if clt.Config == nil {
		err = errors.New("nil Config")
		return
	}

	tk, err := clt.getToken()
	if err != nil {
		return
	}

	return clt.tokenRefresh(tk)
}

func (clt *Client) tokenRefresh(tk *Token) (token *Token, err error) {
	if err = clt.updateToken(tk, clt.Config.RefreshTokenURL(tk.RefreshToken)); err != nil {
		return
	}

	if err = clt.putToken(tk); err != nil {
		return
	}
	token = tk
	return
}

// 从服务器获取新的 token 更新 tk, 通过 code 换取 token 或者刷新 token
func (clt *Client) updateToken(tk *Token, url string) (err error) {
	var result struct {
		Error
		AccessToken  string `json:"access_token"`
		ExpiresIn    int64  `json:"expires_in"`
		RefreshToken string `json:"refresh_token"`
		OpenId       string `json:"openid"`
		UnionId      string `json:"unionid"`
		Scope        string `json:"scope"`
	}
	if err = clt.getJSON(url, &result); err != nil {
		return
	}

	if result.ErrCode != ErrCodeOK {
		return &result.Error
	}

	// 由于网络的延时, 分布式服务器之间的时间可能不是绝对同步, access_token 过期时间留了一个缓冲区;
	switch {
	case result.ExpiresIn > 31556952: // 60*60*24*365.2425
		err = errors.New("expires_in too large: " + strconv.FormatInt(result.ExpiresIn, 10))
		return
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
	default:
		err = errors.New("expires_in too small: " + strconv.FormatInt(result.ExpiresIn, 10))
		return
	}

	tk.AccessToken = result.AccessToken
	tk.CreateAt = time.Now().Unix()
	tk.ExpiresIn = result.ExpiresIn
	if result.RefreshToken != "" {
		tk.RefreshToken = result.RefreshToken
	}
	tk.OpenId = result.OpenId
	tk.UnionId = result.UnionId

	strArr := strings.Split(result.Scope, ",")
	tk.Scopes = make([]string, 0, len(strArr))
	for _, str := range strArr {
		str = strings.TrimSpace(str)
		if str == "" {
			continue
		}
		tk.Scopes = append(tk.Scopes, str)
	}
	return
=======
package oauth2

import (
	"time"
)

type Endpoint interface {
	ExchangeTokenURL(code string) string        // 通过code换取access_token的地址
	RefreshTokenURL(refreshToken string) string // 刷新access_token的地址
}

type TokenStorage interface {
	Token() (*Token, error)
	PutToken(*Token) error
}

type Token struct {
	AccessToken  string `json:"access_token"`            // 网页授权接口调用凭证
	CreatedAt    int64  `json:"created_at"`              // access_token 创建时间, unixtime, 分布式系统要求时间同步, 建议使用 NTP
	ExpiresIn    int64  `json:"expires_in"`              // access_token 接口调用凭证超时时间, 单位: 秒
	RefreshToken string `json:"refresh_token,omitempty"` // 刷新 access_token 的凭证

	OpenId  string `json:"openid,omitempty"`
	UnionId string `json:"unionid,omitempty"`
	Scope   string `json:"scope,omitempty"` // 用户授权的作用域, 使用逗号(,)分隔
}

// Expired 判断 token.AccessToken 是否过期, 过期返回 true, 否则返回 false.
func (token *Token) Expired() bool {
	return time.Now().Unix() >= token.CreatedAt+token.ExpiresIn
>>>>>>> github/v2
}

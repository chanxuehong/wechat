package oauth2

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/chanxuehong/wechat/internal/api"
)

// Exchange 通过 code 换取网页授权 access_token.
//  NOTE: 返回的 token == clt.Token
func (clt *Client) Exchange(code string) (token *Token, err error) {
	if clt.Config == nil {
		err = errors.New("nil Client.Config")
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

// TokenRefresh 刷新 access_token.
//  NOTE:
//  1. 返回的 token == clt.Token
//  2. refreshToken 可以为空.
func (clt *Client) TokenRefresh(refreshToken string) (token *Token, err error) {
	if clt.Config == nil {
		err = errors.New("nil Client.Config")
		return
	}

	var tk *Token
	if refreshToken == "" {
		if tk, err = clt.GetToken(false); err != nil {
			return
		}
		refreshToken = tk.RefreshToken
	} else {
		tk = new(Token)
	}

	if err = clt.updateToken(tk, clt.Config.RefreshTokenURL(refreshToken)); err != nil {
		return
	}
	if err = clt.putToken(tk); err != nil {
		return
	}
	token = tk
	return
}

func (clt *Client) updateToken(tk *Token, url string) (err error) {
	api.DebugPrintGetRequest(url)
	httpResp, err := clt.httpClient().Get(url)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return fmt.Errorf("http.Status: %s", httpResp.Status)
	}

	var result struct {
		Error
		Token
	}
	if err = api.JsonHttpResponseUnmarshal(httpResp.Body, &result); err != nil {
		return
	}
	if result.ErrCode != ErrCodeOK {
		return &result.Error
	}

	// 由于网络的延时 和 分布式服务器之间的时间可能不是绝对同步, access_token 过期时间留了一个缓冲区
	switch {
	case result.ExpiresIn > 31556952: // 60*60*24*365.2425
		return errors.New("expires_in too large: " + strconv.FormatInt(result.ExpiresIn, 10))
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
		return errors.New("expires_in too small: " + strconv.FormatInt(result.ExpiresIn, 10))
	}

	result.Token.CreatedAt = time.Now().Unix()
	*tk = result.Token
	return
}

package oauth2

import (
	"errors"
	"net/http"

	"github.com/chanxuehong/wechat/util"
)

type Client struct {
	Endpoint Endpoint

	// TokenStorage 和 Token 两个字段正常情况下只用指定一个, 如果两个同时被指定了, 优先使用 TokenStorage
	TokenStorage TokenStorage
	Token        *Token // Client 自动将最新的 Token 更新到此字段, 不管 Token 字段一开始是否被指定!!!

	HttpClient *http.Client // 如果 HttpClient == nil 则默认用 util.DefaultHttpClient
}

func (clt *Client) httpClient() *http.Client {
	if clt.HttpClient != nil {
		return clt.HttpClient
	}
	return util.DefaultHttpClient
}

// GetToken 获取 Token, autoRefresh 为 true 时如果 Token 过期则自动刷新.
func (clt *Client) GetToken(autoRefresh bool) (tk *Token, err error) {
	if clt.TokenStorage != nil {
		if tk, err = clt.TokenStorage.Token(); err != nil {
			return
		}
		if tk == nil {
			err = errors.New("incorrect TokenStorage.Token implementation")
			return
		}
		clt.Token = tk // update local
	} else {
		tk = clt.Token
		if tk == nil {
			err = errors.New("nil TokenStorage and nil Token")
			return
		}
	}
	if autoRefresh && tk.Expired() {
		return clt.RefreshToken(tk.RefreshToken)
	}
	return
}

func (clt *Client) putToken(tk *Token) (err error) {
	if clt.TokenStorage != nil {
		if err = clt.TokenStorage.PutToken(tk); err != nil {
			return
		}
	}
	clt.Token = tk
	return
}

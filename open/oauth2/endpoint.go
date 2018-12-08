package oauth2

import (
	mpoauth2 "github.com/chanxuehong/wechat/mp/oauth2"
	"github.com/chanxuehong/wechat/oauth2"
)

var _ oauth2.Endpoint = (*Endpoint)(nil)

type Endpoint mpoauth2.Endpoint

func NewEndpoint(AppId, AppSecret string) *Endpoint {
	return (*Endpoint)(mpoauth2.NewEndpoint(AppId, AppSecret))
}

func (p *Endpoint) ExchangeTokenURL(code string) string {
	return ((*mpoauth2.Endpoint)(p)).ExchangeTokenURL(code)
}

func (p *Endpoint) RefreshTokenURL(refreshToken string) string {
	return ((*mpoauth2.Endpoint)(p)).RefreshTokenURL(refreshToken)
}

package oauth2

import (
	"net/url"

	"github.com/chanxuehong/wechat/oauth2"
)

var _ oauth2.Config = (*Config)(nil)

type Config struct {
	AppId     string
	AppSecret string
}

func NewConfig(AppId, AppSecret string) *Config {
	return &Config{
		AppId:     AppId,
		AppSecret: AppSecret,
	}
}

func (cfg *Config) ExchangeTokenURL(code string) string {
	return "https://api.weixin.qq.com/sns/oauth2/access_token?appid=" + url.QueryEscape(cfg.AppId) +
		"&secret=" + url.QueryEscape(cfg.AppSecret) +
		"&code=" + url.QueryEscape(code) +
		"&grant_type=authorization_code"
}

func (cfg *Config) RefreshTokenURL(refreshToken string) string {
	return "https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=" + url.QueryEscape(cfg.AppId) +
		"&grant_type=refresh_token&refresh_token=" + url.QueryEscape(refreshToken)
}

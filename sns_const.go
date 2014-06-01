package wechat

const (
	snsOAuth2AuthURL         = `https://open.weixin.qq.com/connect/oauth2/authorize`
	snsOAuth2TokenURL        = `https://api.weixin.qq.com/sns/oauth2/access_token`
	snsOAuth2RefreshTokenURL = `https://api.weixin.qq.com/sns/oauth2/refresh_token`
	// https://api.weixin.qq.com/sns/userinfo?access_token=ACCESS_TOKEN&openid=OPENID&lang=zh_CN
	snsUserInfoURLFormat = `https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=%s`
)

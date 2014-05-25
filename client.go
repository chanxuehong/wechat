package wechat

type Client struct {
	appid, appsecret string
	accessToken      accessToken

	//TODO: 以后增加必要的数据
}

func NewClient(appid, appsecret string) *Client {
	return &Client{
		appid:     appid,
		appsecret: appsecret,
	}
}

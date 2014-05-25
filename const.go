package wechat

const (
	// https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=APPID&secret=APPSECRET
	getAccessTokenUrlFormat = `https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s`
	// https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN
	customSendMessageUrlFormat = `https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=%s`
	// https://api.weixin.qq.com/cgi-bin/message/mass/sendall?access_token=ACCESS_TOKEN
	massSendMessageByGroupUrlFormat = `https://api.weixin.qq.com/cgi-bin/message/mass/sendall?access_token=%s`
	// https://api.weixin.qq.com/cgi-bin/message/mass/send?access_token=ACCESS_TOKEN
	massSendMessageByOpenIdUrlFormat = `https://api.weixin.qq.com/cgi-bin/message/mass/send?access_token=%s`
	// https://api.weixin.qq.com//cgi-bin/message/mass/delete?access_token=ACCESS_TOKEN
	massDeleteUrlFormat = `https://api.weixin.qq.com//cgi-bin/message/mass/delete?access_token=%s`
	// https://api.weixin.qq.com/cgi-bin/menu/create?access_token=ACCESS_TOKEN
	menuCreateUrlFormat = `https://api.weixin.qq.com/cgi-bin/menu/create?access_token=%s`
	// https://api.weixin.qq.com/cgi-bin/menu/get?access_token=ACCESS_TOKEN
	menuGetUrlFormat = `https://api.weixin.qq.com/cgi-bin/menu/get?access_token=%s`
	// https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=ACCESS_TOKEN
	menuDeleteUrlFormat = `https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=%s`
)

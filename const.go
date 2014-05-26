package wechat

const (
	// https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=APPID&secret=APPSECRET
	getAccessTokenUrlFormat = `https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s`
	// https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN
	customSendMessageUrlPrefix = `https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=`
	// https://api.weixin.qq.com/cgi-bin/message/mass/sendall?access_token=ACCESS_TOKEN
	massSendMessageByGroupUrlPrefix = `https://api.weixin.qq.com/cgi-bin/message/mass/sendall?access_token=`
	// https://api.weixin.qq.com/cgi-bin/message/mass/send?access_token=ACCESS_TOKEN
	massSendMessageByOpenIdUrlPrefix = `https://api.weixin.qq.com/cgi-bin/message/mass/send?access_token=`
	// https://api.weixin.qq.com//cgi-bin/message/mass/delete?access_token=ACCESS_TOKEN
	massDeleteUrlPrefix = `https://api.weixin.qq.com//cgi-bin/message/mass/delete?access_token=`
	// https://api.weixin.qq.com/cgi-bin/menu/create?access_token=ACCESS_TOKEN
	menuCreateUrlPrefix = `https://api.weixin.qq.com/cgi-bin/menu/create?access_token=`
	// https://api.weixin.qq.com/cgi-bin/menu/get?access_token=ACCESS_TOKEN
	menuGetUrlPrefix = `https://api.weixin.qq.com/cgi-bin/menu/get?access_token=`
	// https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=ACCESS_TOKEN
	menuDeleteUrlPrefix = `https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=`
)

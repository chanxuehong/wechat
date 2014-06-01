package wechat

import (
	"net/url"
)

// !!! 是不是所有的变量都要加 url.QueryEscape ? 知道的告诉我一声 !!!

// https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=APPID&secret=APPSECRET
func clientTokenGetURL(appid, appsecret string) string {
	return "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" +
		appid +
		"&secret=" +
		appsecret
}

// https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN
func clientMessageCustomSendURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/cgi-bin/message/mass/sendall?access_token=ACCESS_TOKEN
func clientMessageMassSendByGroupURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/message/mass/sendall?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/cgi-bin/message/mass/send?access_token=ACCESS_TOKEN
func clientMessageMassSendByOpenIdURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/message/mass/send?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com//cgi-bin/message/mass/delete?access_token=ACCESS_TOKEN
func clientMessageMassDeleteURL(accesstoken string) string {
	return "https://api.weixin.qq.com//cgi-bin/message/mass/delete?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/cgi-bin/menu/create?access_token=ACCESS_TOKEN
func clientMenuCreateURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/menu/create?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/cgi-bin/menu/get?access_token=ACCESS_TOKEN
func clientMenuGetURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/menu/get?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=ACCESS_TOKEN
func clientMenuDeleteURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=" +
		accesstoken
}

// http://file.api.weixin.qq.com/cgi-bin/media/upload?access_token=ACCESS_TOKEN&type=TYPE
func clientMediaUploadURL(accesstoken string, mediatype string) string {
	return "http://file.api.weixin.qq.com/cgi-bin/media/upload?access_token=" +
		accesstoken +
		"&type=" +
		mediatype
}

// http://file.api.weixin.qq.com/cgi-bin/media/get?access_token=ACCESS_TOKEN&media_id=MEDIA_ID
func clientMediaDownloadURL(accesstoken string, mediaid string) string {
	return "http://file.api.weixin.qq.com/cgi-bin/media/get?access_token=" +
		accesstoken +
		"&media_id=" +
		mediaid
}

// https://api.weixin.qq.com/cgi-bin/media/uploadnews?access_token=ACCESS_TOKEN
func clientMediaUploadNewsURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/media/uploadnews?access_token=" +
		accesstoken
}

// https://file.api.weixin.qq.com/cgi-bin/media/uploadvideo?access_token=ACCESS_TOKEN
func clientMediaUploadVideoURL(accesstoken string) string {
	return "https://file.api.weixin.qq.com/cgi-bin/media/uploadvideo?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=TOKEN
func clientQRCodeCreateURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=" +
		accesstoken
}

// https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=TICKET
func clientQRCodeURL(ticket string) string {
	return "https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=" +
		url.QueryEscape(ticket)
}

// https://api.weixin.qq.com/cgi-bin/groups/create?access_token=ACCESS_TOKEN
func clientUserGroupCreateURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/groups/create?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/cgi-bin/groups/get?access_token=ACCESS_TOKEN
func clientUserGroupGetURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/groups/get?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/cgi-bin/groups/update?access_token=ACCESS_TOKEN
func clientUserGroupRenameURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/groups/update?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/cgi-bin/groups/getid?access_token=ACCESS_TOKEN
func clientUserInWhichGroupURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/groups/getid?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/cgi-bin/groups/members/update?access_token=ACCESS_TOKEN
func clientUserMoveToGroupURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/groups/members/update?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/cgi-bin/user/info?access_token=ACCESS_TOKEN&openid=OPENID&lang=zh_CN
func clientUserInfoURL(accesstoken, openid, lang string) string {
	return "https://api.weixin.qq.com/cgi-bin/user/info?access_token=" +
		accesstoken +
		"&openid=" +
		openid +
		"&lang=" +
		lang
}

// https://api.weixin.qq.com/cgi-bin/user/get?access_token=ACCESS_TOKEN&next_openid=NEXT_OPENID
func clientUserGetURL(accesstoken, nextOpenId string) string {
	if nextOpenId == "" {
		return "https://api.weixin.qq.com/cgi-bin/user/get?access_token=" +
			accesstoken
	}
	return "https://api.weixin.qq.com/cgi-bin/user/get?access_token=" +
		accesstoken +
		"&next_openid=" +
		nextOpenId
}

// https://api.weixin.qq.com/cgi-bin/customservice/getrecord?access_token=ACCESS_TOKEN
func clientCSRecordGetURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/customservice/getrecord?access_token=" +
		accesstoken
}

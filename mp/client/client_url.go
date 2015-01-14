// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"net/url"
)

// https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN
func messageCustomSendURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=ACCESS_TOKEN
func messageTemplateSendURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/cgi-bin/message/mass/sendall?access_token=ACCESS_TOKEN
func messageMassSendToGroupURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/message/mass/sendall?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/cgi-bin/message/mass/sendall?access_token=ACCESS_TOKEN
func messageMassSendToAllURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/message/mass/sendall?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/cgi-bin/message/mass/send?access_token=ACCESS_TOKEN
func messageMassSendToUsersURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/message/mass/send?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/cgi-bin/message/mass/preview?access_token=ACCESS_TOKEN
func messageMassPreviewURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/message/mass/preview?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/cgi-bin/message/mass/get?access_token=ACCESS_TOKEN
func messageMassGetStatusURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/message/mass/get?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com//cgi-bin/message/mass/delete?access_token=ACCESS_TOKEN
func messageMassDeleteURL(accesstoken string) string {
	return "https://api.weixin.qq.com//cgi-bin/message/mass/delete?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/cgi-bin/menu/create?access_token=ACCESS_TOKEN
func menuCreateURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/menu/create?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/cgi-bin/menu/get?access_token=ACCESS_TOKEN
func menuGetURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/menu/get?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=ACCESS_TOKEN
func menuDeleteURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=" +
		accesstoken
}

// http://file.api.weixin.qq.com/cgi-bin/media/upload?access_token=ACCESS_TOKEN&type=TYPE
func mediaUploadURL(accesstoken string, mediatype string) string {
	return "http://file.api.weixin.qq.com/cgi-bin/media/upload?access_token=" +
		accesstoken +
		"&type=" +
		mediatype
}

// http://file.api.weixin.qq.com/cgi-bin/media/get?access_token=ACCESS_TOKEN&media_id=MEDIA_ID
func mediaDownloadURL(accesstoken string, mediaid string) string {
	return "http://file.api.weixin.qq.com/cgi-bin/media/get?access_token=" +
		accesstoken +
		"&media_id=" +
		mediaid
}

// https://api.weixin.qq.com/cgi-bin/media/uploadnews?access_token=ACCESS_TOKEN
func mediaCreateNewsURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/media/uploadnews?access_token=" +
		accesstoken
}

// https://file.api.weixin.qq.com/cgi-bin/media/uploadvideo?access_token=ACCESS_TOKEN
func mediaCreateVideoURL(accesstoken string) string {
	return "https://file.api.weixin.qq.com/cgi-bin/media/uploadvideo?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=TOKEN
func qrcodeCreateURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=" +
		accesstoken
}

// https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=TICKET
func qrcodePicURL(ticket string) string {
	return "https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=" +
		url.QueryEscape(ticket)
}

// https://api.weixin.qq.com/cgi-bin/groups/create?access_token=ACCESS_TOKEN
func userGroupCreateURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/groups/create?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/cgi-bin/groups/get?access_token=ACCESS_TOKEN
func userGroupGetURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/groups/get?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/cgi-bin/groups/update?access_token=ACCESS_TOKEN
func userGroupRenameURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/groups/update?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/cgi-bin/groups/getid?access_token=ACCESS_TOKEN
func userInWhichGroupURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/groups/getid?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/cgi-bin/groups/members/update?access_token=ACCESS_TOKEN
func userMoveToGroupURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/groups/members/update?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/cgi-bin/user/info?access_token=ACCESS_TOKEN&openid=OPENID&lang=zh_CN
func userInfoURL(accesstoken, openid, lang string) string {
	return "https://api.weixin.qq.com/cgi-bin/user/info?access_token=" +
		accesstoken +
		"&openid=" +
		openid +
		"&lang=" +
		lang
}

// https://api.weixin.qq.com/cgi-bin/user/get?access_token=ACCESS_TOKEN&next_openid=NEXT_OPENID
func userGetURL(accesstoken, nextOpenId string) string {
	if nextOpenId == "" {
		return "https://api.weixin.qq.com/cgi-bin/user/get?access_token=" +
			accesstoken
	}
	return "https://api.weixin.qq.com/cgi-bin/user/get?access_token=" +
		accesstoken +
		"&next_openid=" +
		nextOpenId
}

// https://api.weixin.qq.com/cgi-bin/user/info/updateremark?access_token=ACCESS_TOKEN
func userUpdateRemarkURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/user/info/updateremark?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/cgi-bin/customservice/getrecord?access_token=ACCESS_TOKEN
func customServiceRecordGetURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/customservice/getrecord?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/cgi-bin/customservice/getkflist?access_token= ACCESS_TOKEN
func customServiceKFListURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/customservice/getkflist?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/cgi-bin/customservice/getonlinekflist?access_token= ACCESS_TOKEN
func customServiceOnlineKFListURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/customservice/getonlinekflist?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/customservice/kfaccount/add?access_token=ACCESS_TOKEN
func customServiceKFAccountAddURL(accesstoken string) string {
	return "https://api.weixin.qq.com/customservice/kfaccount/add?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/customservice/kfaccount/update?access_token=ACCESS_TOKEN
func customServiceKFAccountSetURL(accesstoken string) string {
	return "https://api.weixin.qq.com/customservice/kfaccount/update?access_token=" +
		accesstoken
}

// http://api.weixin.qq.com/customservice/kfacount/uploadheadimg?access_token=ACCESS_TOKEN&kf_account=KFACCOUNT
func customServiceKFAccountUploadHeadImgURL(accesstoken, kfAccount string) string {
	return "http://api.weixin.qq.com/customservice/kfacount/uploadheadimg?access_token=" +
		accesstoken + "&kf_account=" + url.QueryEscape(kfAccount)
}

// https://api.weixin.qq.com/customservice/kfaccount/del?access_token=ACCESS_TOKEN&kf_account=KFACCOUNT
func customServiceKFAccountDeleteURL(accesstoken, kfAccount string) string {
	return "https://api.weixin.qq.com/customservice/kfaccount/del?access_token=" +
		accesstoken + "&kf_account=" + url.QueryEscape(kfAccount)
}

// https://api.weixin.qq.com/cgi-bin/shorturl?access_token=ACCESS_TOKEN
func shortURLURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/shorturl?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/cgi-bin/getcallbackip?access_token=ACCESS_TOKEN
func getCallbackIPURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/getcallbackip?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=ACCESS_TOKEN&type=jsapi
func ticketGetTicketURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=" + accesstoken + "&type=jsapi"
}

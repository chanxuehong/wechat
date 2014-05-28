package wechat

const postJSONContentType = "application/json; charset=utf-8"

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
	// http://file.api.weixin.qq.com/cgi-bin/media/upload?access_token=ACCESS_TOKEN&type=TYPE
	mediaUploadUrlFormat = `http://file.api.weixin.qq.com/cgi-bin/media/upload?access_token=%s&type=%s`
	// http://file.api.weixin.qq.com/cgi-bin/media/get?access_token=ACCESS_TOKEN&media_id=MEDIA_ID
	mediaDownloadUrlFormat = `http://file.api.weixin.qq.com/cgi-bin/media/get?access_token=%s&media_id=%s`
	// https://api.weixin.qq.com/cgi-bin/media/uploadnews?access_token=ACCESS_TOKEN
	mediaUploadNewsUrlPrefix = `https://api.weixin.qq.com/cgi-bin/media/uploadnews?access_token=`
	// https://file.api.weixin.qq.com/cgi-bin/media/uploadvideo?access_token=ACCESS_TOKEN
	mediaUploadVideoUrlPrefix = `https://file.api.weixin.qq.com/cgi-bin/media/uploadvideo?access_token=`
	// https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=TOKEN
	qrcodeCreateUrlPrefix = `https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=`
	// https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=TICKET
	qrcodeUrlPrefix = `https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=`
	// https://api.weixin.qq.com/cgi-bin/groups/create?access_token=ACCESS_TOKEN
	userGroupCreateUrlPrefix = `https://api.weixin.qq.com/cgi-bin/groups/create?access_token=`
	// https://api.weixin.qq.com/cgi-bin/groups/get?access_token=ACCESS_TOKEN
	userGroupGetUrlPrefix = `https://api.weixin.qq.com/cgi-bin/groups/get?access_token=`
	// https://api.weixin.qq.com/cgi-bin/groups/update?access_token=ACCESS_TOKEN
	userGroupRenameUrlPrefix = `https://api.weixin.qq.com/cgi-bin/groups/update?access_token=`
	// https://api.weixin.qq.com/cgi-bin/groups/getid?access_token=ACCESS_TOKEN
	userInWhichGroupUrlPrefix = `https://api.weixin.qq.com/cgi-bin/groups/getid?access_token=`
	// https://api.weixin.qq.com/cgi-bin/groups/members/update?access_token=ACCESS_TOKEN
	userMoveToGroupUrlPrefix = `https://api.weixin.qq.com/cgi-bin/groups/members/update?access_token=`
	// https://api.weixin.qq.com/cgi-bin/user/info?access_token=ACCESS_TOKEN&openid=OPENID&lang=zh_CN
	userInfoUrlFormat = `https://api.weixin.qq.com/cgi-bin/user/info?access_token=%s&openid=%s&lang=%s`
	// https://api.weixin.qq.com/cgi-bin/user/get?access_token=ACCESS_TOKEN&next_openid=NEXT_OPENID
	userGetUrlPrefix = `https://api.weixin.qq.com/cgi-bin/user/get?access_token=`
)

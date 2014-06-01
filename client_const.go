package wechat

const postJSONContentType = "application/json; charset=utf-8"

const (
	// https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=APPID&secret=APPSECRET
	clientTokenGetUrlFormat = `https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s`
	// https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN
	clientMessageCustomSendUrlPrefix = `https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=`
	// https://api.weixin.qq.com/cgi-bin/message/mass/sendall?access_token=ACCESS_TOKEN
	clientMessageMassSendByGroupUrlPrefix = `https://api.weixin.qq.com/cgi-bin/message/mass/sendall?access_token=`
	// https://api.weixin.qq.com/cgi-bin/message/mass/send?access_token=ACCESS_TOKEN
	clientMessageMassSendByOpenIdUrlPrefix = `https://api.weixin.qq.com/cgi-bin/message/mass/send?access_token=`
	// https://api.weixin.qq.com//cgi-bin/message/mass/delete?access_token=ACCESS_TOKEN
	clientMessageMassDeleteUrlPrefix = `https://api.weixin.qq.com//cgi-bin/message/mass/delete?access_token=`
	// https://api.weixin.qq.com/cgi-bin/menu/create?access_token=ACCESS_TOKEN
	clientMenuCreateUrlPrefix = `https://api.weixin.qq.com/cgi-bin/menu/create?access_token=`
	// https://api.weixin.qq.com/cgi-bin/menu/get?access_token=ACCESS_TOKEN
	clientMenuGetUrlPrefix = `https://api.weixin.qq.com/cgi-bin/menu/get?access_token=`
	// https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=ACCESS_TOKEN
	clientMenuDeleteUrlPrefix = `https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=`
	// http://file.api.weixin.qq.com/cgi-bin/media/upload?access_token=ACCESS_TOKEN&type=TYPE
	clientMediaUploadUrlFormat = `http://file.api.weixin.qq.com/cgi-bin/media/upload?access_token=%s&type=%s`
	// http://file.api.weixin.qq.com/cgi-bin/media/get?access_token=ACCESS_TOKEN&media_id=MEDIA_ID
	clientMediaDownloadUrlFormat = `http://file.api.weixin.qq.com/cgi-bin/media/get?access_token=%s&media_id=%s`
	// https://api.weixin.qq.com/cgi-bin/media/uploadnews?access_token=ACCESS_TOKEN
	clientMediaUploadNewsUrlPrefix = `https://api.weixin.qq.com/cgi-bin/media/uploadnews?access_token=`
	// https://file.api.weixin.qq.com/cgi-bin/media/uploadvideo?access_token=ACCESS_TOKEN
	clientMediaUploadVideoUrlPrefix = `https://file.api.weixin.qq.com/cgi-bin/media/uploadvideo?access_token=`
	// https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=TOKEN
	clientQRCodeCreateUrlPrefix = `https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=`
	// https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=TICKET
	clientQRCodeUrlPrefix = `https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=`
	// https://api.weixin.qq.com/cgi-bin/groups/create?access_token=ACCESS_TOKEN
	clientUserGroupCreateUrlPrefix = `https://api.weixin.qq.com/cgi-bin/groups/create?access_token=`
	// https://api.weixin.qq.com/cgi-bin/groups/get?access_token=ACCESS_TOKEN
	clientUserGroupGetUrlPrefix = `https://api.weixin.qq.com/cgi-bin/groups/get?access_token=`
	// https://api.weixin.qq.com/cgi-bin/groups/update?access_token=ACCESS_TOKEN
	clientUserGroupRenameUrlPrefix = `https://api.weixin.qq.com/cgi-bin/groups/update?access_token=`
	// https://api.weixin.qq.com/cgi-bin/groups/getid?access_token=ACCESS_TOKEN
	clientUserInWhichGroupUrlPrefix = `https://api.weixin.qq.com/cgi-bin/groups/getid?access_token=`
	// https://api.weixin.qq.com/cgi-bin/groups/members/update?access_token=ACCESS_TOKEN
	clientUserMoveToGroupUrlPrefix = `https://api.weixin.qq.com/cgi-bin/groups/members/update?access_token=`
	// https://api.weixin.qq.com/cgi-bin/user/info?access_token=ACCESS_TOKEN&openid=OPENID&lang=zh_CN
	clientUserInfoUrlFormat = `https://api.weixin.qq.com/cgi-bin/user/info?access_token=%s&openid=%s&lang=%s`
	// https://api.weixin.qq.com/cgi-bin/user/get?access_token=ACCESS_TOKEN&next_openid=NEXT_OPENID
	clientUserGetUrlPrefix = `https://api.weixin.qq.com/cgi-bin/user/get?access_token=`
	// https://api.weixin.qq.com/cgi-bin/customservice/getrecord?access_token=ACCESS_TOKEN
	clientCSRecordGetUrlPrefix = `https://api.weixin.qq.com/cgi-bin/customservice/getrecord?access_token=`
)

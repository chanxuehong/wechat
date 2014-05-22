package message

const (
	// 请求消息类型
	RQST_MSG_TYPE_TEXT     = "text"     // 文本消息
	RQST_MSG_TYPE_IMAGE    = "image"    // 图片消息
	RQST_MSG_TYPE_VOICE    = "voice"    // 语音消息
	RQST_MSG_TYPE_VIDEO    = "video"    // 视频消息
	RQST_MSG_TYPE_LOCATION = "location" // 地理位置消息
	RQST_MSG_TYPE_LINK     = "link"     // 链接消息
	RQST_MSG_TYPE_EVENT    = "event"    // 事件推送

	RQST_EVENT_TYPE_SUBSCRIBE         = "subscribe"         // 订阅, 包括点击订阅和扫描二维码
	RQST_EVENT_TYPE_UNSUBSCRIBE       = "unsubscribe"       // 取消订阅
	RQST_EVENT_TYPE_SCAN              = "SCAN"              // 已经订阅用户扫描二维码事件
	RQST_EVENT_TYPE_LOCATION          = "LOCATION"          // 上报地理位置事件
	RQST_EVENT_TYPE_CLICK             = "CLICK"             // 点击菜单拉取消息时的事件推送
	RQST_EVENT_TYPE_VIEW              = "VIEW"              // 点击菜单跳转链接时的事件推送
	RQST_EVENT_TYPE_MASSSENDJOBFINISH = "MASSSENDJOBFINISH" // 微信服务器推送群发结果
)

const (
	// 回复消息类型
	RESP_MSG_TYPE_TEXT  = "text"  // 文本消息
	RESP_MSG_TYPE_IMAGE = "image" // 图片消息
	RESP_MSG_TYPE_VOICE = "voice" // 语音消息
	RESP_MSG_TYPE_VIDEO = "video" // 视频消息
	RESP_MSG_TYPE_MUSIC = "music" // 音乐消息
	RESP_MSG_TYPE_NEWS  = "news"  // 图文消息
)

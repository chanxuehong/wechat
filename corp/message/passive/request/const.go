// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package request

const (
	// 微信服务器推送过来的消息类型
	MSG_TYPE_TEXT     = "text"
	MSG_TYPE_IMAGE    = "image"
	MSG_TYPE_VOICE    = "voice"
	MSG_TYPE_VIDEO    = "video"
	MSG_TYPE_LOCATION = "location"
	MSG_TYPE_EVENT    = "event"
)

const (
	// 微信服务器推送过来的事件类型
	EVENT_TYPE_SUBSCRIBE          = "subscribe"          // 关注
	EVENT_TYPE_UNSUBSCRIBE        = "unsubscribe"        // 取消关注
	EVENT_TYPE_CLICK              = "CLICK"              // 点击菜单拉取消息时的事件推送
	EVENT_TYPE_VIEW               = "VIEW"               // 点击菜单跳转链接时的事件推送
	EVENT_TYPE_LOCATION           = "LOCATION"           // 上报地理位置事件
	EVENT_TYPE_SCANCODE_PUSH      = "scancode_push"      // scancode_push：扫码推事件的事件推送
	EVENT_TYPE_SCANCODE_WAITMSG   = "scancode_waitmsg"   // scancode_waitmsg：扫码推事件且弹出“消息接收中”提示框的事件推送
	EVENT_TYPE_PIC_SYSPHOTO       = "pic_sysphoto"       // pic_sysphoto：弹出系统拍照发图的事件推送
	EVENT_TYPE_PIC_PHOTO_OR_ALBUM = "pic_photo_or_album" // pic_photo_or_album：弹出拍照或者相册发图的事件推送
	EVENT_TYPE_PIC_WEIXIN         = "pic_weixin"         // pic_weixin：弹出微信相册发图器的事件推送
	EVENT_TYPE_LOCATION_SELECT    = "location_select"    // location_select：弹出地理位置选择器的事件推送
)

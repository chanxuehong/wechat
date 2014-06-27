// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package response

const (
	NewsArticleCountLimit = 10 // 被动回复图文消息的文章数据最大数
)

const (
	// 被动回复的消息类型
	MSG_TYPE_TEXT                      = "text"                      // 文本消息
	MSG_TYPE_IMAGE                     = "image"                     // 图片消息
	MSG_TYPE_VOICE                     = "voice"                     // 语音消息
	MSG_TYPE_VIDEO                     = "video"                     // 视频消息
	MSG_TYPE_MUSIC                     = "music"                     // 音乐消息
	MSG_TYPE_NEWS                      = "news"                      // 图文消息
	MSG_TYPE_TRANSFER_CUSTOMER_SERVICE = "transfer_customer_service" // 将消息转发到多客服
)

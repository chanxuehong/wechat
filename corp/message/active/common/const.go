// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package common

const (
	NewsArticleCountLimit = 10

	NewsArticleShowCoverPicTrue  = 1
	NewsArticleShowCoverPicFalse = 0
)

const (
	MSG_TYPE_TEXT   = "text"
	MSG_TYPE_IMAGE  = "image"
	MSG_TYPE_VOICE  = "voice"
	MSG_TYPE_VIDEO  = "video"
	MSG_TYPE_FILE   = "file"
	MSG_TYPE_NEWS   = "news"
	MSG_TYPE_MPNEWS = "mpnews"
)

const (
	MSG_SAFE_TRUE  = 1
	MSG_SAFE_FALSE = 0
)

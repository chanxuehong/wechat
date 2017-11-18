// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://gopkg.in/chanxuehong/wechat.v1 for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/v1/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package corp

import (
	"net/http"
	"time"
)

// 一般请求的 http.Client
var TextHttpClient = &http.Client{
	Timeout: 60 * time.Second,
}

// 多媒体上传下载请求的 http.Client
var MediaHttpClient = &http.Client{
	Timeout: 300 * time.Second, // 因为目前微信支持最大的文件是 10MB, 请求超时时间保守设置为 300 秒
}

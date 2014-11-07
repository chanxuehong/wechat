// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

// copy from github.com/chanxuehong/wechat/mp/client/buffer_pool.go

package merchant

import (
	"bytes"
	"sync"
)

// 用于 Client 普通的文本操作
var textBufferPool = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 0, 16<<10)) // 默认 16KB
	},
}

// 用于 Client 多媒体操作
var mediaBufferPool = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 0, 10<<20)) // 默认 10MB
	},
}

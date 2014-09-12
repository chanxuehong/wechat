// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"bytes"
	"sync"
)

// 用于 Client 多媒体操作
var mediaBufferPool = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 10<<20)) // 默认 10MB
	},
}

// 用于 Client 普通的文本操作
var textBufferPool = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 64<<10)) // 默认 64KB
	},
}

func init() {
	// 预分配
	mediaBuf := mediaBufferPool.Get()
	mediaBufferPool.Put(mediaBuf)
	textBuf := textBufferPool.Get()
	textBufferPool.Put(textBuf)
}

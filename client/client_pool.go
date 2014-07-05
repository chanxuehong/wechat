// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"bytes"
)

func newBuffer() interface{} {
	return bytes.NewBuffer(make([]byte, 2<<20)) // 默认 2MB
}

func (c *Client) getBufferFromPool() (buf *bytes.Buffer) {
	buf = c.bufferPool.Get().(*bytes.Buffer)
	buf.Reset() // important!
	return
}

func (c *Client) putBufferToPool(buf *bytes.Buffer) {
	c.bufferPool.Put(buf)
}

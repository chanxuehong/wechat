package wechat

import (
	"bytes"
)

func newBuffer() interface{} {
	return bytes.NewBuffer(make([]byte, 1<<20)) // 默认 1M
}

func (c *Client) getBuffer() *bytes.Buffer {
	buf := c.bufferPool.Get().(*bytes.Buffer)
	buf.Reset() // important!
	return buf
}

func (c *Client) putBuffer(buf *bytes.Buffer) {
	c.bufferPool.Put(buf)
}

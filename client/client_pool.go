package client

import (
	"bytes"
)

func newBuffer() interface{} {
	return bytes.NewBuffer(make([]byte, 1<<20)) // 默认 1MB
}

func (c *Client) getBufferFromPool() *bytes.Buffer {
	buf := c.bufferPool.Get().(*bytes.Buffer)
	buf.Reset() // important!
	return buf
}

func (c *Client) putBufferToPool(buf *bytes.Buffer) {
	c.bufferPool.Put(buf)
}

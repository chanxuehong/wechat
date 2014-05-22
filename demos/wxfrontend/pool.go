package wxfrontend

import (
	"github.com/chanxuehong/util/pool"
	"github.com/chanxuehong/wechat/message"
)

var (
	requestMsgPool *pool.Pool // go1.3 有新的 sync.Pool
)

func newRequestMsg() interface{} {
	return new(message.RequestMsg)
}

// 工厂函数
func getRequestMsg() *message.RequestMsg {
	msg := requestMsgPool.Get().(*message.RequestMsg)

	return msg.Zero() // important!
}
func putRequestMsg(rqstMsg *message.RequestMsg) {
	requestMsgPool.Put(rqstMsg)
}

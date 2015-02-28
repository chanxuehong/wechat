// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay

type MessageServer interface {
	AppId() string
	MchId() string
	APIKey() string // API密钥

	MessageHandler() MessageHandler // 获取 MessageHandler
}

var _ MessageServer = new(DefaultMessageServer)

type DefaultMessageServer struct {
	appId  string
	mchId  string
	apiKey string

	messageHandler MessageHandler
}

func NewDefaultMessageServer(appId, mchId, apiKey string, handler MessageHandler) *DefaultMessageServer {
	if handler == nil {
		panic("pay: nil MessageHandler")
	}

	return &DefaultMessageServer{
		appId:          appId,
		mchId:          mchId,
		apiKey:         apiKey,
		messageHandler: handler,
	}
}

func (srv *DefaultMessageServer) AppId() string {
	return srv.appId
}
func (srv *DefaultMessageServer) MchId() string {
	return srv.mchId
}
func (srv *DefaultMessageServer) APIKey() string {
	return srv.apiKey
}
func (srv *DefaultMessageServer) MessageHandler() MessageHandler {
	return srv.messageHandler
}

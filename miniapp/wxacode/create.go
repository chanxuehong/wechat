package wxacode

import (
	"github.com/chanxuehong/wechat/mp/core"
)

func Create(clt *core.Client, request *QrcodeRequest) (data []byte, err error) {
	var result struct {
		core.Error
		// Buffer 图片 Buffer
		Buffer []byte `json:"buffer,omitempty"`
		// ContentType content-type
		ContentType string `json:"content_type,omitempty"`
	}
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/wxaapp/createwxaqrcode?access_token="
	if err = clt.PostJSON(incompleteURL, request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	data = result.Buffer
	return
}

package wxacode

import (
	"encoding/json"

	"github.com/bububa/wechat/mp/core"
)

func Get(clt *core.Client, request *QrcodeRequest) (data []byte, err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxa/getwxacode?access_token="
	data, err = PostJSON(clt, incompleteURL, &request)
	if err != nil {
		return
	}
	var result core.Error
	json.Unmarshal(data, &result)
	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	return
}

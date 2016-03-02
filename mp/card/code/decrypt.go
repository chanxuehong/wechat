package code

import (
	"github.com/chanxuehong/wechat/mp/core"
)

// Code解码接口
func Decrypt(clt *core.Client, encryptCode string) (code string, err error) {
	request := struct {
		EncryptCode string `json:"encrypt_code"`
	}{
		EncryptCode: encryptCode,
	}

	var result struct {
		core.Error
		Code string `json:"code"`
	}

	incompleteURL := "https://api.weixin.qq.com/card/code/decrypt?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	code = result.Code
	return
}

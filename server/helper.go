// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package server

import (
	"encoding/base64"
	"errors"
)

func AESKeyEncode(AESKey []byte) (encodedAESKey string, err error) {
	if len(AESKey) != 32 {
		err = errors.New("the length of AESKey must be equal to 32")
		return
	}
	tmp := base64.StdEncoding.EncodeToString(AESKey)
	encodedAESKey = tmp[:len(tmp)-1]
	return
}

func AESKeyDecode(encodedAESKey string) (AESKey []byte, err error) {
	if len(encodedAESKey) != 43 {
		err = errors.New("the length of encodedAESKey must be equal to 43")
		return
	}
	return base64.StdEncoding.DecodeString(encodedAESKey + "=")
}

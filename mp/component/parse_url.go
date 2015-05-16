// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package component

import (
	"errors"
	"net/url"
)

func parsePostURLQuery(queryValues url.Values) (timestamp, nonce, encryptType, msgSignature string, err error) {
	timestamp = queryValues.Get("timestamp")
	if timestamp == "" {
		err = errors.New("timestamp is empty")
		return
	}

	nonce = queryValues.Get("nonce")
	if nonce == "" {
		err = errors.New("nonce is empty")
		return
	}

	encryptType = queryValues.Get("encrypt_type")
	if encryptType == "" {
		err = errors.New("encrypt_type is empty")
		return
	}

	msgSignature = queryValues.Get("msg_signature")
	if msgSignature == "" {
		err = errors.New("msg_signature is empty")
		return
	}

	return
}

func parseGetURLQuery(queryValues url.Values) (signature, timestamp, nonce, echostr string, err error) {
	signature = queryValues.Get("signature")
	if signature == "" {
		err = errors.New("signature is empty")
		return
	}

	timestamp = queryValues.Get("timestamp")
	if timestamp == "" {
		err = errors.New("timestamp is empty")
		return
	}

	nonce = queryValues.Get("nonce")
	if nonce == "" {
		err = errors.New("nonce is empty")
		return
	}

	echostr = queryValues.Get("echostr")
	if echostr == "" {
		err = errors.New("echostr is empty")
		return
	}

	return
}

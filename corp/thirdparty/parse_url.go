// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package thirdparty

import (
	"errors"
	"net/url"
)

func parsePostURLQuery(urlValues url.Values) (msgSignature, timestamp, nonce string, err error) {
	msgSignature = urlValues.Get("msg_signature")
	if msgSignature == "" {
		err = errors.New("msg_signature is empty")
		return
	}

	timestamp = urlValues.Get("timestamp")
	if timestamp == "" {
		err = errors.New("timestamp is empty")
		return
	}

	nonce = urlValues.Get("nonce")
	if nonce == "" {
		err = errors.New("nonce is empty")
		return
	}

	return
}

func parseGetURLQuery(urlValues url.Values) (msgSignature, timestamp, nonce, echostr string, err error) {
	msgSignature = urlValues.Get("msg_signature")
	if msgSignature == "" {
		err = errors.New("msg_signature is empty")
		return
	}

	timestamp = urlValues.Get("timestamp")
	if timestamp == "" {
		err = errors.New("timestamp is empty")
		return
	}

	nonce = urlValues.Get("nonce")
	if nonce == "" {
		err = errors.New("nonce is empty")
		return
	}

	echostr = urlValues.Get("echostr")
	if echostr == "" {
		err = errors.New("echostr is empty")
		return
	}

	return
}

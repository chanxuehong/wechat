// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package mp

import (
	"errors"
	"net/url"
)

func parsePostURLQuery(urlValues url.Values) (signature, timestamp, nonce,
	encryptType, msgSignature string, err error) {

	signature = urlValues.Get("signature")
	if signature == "" {
		err = errors.New("signature is empty")
		return
	}

	timestamp = urlValues.Get("timestamp")
	if timestamp == "" {
		err = errors.New("timestamp is empty")
		return
	}
	
	// verify the timestamp
	if err = util.VtimestampStr(timestampStr, 30); err !=nil {
		return
	}

	nonce = urlValues.Get("nonce")
	if nonce == "" {
		err = errors.New("nonce is empty")
		return
	}

	encryptType = urlValues.Get("encrypt_type")
	msgSignature = urlValues.Get("msg_signature")

	return
}

func parseGetURLQuery(urlValues url.Values) (signature, timestamp, nonce, echostr string, err error) {
	signature = urlValues.Get("signature")
	if signature == "" {
		err = errors.New("signature is empty")
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

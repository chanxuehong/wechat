// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package server

import (
	"errors"
	"net/url"
)

func parsePostURLQuery(URL *url.URL) (signature, timestamp, nonce, encryptType, msgSignature string, err error) {
	if URL == nil {
		err = errors.New("URL == nil")
		return
	}

	urlValues, err := url.ParseQuery(URL.RawQuery)
	if err != nil {
		return
	}

	signature = urlValues.Get("signature")
	//if signature == "" {
	//	err = errors.New("signature is empty")
	//	return
	//}

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

	encryptType = urlValues.Get("encrypt_type")
	msgSignature = urlValues.Get("msg_signature")

	return
}

func parsePostURLQueryEx(URL *url.URL) (agentkey, signature, timestamp, nonce, encryptType, msgSignature string, err error) {
	if URL == nil {
		err = errors.New("URL == nil")
		return
	}

	urlValues, err := url.ParseQuery(URL.RawQuery)
	if err != nil {
		return
	}

	agentkey = urlValues.Get(URLQueryAgentKeyName)
	if agentkey == "" {
		err = errors.New(URLQueryAgentKeyName + " is empty")
		return
	}

	signature = urlValues.Get("signature")
	//if signature == "" {
	//	err = errors.New("signature is empty")
	//	return
	//}

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

	encryptType = urlValues.Get("encrypt_type")
	msgSignature = urlValues.Get("msg_signature")

	return
}

func parseGetURLQuery(URL *url.URL) (signature, timestamp, nonce, echostr string, err error) {
	if URL == nil {
		err = errors.New("URL == nil")
		return
	}

	urlValues, err := url.ParseQuery(URL.RawQuery)
	if err != nil {
		return
	}

	signature = urlValues.Get("signature")
	//if signature == "" {
	//	err = errors.New("signature is empty")
	//	return
	//}

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

func parseGetURLQueryEx(URL *url.URL) (agentkey, signature, timestamp, nonce, echostr string, err error) {
	if URL == nil {
		err = errors.New("URL == nil")
		return
	}

	urlValues, err := url.ParseQuery(URL.RawQuery)
	if err != nil {
		return
	}

	agentkey = urlValues.Get(URLQueryAgentKeyName)
	if agentkey == "" {
		err = errors.New(URLQueryAgentKeyName + " is empty")
		return
	}

	signature = urlValues.Get("signature")
	//if signature == "" {
	//	err = errors.New("signature is empty")
	//	return
	//}

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

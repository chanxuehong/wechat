// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"bytes"
	"io/ioutil"
)

var _test_client = func() *Client {
	// 填入正确的 appid, appsecret
	clt := NewClient("appid", "appsecret", nil)

	// 预热
	buf := clt.getBufferFromPool()
	clt.putBufferToPool(buf)

	return clt
}()

// 比较两个文件是否相等
func fileEqual(filepath1, filepath2 string) (bool, error) {
	// 因为文件不大, 一次性读入内存
	b1, err := ioutil.ReadFile(filepath1)
	if err != nil {
		return false, err
	}
	b2, err := ioutil.ReadFile(filepath2)
	if err != nil {
		return false, err
	}

	return bytes.Equal(b1, b2), nil
}

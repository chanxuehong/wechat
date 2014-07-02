// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"testing"
)

// 获取当前的 token
func TestToken(t *testing.T) {
	tk, err := _test_client.Token()
	if err != nil {
		t.Error(err)
		return
	}

	if tk == "" {
		t.Error(`token == ""`)
		return
	}
}

func TestTokenRefresh(t *testing.T) {
	tk0, err := _test_client.TokenRefresh()
	if err != nil {
		t.Error(err)
		return
	}

	tk1, err := _test_client.Token()
	if err != nil || tk0 != tk1 {
		t.Error("逻辑错误: TokenRefresh 成功运行, 但是没有更新当前的 token 值.")
		return
	}
}

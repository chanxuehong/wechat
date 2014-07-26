// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package tokencache

import (
	"encoding/json"
	"errors"
	"github.com/bradfitz/gomemcache/memcache"
)

type MemCacheTokenCache struct {
	appid    string
	mcClient *memcache.Client
}

func NewMemCacheTokenCache(appid string, mcServer ...string) *MemCacheTokenCache {
	mcClient := memcache.New(mcServer...)
	return &MemCacheTokenCache{
		appid:    appid,
		mcClient: mcClient,
	}
}

// 正常情况下 ErrMsg == ""
type Token struct {
	Value  string `json:"value"`
	ErrMsg string `json:"errmsg"`
}

func (clt *MemCacheTokenCache) Token() (token string, err error) {
	item, err := clt.mcClient.Get(clt.appid)
	if err != nil {
		return
	}

	var tk Token
	if err = json.Unmarshal(item.Value, &tk); err != nil {
		return
	}

	if tk.ErrMsg != "" {
		err = errors.New(tk.ErrMsg)
		return
	}

	if tk.Value == "" {
		err = errors.New("token is empty")
		return
	}

	token = tk.Value
	return
}

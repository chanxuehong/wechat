package qrcode

import (
	"github.com/bububa/wechat/mp/base"
	"github.com/bububa/wechat/mp/core"
)

// ShortURL 将一条长链接转成短链接.
func ShortURL(clt *core.Client, longURL string) (shortURL string, err error) {
	return base.ShortURL(clt, longURL)
}

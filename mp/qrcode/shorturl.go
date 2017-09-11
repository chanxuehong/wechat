package qrcode

import (
	"github.com/mingjunyang/wechat.v2/mp/base"
	"github.com/mingjunyang/wechat.v2/mp/core"
)

// ShortURL 将一条长链接转成短链接.
func ShortURL(clt *core.Client, longURL string) (shortURL string, err error) {
	return base.ShortURL(clt, longURL)
}

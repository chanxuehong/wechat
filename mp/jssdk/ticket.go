package jssdk

import (
	"errors"
	"strconv"

	"github.com/chanxuehong/wechat/mp/core"
)

// updateTicket 从微信服务器获取新的 jsapi_ticket 并存入缓存, 同时返回该 jsapi_ticket.
func UpdateTicket(ctl *core.Client) (ticket string, expiresIn int64, err error) {
	var incompleteURL = "https://api.weixin.qq.com/cgi-bin/ticket/getticket?type=jsapi&access_token="
	var result struct {
		core.Error
		jsapiTicket
	}
	if err = ctl.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}

	// 由于网络的延时, jsapi_ticket 过期时间留有一个缓冲区
	switch {
	case result.ExpiresIn > 31556952: // 60*60*24*365.2425
		err = errors.New("expires_in too large: " + strconv.FormatInt(result.ExpiresIn, 10))
		return
	case result.ExpiresIn > 60*60:
		result.ExpiresIn -= 60 * 10
	case result.ExpiresIn > 60*30:
		result.ExpiresIn -= 60 * 5
	case result.ExpiresIn > 60*5:
		result.ExpiresIn -= 60
	case result.ExpiresIn > 60:
		result.ExpiresIn -= 10
	default:
		err = errors.New("expires_in too small: " + strconv.FormatInt(result.ExpiresIn, 10))
		return
	}
	return result.jsapiTicket.Ticket, result.jsapiTicket.ExpiresIn, nil
}

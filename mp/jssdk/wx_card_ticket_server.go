// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package jssdk

import (
	"errors"
	"strconv"
	"sync"
	"time"

	"github.com/chanxuehong/wechat/mp"
)

var _ TicketServer = (*WxCardTicketServer)(nil)

// TicketServer 的简单实现.
//  NOTE:
//  1. 用于单进程环境.
//  2. 因为 WxCardTicketServer 同时也是一个简单的中控服务器, 而不是仅仅实现 TicketServer 接口,
//     所以整个系统只能存在一个 WxCardTicketServer 实例!
type WxCardTicketServer struct {
	mpClient *mp.Client

	resetTickerChan chan time.Duration // 用于重置 ticketDaemon 里的 ticker

	ticketGet struct {
		sync.Mutex
		LastTicketInfo ticketInfo // 最后一次成功从微信服务器获取的 jsapi_ticket 信息
		LastTimestamp  int64      // 最后一次成功从微信服务器获取 jsapi_ticket 的时间戳
	}

	ticketCache struct {
		sync.RWMutex
		Ticket string
	}
}

// 创建一个新的 WxCardTicketServer.
func NewWxCardTicketServer(clt *mp.Client) (srv *WxCardTicketServer) {
	if clt == nil {
		panic("nil mp.Client")
	}

	srv = &WxCardTicketServer{
		mpClient:        clt,
		resetTickerChan: make(chan time.Duration),
	}

	go srv.ticketDaemon(time.Hour * 24) // 启动 tokenDaemon
	return
}

func (srv *WxCardTicketServer) TagB38894EBFE9911E4BE17A4DB30FED8E1() {}

func (srv *WxCardTicketServer) Ticket() (ticket string, err error) {
	srv.ticketCache.RLock()
	ticket = srv.ticketCache.Ticket
	srv.ticketCache.RUnlock()

	if ticket != "" {
		return
	}
	return srv.TicketRefresh()
}

func (srv *WxCardTicketServer) TicketRefresh() (ticket string, err error) {
	ticketInfo, cached, err := srv.getTicket()
	if err != nil {
		return
	}
	if !cached {
		srv.resetTickerChan <- time.Duration(ticketInfo.ExpiresIn) * time.Second
	}
	ticket = ticketInfo.Ticket
	return
}

func (srv *WxCardTicketServer) ticketDaemon(tickDuration time.Duration) {
NEW_TICK_DURATION:
	ticker := time.NewTicker(tickDuration)

	for {
		select {
		case tickDuration = <-srv.resetTickerChan:
			ticker.Stop()
			goto NEW_TICK_DURATION

		case <-ticker.C:
			ticketInfo, cached, err := srv.getTicket()
			if err != nil {
				break
			}
			if !cached {
				newTickDuration := time.Duration(ticketInfo.ExpiresIn) * time.Second
				if tickDuration != newTickDuration {
					tickDuration = newTickDuration
					ticker.Stop()
					goto NEW_TICK_DURATION
				}
			}
		}
	}
}

// 从微信服务器获取 jsapi_ticket.
//  同一时刻只能一个 goroutine 进入, 防止没必要的重复获取.
func (srv *WxCardTicketServer) getTicket() (ticket ticketInfo, cached bool, err error) {
	srv.ticketGet.Lock()
	defer srv.ticketGet.Unlock()

	timeNowUnix := time.Now().Unix()

	// 在收敛周期内直接返回最近一次获取的 jsapi_ticket, 这里的收敛时间设定为4秒
	if n := srv.ticketGet.LastTimestamp; n <= timeNowUnix && timeNowUnix < n+4 {
		// 因为只有成功获取后才会更新 srv.tokenGet.LastTimestamp, 所以这些都是有效数据
		ticket = ticketInfo{
			Ticket:    srv.ticketGet.LastTicketInfo.Ticket,
			ExpiresIn: srv.ticketGet.LastTicketInfo.ExpiresIn - timeNowUnix + n,
		}
		cached = true
		return
	}

	var result struct {
		mp.Error
		ticketInfo
	}

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/ticket/getticket?type=wx_card&access_token="
	if err = srv.mpClient.GetJSON(incompleteURL, &result); err != nil {
		srv.ticketCache.Lock()
		srv.ticketCache.Ticket = ""
		srv.ticketCache.Unlock()
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		srv.ticketCache.Lock()
		srv.ticketCache.Ticket = ""
		srv.ticketCache.Unlock()

		err = &result.Error
		return
	}

	// 由于网络的延时, jsapi_ticket 过期时间留了一个缓冲区
	switch {
	case result.ExpiresIn > 31556952: // 60*60*24*365.2425
		srv.ticketCache.Lock()
		srv.ticketCache.Ticket = ""
		srv.ticketCache.Unlock()

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
		srv.ticketCache.Lock()
		srv.ticketCache.Ticket = ""
		srv.ticketCache.Unlock()

		err = errors.New("expires_in too small: " + strconv.FormatInt(result.ExpiresIn, 10))
		return
	}

	srv.ticketGet.LastTicketInfo = result.ticketInfo
	srv.ticketGet.LastTimestamp = timeNowUnix

	srv.ticketCache.Lock()
	srv.ticketCache.Ticket = result.ticketInfo.Ticket
	srv.ticketCache.Unlock()

	ticket = result.ticketInfo
	return
}

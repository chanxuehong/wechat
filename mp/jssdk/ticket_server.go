package jssdk

import (
	"errors"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/chanxuehong/wechat/mp/core"
)

// jsapi_ticket 中控服务器接口.
type TicketServer interface {
	Ticket() (ticket string, err error)        // 请求中控服务器返回缓存的 jsapi_ticket
	RefreshTicket() (ticket string, err error) // 请求中控服务器刷新 jsapi_ticket
	IIDB04E44A0E1DC11E5ADCEA4DB30FED8E1()      // 接口标识, 没有实际意义
}

var _ TicketServer = (*DefaultTicketServer)(nil)

// DefaultTicketServer 实现了 TicketServer 接口.
//  NOTE:
//  1. 用于单进程环境.
//  2. 因为 DefaultTicketServer 同时也是一个简单的中控服务器, 而不是仅仅实现 TicketServer 接口,
//     所以整个系统只能存在一个 DefaultTicketServer 实例!
type DefaultTicketServer struct {
	coreClient *core.Client

	ticker          *time.Ticker       // 用于定时更新 jsapi_ticket
	resetTickerChan chan time.Duration // 用于重置 ticker

	ticketGet struct {
		sync.Mutex
		lastTicket    jsapiTicket // 最近一次成功更新 jsapi_ticket 的数据信息
		lastTimestamp time.Time   // 最近一次成功更新 jsapi_ticket 的时间戳
	}

	ticketCache struct {
		sync.RWMutex
		ticket string // 保存有效的 jsapi_ticket, 当更新 jsapi_ticket 失败时此字段置空
	}
}

// 微信服务器返回的 jsapi_ticket 的数据结构.
type jsapiTicket struct {
	Ticket    string `json:"ticket"`
	ExpiresIn int64  `json:"expires_in"`
}

// NewDefaultTicketServer 创建一个新的 DefaultTicketServer.
func NewDefaultTicketServer(clt *core.Client) (srv *DefaultTicketServer) {
	if clt == nil {
		panic("nil core.Client")
	}
	srv = &DefaultTicketServer{
		coreClient:      clt,
		ticker:          nil,
		resetTickerChan: make(chan time.Duration),
	}

	go srv.ticketUpdateDaemon(time.Hour * 24 * time.Duration(100+rand.Int63n(200)))
	return
}

func (srv *DefaultTicketServer) IIDB04E44A0E1DC11E5ADCEA4DB30FED8E1() {}

func (srv *DefaultTicketServer) Ticket() (ticket string, err error) {
	srv.ticketCache.RLock()
	ticket = srv.ticketCache.ticket
	srv.ticketCache.RUnlock()

	if ticket != "" {
		return
	}
	return srv.RefreshTicket()
}

func (srv *DefaultTicketServer) RefreshTicket() (ticket string, err error) {
	jsapiTicket, cached, err := srv.refreshTicket()
	if err != nil {
		return
	}
	if !cached {
		srv.resetTickerChan <- time.Duration(jsapiTicket.ExpiresIn) * time.Second
	}
	ticket = jsapiTicket.Ticket
	return
}

func (srv *DefaultTicketServer) ticketUpdateDaemon(initTickDuration time.Duration) {
	tickDuration := initTickDuration

NEW_TICK_DURATION:
	srv.ticker = time.NewTicker(tickDuration)
	for {
		select {
		case tickDuration = <-srv.resetTickerChan:
			srv.ticker.Stop()
			goto NEW_TICK_DURATION

		case <-srv.ticker.C:
			jsapiTicket, cached, err := srv.refreshTicket()
			if err != nil {
				break
			}
			if !cached {
				newTickDuration := time.Duration(jsapiTicket.ExpiresIn) * time.Second
				if tickDuration != newTickDuration {
					tickDuration = newTickDuration
					srv.ticker.Stop()
					goto NEW_TICK_DURATION
				}
			}
		}
	}
}

// refreshTicket 从微信服务器获取新的 jsapi_ticket 并存入缓存, 同时返回该 jsapi_ticket.
func (srv *DefaultTicketServer) refreshTicket() (ticket jsapiTicket, cached bool, err error) {
	srv.ticketGet.Lock()
	defer srv.ticketGet.Unlock()

	timeNow := time.Now()

	// 如果在收敛周期内则直接返回最近一次获取的 jsapi_ticket!
	//
	// 当 jsapi_ticket 失效时, SDK 内部会尝试刷新 jsapi_ticket, 这时在短时间(一个收敛周期)内可能会有多个 goroutine 同时进行刷新,
	// 实际上我们没有必要也不能让这些操作都去微信服务器获取新的 jsapi_ticket, 而只用返回同一个 jsapi_ticket 即可.
	// 因为 jsapi_ticket 缓存在内存里, 所以收敛周期为3个http周期, 我们这里设定为3秒.
	if d := timeNow.Sub(srv.ticketGet.lastTimestamp); 0 <= d && d < time.Second*3 {
		ticket = srv.ticketGet.lastTicket
		cached = true
		return
	}

	var incompleteURL = "https://api.weixin.qq.com/cgi-bin/ticket/getticket?type=jsapi&access_token="
	var result struct {
		core.Error
		jsapiTicket
	}
	if err = srv.coreClient.GetJSON(incompleteURL, &result); err != nil {
		srv.ticketCache.Lock()
		srv.ticketCache.ticket = ""
		srv.ticketCache.Unlock()
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		srv.ticketCache.Lock()
		srv.ticketCache.ticket = ""
		srv.ticketCache.Unlock()
		err = &result.Error
		return
	}

	// 由于网络的延时, jsapi_ticket 过期时间留有一个缓冲区
	switch {
	case result.ExpiresIn > 31556952: // 60*60*24*365.2425
		srv.ticketCache.Lock()
		srv.ticketCache.ticket = ""
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
		srv.ticketCache.ticket = ""
		srv.ticketCache.Unlock()
		err = errors.New("expires_in too small: " + strconv.FormatInt(result.ExpiresIn, 10))
		return
	}

	// 更新 ticketGet
	srv.ticketGet.lastTicket = result.jsapiTicket
	srv.ticketGet.lastTimestamp = timeNow

	// 更新 ticketCache
	srv.ticketCache.Lock()
	srv.ticketCache.ticket = result.Ticket
	srv.ticketCache.Unlock()

	ticket = result.jsapiTicket
	return
}

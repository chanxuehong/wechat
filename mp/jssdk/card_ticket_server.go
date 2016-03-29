package jssdk

import (
	"errors"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/chanxuehong/wechat/mp/core"
)

// 卡劵 api_ticket 中控服务器接口.
type CardTicketServer interface {
	Ticket() (ticket string, err error)        // 请求中控服务器返回缓存的卡劵 api_ticket
	RefreshTicket() (ticket string, err error) // 请求中控服务器刷新卡劵 api_ticket
	IIDB9BDD0A1E1DC11E5844AA4DB30FED8E1()      // 接口标识, 没有实际意义
}

var _ CardTicketServer = (*DefaultCardTicketServer)(nil)

// DefaultCardTicketServer 实现了 CardTicketServer 接口.
//  NOTE:
//  1. 用于单进程环境.
//  2. 因为 DefaultCardTicketServer 同时也是一个简单的中控服务器, 而不是仅仅实现 CardTicketServer 接口,
//     所以整个系统只能存在一个 DefaultCardTicketServer 实例!
type DefaultCardTicketServer struct {
	coreClient *core.Client

	ticker          *time.Ticker       // 用于定时更新卡劵 api_ticket
	resetTickerChan chan time.Duration // 用于重置 ticker

	ticketGet struct {
		sync.Mutex
		lastTicket    cardApiTicket // 最近一次成功更新卡劵 api_ticket 的数据信息
		lastTimestamp time.Time     // 最近一次成功更新卡劵 api_ticket 的时间戳
	}

	ticketCache struct {
		sync.RWMutex
		ticket string // 保存有效的卡劵 api_ticket, 当更新卡劵 api_ticket 失败时此字段置空
	}
}

// 微信服务器返回的卡劵 api_ticket 的数据结构.
type cardApiTicket struct {
	Ticket    string `json:"ticket"`
	ExpiresIn int64  `json:"expires_in"`
}

// NewDefaultCardTicketServer 创建一个新的 DefaultCardTicketServer.
func NewDefaultCardTicketServer(clt *core.Client) (srv *DefaultCardTicketServer) {
	if clt == nil {
		panic("nil core.Client")
	}
	srv = &DefaultCardTicketServer{
		coreClient:      clt,
		ticker:          nil,
		resetTickerChan: make(chan time.Duration),
	}

	go srv.ticketUpdateDaemon(time.Hour * 24 * time.Duration(100+rand.Int63n(200)))
	return
}

func (srv *DefaultCardTicketServer) IIDB9BDD0A1E1DC11E5844AA4DB30FED8E1() {}

func (srv *DefaultCardTicketServer) Ticket() (ticket string, err error) {
	srv.ticketCache.RLock()
	ticket = srv.ticketCache.ticket
	srv.ticketCache.RUnlock()

	if ticket != "" {
		return
	}
	return srv.RefreshTicket()
}

func (srv *DefaultCardTicketServer) RefreshTicket() (ticket string, err error) {
	cardApiTicket, cached, err := srv.refreshTicket()
	if err != nil {
		return
	}
	if !cached {
		srv.resetTickerChan <- time.Duration(cardApiTicket.ExpiresIn) * time.Second
	}
	ticket = cardApiTicket.Ticket
	return
}

func (srv *DefaultCardTicketServer) ticketUpdateDaemon(initTickDuration time.Duration) {
	tickDuration := initTickDuration

NEW_TICK_DURATION:
	srv.ticker = time.NewTicker(tickDuration)
	for {
		select {
		case tickDuration = <-srv.resetTickerChan:
			srv.ticker.Stop()
			goto NEW_TICK_DURATION

		case <-srv.ticker.C:
			cardApiTicket, cached, err := srv.refreshTicket()
			if err != nil {
				break
			}
			if !cached {
				newTickDuration := time.Duration(cardApiTicket.ExpiresIn) * time.Second
				if tickDuration != newTickDuration {
					tickDuration = newTickDuration
					srv.ticker.Stop()
					goto NEW_TICK_DURATION
				}
			}
		}
	}
}

// refreshTicket 从微信服务器获取新的卡劵 api_ticket 并存入缓存, 同时返回该卡劵 api_ticket.
func (srv *DefaultCardTicketServer) refreshTicket() (ticket cardApiTicket, cached bool, err error) {
	srv.ticketGet.Lock()
	defer srv.ticketGet.Unlock()

	timeNow := time.Now()

	// 如果在收敛周期内则直接返回最近一次获取的卡劵 api_ticket!
	//
	// 当卡劵 api_ticket 失效时, SDK 内部会尝试刷新卡劵 api_ticket, 这时在短时间(一个收敛周期)内可能会有多个 goroutine 同时进行刷新,
	// 实际上我们没有必要也不能让这些操作都去微信服务器获取新的卡劵 api_ticket, 而只用返回同一个卡劵 api_ticket 即可.
	// 因为卡劵 api_ticket 缓存在内存里, 所以收敛周期为3个http周期, 我们这里设定为3秒.
	if d := timeNow.Sub(srv.ticketGet.lastTimestamp); 0 <= d && d < time.Second*3 {
		ticket = srv.ticketGet.lastTicket
		cached = true
		return
	}

	var incompleteURL = "https://api.weixin.qq.com/cgi-bin/ticket/getticket?type=wx_card&access_token="
	var result struct {
		core.Error
		cardApiTicket
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

	// 由于网络的延时, 卡劵 api_ticket 过期时间留有一个缓冲区
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
	srv.ticketGet.lastTicket = result.cardApiTicket
	srv.ticketGet.lastTimestamp = timeNow

	// 更新 ticketCache
	srv.ticketCache.Lock()
	srv.ticketCache.ticket = result.Ticket
	srv.ticketCache.Unlock()

	ticket = result.cardApiTicket
	return
}

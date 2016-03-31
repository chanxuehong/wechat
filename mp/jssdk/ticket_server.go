package jssdk

import (
	"errors"
	"math/rand"
	"strconv"
	"sync/atomic"
	"time"
	"unsafe"

	"github.com/chanxuehong/wechat/mp/core"
)

// jsapi_ticket 中控服务器接口.
type TicketServer interface {
	Ticket() (ticket string, err error)                            // 请求中控服务器返回缓存的 jsapi_ticket
	RefreshTicket(currentTicket string) (ticket string, err error) // 请求中控服务器刷新 jsapi_ticket
	IIDB04E44A0E1DC11E5ADCEA4DB30FED8E1()                          // 接口标识, 没有实际意义
}

var _ TicketServer = (*DefaultTicketServer)(nil)

// DefaultTicketServer 实现了 TicketServer 接口.
//  NOTE:
//  1. 用于单进程环境.
//  2. 因为 DefaultTicketServer 同时也是一个简单的中控服务器, 而不是仅仅实现 TicketServer 接口,
//     所以整个系统只能存在一个 DefaultTicketServer 实例!
type DefaultTicketServer struct {
	coreClient *core.Client

	refreshTicketRequestChan  chan string              // chan currentTicket
	refreshTicketResponseChan chan refreshTicketResult // chan {ticket, err}

	ticketCache unsafe.Pointer // *jsapiTicket
}

// NewDefaultTicketServer 创建一个新的 DefaultTicketServer.
func NewDefaultTicketServer(clt *core.Client) (srv *DefaultTicketServer) {
	if clt == nil {
		panic("nil core.Client")
	}
	srv = &DefaultTicketServer{
		coreClient:                clt,
		refreshTicketRequestChan:  make(chan string),
		refreshTicketResponseChan: make(chan refreshTicketResult),
	}

	go srv.ticketUpdateDaemon(time.Hour * 24 * time.Duration(100+rand.Int63n(200)))
	return
}

func (srv *DefaultTicketServer) IIDB04E44A0E1DC11E5ADCEA4DB30FED8E1() {}

func (srv *DefaultTicketServer) Ticket() (ticket string, err error) {
	if p := (*jsapiTicket)(atomic.LoadPointer(&srv.ticketCache)); p != nil {
		return p.Ticket, nil
	}
	return srv.RefreshTicket("")
}

type refreshTicketResult struct {
	ticket string
	err    error
}

func (srv *DefaultTicketServer) RefreshTicket(currentTicket string) (ticket string, err error) {
	srv.refreshTicketRequestChan <- currentTicket
	rslt := <-srv.refreshTicketResponseChan
	return rslt.ticket, rslt.err
}

func (srv *DefaultTicketServer) ticketUpdateDaemon(initTickDuration time.Duration) {
	tickDuration := initTickDuration

NEW_TICK_DURATION:
	ticker := time.NewTicker(tickDuration)
	for {
		select {
		case currentTicket := <-srv.refreshTicketRequestChan:
			jsapiTicket, cached, err := srv.updateTicket(currentTicket)
			if err != nil {
				srv.refreshTicketResponseChan <- refreshTicketResult{err: err}
				break
			}
			srv.refreshTicketResponseChan <- refreshTicketResult{ticket: jsapiTicket.Ticket}
			if !cached {
				tickDuration = time.Duration(jsapiTicket.ExpiresIn) * time.Second
				ticker.Stop()
				goto NEW_TICK_DURATION
			}

		case <-ticker.C:
			jsapiTicket, _, err := srv.updateTicket("")
			if err != nil {
				break
			}
			newTickDuration := time.Duration(jsapiTicket.ExpiresIn) * time.Second
			if abs(tickDuration-newTickDuration) > time.Second*5 {
				tickDuration = newTickDuration
				ticker.Stop()
				goto NEW_TICK_DURATION
			}
		}
	}
}

func abs(x time.Duration) time.Duration {
	if x >= 0 {
		return x
	}
	return -x
}

type jsapiTicket struct {
	Ticket    string `json:"ticket"`
	ExpiresIn int64  `json:"expires_in"`
}

// updateTicket 从微信服务器获取新的 jsapi_ticket 并存入缓存, 同时返回该 jsapi_ticket.
func (srv *DefaultTicketServer) updateTicket(currentTicket string) (ticket *jsapiTicket, cached bool, err error) {
	if currentTicket != "" {
		if p := (*jsapiTicket)(atomic.LoadPointer(&srv.ticketCache)); p != nil && currentTicket != p.Ticket {
			return p, true, nil // 无需更改 p.ExpiresIn 参数值, cached == true 时用不到
		}
	}

	var incompleteURL = "https://api.weixin.qq.com/cgi-bin/ticket/getticket?type=jsapi&access_token="
	var result struct {
		core.Error
		jsapiTicket
	}
	if err = srv.coreClient.GetJSON(incompleteURL, &result); err != nil {
		atomic.StorePointer(&srv.ticketCache, nil)
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		atomic.StorePointer(&srv.ticketCache, nil)
		err = &result.Error
		return
	}

	// 由于网络的延时, jsapi_ticket 过期时间留有一个缓冲区
	switch {
	case result.ExpiresIn > 31556952: // 60*60*24*365.2425
		atomic.StorePointer(&srv.ticketCache, nil)
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
		atomic.StorePointer(&srv.ticketCache, nil)
		err = errors.New("expires_in too small: " + strconv.FormatInt(result.ExpiresIn, 10))
		return
	}

	ticketCopy := result.jsapiTicket
	atomic.StorePointer(&srv.ticketCache, unsafe.Pointer(&ticketCopy))
	ticket = &ticketCopy
	return
}

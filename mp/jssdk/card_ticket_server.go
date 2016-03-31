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

// 卡劵 api_ticket 中控服务器接口.
type CardTicketServer interface {
	Ticket() (ticket string, err error)                            // 请求中控服务器返回缓存的卡劵 api_ticket
	RefreshTicket(currentTicket string) (ticket string, err error) // 请求中控服务器刷新卡劵 api_ticket
	IIDB9BDD0A1E1DC11E5844AA4DB30FED8E1()                          // 接口标识, 没有实际意义
}

var _ CardTicketServer = (*DefaultCardTicketServer)(nil)

// DefaultCardTicketServer 实现了 CardTicketServer 接口.
//  NOTE:
//  1. 用于单进程环境.
//  2. 因为 DefaultCardTicketServer 同时也是一个简单的中控服务器, 而不是仅仅实现 CardTicketServer 接口,
//     所以整个系统只能存在一个 DefaultCardTicketServer 实例!
type DefaultCardTicketServer struct {
	coreClient *core.Client

	refreshTicketRequestChan  chan string              // chan currentTicket
	refreshTicketResponseChan chan refreshTicketResult // chan {ticket, err}

	ticketCache unsafe.Pointer // *cardApiTicket
}

// NewDefaultCardTicketServer 创建一个新的 DefaultCardTicketServer.
func NewDefaultCardTicketServer(clt *core.Client) (srv *DefaultCardTicketServer) {
	if clt == nil {
		panic("nil core.Client")
	}
	srv = &DefaultCardTicketServer{
		coreClient:                clt,
		refreshTicketRequestChan:  make(chan string),
		refreshTicketResponseChan: make(chan refreshTicketResult),
	}

	go srv.ticketUpdateDaemon(time.Hour * 24 * time.Duration(100+rand.Int63n(200)))
	return
}

func (srv *DefaultCardTicketServer) IIDB9BDD0A1E1DC11E5844AA4DB30FED8E1() {}

func (srv *DefaultCardTicketServer) Ticket() (ticket string, err error) {
	if p := (*cardApiTicket)(atomic.LoadPointer(&srv.ticketCache)); p != nil {
		return p.Ticket, nil
	}
	return srv.RefreshTicket("")
}

//type refreshTicketResult struct {
//	ticket string
//	err    error
//}

func (srv *DefaultCardTicketServer) RefreshTicket(currentTicket string) (ticket string, err error) {
	srv.refreshTicketRequestChan <- currentTicket
	rslt := <-srv.refreshTicketResponseChan
	return rslt.ticket, rslt.err
}

func (srv *DefaultCardTicketServer) ticketUpdateDaemon(initTickDuration time.Duration) {
	tickDuration := initTickDuration

NEW_TICK_DURATION:
	ticker := time.NewTicker(tickDuration)
	for {
		select {
		case currentTicket := <-srv.refreshTicketRequestChan:
			cardApiTicket, cached, err := srv.updateTicket(currentTicket)
			if err != nil {
				srv.refreshTicketResponseChan <- refreshTicketResult{err: err}
				break
			}
			srv.refreshTicketResponseChan <- refreshTicketResult{ticket: cardApiTicket.Ticket}
			if !cached {
				tickDuration = time.Duration(cardApiTicket.ExpiresIn) * time.Second
				ticker.Stop()
				goto NEW_TICK_DURATION
			}

		case <-ticker.C:
			cardApiTicket, _, err := srv.updateTicket("")
			if err != nil {
				break
			}
			newTickDuration := time.Duration(cardApiTicket.ExpiresIn) * time.Second
			if abs(tickDuration-newTickDuration) > time.Second*5 {
				tickDuration = newTickDuration
				ticker.Stop()
				goto NEW_TICK_DURATION
			}
		}
	}
}

//func abs(x time.Duration) time.Duration {
//	if x >= 0 {
//		return x
//	}
//	return -x
//}

type cardApiTicket struct {
	Ticket    string `json:"ticket"`
	ExpiresIn int64  `json:"expires_in"`
}

// updateTicket 从微信服务器获取新的卡劵 api_ticket 并存入缓存, 同时返回该卡劵 api_ticket.
func (srv *DefaultCardTicketServer) updateTicket(currentTicket string) (ticket *cardApiTicket, cached bool, err error) {
	if currentTicket != "" {
		if p := (*cardApiTicket)(atomic.LoadPointer(&srv.ticketCache)); p != nil && currentTicket != p.Ticket {
			return p, true, nil // 无需更改 p.ExpiresIn 参数值, cached == true 时用不到
		}
	}

	var incompleteURL = "https://api.weixin.qq.com/cgi-bin/ticket/getticket?type=wx_card&access_token="
	var result struct {
		core.Error
		cardApiTicket
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

	// 由于网络的延时, 卡劵 api_ticket 过期时间留有一个缓冲区
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

	ticketCopy := result.cardApiTicket
	atomic.StorePointer(&srv.ticketCache, unsafe.Pointer(&ticketCopy))
	ticket = &ticketCopy
	return
}

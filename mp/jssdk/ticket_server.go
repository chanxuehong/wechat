<<<<<<< HEAD
// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

=======
>>>>>>> github/v2
package jssdk

import (
	"errors"
<<<<<<< HEAD
=======
	"math/rand"
>>>>>>> github/v2
	"strconv"
	"sync"
	"time"

<<<<<<< HEAD
	"github.com/chanxuehong/wechat/mp"
=======
	"github.com/chanxuehong/wechat/mp/core"
>>>>>>> github/v2
)

// jsapi_ticket 中控服务器接口.
type TicketServer interface {
<<<<<<< HEAD
	// 从中控服务器获取被缓存的 jsapi_ticket.
	Ticket() (string, error)

	// 请求中控服务器到微信服务器刷新 jsapi_ticket.
	//
	//  高并发场景下某个时间点可能有很多请求(比如缓存的 jsapi_ticket 刚好过期时), 但是我们
	//  不期望也没有必要让这些请求都去微信服务器获取 jsapi_ticket(有可能导致api超过调用限制),
	//  实际上这些请求只需要一个新的 jsapi_ticket 即可, 所以建议 TicketServer 从微信服务器
	//  获取一次 jsapi_ticket 之后的至多5秒内(收敛时间, 视情况而定, 理论上至多5个http或tcp周期)
	//  再次调用该函数不再去微信服务器获取, 而是直接返回之前的结果.
	TicketRefresh() (string, error)

	// 没有实际意义, 接口标识
	TagB38894EBFE9911E4BE17A4DB30FED8E1()
=======
	Ticket() (ticket string, err error)        // 请求中控服务器返回缓存的 jsapi_ticket
	TicketRefresh() (ticket string, err error) // 请求中控服务器刷新 jsapi_ticket
	IIDB04E44A0E1DC11E5ADCEA4DB30FED8E1()      // 接口标识, 没有实际意义
>>>>>>> github/v2
}

var _ TicketServer = (*DefaultTicketServer)(nil)

<<<<<<< HEAD
// TicketServer 的简单实现.
=======
// DefaultTicketServer 实现了 TicketServer 接口.
>>>>>>> github/v2
//  NOTE:
//  1. 用于单进程环境.
//  2. 因为 DefaultTicketServer 同时也是一个简单的中控服务器, 而不是仅仅实现 TicketServer 接口,
//     所以整个系统只能存在一个 DefaultTicketServer 实例!
type DefaultTicketServer struct {
<<<<<<< HEAD
	mpClient *mp.Client

	resetTickerChan chan time.Duration // 用于重置 ticketDaemon 里的 ticker

	ticketGet struct {
		sync.Mutex
		LastTicketInfo ticketInfo // 最后一次成功从微信服务器获取的 jsapi_ticket 信息
		LastTimestamp  int64      // 最后一次成功从微信服务器获取 jsapi_ticket 的时间戳
=======
	coreClient *core.Client

	ticker          *time.Ticker       // 用于定时更新 jsapi_ticket
	resetTickerChan chan time.Duration // 用于重置 ticker

	ticketGet struct {
		sync.Mutex
		lastTicket    jsapiTicket // 最近一次成功更新 jsapi_ticket 的数据信息
		lastTimestamp time.Time   // 最近一次成功更新 jsapi_ticket 的时间戳
>>>>>>> github/v2
	}

	ticketCache struct {
		sync.RWMutex
<<<<<<< HEAD
		Ticket string
	}
}

// 创建一个新的 DefaultTicketServer.
func NewDefaultTicketServer(clt *mp.Client) (srv *DefaultTicketServer) {
	if clt == nil {
		panic("nil mp.Client")
	}

	srv = &DefaultTicketServer{
		mpClient:        clt,
		resetTickerChan: make(chan time.Duration),
	}

	go srv.ticketDaemon(time.Hour * 24) // 启动 tokenDaemon
	return
}

func (srv *DefaultTicketServer) TagB38894EBFE9911E4BE17A4DB30FED8E1() {}

func (srv *DefaultTicketServer) Ticket() (ticket string, err error) {
	srv.ticketCache.RLock()
	ticket = srv.ticketCache.Ticket
=======
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
>>>>>>> github/v2
	srv.ticketCache.RUnlock()

	if ticket != "" {
		return
	}
	return srv.TicketRefresh()
}

func (srv *DefaultTicketServer) TicketRefresh() (ticket string, err error) {
<<<<<<< HEAD
	ticketInfo, cached, err := srv.getTicket()
=======
	jsapiTicket, cached, err := srv.ticketRefresh()
>>>>>>> github/v2
	if err != nil {
		return
	}
	if !cached {
<<<<<<< HEAD
		srv.resetTickerChan <- time.Duration(ticketInfo.ExpiresIn) * time.Second
	}
	ticket = ticketInfo.Ticket
	return
}

func (srv *DefaultTicketServer) ticketDaemon(tickDuration time.Duration) {
NEW_TICK_DURATION:
	ticker := time.NewTicker(tickDuration)

	for {
		select {
		case tickDuration = <-srv.resetTickerChan:
			ticker.Stop()
			goto NEW_TICK_DURATION

		case <-ticker.C:
			ticketInfo, cached, err := srv.getTicket()
=======
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
			jsapiTicket, cached, err := srv.ticketRefresh()
>>>>>>> github/v2
			if err != nil {
				break
			}
			if !cached {
<<<<<<< HEAD
				newTickDuration := time.Duration(ticketInfo.ExpiresIn) * time.Second
				if tickDuration != newTickDuration {
					tickDuration = newTickDuration
					ticker.Stop()
=======
				newTickDuration := time.Duration(jsapiTicket.ExpiresIn) * time.Second
				if tickDuration != newTickDuration {
					tickDuration = newTickDuration
					srv.ticker.Stop()
>>>>>>> github/v2
					goto NEW_TICK_DURATION
				}
			}
		}
	}
}

<<<<<<< HEAD
type ticketInfo struct {
	Ticket    string `json:"ticket"`
	ExpiresIn int64  `json:"expires_in"` // 有效时间, seconds
}

// 从微信服务器获取 jsapi_ticket.
//  同一时刻只能一个 goroutine 进入, 防止没必要的重复获取.
func (srv *DefaultTicketServer) getTicket() (ticket ticketInfo, cached bool, err error) {
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
=======
// ticketRefresh 从微信服务器获取新的 jsapi_ticket 并存入缓存, 同时返回该 jsapi_ticket.
func (srv *DefaultTicketServer) ticketRefresh() (ticket jsapiTicket, cached bool, err error) {
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
>>>>>>> github/v2
		cached = true
		return
	}

<<<<<<< HEAD
	var result struct {
		mp.Error
		ticketInfo
	}

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/ticket/getticket?type=jsapi&access_token="
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

=======
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
>>>>>>> github/v2
		err = &result.Error
		return
	}

<<<<<<< HEAD
	// 由于网络的延时, jsapi_ticket 过期时间留了一个缓冲区
	switch {
	case result.ExpiresIn > 31556952: // 60*60*24*365.2425
		srv.ticketCache.Lock()
		srv.ticketCache.Ticket = ""
		srv.ticketCache.Unlock()

=======
	// 由于网络的延时, jsapi_ticket 过期时间留有一个缓冲区
	switch {
	case result.ExpiresIn > 31556952: // 60*60*24*365.2425
		srv.ticketCache.Lock()
		srv.ticketCache.ticket = ""
		srv.ticketCache.Unlock()
>>>>>>> github/v2
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
<<<<<<< HEAD
		srv.ticketCache.Ticket = ""
		srv.ticketCache.Unlock()

=======
		srv.ticketCache.ticket = ""
		srv.ticketCache.Unlock()
>>>>>>> github/v2
		err = errors.New("expires_in too small: " + strconv.FormatInt(result.ExpiresIn, 10))
		return
	}

<<<<<<< HEAD
	srv.ticketGet.LastTicketInfo = result.ticketInfo
	srv.ticketGet.LastTimestamp = timeNowUnix

	srv.ticketCache.Lock()
	srv.ticketCache.Ticket = result.ticketInfo.Ticket
	srv.ticketCache.Unlock()

	ticket = result.ticketInfo
=======
	// 更新 ticketGet
	srv.ticketGet.lastTicket = result.jsapiTicket
	srv.ticketGet.lastTimestamp = timeNow

	// 更新 ticketCache
	srv.ticketCache.Lock()
	srv.ticketCache.ticket = result.Ticket
	srv.ticketCache.Unlock()

	ticket = result.jsapiTicket
>>>>>>> github/v2
	return
}

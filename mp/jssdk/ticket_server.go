// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package jssdk

import (
	"errors"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/chanxuehong/wechat/mp"
)

// jsapi_ticket 中控服务器接口.
type TicketServer interface {
	// 从中控服务器获取被缓存的 jsapi_ticket.
	Ticket() (ticket string, err error)

	// 请求中控服务器到微信服务器刷新 jsapi_ticket.
	//
	//  高并发场景下某个时间点可能有很多请求(比如缓存的jsapi_ticket刚好过期时), 但是我们
	//  不期望也没有必要让这些请求都去微信服务器获取 jsapi_ticket(有可能导致api超过调用限制),
	//  实际上这些请求只需要一个新的 jsapi_ticket 即可, 所以建议 TokenServer 从微信服务器
	//  获取一次 jsapi_ticket 之后的至多5秒内(收敛时间, 视情况而定, 理论上至多5个http或tcp周期)
	//  再次调用该函数不再去微信服务器获取, 而是直接返回之前的结果.
	TicketRefresh() (ticket string, err error)
}

var _ TicketServer = new(DefaultTicketServer)

// TicketServer 的简单实现.
//  NOTE:
//  1. 用于单进程环境.
//  2. 因为 DefaultTicketServer 同时也是一个简单的中控服务器, 而不是仅仅实现 TicketServer 接口,
//     所以整个系统只能存在一个 DefaultTicketServer 实例!
type DefaultTicketServer struct {
	mp.WechatClient

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

// 创建一个新的 DefaultTicketServer.
//  如果 httpClient == nil 则默认使用 http.DefaultClient.
func NewDefaultTicketServer(tokenServer mp.TokenServer, httpClient *http.Client) (srv *DefaultTicketServer) {
	if tokenServer == nil {
		panic("nil tokenServer")
	}
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	srv = &DefaultTicketServer{
		WechatClient: mp.WechatClient{
			TokenServer: tokenServer,
			HttpClient:  httpClient,
		},
		resetTickerChan: make(chan time.Duration),
	}

	// 获取 jsapi_ticket 并启动 goroutine ticketDaemon
	ticketInfo, cached, err := srv.getTicket()
	if err != nil {
		panic(err)
	}
	if !cached {
		srv.ticketCache.Ticket = ticketInfo.Ticket
		go srv.ticketDaemon(time.Duration(ticketInfo.ExpiresIn) * time.Second)
	}
	return
}

func (srv *DefaultTicketServer) Ticket() (ticket string, err error) {
	srv.ticketCache.RLock()
	ticket = srv.ticketCache.Ticket
	srv.ticketCache.RUnlock()

	if ticket != "" {
		return
	}
	return srv.TicketRefresh()
}

func (srv *DefaultTicketServer) TicketRefresh() (ticket string, err error) {
	ticketInfo, cached, err := srv.getTicket()
	if err != nil {
		srv.ticketCache.Lock()
		srv.ticketCache.Ticket = ""
		srv.ticketCache.Unlock()
		return
	}
	if !cached {
		srv.ticketCache.Lock()
		srv.ticketCache.Ticket = ticketInfo.Ticket
		srv.ticketCache.Unlock()

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
			if err != nil {
				srv.ticketCache.Lock()
				srv.ticketCache.Ticket = ""
				srv.ticketCache.Unlock()
				break
			}
			if !cached {
				srv.ticketCache.Lock()
				srv.ticketCache.Ticket = ticketInfo.Ticket
				srv.ticketCache.Unlock()

				newTickDuration := time.Duration(ticketInfo.ExpiresIn) * time.Second
				if tickDuration != newTickDuration {
					ticker.Stop()
					tickDuration = newTickDuration
					goto NEW_TICK_DURATION
				}
			}
		}
	}
}

type ticketInfo struct {
	Ticket    string `json:"ticket"`
	ExpiresIn int64  `json:"expires_in"` // 有效时间, seconds
}

// 从微信服务器获取 jsapi_ticket.
func (srv *DefaultTicketServer) getTicket() (ticket ticketInfo, cached bool, err error) {
	srv.ticketGet.Lock()
	defer srv.ticketGet.Unlock()

	timeNowUnix := time.Now().Unix()

	// 在收敛周期内直接返回最近一次获取的 jsapi_ticket, 这里的收敛时间设定为4秒
	if n := srv.ticketGet.LastTimestamp; timeNowUnix >= n && timeNowUnix < n+4 {
		ticket = ticketInfo{
			Ticket:    srv.ticketGet.LastTicketInfo.Ticket,
			ExpiresIn: srv.ticketGet.LastTicketInfo.ExpiresIn + n - timeNowUnix,
		}
		cached = true
		return
	}

	var result struct {
		mp.Error
		ticketInfo
	}

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/ticket/getticket?type=jsapi&access_token="
	if err = srv.GetJSON(incompleteURL, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}

	// 由于网络的延时, jsapi_ticket 过期时间留了一个缓冲区
	switch {
	case result.ExpiresIn > 60*60:
		result.ExpiresIn -= 60 * 10
	case result.ExpiresIn > 60*30:
		result.ExpiresIn -= 60 * 5
	case result.ExpiresIn > 60*5:
		result.ExpiresIn -= 60
	case result.ExpiresIn > 60:
		result.ExpiresIn -= 10
	case result.ExpiresIn > 0:
	default:
		err = errors.New("invalid expires_in: " + strconv.FormatInt(result.ExpiresIn, 10))
		return
	}

	srv.ticketGet.LastTicketInfo = result.ticketInfo
	srv.ticketGet.LastTimestamp = timeNowUnix
	ticket = result.ticketInfo
	return
}

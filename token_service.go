// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package wechat

// access token 伺服接口, 用于集中式获取 access token 场景, see token_service.png
type TokenService interface {
	// 获取 access token, 该 token 一般缓存在某个地方.
	// 正常情况下 token != "" && err == nil, 否则 token == "" && err != nil
	// NOTE: 该方法一定要功能上实现!
	Token() (token string, err error)

	// 从微信服务器获取新的 access token.
	// 正常情况下 token != "" && err == nil, 否则 token == "" && err != nil
	//  NOTE:
	//  1. 一般情况下无需调用该函数, 请使用 Token() 获取 access token.
	//  2. 该方法可以选择是否功能上实现, 如果没有需求可以在语法上实现即可!
	//  3. 即使 access token 过期(错误代码 40001, 正常情况下不会出现),
	//     也请谨慎调用 TokenRefresh, 建议直接返回错误! 因为很有可能高并发情况下造成雪崩效应!
	//  4. 再次强调, 调用这个函数你应该知道发生了什么!!!
	TokenRefresh() (token string, err error)
}

// +build wechatdebug

package retry

import (
	"log"
)

// access_token 过期重试之前打印相应信息
func DebugPrintError(errcode int64, errmsg string, token string) {
	log.Printf("[WECHAT_DEBUG] [RETRY] errcode: %d, errmsg: %s\n", errcode, errmsg)
	log.Println("[WECHAT_DEBUG] [RETRY] current token:", token)
}

// access_token 过期重试过程中打印新的 access_token
func DebugPrintNewToken(token string) {
	log.Println("[WECHAT_DEBUG] [RETRY] new token:", token)
}

// access_token 过期重试失败打印对应的 access_token
func DebugPrintFallthrough(token string) {
	log.Println("[WECHAT_DEBUG] [RETRY] fallthrough, current token:", token)
}

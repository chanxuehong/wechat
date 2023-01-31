package retry

import (
	"log"
)

func DebugPrintError(errcode int64, errmsg string, token string, debug bool) {
	if !debug {
		return
	}
	const format = "[WECHAT_DEBUG] [API] [RETRY] errcode: %d, errmsg: %s\n" +
		"current token: %s\n"
	log.Printf(format, errcode, errmsg, token)
}

func DebugPrintNewToken(token string, debug bool) {
	if !debug {
		return
	}
	log.Println("[WECHAT_DEBUG] [API] [RETRY] new token:", token)
}

func DebugPrintFallthrough(token string, debug bool) {
	if !debug {
		return
	}
	log.Println("[WECHAT_DEBUG] [API] [RETRY] fallthrough, current token:", token)
}

// +build !wechat_debug

package retry

func DebugPrintError(errcode int64, errmsg string, token string) {}

func DebugPrintNewToken(token string) {}

func DebugPrintFallthrough(token string) {}

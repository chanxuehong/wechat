// +build !wechatdebug

package retry

// access_token 过期重试之前打印相应信息
func DebugPrintError(errcode int64, errmsg string, token string) {}

// access_token 过期重试过程中打印新的 access_token
func DebugPrintNewToken(token string) {}

// access_token 过期重试失败打印对应的 access_token
func DebugPrintFallthrough(token string) {}

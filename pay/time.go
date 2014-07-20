// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay

import (
	"time"
)

var beijingLocation = time.FixedZone("GMT+8", 60*60*8)

// 格式化时间到 yyyyMMDDHHmmss, GMT+8
func FormatTime(t time.Time) string {
	return t.In(beijingLocation).Format("20060102150405")
}

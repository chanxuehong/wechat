package core

import (
	"time"

	"github.com/chanxuehong/wechat/util"
)

// FormatTime 将参数 t 格式化成北京时间 yyyyMMddHHmmss.
func FormatTime(t time.Time) string {
	return t.In(util.BeijingLocation).Format("20060102150405")
}

// ParseTime 将北京时间 yyyyMMddHHmmss 字符串解析到 time.Time.
func ParseTime(value string) (time.Time, error) {
	return time.ParseInLocation("20060102150405", value, util.BeijingLocation)
}

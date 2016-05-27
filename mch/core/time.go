package core

import (
	"time"

	timex "github.com/chanxuehong/time"
)

// FormatTime 将参数 t 格式化成 北京时间yyyyMMddHHmmss
func FormatTime(t time.Time) string {
	return t.In(timex.BeijingLocation).Format("20060102150405")
}

// ParseTime 将 北京时间yyyyMMddHHmmss 字符串解析到 time.Time
func ParseTime(value string) (time.Time, error) {
	return time.ParseInLocation("20060102150405", value, timex.BeijingLocation)
}

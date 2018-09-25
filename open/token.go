package open

import (
	"time"
)

//IDurationToken 用于一般性存储含有过期时间的token
type IDurationToken interface {
	Value() (string, error)
	Put(val string, expire time.Duration) error
	Expired() bool
}

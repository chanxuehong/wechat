package core

import (
	"sync/atomic"
)

const (
	startedCheckerInitialValue = uintptr(0)
	startedCheckerStartedValue = ^uintptr(0)
)

// 对于要求不高的场景, 可以嵌入 startedChecker 到其他类型来检查是否在启动服务后设置参数.
// 对于要求高的场景, 需要用到互斥锁.
type startedChecker uintptr

func (p *startedChecker) start() {
	if uintptr(*p) == startedCheckerInitialValue {
		atomic.CompareAndSwapUintptr((*uintptr)(p), startedCheckerInitialValue, startedCheckerStartedValue)
	}
}

func (v startedChecker) check() {
	if uintptr(v) != startedCheckerInitialValue {
		panic("the service has been started.")
	}
}

package core

import (
	"sync/atomic"
)

const (
	startedCheckerInitialValue = uintptr(0)
	startedCheckerStartedValue = ^uintptr(0)
)

// 正常情况下实例的配置方法和服务方法是不能并行执行的(这两种方法竞争同一份配置数据), 一般我们有两种方案:
// 1. 用互斥锁
// 2. 明确文档告知该实例不是并行安全的
// 对于第2种场景, 很多程序员有可能不小心并行执行了该实例的配置方法和服务方法, 那有没有好的解决方案呢?
// 其实在大部分场景下, 实例可以先配置, 投入服务后就没有必要修改其配置了(如果有必要修改的只能用互斥锁了),
// startedChecker 就是为这种场景设计的, 能在很大程度上保证数据安全(不是绝对, 极小的概率下会出现数据竞争).
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

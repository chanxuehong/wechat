package core

import (
	"sync/atomic"
)

const (
	startedCheckerInitialValue = uintptr(0)
	startedCheckerStartedValue = ^uintptr(0)
)

type startedChecker uintptr

func (p *startedChecker) start() {
	if uintptr(*p) != startedCheckerStartedValue {
		atomic.CompareAndSwapUintptr((*uintptr)(p), startedCheckerInitialValue, startedCheckerStartedValue)
	}
}

func (v startedChecker) check() {
	if uintptr(v) != startedCheckerInitialValue {
		panic("the service has been started.")
	}
}

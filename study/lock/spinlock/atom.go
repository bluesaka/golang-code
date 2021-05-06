package spinlock

import (
	"sync/atomic"
)

type SpinAtom int32

func (l *SpinAtom) Lock() {
	for !atomic.CompareAndSwapInt32((*int32)(l), 0, 1) {
		//log.Println("spin atom")
	}
}

func (l *SpinAtom) Unlock() {
	atomic.StoreInt32((*int32)(l), 0)
}

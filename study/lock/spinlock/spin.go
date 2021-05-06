/**
自旋锁，是为防止多处理器并发，保护共享资源而提出一种锁机制
和互斥锁类似，在任何时刻，最多只有一个锁持有者

自旋锁和互斥锁在调度机制上略有不同，
互斥锁如果资源已经被占用，则其他申请者会进入睡眠状态放弃cpu，
而自旋锁则是不断循环并测试锁的状态(忙等待 busy waiting)，这样就一直占用cpu
*/
package spinlock

type Spin int32

func (l *Spin) Lock() {
	lock((*int32)(l), 0, 1)
}

func (l *Spin) Unlock() {
	unlock((*int32)(l), 0)
}

// @link https://www.zhihu.com/zvideo/1357059453523681280
// spin_amd64.s or spin_386.s based on os
func lock(ptr *int32, o, n int32)
func unlock(ptr *int32, n int32)

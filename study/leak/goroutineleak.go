/**
产生goroutine leak（协程泄漏）的原因可能有以下几种：
- goroutine由于channel的读/写端退出而一直阻塞，导致goroutine一直占用资源，而无法退出
- goroutine进入死循环中，导致资源一直无法释放

goroutine终止的场景:
- 当一个goroutine完成它的工作
- 由于发生了没有处理的错误
- 有其他的协程告诉它终止
*/

package leak

import (
	"context"
	"fmt"
	"time"
)

func GoroutineLeakTimeout(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	// 如果done channel是无缓冲的，那么在超时退出后，done的写入会一直阻塞，导致这个函数每次调用都会占用一个goroutine
	// 设置缓冲大小后，done可以写入而不卡住goroutine x
	done := make(chan struct{})
	//done := make(chan struct{}, 1)

	go func() {
		time.Sleep(time.Second * 2)
		done <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		return fmt.Errorf("timeout")
	case <-done:
		return nil
	}
}

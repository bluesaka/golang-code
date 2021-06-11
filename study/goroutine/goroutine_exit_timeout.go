/**
如何退出协程（超时场景）
@link https://geektutu.com/post/hpg-timeout-goroutine.html
*/
package goroutine

import (
	"fmt"
	"log"
	"runtime"
	"testing"
	"time"
)

func DoSleep(done chan struct{}) {
	time.Sleep(time.Second)
	done <- struct{}{}
}

/**
done是无缓冲channel，超时后主协程退出，而done没有接收者且无缓冲区，发送者会一直阻塞，导致协程不能退出
*/
func Timeout(f func(chan struct{})) error {
	done := make(chan struct{})
	go f(done)

	select {
	case <-done:
		log.Println("done")
		return nil
	case <-time.After(time.Millisecond):
		return fmt.Errorf("timeout")
	}
}

func test(t *testing.T, f func(chan struct{})) {
	t.Helper()
	for i := 0; i < 1000; i++ {
		Timeout(f)
	}
	time.Sleep(time.Second * 2)
	t.Log("goroutine num:", runtime.NumGoroutine())
}

// ---------------------------------------------------------------------------------------------------------
// 创建有缓冲区的channel避免协程无法退出
// 有缓冲区，即使没有接收方，发送方也不会发生阻塞
func TimeoutWithBuffer(f func(chan struct{})) error {
	done := make(chan struct{}, 1)
	go f(done)

	select {
	case <-done:
		log.Println("done")
		return nil
	case <-time.After(time.Millisecond):
		return fmt.Errorf("timeout")
	}
}

func testWithBuffer(t *testing.T, f func(chan struct{})) {
	t.Helper()
	for i := 0; i < 1000; i++ {
		TimeoutWithBuffer(f)
	}
	time.Sleep(time.Second * 2)
	t.Log("goroutine num:", runtime.NumGoroutine())
}

// ---------------------------------------------------------------------------------------------------------
// 使用select尝试向done channel发送信号，如果发送失败，则说明缺少接收者，说明超时了，那么直接退出即可。
func DoSleepWithSelect(done chan struct{}) {
	time.Sleep(time.Second)
	select {
	case done <- struct{}{}:
	default:
		return
	}
}

// ---------------------------------------------------------------------------------------------------------
// 更复杂的场景：一个任务分两个阶段执行，只检测第一阶段是否超时，若没有超时则执行第二阶段，若超时则终止
func Do2Phases(p1, p2 chan struct{}) {
	time.Sleep(time.Second)
	select {
	case p1 <- struct{}{}:
	default:
		return
	}

	time.Sleep(time.Second)
	p2 <- struct{}{}
}

func TimeoutFirstPhase() error {
	p1, p2 := make(chan struct{}), make(chan struct{})
	go Do2Phases(p1, p2)

	select {
	case <-p1:
		<-p2
		log.Println("done")
		return nil
	case <-time.After(time.Millisecond):
		return fmt.Errorf("timeout")
	}
}

func test2PhasesTimeout(t *testing.T) {
	for i := 0; i < 1000; i++ {
		TimeoutFirstPhase()
	}
	time.Sleep(time.Second * 3)
	t.Log("goroutine num:", runtime.NumGoroutine())
}

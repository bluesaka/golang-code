/**
如何退出协程（其他场景）
@link https://geektutu.com/post/hpg-exit-goroutine.html
*/
package goroutine

import (
	"log"
	"time"
)

// ch 未关闭，导致协程无法退出
func DoTask(ch chan int) {
	for {
		select {
		case i := <- ch:
			time.Sleep(time.Millisecond)
			log.Println("i:", i)
		}
	}
}

func SendTask() {
	ch := make(chan int, 10)
	go DoTask(ch)
	for i := 0; i< 1000; i++ {
		ch <- i
	}
}

// ---------------------------------------------------------------------------------------------------------
// 关闭channel，使协程正确退出
func DoTaskWithClose(ch chan int) {
	for {
		select {
		case i, ok := <-ch:
			if !ok {
				log.Println("chan closed")
				return
			}
			time.Sleep(time.Millisecond)
			log.Println("i:", i)
		}
	}
}

func SendTaskWithClose() {
	ch := make(chan int, 10)
	go DoTaskWithClose(ch)
	for i := 0; i< 1000; i++ {
		ch <- i
	}
	close(ch)
}
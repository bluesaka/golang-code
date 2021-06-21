package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	m2()
}

func m1() {
	go func() {
		select {}
	}()
	time.Sleep(time.Second)
	fmt.Println("main done")
}

/**
go1.14以下的版本会死循环，卡住，因为，即使主goroutine抢到了P，也会因为sleep，让出cpu，继而go func开始无限执行。
go1.14及以上版本，因为有sysmon抢占，即使第一个go func抢到了P，也会因为执行超过10ms，被踢出P的local队列。执行主goroutine的打印，然后退出，不会卡住。
 */
func m2() {
	runtime.GOMAXPROCS(1)
	go func() {
		select {}
	}()
	time.Sleep(time.Second)
	fmt.Println("main done")
}

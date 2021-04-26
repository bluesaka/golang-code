package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	mutex4()
}

/**
不加Lock情况下，期望输出100000，实际输出4万多，
存在goroutine抢占资源，非线程安全
*/
func mutex1() {
	var count = 0
	var n = 10
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 10000; i++ {
				count++
			}
		}()
	}
	wg.Wait()
	log.Println(count)
}

/**
加
加了Lock情况下，输出100000，符合预期
Lock()和Unlock()之间的代码段成为资源的临界区，是线程安全的，任何一个时间点只能有一个goroutine执行该代码段
*/
func mutex2() {
	var count = 0
	var n = 10
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(10)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 10000; i++ {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	log.Println(count)
}

/**
并发读示例
Mutex在大并发的情况下，会造成所等待，对性能的影响较大
RWMutex 单写多读的锁适用于读多写少的场景
*/
func mutex3() {
	var m sync.RWMutex
	go read(&m, 1)
	go read(&m, 2)
	go read(&m, 3)
	time.Sleep(time.Second * 2)
}

func read(m *sync.RWMutex, i int) {
	log.Println(i, "read start")
	m.RLock()
	defer m.RUnlock()
	log.Println(i, "reading")
	time.Sleep(time.Second)
	log.Println(i, "read end")
}

var m4count = 0

/**
Lock()加写锁、Unlock()释放写锁
如果加写锁前已经有其他的读锁和写锁，则Lock()会阻塞，直到该锁可用

RLock()加读锁、RUnlock()释放读锁
如果存在写锁，则无法加读锁；当只有读锁或没有锁时，可以加读锁，读锁可以加多个
*/
func mutex4() {
	var m sync.RWMutex
	for i := 1; i <= 3; i++ {
		go read2(&m, i)
	}
	for i := 1; i <= 3; i++ {
		go write2(&m, i)
	}
	time.Sleep(time.Second * 10)
	log.Println("final count:", m4count)
}

func read2(m *sync.RWMutex, i int) {
	log.Println(i, "read start")
	m.RLock()
	defer m.RUnlock()
	log.Println(i, "reading count:", m4count)
	time.Sleep(time.Second)
	log.Println(i, "read end")
}

func write2(m *sync.RWMutex, i int) {
	log.Println(i, "write start")
	m.Lock()
	defer m.Unlock()
	m4count++
	log.Println(i, "writing count:", m4count)
	time.Sleep(time.Second)
	log.Println(i, "write end")
}

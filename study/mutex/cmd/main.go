/**
Mutex 互斥锁

两种模式：
正常模式 饥饿模式

@link https://www.zhihu.com/zvideo/1359608861503242241

一个尝试获得锁的goroutine会先自旋几次(4次)，若仍未取得锁，则通过信号量排队等待，所有的等待者都会以FIFO先进先出的顺序进行排队，
当锁被释放后，第一个等待者被唤醒后不会直接获得锁，而是会和后来者竞争，也就是和那些处于自旋阶段尚未排队的goroutine，
而这种情况下，后来者更有优势，一是因为后来者正在cpu上运行，二是后来者通常有多个，而被唤醒的goroutine每次只有一个，
当被唤醒的goroutine此次未获得锁，则会被重新插入到排队队列的头部而不是尾部。

`正常模式` 切换为 `饥饿模式`
当一个goroutine本次加锁等待时间超过1ms后，它会把当前Mutex从正常模式切换为饥饿模式。

在饥饿模式下，释放的锁会直接传递给排队头部的goroutine，后来者不会自旋，也不会尝试获得锁，会直接插入到排队尾部进行等待

`饥饿模式` 切换为 `正常模式` 的两种情况：
1. 获得锁的goroutine的等待时间小于1ms
2. 排队的队列空了
 */

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

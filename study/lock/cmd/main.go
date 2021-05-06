package main

import (
	"flag"
	"go-code/study/lock/spinlock"
	"log"
	"time"
)

type Locker interface {
	Lock()
	Unlock()
}

func main() {
	spinlock1()
}

func spinlock1() {
	var p bool
	flag.BoolVar(&p, "pause", false, "spin lock with cpu pause")
	flag.Parse()

	var l Locker
	if p {
		// 带有中断cpu的自旋锁，不占用cpu
		l = new(spinlock.Spin)
	} else {
		// 自旋锁会不断的检测锁占用cpu，会造成cpu跑满的情况
		l = new(spinlock.SpinAtom)
	}

	var n int
	for i := 0; i < 2; i++ {
		go routine(i, &n, l, time.Second)
	}

	select {}
}

func routine(i int, n *int, l Locker, d time.Duration) {
	for {
		func() {
			l.Lock()
			defer l.Unlock()
			*n++
			log.Printf("goroutine: %d, n: %d\n", i, *n)
			time.Sleep(d)
		}()
	}
}

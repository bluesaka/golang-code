package main

import (
	"log"
	"runtime"
	"time"
)

func main() {
	chanLoop2()
}

func noBufferChan() {
	// 无缓冲区，读写都会阻塞
	c1 := make(chan int)
	// 写阻塞
	c1 <- 1
	// this will cause deadlock
	// 读阻塞，会去找写入的值，但是上面的写入在此代码执行之前，无法后退，导致死锁
	log.Println(<-c1)
}

func bufferChan() {
	c1 := make(chan int, 1)
	c1 <- 1
	// 超过缓冲区会死锁
	// c1 <- 2
	// output: 1
	log.Println(<-c1)
}

func noBufferChanWithGoroutine() {
	// 无缓冲区，读写都会阻塞
	c1 := make(chan int)
	// 写阻塞
	go func() {
		c1 <- 1
	}()
	// this will cause deadlock
	// debug调试发现，程序会先执行到这里，读阻塞，然后执行上面的协程，输出channel的值
	log.Println(<-c1)
}

func noBufferChanWithLoop() {
	c1 := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c1 <- i
		}
	}()
	for i := 0; i < 10; i++ {
		// 顺序输出 0 - 9
		log.Println(<-c1)
	}
}

func bufferChanWithLoop() {
	c1 := make(chan int, 5)
	go func() {
		for i := 0; i < 10; i++ {
			c1 <- i
		}
	}()
	for i := 0; i < 10; i++ {
		// 顺序输出 0 - 9，不过和 上面noBufferChanWithLoop无缓冲区的channel不一样
		// debug断点调试发现， c1 <- 1会先塞满5个值，然后再取一个塞一个，直到全部执行完毕
		log.Println(<-c1)
	}
}

func oneWayChan() {
	c1 := make(chan int, 5)
	// 只读channel (receive-only)
	var readChan <-chan int = c1
	// readChan <- 1 // error

	// 只写channel（send-only）
	var writeChan chan<- int = c1
	// <-writeChan // error
	writeChan <- 1
	log.Println(<-readChan)
}

func closeChan() {
	c1 := make(chan int, 3)
	c1 <- 1
	c1 <- 2
	// close channel
	close(c1)
	// panic: send on closed channel
	c1 <- 3
	log.Println(<-c1)
	log.Println(<-c1)
	log.Println(<-c1)
}

func chanLoop() {
	c1 := make(chan int, 3)
	c1 <- 1
	c1 <- 2
	c1 <- 3
	// 需要close channel才能range channel
	close(c1)
	for v := range c1 {
		log.Println(v)
	}
}

func chanLoop2() {
	c := make(chan int)
	go func() {
		for v := range c {
			log.Println(v)
		}
	}()
	go func() {
		c <- 1
		c <- 2
	}()
	time.Sleep(time.Second)
	log.Println("num goroutine:", runtime.NumGoroutine())
}

func chanSelect() {
	c1 := make(chan int, 1)
	c2 := make(chan int, 1)
	c3 := make(chan int, 1)
	c1 <- 1
	c2 <- 2
	c3 <- 3
	// case random
	select {
	case c1 := <-c1:
		log.Println("c1:", c1)
	case c2 := <-c2:
		log.Println("c2:", c2)
	case c3 := <-c3:
		log.Println("c3:", c3)
	default:
		log.Println("none")
	}
}

func chan1() {
	naturals := make(chan int)
	squares := make(chan int)

	// counter
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// squarer
	go func() {
		for {
			//x := <-naturals
			x, ok := <-naturals
			if !ok {
				break
			}
			squares <- x * x
		}
		close(squares)
	}()

	// printer
	//for {
	//	log.Println(<-squares)
	//}
	for x := range squares {
		log.Println(x)
	}

}

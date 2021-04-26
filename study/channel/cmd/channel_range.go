package main

import (
	"log"
	"strconv"
	"time"
)

func main() {
	chanRange2()
}

/**
此示例中channel被正确关闭close
*/
func chanRange2() {
	ch := make(chan string)
	go makeCake2(ch)
	receiveCake2(ch)
}

/**
此示例中channel没有被关闭，只是随着main函数退出，所有goroutine被关闭，该语句被结束了而已

正确的做法：
- 发送者数据发送完毕后，关闭channel
- 接受者数据接受完毕后，终止程序
- 去除main函数time.Sleep
*/
func chanRange1() {
	ch := make(chan string)
	go makeCake1(ch)
	go receiveCake1(ch)
	time.Sleep(1e9)
}

func makeCake2(ch chan string) {
	// 关闭channel
	// 如果channel已经关闭，继续往channel发送数据会panic:send on closed channel
	// 如果关闭一个已经关闭的channel也会panic: close of closed channel
	// channel被关闭后，仍然可以从中读取已经发送的数据，数据读取完毕后，将读取到零值
	defer close(ch)
	for i := 1; i <= 5; i++ {
		ch <- "cake" + strconv.Itoa(i)
	}
}

func receiveCake2(ch <-chan string) {
	//range会一直阻塞当前协程，如果其他协程中调用了close(ch)，name就会跳出for range循环
	for cake := range ch {
		log.Println("receive cake:", cake)
	}

	// 此语句等价于上面的for range循环
	for {
		if cake, ok := <-ch; ok {
			log.Println("receive2 cake:", cake)
		} else {
			break
		}
	}

}

func makeCake1(ch chan string) {
	for i := 1; i <= 5; i++ {
		ch <- "cake" + strconv.Itoa(i)
	}
}

func receiveCake1(ch <-chan string) {
	// range会一直阻塞当前协程，如果其他协程中调用了close(ch)，name就会跳出for range循环
	for cake := range ch {
		log.Println("receive cake:", cake)
	}
}

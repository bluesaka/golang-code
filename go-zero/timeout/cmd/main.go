/**
 * 超时
 * @link https://mp.weixin.qq.com/s/5Q05d6OwvS-Zj-yNGwoh8A
 */

package main

import (
	"context"
	"fmt"
	"go-code/go-zero/timeout"
	"log"
	"runtime"
	"sync"
	"time"
)

func main() {
	test1()
	//test2()
	//test3()
}

func test1() {
	log.Println("start")
	total := 10
	var wg sync.WaitGroup
	wg.Add(total)
	now := time.Now()
	for i := 0; i < total; i++ {
		//go func() {
		//	defer wg.Done()
		//	if err := timeout.RequestWork2(context.Background(), nil); err != nil {
		//		//log.Println("err:", err)
		//	}
		//}()
		go func(i int) {
			defer wg.Done()
			ctx := context.WithValue(context.Background(), "id", i)
			//log.Println("g id:", i)
			if err := timeout.RequestWork2(ctx, nil); err != nil {
				//log.Println("err:", err)
			}
		}(i)
	}
	wg.Wait()
	log.Println("elapsed:", time.Since(now))
	time.Sleep(time.Second * 9)
	log.Println("goroutine num:", runtime.NumGoroutine())
}

func test2() {
	fmt.Println("start")
	total := 1000
	var wg sync.WaitGroup
	wg.Add(total)
	now := time.Now()
	for i := 0; i < total; i++ {
		go func() {
			defer func() {
				if p := recover(); p != nil {
					fmt.Println("oops, panic")
				}
			}()

			defer wg.Done()
			if err := timeout.RequestWork3(context.Background(), nil); err != nil {
				//fmt.Println("err:", err)
			}
		}()
	}
	wg.Wait()
	fmt.Println("elapsed:", time.Since(now))
	fmt.Println("goroutine num:", runtime.NumGoroutine())
}

func test3() {
	fmt.Println("start")
	total := 10
	var wg sync.WaitGroup
	wg.Add(total)
	now := time.Now()
	for i := 0; i < total; i++ {
		go func() {
			defer func() {
				if p := recover(); p != nil {
					fmt.Println("oops, panic")
				}
			}()

			defer wg.Done()
			if err := timeout.RequestWork4(context.Background(), nil); err != nil {
				//fmt.Println("err:", err)
			}
		}()
	}
	wg.Wait()
	fmt.Println("elapsed:", time.Since(now))
	time.Sleep(time.Second * 20)
	fmt.Println("goroutine num:", runtime.NumGoroutine())
}

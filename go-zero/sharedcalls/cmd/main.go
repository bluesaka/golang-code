package main

import (
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/core/syncx"
	"go-code/go-zero/sharedcalls"
	"log"
	"sync"
	"time"
)

func main() {
	//test()
	//test2()
	test3()
}

func test() {
	count := 5
	var wg sync.WaitGroup
	wg.Add(count)
	sharedCalls := syncx.NewSharedCalls()

	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			val, err := sharedCalls.Do("testkey", func() (interface{}, error) {
				time.Sleep(time.Second)
				return stringx.RandId(), nil
			})
			if err != nil {
				log.Println(err)
			} else {
				log.Println(val)
			}
		}()
	}

	wg.Wait()
}

func test2() {
	count := 5
	var wg sync.WaitGroup
	wg.Add(count)

	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			log.Println(stringx.RandId())
		}()
	}

	wg.Wait()
}

func test3() {
	count := 5
	var wg sync.WaitGroup
	wg.Add(count)
	sharedCalls := sharedcalls.NewMySharedCall()

	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			val, err := sharedCalls.Do("testkey2", func() (interface{}, error) {
				time.Sleep(time.Second)
				return stringx.RandId(), nil
			})
			if err != nil {
				log.Println(err)
			} else {
				log.Println(val)
			}
		}()
	}

	wg.Wait()
}

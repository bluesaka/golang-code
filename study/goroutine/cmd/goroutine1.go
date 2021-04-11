package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	runOrder()
}

func runSimple() {
	log.Println("runSimple start...")
	// this will not output
	go func() {
		log.Println("runSimple processing")
	}()
	log.Println("runSimple end...")
}

func runSimpleWithSleep() {
	log.Println("runSimpleWithSleep start...")
	// this will not output
	go func() {
		log.Println("runSimpleWithSleep processing")
	}()
	time.Sleep(time.Second)
	log.Println("runSimpleWithSleep end...")
}

func runWithLoop() {
	log.Println("runWithLoop start...")
	// 我的机器在20多、30多执行了该协程
	go func() {
		log.Println("runWithLoop processing")
	}()
	for i := 0; i < 100; i++ {
		log.Println(i)
	}
	log.Println("runWithLoop end...")
}

func runWithWaitGroup() {
	log.Println("runWithWaitGroup start...")
	var wg sync.WaitGroup
	wg.Add(1)
	go runWithWaitGroup2(&wg)
	wg.Wait()
	log.Println("runWithWaitGroup end")
}

func runWithWaitGroup2(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)
	log.Println("runWithWaitGroup processing...")
}

// 顺序打印出 dog, cat, fish，共10次
func runWithGoroutineOrder() {
	dogChan := make(chan bool, 1)
	catChan := make(chan bool, 1)
	fishChan := make(chan bool, 1)
	dogChan <- true

	var wg sync.WaitGroup
	wg.Add(30)

	for i := 0; i < 10; i++ {
		go printDog(&wg, dogChan, catChan)
		go printCat(&wg, catChan, fishChan)
		go printFish(&wg, fishChan, dogChan)
	}

	wg.Wait()
}

func printDog(wg *sync.WaitGroup, dogChan, catChan chan bool) {
	if ok := <-dogChan; ok {
		defer wg.Done()
		log.Println("dog")
		catChan <- true
	}
}

func printCat(wg *sync.WaitGroup, catChan, fishChan chan bool) {
	if ok := <-catChan; ok {
		defer wg.Done()
		log.Println("cat")
		fishChan <- true
	}
}

func printFish(wg *sync.WaitGroup, fishChan, dogChan chan bool) {
	if ok := <-fishChan; ok {
		defer wg.Done()
		log.Println("fish")
		dogChan <- true
	}
}

func runOrder() {
	a := make(chan bool, 1)
	b := make(chan bool, 1)
	done := make(chan bool, 1)

	go func() {
		s := []int{1, 2, 3, 4}
		for i := 0; i < len(s); i++ {
			if ok := <-a; ok {
				log.Println(s[i])
				b <- true
			}
		}
	}()

	go func() {
		defer func() {
			close(done)
		}()
		s := []byte{'a', 'b', 'c', 'd'}
		for i := 0; i < len(s); i++ {
			if ok := <-b; ok {
				log.Printf(string(s[i]))
				a <- true
			}
		}
	}()
	a <- true
	<-done
}


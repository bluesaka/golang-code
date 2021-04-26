/**
atomic
原子操作，减少data race数据竞争
CAS (Compare And Swap)

data race
数据竞争
go run -race main.go
*/

package main

import (
	"log"
	"sync"
	"sync/atomic"
)

func main() {
	//atomic1()
	//dataRace)
	//dataRaceWithLock()
	//dataRaceWithLock()
}

func atomic1() {
	var wg sync.WaitGroup
	var count, count2 int64

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count++
			atomic.AddInt64(&count2, 1)
		}()
	}
	wg.Wait()
	log.Println("count:", count)
	log.Println("count2:", count2)
}

type data struct {
	a []int
}

func dataRace() {
	var s []int
	go func() {
		i := 0
		for {
			i++
			s = []int{i, i + 1, i + 2, i + 3, i + 4}
		}
	}()

	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			log.Println(s)
		}()
	}
	wg.Wait()
}

func dataRaceWithLock() {
	var s []int
	var mu sync.RWMutex

	go func() {
		i := 0
		for {
			i++
			mu.Lock()
			s = []int{i, i + 1, i + 2, i + 3, i + 4}
			mu.Unlock()
		}
	}()

	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.RLock()
			log.Println(s)
			mu.RUnlock()
		}()
	}
	wg.Wait()
}

func dataRaceWithAtomic() {
	var s []int
	var v atomic.Value

	go func() {
		i := 0
		for {
			i++
			s = []int{i, i + 1, i + 2, i + 3, i + 4}
			v.Store(s)
		}
	}()

	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			val := v.Load()
			log.Println(val)
		}()
	}
	wg.Wait()
}

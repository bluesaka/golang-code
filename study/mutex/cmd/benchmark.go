/**
2021/04/07 14:37:56 testRwMutexReadOnly cost: 428.214707ms
2021/04/07 14:37:58 testMutexReadOnly cost: 1.579857191s
2021/04/07 14:37:58 testSyncMapReadOnly cost: 118.397003ms

2021/04/07 14:38:01 testRwMutexWriteOnly cost: 2.396842978s
2021/04/07 14:38:03 testMutexWriteOnly cost: 1.949537436s
2021/04/07 14:38:07 testSyncMapWriteOnly cost: 4.634974354s

2021/04/07 14:38:08 testRwMutexReadWrite cost: 1.082658942s
2021/04/07 14:38:10 testMutexReadWrite cost: 1.884788924s
2021/04/07 14:38:10 testSyncMapReadWrite cost: 276.691676ms

2021/04/07 14:38:11 testRwMutexReadWrite cost: 692.902331ms
2021/04/07 14:38:13 testMutexReadWrite cost: 1.718180202s
2021/04/07 14:38:13 testSyncMapReadWrite cost: 138.800543ms



总结：
只读场景：sync.Map > RwMutex >> Mutex
只写场景：Mutex > RwMutex >> sync.Map
读写场景（读80% 写20%）：sync.Map >> RwMutex > Mutex
读写场景（读98% 写2%）：sync.Map > RwMutex >> Mutex
*/
package main

import (
	"log"
	"sync"
	"time"
)

var (
	num  = 10000
	gnum = 1000
)

func main() {
	count := 10000
	d1 := 5
	d2 := 50
	testRwMutexReadOnly(count)
	testMutexReadOnly(count)
	testSyncMapReadOnly(count)

	testRwMutexWriteOnly(count)
	testMutexWriteOnly(count)
	testSyncMapWriteOnly(count)

	testRwMutexReadWrite(count, d1)
	testMutexReadWrite(count, d1)
	testSyncMapReadWrite(count, d1)

	testRwMutexReadWrite(count, d2)
	testMutexReadWrite(count, d2)
	testSyncMapReadWrite(count, d2)
}

func testRwMutexReadOnly(count int) {
	t1 := time.Now()
	var wg sync.WaitGroup
	var r = newRwMutex(count)
	wg.Add(gnum)
	for i := 0; i < gnum; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < num; j++ {
				r.get(j)
			}
		}()
	}
	wg.Wait()
	log.Println("testRwMutexReadOnly cost:", time.Now().Sub(t1).String())
}

func testMutexReadOnly(count int) {
	t1 := time.Now()
	var wg sync.WaitGroup
	var r = newMutex(count)
	wg.Add(gnum)
	for i := 0; i < gnum; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < num; j++ {
				r.get(j)
			}
		}()
	}
	wg.Wait()
	log.Println("testMutexReadOnly cost:", time.Now().Sub(t1).String())
}

func testSyncMapReadOnly(count int) {
	t1 := time.Now()
	var wg sync.WaitGroup
	var r = newSyncMap(count)
	wg.Add(gnum)
	for i := 0; i < gnum; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < num; j++ {
				r.Load(j)
			}
		}()
	}
	wg.Wait()
	log.Println("testSyncMapReadOnly cost:", time.Now().Sub(t1).String())
}

func testRwMutexWriteOnly(count int) {
	t1 := time.Now()
	var w = &sync.WaitGroup{}
	var m = newRwMutex(count)
	w.Add(gnum)
	for i := 0; i < gnum; i++ {
		go func() {
			defer w.Done()
			for in := 0; in < num; in++ {
				m.set(in, in)
			}
		}()
	}
	w.Wait()
	log.Println("testRwMutexWriteOnly cost:", time.Now().Sub(t1).String())
}

func testMutexWriteOnly(count int) {
	t1 := time.Now()
	var w = &sync.WaitGroup{}
	var m = newMutex(count)
	w.Add(gnum)
	for i := 0; i < gnum; i++ {
		go func() {
			defer w.Done()
			for in := 0; in < num; in++ {
				m.set(in, in)
			}
		}()
	}
	w.Wait()
	log.Println("testMutexWriteOnly cost:", time.Now().Sub(t1).String())
}

func testSyncMapWriteOnly(count int) {
	t1 := time.Now()
	var w = &sync.WaitGroup{}
	var m = newSyncMap(count)
	w.Add(gnum)
	for i := 0; i < gnum; i++ {
		go func() {
			defer w.Done()
			for in := 0; in < num; in++ {
				m.Store(in, in)
			}
		}()
	}
	w.Wait()
	log.Println("testSyncMapWriteOnly cost:", time.Now().Sub(t1).String())
}

func testRwMutexReadWrite(count, div int) {
	t1 := time.Now()
	var w = &sync.WaitGroup{}
	var m = newRwMutex(count)
	w.Add(gnum)
	for i := 0; i < gnum; i++ {
		if i%div != 0 {
			go func() {
				defer w.Done()
				for in := 0; in < num; in++ {
					m.get(in)
				}
			}()
		} else {
			go func() {
				defer w.Done()
				for in := 0; in < num; in++ {
					m.set(in, in)
				}
			}()
		}
	}
	w.Wait()
	log.Println("testRwMutexReadWrite cost:", time.Now().Sub(t1).String())
}

func testMutexReadWrite(count, div int) {
	t1 := time.Now()
	var w = &sync.WaitGroup{}
	var m = newMutex(count)
	w.Add(gnum)
	for i := 0; i < gnum; i++ {
		if i%div != 0 {
			go func() {
				defer w.Done()
				for in := 0; in < num; in++ {
					m.get(in)
				}
			}()
		} else {
			go func() {
				defer w.Done()
				for in := 0; in < num; in++ {
					m.set(in, in)
				}
			}()
		}
	}
	w.Wait()
	log.Println("testMutexReadWrite cost:", time.Now().Sub(t1).String())
}

func testSyncMapReadWrite(count, div int) {
	t1 := time.Now()
	var w = &sync.WaitGroup{}
	var m = newSyncMap(count)
	w.Add(gnum)
	for i := 0; i < gnum; i++ {
		if i%div != 0 {
			go func() {
				defer w.Done()
				for in := 0; in < num; in++ {
					m.Load(in)
				}
			}()
		} else {
			go func() {
				defer w.Done()
				for in := 0; in < num; in++ {
					m.Store(in, in)
				}
			}()
		}
	}
	w.Wait()
	log.Println("testSyncMapReadWrite cost:", time.Now().Sub(t1).String())
}

//type IGet interface {
//	get(i int) int
//}

type rwMutex struct {
	mu *sync.RWMutex
	m  map[int]int
}

func (r *rwMutex) get(i int) int {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.m[i]
}

func (r *rwMutex) set(k, v int) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.m[k] = v
}

type mutex struct {
	mu *sync.Mutex
	m  map[int]int
}

func (r *mutex) get(i int) int {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.m[i]
}

func (r *mutex) set(k, v int) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.m[k] = v
}

func newRwMutex(count int) *rwMutex {
	//var t = &rwMutex{}
	var t = new(rwMutex)
	t.mu = &sync.RWMutex{}
	t.m = make(map[int]int, count)
	for i := 0; i < count; i++ {
		t.m[i] = 0
	}
	return t
}

func newMutex(count int) *mutex {
	var t = new(mutex)
	t.mu = &sync.Mutex{}
	t.m = make(map[int]int, count)
	for i := 0; i < count; i++ {
		t.m[i] = 0
	}
	return t
}

func newSyncMap(count int) *sync.Map {
	t := &sync.Map{}
	for i := 0; i < count; i++ {
		t.Store(i, 0)
	}
	return t
}

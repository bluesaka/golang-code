package _map

import (
	"fmt"
	"runtime"
)

/**
https://github.com/golang/go/issues/20135
 */
func MapGCIssue() {
	v := struct{}{}
	a := make(map[int]struct{})

	for i := 0; i < 10000; i++ {
		a[i] = v
	}
	fmt.Println("len:", len(a))
	runtime.GC()
	fmt.Println("len:", len(a))
	printMemStats("After Map Add 10000")

	for i := 0; i < 9999; i++ {
		delete(a, i)
	}
	fmt.Println("len:", len(a))
	runtime.GC()
	fmt.Println("len:", len(a))
	printMemStats("After Map Delete 9999")

	for i := 0; i < 10000-1; i++ {
		a[i] = v
	}
	fmt.Println("len:", len(a))
	runtime.GC()
	fmt.Println("len:", len(a))
	printMemStats("After Map Add 9999 again")

	a = nil
	fmt.Println("len:", len(a))
	runtime.GC()
	fmt.Println("len:", len(a))
	printMemStats("After Map Set nil")
}

func printMemStats(mag string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%v：memory = %vKB, GC Times = %v\n", mag, m.Alloc/1024, m.NumGC)
}

// MapRange map的遍历顺序是随机的
func MapRange() {
	m := make(map[int]string, 9)
	m[1] = "a"
	m[2] = "b"
	m[3] = "c"

	for _, v := range m {
		fmt.Println(v)
	}
}

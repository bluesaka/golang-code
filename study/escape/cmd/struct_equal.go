package main

import "fmt"

type House struct{}

func main() {
	t()
	t2()
	t3()
}

func t() {
	a := &House{}
	b := &House{}
	// 未发生内存逃逸，所以不相等
	println(a, b, a == b) // 0xc000098f78 0xc000098f78 false
}

func t2() {
	a := &House{}
	b := &House{}
	// 内存逃逸后相等，因为都指向zerobase
	fmt.Println(a, b, a == b) // &{} &{} true
}

func t3() {
	a := &House{}
	b := &House{}
	// 发生了内存逃逸，所以两个都相等
	fmt.Println(a, b, a == b) // &{} &{} true
	println(a, b, a == b) // 0x119d408 0x119d408 true
}

/**
内存逃逸导致struct比较的问题

go run -gcflags="-m -l" struct_equal.go

./struct_equal.go:14:7: &House literal does not escape
./struct_equal.go:15:7: &House literal does not escape
./struct_equal.go:21:7: &House literal escapes to heap
./struct_equal.go:22:7: &House literal escapes to heap
./struct_equal.go:24:13: ... argument does not escape
./struct_equal.go:24:22: a == b escapes to heap
./struct_equal.go:28:7: &House literal escapes to heap
./struct_equal.go:29:7: &House literal escapes to heap
./struct_equal.go:31:13: ... argument does not escape
./struct_equal.go:31:22: a == b escapes to heap

*/
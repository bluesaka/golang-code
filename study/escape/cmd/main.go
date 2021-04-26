/**
内存逃逸
 - 指针逃逸
 - 栈空间不足逃逸
 - 动态类型逃逸

go build -gcflags=-m

$ go build -gcflags -m main.go
# command-line-arguments
./main.go:16:6: can inline test
./main.go:12:6: can inline main
./main.go:13:6: inlining call to test
./main.go:17:2: moved to heap: i
*/
package main

import "fmt"

func main() {
	test()
	// 动态类型逃逸，println打印的类型是 ...interface{}
	fmt.Println("a")
}

// 指针逃逸
func test() *int {
	i := 1
	return &i
}

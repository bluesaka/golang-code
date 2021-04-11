package main

import "fmt"

// go install
// 会在$GOPATH/bin下生成cmd可执行文件
// 执行 $GOPATH/bin/cmd 输出 hello world
func main() {
	fmt.Println("hello world")
}

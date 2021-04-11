package main

import "go-code/study/golint"

func main() {
	test(1)
	golint.Say(2)
}

// 代码检查工具
// golint main.go / golint cmd
// func parameter Id should be ID
func test(Id int) int {
	return Id
}

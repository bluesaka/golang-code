/**
断言assert

类型转换推荐使用cast库  `github.com/spf13/cast`
*/
package main

import "fmt"

func main() {
	assert2()
}

func assert1() {
	var a interface{}
	// panic: interface conversion: interface {} is nil, not string
	fmt.Println("string assert:", a.(string))
}

func assert2() {
	var a interface{}
	a = "string a"
	// ok
	fmt.Println("string assert:", a.(string))

	value, ok := a.(string)
	if !ok {
		fmt.Println("a is not string type")
		return
	}
	fmt.Println("value is:", value)
}

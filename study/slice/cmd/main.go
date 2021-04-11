package main

import (
	"fmt"
	"reflect"
)

func main() {
	m := []int{1, 2, 3}
	fmt.Println(reflect.TypeOf(m))
	fmt.Println(m)

	m1 := [3]int{1, 2, 3}
	fmt.Println(reflect.TypeOf(m1))
	fmt.Println(m1)

	m2 := [...]int{1, 2, 3}
	fmt.Println(reflect.TypeOf(m2))
	fmt.Println(m2)

	m3 := [3]string{"1,2,3"}
	fmt.Println(reflect.TypeOf(m3))
	fmt.Println(m3)

	m4 := [3]string{"1", "2", "3"}
	fmt.Println(reflect.TypeOf(m4))
	fmt.Println(m4)
}

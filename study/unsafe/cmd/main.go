package main

import (
	"fmt"
	"go-code/study/unsafe"
)

func main() {
	//pointer1()
	unsafe.SliceTest()
	unsafe.MapTest()
}

func pointer1() {
	a := 3
	double(&a)
	fmt.Println("a =", a) // a = 6

	p := &a
	double(p)
	fmt.Println("a =", a, p == nil) // a = 12 false
}

func double(x *int) {
	*x += *x
	/**
	x指向的值变为原来两倍，而x本身指向nil不影响外层的a

	&a -> 3

	&a -> 3
	x -> 3

	&a -> 6
	x -> 6

	&a -> 6
	x -> nil
	 */
	x = nil
}
package main

import "fmt"

type AAA int
const (
	A1 AAA = iota
	A2
)

func main() {
	a := uint(1)
	b := uint(2)
	fmt.Println(a-b) // 18446744073709551615

	a1 := uint32(1)
	b1 := uint32(2)
	fmt.Println(a1-b1) // 4294967295

	a2 := uint64(1)
	b2 := uint64(2)
	fmt.Println(a2-b2) // 18446744073709551615

	fmt.Println(A1)
	fmt.Println(A2)
}
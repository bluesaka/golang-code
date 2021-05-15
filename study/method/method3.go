package method

import (
	"fmt"
	"reflect"
)

func VariadicFunc(i int, j ...int) {
	fmt.Println("i is:", i)                      // i is: 1
	fmt.Println("j is:", j)                      // j is: [2 3]
	fmt.Println("j type is:", reflect.TypeOf(j)) // []int

	for _, v := range j {
		fmt.Println("v is:", v)
	}
}

func Test5() {
	VariadicFunc(1, 2, 3)
}

/**
@link https://www.zhihu.com/zvideo/1331923200800632832
 */
package method

import (
	"fmt"
	"reflect"
)

type A struct {
	name string
}

func (a A) GetName() string {
	a.name = "Hi " + a.name
	return a.name
}

func NameOfA(a A) string {
	a.name = "Hi " + a.name
	return a.name
}

func NameTest1() {
	a := A{name: "mike"}
	fmt.Println(a.GetName())
	fmt.Println(A.GetName(a))
}

func NameTest2() {
	a1 := reflect.TypeOf(A.GetName)
	a2 := reflect.TypeOf(NameOfA)

	fmt.Println("a1 type:", a1)        // a1 type: func(method.A) string
	fmt.Println("a2 type:", a2)        // a2 type: func(method.A) string
	fmt.Println("a1 == a2:", a1 == a2) // true
}

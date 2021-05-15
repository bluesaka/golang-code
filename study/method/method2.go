/**
@link https://www.zhihu.com/zvideo/1331923200800632832
*/
package method

import (
	"fmt"
)

type A2 struct {
	name string
}

func (a A2) GetName() string {
	return a.name
}

// GetFunc func closure
func GetFunc() func() string {
	a := A2{name: "mike in GetFunc"}
	return a.GetName
}

// GetFunc2 equals to GetFunc
func GetFunc2() func() string {
	a := A2{name: "mike in GetFunc2"}
	return func() string {
		return A2.GetName(a)
	}
}

func NameTest3() {
	a := A2{name: "mike in NameTest3"}

	f1 := a.GetName
	fmt.Println(f1()) // mike in NameTest3

	f2 := GetFunc()
	fmt.Println(f2()) // mike in GetFunc
}

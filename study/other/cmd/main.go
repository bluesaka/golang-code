package main

import (
	"fmt"
	"sort"
)

type user struct {
	ID int
}

func main() {
	var str []string
	str = append(str, "b1")
	str = append(str, "a1")
	str = append(str, "A2")
	fmt.Println(str)
	sort.Strings(str)
	fmt.Println(str)
}

func t() (user, error) {
	return user{ID: 678}, nil
}

func t2(params map[string]string) {
	params["bbb"] = "333"
}

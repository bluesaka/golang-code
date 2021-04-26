package main

import (
	json2 "go-code/study/json"
	"log"
)

func main() {
	u := json2.User{
		Name: "hello",
		Age:  18,
	}
	b, _ := json2.Marshal(u)
	log.Println(b)

	user := json2.Unmarshal(b)
	log.Printf("%+v\n", user)
}

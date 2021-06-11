package main

import (
	"go-code/study/goroutine"
	"log"
)

func main() {
	err := goroutine.Timeout(goroutine.DoSleep)
	log.Println("err:", err)
}

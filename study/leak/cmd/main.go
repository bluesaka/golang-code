package main

import (
	"context"
	"go-code/study/leak"
	"log"
	"runtime"
	"time"
)

func main() {
	log.Println("start")
	err := leak.GoroutineLeakTimeout(context.Background())
	log.Println(err)
	log.Println("goroutine num:", runtime.NumGoroutine())
	time.Sleep(time.Second * 3)
	log.Println("goroutine num:", runtime.NumGoroutine())
}

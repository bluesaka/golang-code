package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hello world")
		w.Write([]byte("hello world"))
	})

	log.Println("server starting on :8000")
	go http.ListenAndServe(":8000", nil)

	// 捕捉信号
	log.Println("listening signal")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("server quit")
}
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	srv := &http.Server{
		Addr: ":8000",
	}
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hello world")
		w.Write([]byte("hello world"))
	})

	log.Println("server starting on :8000")
	go srv.ListenAndServe()

	// 捕捉信号
	log.Println("listening signal")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// 优雅关闭服务
	log.Println("server shutdown")
	srv.Shutdown(context.Background())

	// 等待10秒处理未完成的任务
	log.Println("sleep 10 seconds for unfinished service")
	time.Sleep(10 * time.Second)
	log.Println("server exited")
}

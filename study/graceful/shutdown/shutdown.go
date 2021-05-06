/**
graceful shutdown 优雅停止
使用channel监听信号量
在服务停止前做一些操作，如关闭redis连接池、mysql连接池等
 */

package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Println("server start...")

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("server shutdown...")

	log.Println("do something")

	log.Println("server exited")


}
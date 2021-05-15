package main

import (
	"context"
	"flag"
	"github.com/gin-gonic/gin"
	"log"
	"my-gin/routers"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	port := flag.String("port", "6001", "port")
	flag.Parse()

	if true {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	routers.SetRouters(r)

	srv := &http.Server{
		Addr:           ":" + *port,
		Handler:        r,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("server start on " + *port)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Server Shutdown: ", err)
	}

	log.Println("server exit")
}

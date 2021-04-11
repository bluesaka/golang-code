package main

import (
	"context"
	"log"
	"time"
)

func main() {
	timeout(context.Background())
}

func timeout(c context.Context) {
	ctx, cancel := context.WithTimeout(c, time.Second * 5)
	defer cancel()

	select {
	case <-ctx.Done():
		log.Println("timeout")
		return
	}
}

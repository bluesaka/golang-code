package timeout

import (
	"context"
	"log"
	"testing"
	"time"
)

func TestDeadline(t *testing.T) {
	ctx, cancel := ShrinkDeadline(context.Background(), time.Second*5)
	defer cancel()

	log.Println("start")

	done := make(chan struct{}, 1)
	go func() {
		doTask(ctx)
		done <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		log.Println("timeout")
	case <-done:
		log.Println("done")
	}

	log.Println("end")
}

func doTask(ctx context.Context) {
	taskMysql(ctx)
	taskHttp(ctx)
	taskRpc(ctx)
}

func taskMysql(ctx context.Context) bool {
	t := time.Second * 2
	ctx, cancel := ShrinkDeadline(ctx, t)
	defer cancel()
	time.Sleep(t)
	log.Println("taskMysql done")
	return true
}

func taskHttp(ctx context.Context) bool {
	t := time.Second * 2
	ctx, cancel := ShrinkDeadline(ctx, t)
	defer cancel()
	time.Sleep(t)
	log.Println("taskHttp done")
	return true
}

func taskRpc(ctx context.Context) bool {
	t := time.Second * 1
	ctx, cancel := ShrinkDeadline(ctx, t)
	defer cancel()
	time.Sleep(t)
	log.Println("taskRpc done")
	return true
}

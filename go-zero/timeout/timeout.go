package timeout

/**
 * 超时
 * @link https://mp.weixin.qq.com/s/5Q05d6OwvS-Zj-yNGwoh8A
 */

import (
	"context"
	"log"
	"time"

	"github.com/tal-tech/go-zero/core/contextx"
)

func hardWork(job interface{}) error {
	//time.Sleep(time.Minute)
	time.Sleep(time.Second * 10)
	log.Println("hardWork done")
	return nil
}

func hardWork2(job interface{}) error {
	panic("panic")
	time.Sleep(time.Minute)
	return nil
}

func hardWork4(ctx context.Context, job interface{}) error {
	//time.Sleep(time.Minute)
	time.Sleep(time.Second * 10)
	log.Println("goroutine id", ctx.Value("id"), "hardWork done")
	return nil
}

func RequestWork(ctx context.Context, job interface{}) error {
	return hardWork(job)
}

func RequestWork2(ctx context.Context, job interface{}) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	// create channel with buffer size 1 to avoid goroutine leak
	/**
	 * 如果不设置缓冲大小，当2s超时后函数退出，done <- hardWork4(ctx, job)会一直卡这些不进去，
	 * 导致每个RequestWork2请求都会一直占用一个goroutine
	 * 设置缓冲大小后，不论是否超时，done <- hardWork4(ctx, job)都能写入而不卡住goroutine
	 * 向一个没goroutine接受的channel写数据，是可以的
	 */
	done := make(chan error, 1)
	go func() {
		done <- hardWork4(ctx, job)
	}()

	select {
	case err := <-done:
		log.Println("goroutine id", ctx.Value("id"), "done")
		return err
	case <-ctx.Done():
		log.Println("goroutine id", ctx.Value("id"), "timeout")
		return ctx.Err()
	}
}

func RequestWork3(ctx context.Context, job interface{}) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	done := make(chan error, 1)
	panicChan := make(chan interface{}, 1)
	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()

		done <- hardWork2(job)
	}()

	select {
	case err := <-done:
		return err
	case p := <-panicChan:
		panic(p)
	case <-ctx.Done():
		return ctx.Err()
	}
}

// ------------------------------------------------------------------------------------------------
func hardWork3(job interface{}) error {
	time.Sleep(time.Second * 10)
	return nil
}

func RequestWork4(ctx context.Context, job interface{}) error {
	ctx, cancel := contextx.ShrinkDeadline(ctx, time.Second*2)
	defer cancel()

	done := make(chan error, 1)
	panicChan := make(chan interface{}, 1)
	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()

		done <- hardWork2(job)
	}()

	select {
	case err := <-done:
		return err
	case p := <-panicChan:
		panic(p)
	case <-ctx.Done():
		return ctx.Err()
	}
}

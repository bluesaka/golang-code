/**
https://zeromicro.github.io/go-zero/tokenlimit.html

go-zero 中的 tokenlimit 限流方案适用于瞬时流量冲击，现实请求场景并不以恒定的速率。
令牌桶相当预请求，当真实的请求到达不至于瞬间被打垮。
当流量冲击到一定程度，则才会按照预定速率进行消费。

但是生产token上，不能按照当时的流量情况作出动态调整，不够灵活，还可以进行进一步优化。
此外可以参考Token bucket WIKI 中提到分层令牌桶，根据不同的流量带宽，分至不同排队中。

*/
package main

import (
	"github.com/tal-tech/go-zero/core/limit"
	"github.com/tal-tech/go-zero/core/stores/redis"
	"log"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

const (
	burst   = 100
	rate    = 100
	seconds = 5
)

func main() {
	tokenLimit()
}

func tokenLimit() {
	log.Println("start...")
	tl := limit.NewTokenLimiter(rate, burst, redis.NewRedis("localhost:6379", "node"), "tokenlimit")
	timer := time.NewTimer(time.Second * seconds)
	quit := make(chan struct{}, 1)
	defer timer.Stop()
	go func() {
		<-timer.C
		close(quit)
		//quit <- struct{}{}
	}()

	var allowed, denied int32
	var wg sync.WaitGroup
	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			for {
				select {
				case <-quit:
					wg.Done()
					return
				default:
					if tl.Allow() {
						atomic.AddInt32(&allowed, 1)
					} else {
						atomic.AddInt32(&denied, 1)
					}
				}
			}
		}()
	}
	wg.Wait()
	log.Printf("allowed: %d, denied: %d, qps: %d\n", allowed, denied, (allowed+denied)/seconds)
}

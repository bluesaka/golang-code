/**
https://zeromicro.github.io/go-zero/periodlimit.html

go-zero 中的 periodlimit 限流方案是基于 redis 计数器，通过调用 redis lua script ，保证计数过程的原子性，
同时保证在分布式的情况下计数是正常的。但是这种方案存在缺点，因为它要记录时间窗口内的所有行为记录，
如果这个量特别大的时候，内存消耗会变得非常严重。
*/
package main

import (
	"github.com/tal-tech/go-zero/core/limit"
	"github.com/tal-tech/go-zero/core/stores/redis"
	"log"
	"time"
)

const (
	seconds = 2
	total   = 100
	quota   = 10
)

func main() {
	periodLimit()
}

func periodLimit() {
	pl := limit.NewPeriodLimit(seconds, quota, redis.NewRedis("127.0.0.1:6379", redis.NodeType, ""), "periodlimit")
	key := "first"
	//code, err := pl.Take(key)
	//if err != nil {
	//	logx.Error(err)
	//	return true
	//}

	for i := 0; i < 5; i++ {
		n := 0
		switch i {
		case 0:
			n = 2
		case 1:
			n = 20
		case 2:
			n = 8
		case 3:
			n = 5
		case 4:
			n = 6
		}
		for j := 0; j < n; j++ {
			code, err := pl.Take(key)
			if err != nil {
				panic(err)
			}
			switch code {
			case limit.OverQuota:
				log.Println("over")
			case limit.Allowed:
				log.Println("allow")
			case limit.HitQuota:
				log.Println("hit")
			default:
				log.Println("default")
			}
		}
		time.Sleep(time.Second)
	}
}

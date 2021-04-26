/**
https://zeromicro.github.io/go-zero/periodlimit.html

go-zero 中的 periodlimit 限流方案是基于 redis 计数器，通过调用 redis lua script ，保证计数过程的原子性，
同时保证在分布式的情况下计数是正常的。但是这种方案存在缺点，因为它要记录时间窗口内的所有行为记录，
如果这个量特别大的时候，内存消耗会变得非常严重。
*/
package main

import (
	"github.com/tal-tech/go-zero/core/limit"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/stores/redis"
)

const (
	seconds = 1
	total   = 100
	quota   = 5
)

func main() {
	periodLimit()
}

func periodLimit() bool {
	pl := limit.NewPeriodLimit(seconds, quota, redis.NewRedis("127.0.0.1:6379", redis.NodeType, ""), "periodlimit")
	key := "first"
	code, err := pl.Take(key)
	if err != nil {
		logx.Error(err)
		return true
	}

	switch code {
	case limit.OverQuota:
		logx.Errorf("OverQuota key: %v", key)
		return false
	case limit.Allowed:
		logx.Infof("AllowedQuota key: %v", key)
		return true
	case limit.HitQuota:
		logx.Errorf("HitQuota key: %v", key)
		// todo: maybe we need to let users know they hit the quota
		return false
	default:
		logx.Errorf("DefaultQuota key: %v", key)
		// unknown response, we just let the sms go
		return true
	}
}

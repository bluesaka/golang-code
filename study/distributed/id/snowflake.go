package id

import (
	"errors"
	"strconv"
	"sync"
	"time"
)

const (
	sequenceBit = 12                       // 序号部分 12位
	sequenceMax = -1 ^ (-1 << sequenceBit) // 序号最大值 2^12 - 1 = 4095
	nodeBit     = 10                       // 机器节点部分 10位
	nodeMax     = -1 ^ (-1 << nodeBit)     // 机器节点最大值 2^10 - 1 = 1023
)

type Snowflake struct {
	mu    sync.Mutex
	node  int
	epoch time.Time
	time  int64
	step  int64
}

// NewSnowflake returns Snowflake object
func NewSnowflake(node int) (*Snowflake, error) {
	if node < 0 || node > nodeMax {
		return nil, errors.New("node must be between 0 and " + strconv.Itoa(nodeMax))
	}

	n := &Snowflake{}
	n.node = node
	n.epoch = time.Now()

	return n, nil
}

// Generate generate id
func (s *Snowflake) Generate() int64 {
	s.mu.Lock()

	now := time.Since(s.epoch).Nanoseconds() / int64(time.Millisecond)
	if s.time == now {
		s.step++
		// 当前毫秒内生成的序号已经超出最大范围，等待下一毫秒重新生成
		if s.step > sequenceMax {
			for now <= s.time {
				now = time.Since(s.epoch).Nanoseconds() / int64(time.Millisecond)
				s.step = 0
			}
		}
	} else {
		s.step = 0
	}

	// id = (毫秒时间戳差值 << 22) | (节点值 << 12) | 序列号
	id := (now << (nodeBit + sequenceBit)) | int64(s.node<<sequenceBit) | s.step
	s.time = now
	s.mu.Unlock()
	return id
}

package heartbeat

import (
	"fmt"
	"log"
	"sync"
	"time"
)

var (
	addrList = make(map[string]time.Time)
	mu       sync.Mutex
)

func Listen() {
	mq := NewRabbitMQ()
	defer mq.Close()

	msgList := mq.Consume(queueName)
	for msg := range msgList {
		mu.Lock()
		fmt.Printf("receive msg: %v\n", string(msg.Body))
		addrList[string(msg.Body)] = time.Now()
		mu.Unlock()
	}
}

func removeExpiredAddr() {
	for {
		mu.Lock()
		for addr, lastTime := range addrList {
			if lastTime.Add(time.Second * 10).Before(time.Now()) {
				delete(addrList, addr)
				log.Printf("addr: %s, lastTime: %v removed\n", addr, lastTime)
			}
		}
		mu.Unlock()
		time.Sleep(time.Second)
	}
}

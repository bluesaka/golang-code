package main

import (
	"context"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"log"
	"strconv"
)

func main() {
	//os.Setenv("ROCKETMQ_GO_LOG_LEVEL", "INFO")
	p, err := rocketmq.NewProducer(
		producer.WithNameServer([]string{"127.0.0.1:9876"}),
		//producer.WithNsResolver(primitive.NewPassthroughResolver([]string{"127.0.0.1:9876"})),
		producer.WithRetry(2),
		producer.WithGroupName("GroupName"),
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 3; i++ {
		msg := primitive.NewMessage("topic_a", []byte("hello RocketMQ Go client~ "+strconv.Itoa(i)))
		res, err := p.SendSync(context.Background(), msg)
		if err != nil {
			log.Printf("send message error: %s\n", err)
		} else {
			log.Printf("send message success, result: %s\n", res.String())
		}
	}

	err = p.Shutdown()
	if err != nil {
		log.Printf("shutdown producer error: %s", err.Error())
	}
}

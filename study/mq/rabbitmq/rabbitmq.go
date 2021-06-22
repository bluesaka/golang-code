/**
RabbitMQ
github.com/streadway/amqp
*/
package rabbitmq

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/streadway/amqp"
	"log"
	"strconv"
	"time"
)

const (
	QueueName     = "go-queue-test"
	DelayQueue    = "go-delay-queue-test"
	DeadQueue     = "go-dead-queue-test"
	DelayExchange = "go-delay-exchange-test"
	TTL           = "10000"
)

// Producer producer
func Producer() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		log.Fatalf("amqp dial error: %v\n", err)
	}
	defer conn.Close()

	// channel
	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("open channel error: %v\n", err)
	}
	defer channel.Close()

	// queue declare
	queue, err := channel.QueueDeclare(QueueName, true, false, false, false, nil)
	if err != nil {
		log.Fatalf("queue declare error: %v\n", err)
	}
	log.Printf("queue: %+v\n", queue)

	// publish message
	for i := 1; i <= 10; i++ {
		err = channel.Publish("", QueueName, false, false, amqp.Publishing{
			ContentType:  "text/plain",
			Body:         []byte("hello world " + strconv.Itoa(i)),
			DeliveryMode: 2,
		})
		if err != nil {
			log.Fatalf("publish msg error: %v\n", err)
		}
	}

	log.Println("success")
}

// Consumer consumer
func Consumer() {
	consume(QueueName)
}

// ProducerDelay producer delay
func ProducerDelay() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		log.Fatalf("amqp dial error: %v\n", err)
	}
	defer conn.Close()

	// channel
	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("open channel error: %v\n", err)
	}
	defer channel.Close()

	args := amqp.Table{
		"x-dead-letter-exchange":    DelayExchange, // 死信队列 => 延时交换机
		"x-dead-letter-routing-key": DeadQueue,     // 路由键默认是队列名
		//"x-message-ttl": 10000, // 队列的所有消息过期时间，优先级小于消息指定的Expiration过期时间
	}

	// 声明死信队列
	if _, err := channel.QueueDeclare(DeadQueue, true, false, false, false, args); err != nil {
		log.Fatalf("DeadQueue declare error: %v\n", err)
	}

	// 声明延时队列
	if _, err := channel.QueueDeclare(DelayQueue, true, false, false, false, nil); err != nil {
		log.Fatalf("DelayQueue declare error: %v\n", err)
	}

	// 声明延时交换机
	if err := channel.ExchangeDeclare(DelayExchange, "direct", true, false, false, false, nil); err != nil {
		log.Fatalf("DelayExchange declare error: %v\n", err)
	}

	// 绑定死信队列+延时队列+延时交换机
	if err := channel.QueueBind(DelayQueue, DeadQueue, DelayExchange, false, nil); err != nil {
		log.Fatalf("QueueBind error: %v\n", err)
	}

	data, err := jsoniter.Marshal(map[string]interface{}{
		"name": "abc",
		"time": time.Now().Format("2006-01-02 15:04:05"),
	})
	if err != nil {
		log.Fatalf("Marshal error: %v\n", err)
	}

	err = channel.Publish("", DeadQueue, false, false, amqp.Publishing{
		ContentType:  "text/plain",
		Body:         data,
		DeliveryMode: 2,
		Expiration:   TTL, // 过期时间，单位ms
	})
	if err != nil {
		log.Fatalf("publish msg error: %v\n", err)
	}

	log.Println("success")
}

func ConsumerDelay() {
	consume(DelayQueue)
}

// consume do consume
func consume(queueName string) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		log.Fatalf("amqp dial error: %v\n", err)
	}
	defer conn.Close()

	// channel
	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("open channel error: %v\n", err)
	}
	defer channel.Close()

	// queue declare
	queue, err := channel.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		log.Fatalf("queue declare error: %v\n", err)
	}
	log.Printf("queue: %+v\n", queue)

	msgChan, err := channel.Consume(queueName, "", false, false, false, false, nil)
	if err != nil {
		log.Fatalf("consume error: %v", err)
	}

	go func() {
		for msg := range msgChan {
			log.Printf("message: %+v, body string: %s\n", msg, string(msg.Body))
			if err := msg.Ack(true); err != nil {
				log.Println("ack error:", err)
			}
		}
	}()

	log.Println("connected to RabbitMQ, waiting for messages")
	select {}
}

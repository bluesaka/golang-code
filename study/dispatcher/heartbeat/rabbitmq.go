package heartbeat

import (
	"github.com/streadway/amqp"
	"log"
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewRabbitMQ() *RabbitMQ {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		log.Fatalf("amqp dial error: %v\n", err)
	}

	// channel
	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("open channel error: %v\n", err)
	}

	_, err = channel.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		log.Fatalf("queue declare error: %v\n", err)
	}

	return &RabbitMQ{
		conn:    conn,
		channel: channel,
	}
}

func (mq *RabbitMQ) Produce(queueName, msg string) {
	err := mq.channel.Publish("", queueName, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(msg),
	})
	if err != nil {
		log.Fatal(err)
	}
}

func (mq *RabbitMQ) Consume(queueName string) <-chan amqp.Delivery {
	msg, err := mq.channel.Consume(queueName, "", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}
	return msg
}

func (mq *RabbitMQ) Close() {
	if err := mq.channel.Close(); err != nil {
		log.Fatal(err)
	}
	if err := mq.conn.Close(); err != nil {
		log.Fatal(err)
	}
}

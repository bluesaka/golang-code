/**
Kafka
github.com/Shopify/sarama
*/
package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"sync"
	"time"
)

var (
	address = []string{"localhost:9092"}
	topic   = "go-test-topic"
	timeFormat = "2006-01-02 15:04:05"
)

func Producer() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll // 发送完数据需要leader和follower都确认
	//config.Producer.Partitioner = sarama.NewRandomPartitioner // 随机一个partition
	config.Producer.Partitioner = sarama.NewManualPartitioner // 选择message中的Partition，前提是要kafka的topic中有该分区
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回
	config.Producer.Return.Errors = true
	//config.Version = sarama.V0_11_0_2

	producer, err := sarama.NewAsyncProducer(address, config)
	if err != nil {
		log.Fatalf("kafka new producer error: %v", err)
	}
	defer producer.AsyncClose()

	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Key:       sarama.StringEncoder("go_test_key"),
		Partition: 1, //搭配sarama.NewManualPartitioner手动指定分区，前提是该topic有该分区，不会会报错`kafka: partitioner returned an invalid partition index`
		Timestamp: time.Now(),
	}

	fmt.Println("input msg")
	var value string
	for {
		fmt.Scanln(&value)
		msg.Value = sarama.ByteEncoder(value)
		fmt.Printf("stdin: %s\n", value)

		producer.Input() <- msg

		select {
		case success := <-producer.Successes():
			fmt.Printf("produce success, partition: %d, offset: %d, timestamp: %s\n", success.Partition, success.Offset, success.Timestamp.Format(timeFormat))
		case err := <-producer.Errors():
			fmt.Printf("produce error: %v\n", err.Err)
		}
	}
}

func Consumer() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	consumer, err := sarama.NewConsumer(address, config)
	if err != nil {
		log.Fatalf("kafka new consumer error: %v\n", err)
	}
	defer consumer.Close()

	topicConsumer(topic, consumer)
}

func topicConsumer(topic string, consumer sarama.Consumer) {
	// topic的分区列表，每个分区都有自己的offset信息，根据情况为每个分区创建消费者进行消费
	partitionList, err := consumer.Partitions(topic)
	if err != nil {
		log.Fatalf("consumer get partition list error: %v\n", err)
	}
	if len(partitionList) == 0 {
		log.Printf("topic: %s has no partition\n", topic)
		return
	}
	fmt.Println("partitionList:", partitionList)

	var wg sync.WaitGroup
	wg.Add(len(partitionList))

	for partition := range partitionList {
		//partitionConsumer, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetOldest)
		//可以使用redis等保存offset，并根据业务做相关幂等处理
		partitionConsumer, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			log.Fatalf("kafka ConsumePartition error: %v\n", err)
		}

		go func() {
			defer wg.Done()
			defer partitionConsumer.Close()
			for {
				select {
				case msg := <-partitionConsumer.Messages():
					fmt.Printf("consumer msg offset: %d, partition: %d, timestamp: %s, value: %s\n",
						msg.Offset, msg.Partition, msg.Timestamp.Format(timeFormat), string(msg.Value))
				case err := <-partitionConsumer.Errors():
					fmt.Printf("consume error: %v\n", err)
				}
			}
		}()
	}

	wg.Wait()
}
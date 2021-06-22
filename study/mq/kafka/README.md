### Kafka

#### 相关定义
- Topic：指Kafka的消息分类
- Partition：Topic物理上的分组，一个topic可以分为多个partition，每个partition是一个有序的队列。partition中的每条消息都会被分配一个有序的id（offset）
- Message：消息
- Producer：消息生产者
- Consumer：消息消费者
- Broker：代理，Kafka集群中的一台或多台服务器统称为broker

#### 常用命令
```
kafka基于zookeeper，zk默认端口2181， kafka默认端口9092

# 查看topic list
kafka-topics --zookeeper 127.0.0.1:2181 --list

# 查看某一个topic的信息
kafka-topics --zookeeper 127.0.0.1:2181 --topic go-test-topic --describe

# 创建topic
kafka-topics --zookeeper 127.0.0.1:2181 --create --replication-factor 1 --partitions 1 --topic test

# 修改topic的partition分区数
kafka-topics --zookeeper 127.0.0.1:2181 --topic go-test-topic --alter --partitions 2

# 删除topic
kafka-topics --zookeeper 127.0.0.1:2181 --delete --topic test

# 生产消息
kafka-console-producer --broker-list 127.0.0.1:9092 --topic test

# 消费消息
kafka-console-consumer --bootstrap-server 127.0.0.1:9092 --topic test
```

#### 工具
```
kafka-tool
```
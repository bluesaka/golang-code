### RocketMQ

#### 相关定义
- Topic：消息主题，一级消息类型，生产者向其发送消息
- 生产者：消息发布者，生产并发送消息至Topic
- 消费者：消息订阅者，从Topic接受并消费消息
- 消息：数据
- 消息属性：生产者可以为消息定义属性，包含Message Key和Tag
- Group：一类生产者或消费者，这类生产者或消费者通常生产或消费同一类消息，且消息发布或订阅的逻辑一致

#### Command
```
RockerMQ的NameServer默认端口是9876

# Start NameServer
nohup sh bin/mqnamesrv &

# Start Broker
nohup sh bin/mqbroker -n localhost:9876 &

# Shutdown Server
sh bin/mqshutdown broker
sh bin/mqshutdown namesrv

# 查看集群列表
mqadmin clusterlist -n 127.0.0.1:9876
```

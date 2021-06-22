### RabbitMQ

#### 相关定义
- Broker： 消息队列服务器
- Exchange： 消息交换机，它指定消息按什么规则，路由到哪个队列
- Queue： 消息队列载体，每个消息都会被投入到一个或多个队列
- Binding： 绑定，它的作用就是把exchange和queue按照路由规则绑定起来
- Routing Key： 路由键，exchange根据这个关键字进行消息投递
- VHost： 虚拟主机，一个broker里可以开设多个vhost，用作不同用户的权限分离。
- Producer： 消息生产者，就是投递消息的程序
- Consumer： 消息消费者，就是接受消息的程序
- Channel： 消息通道，在客户端的每个连接里，可建立多个channel，每个channel代表一个会话任务

### Exchange交换机
- direct
- fanout
- topic
- headers

```
direct直连交换机，是RabbitMQ Broker的默认类型，可以不指定路由键RoutingKey（默认是与Queue同名），通过消息携带的路由键将消息投递到对应的队列中
            message(RoutingKey:xxx)                               match RoutingKey
Producer            ==>                 direct exchange      ==>      Queue       ==>     Consumer


Fanout扇形交换机，会忽略RoutingKey，加消息广播到所有绑定的队列中
          message (Without RoutingKey)                  all bind queue
Producer            ==>                 fanout exchange      ==>      Queue       ==>     Consumer


Topic主题交换机，将消息投递到RoutingKey和BindKey都匹配的队列中，匹配规则有如下特点：
1. RoutingKey包含.字符串，如 log.info log.error
2. BindingKey(Queue bind routingKey)也会包含.字符串进行匹配，如 log.*  log.#，*匹配一个单词，#匹配0个或多个单词
```

#### 延时队列
```
Producer => RoutingKey  => DeadExchange => DeadQueue  => DelayRoutingKey => DelayExchange => DelayQueue => Consumer
                                    (TTL时间内未被消费成为死信)
生产者    => 路由键      => 死信交换机     => 死信队列    => 延时路由键       => 延时交换机     => 延时队列    => 消费者
```

#### 工作队列
```
工作队列用来将耗时的任务分发给多个消费者，消息默认是通过轮询机制，分发给各个消费者
```

#### 端口
```
RabbitMQ默认端口5672，admin默认端口15672
```
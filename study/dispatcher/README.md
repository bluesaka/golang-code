### 分布式任务调度系统
- 负载均衡
- 任务协调
- 节点选举
- 健康检查
- 心跳检查

#### 负载均衡策略
- 随机
- 轮询
- 加权轮询
- hash
- 一致性hash
- 最短响应时间
- 最小连接数

##### 加权轮询
```
Nginx负载均衡配置：

http {    
    upstream cluster {    
        server a weight=1;    
        server b weight=2;    
        server c weight=4;    
    }    
    ...  
}   

根据以上配置，Nginx每接受7个客户端请求，就会把其中1个转发到a，2个转发到b，4个转发到c
加权轮询算法生成的节点序列为 [c, c, b, c, a, b, c]
```


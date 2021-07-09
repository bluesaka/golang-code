### Prometheus

>Prometheus(普罗米修斯)是由SoundCloud使用Go语言开发的开源时序列数据库(TSDB). 
>将采集到的样本数据以时间序列（time-series）的方式保存在内存数据库中，并且定时保存到硬盘上。
>time-series是按照时间戳和值的序列顺序存放的，我们称之为向量(vector)

#### 四种度量指标(metric)类型

- Counter
    > Counter 计数器是累计度量指标，只增不减，主要用于统计服务的请求数等
- Gauge
    > Gauge 是一个度量指标，可增可减，常用于温度、内存、并发请求等统计
- Histogram
    > Histogram柱状图，有三种作用：
    - 1.对每个采样点进行统计，达到各个分类值中(bucket)
    - 2.对采样点的值累计和(sum)
    - 3.对采样点的次数累计和(count)
- Summary
    > 与 Histogram柱状图类似，summary是采样点氛围图统计，常用于请求持续时间和响应大小，有三种左右：
    - 1.对于每个采样点进行统计，并形成分位图
    - 2.统计班上所有同学的总成绩(sum)
    - 3.统计班上同学的考试总人数(count)

#### Counter
```
# TYPE promhttp_metric_handler_requests_total counter
promhttp_metric_handler_requests_total{code="200"} 54
promhttp_metric_handler_requests_total{code="500"} 0
promhttp_metric_handler_requests_total{code="503"} 0
```

#### Gauge
```
# TYPE go_goroutines gauge
go_goroutines 22

# TYPE go_threads gauge
go_threads 14

```

#### Histogram
```
# TYPE http_server_requests_duration_ms histogram
http_server_requests_duration_ms_bucket{instance="127.0.0.1:8090",job="topic_api",path="/topic/lists",le="5"} 3
http_server_requests_duration_ms_bucket{instance="127.0.0.1:8090",job="topic_api",path="/topic/lists",le="10"} 3
http_server_requests_duration_ms_bucket{instance="127.0.0.1:8090",job="topic_api",path="/topic/lists",le="25"} 3
http_server_requests_duration_ms_bucket{instance="127.0.0.1:8090",job="topic_api",path="/topic/lists",le="50"} 3
http_server_requests_duration_ms_bucket{instance="127.0.0.1:8090",job="topic_api",path="/topic/lists",le="100"} 3
http_server_requests_duration_ms_bucket{instance="127.0.0.1:8090",job="topic_api",path="/topic/lists",le="250"} 3
http_server_requests_duration_ms_bucket{instance="127.0.0.1:8090",job="topic_api",path="/topic/lists",le="500"} 3
http_server_requests_duration_ms_bucket{instance="127.0.0.1:8090",job="topic_api",path="/topic/lists",le="1000"} 3
http_server_requests_duration_ms_bucket{instance="127.0.0.1:8090",job="topic_api",path="/topic/lists",le="+Inf"} 3
http_server_requests_duration_ms_sum{instance="127.0.0.1:8090",job="topic_api",path="/topic/list"} 0
http_server_requests_duration_ms_count{instance="127.0.0.1:8090",job="topic_api",path="/topic/list"} 3

代码示例：
import prom "github.com/prometheus/client_golang/prometheus"
cfg := {
    Namespace:  "http_server"
    Subsystem:  "requests",
    Name:       "duration_ms",
    Help:       "http server requests duration(ms).",
    Labels:     []string{"path", "code"},
    Buckets:    []float64{5, 10, 25, 50, 100, 250, 500, 1000},
}
vec := prom.NewHistogramVec(prom.HistogramOpts{
    Namespace: cfg.Namespace,
    Subsystem: cfg.Subsystem,
    Name:      cfg.Name,
    Help:      cfg.Help,
    Buckets:   cfg.Buckets,
}, cfg.Labels)
value := int64(123)
vec.Observe(value, path, strconv.Itoa(code))

metric name 由 Namespace + "_" + Subsystem + "_" + Name 组成

label也可以在prometheus.yml添加
static_configs:
  labels:
    app: topic
    env: test

```

#### Summary
```
# TYPE go_gc_duration_seconds summary
go_gc_duration_seconds{quantile="0"} 4.1783e-05
go_gc_duration_seconds{quantile="0.25"} 9.5368e-05
go_gc_duration_seconds{quantile="0.5"} 0.000127167
go_gc_duration_seconds{quantile="0.75"} 0.000155352
go_gc_duration_seconds{quantile="1"} 0.000155352
go_gc_duration_seconds_sum 0.00041967
go_gc_duration_seconds_count 4
```


#### 查询数据类型：
- 瞬时数据 (Instant Vector)：包含一组时序，每个时序只有一个点，例如：http_requests_total
- 区间数据 (Range Vector)：包含一组时序，每个时序有多个点，例如：http_requests_total[5m]
- 纯量数据 (Scalar)：纯量只有一个数字，没有时序，例如：count(http_requests_total)


#### PromQL

> PromQL(Prometheus Query Language) 是 Prometheus 的数据查询 DSL(Domain-specific Language) 语言

```
<----------------------- metric -----------------------><- timestamp -><- value ->
go_goroutines{instance="127.0.0.1:8090",job="topic_api"} @1625673600  => 33

go_goroutines 是metric name
instance、job 是label标签
可使用 go_goroutines{job="topic_api"} 进行查询

https://prometheus.io/docs/prometheus/latest/querying/basics/

# 通过metric name查询
http_requests_total

# 通过label查询
{level="info"}

# 通过metric name + label查询
http_requests_total{level="info"}

# 运算符
http_requests_total{code!="200"} // 查询 code 不为 "200" 的数据
http_requests_total{code=~"2.."} // 查询 code 为 "2xx" 的数据
http_requests_total{code!~"2.."} // 查询 code 不为 "2xx" 的数据

// 时间范围 y年、w周、d天、h时、m分、s秒、ms毫秒
http_requests_total[2y]
http_requests_total[2w]
http_requests_total[2d]
http_requests_total[2h]
http_requests_total[2m]
http_requests_total[2s]

# 统计查询
count(http_requests_total)
sum(http_requests_total)
topk(10, http_requests_total) // topk，这里查询top 10
```

#### 配置
```
默认端口 9090
默认配置文件 /usr/local/etc/prometheus.yml
默认参数配置 /usr/local/etc/prometheus.args
其中 --storage.tsdb.path /usr/local/var/prometheus 为数据存储目录

启动服务：prometheus --config.file=/usr/local/etc/prometheus.yml

配置数据采集，采集源来自app.go
scrape_configs:
- job_name: myapp_test
  scrape_interval: 10s
  static_configs:
  - targets:
    - localhost:9000

```

#### 存储
```
Prometheus 使用本地存储，所以集群化需要借助其他工具如InfluxDB等
采用WAL(write-ahead logging)机制，会将数据存储到内存中并记录一份日志，2小时之后自动写入磁盘
数据存储地址：/usr/local/etc/prometheus.args 中的 --storage.tsdb.path /usr/local/etc/prometheus.yml
```
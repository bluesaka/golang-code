## 链路追踪

### 常用的链路追踪系统
- Skywalking
- 阿里 鹰眼
- 大众点评 CAT
- Twitter Zipkin
- Uber Jaeger

### 规范

- OpenTracing
- OpenCensus
- OpenTelemetry = OpenTracing + OpenCensus

```
链路追踪领域重要分为两派，一派是以CNCF为主的OpenTracing，如jaeger、zipkin，
另一派是以谷歌和微软组成的OpenCensus，
之后诞生了OpenTelemetry，兼容OpenCensus和OpenTracing，可以让使用者无需改动或者很小的改动就可以接入OpenTelemetry。

        OpenTracing                             OpenCensus
- One "vertical" (Tracing)              - Many "verticals" (Tracing, Metrics)
- One "layer" (API)                     - Many "layers" (API, impl, infra)
- "Looser" coupling (small scope)       - "Tighter" coupling (framework-y)
- Lots of languages (~12)               - Many languages (5 in beta)
- Broad adoption                        - Broad adoption
- (FYI: Already part of CNCF)

```


### OpenTracing
> OpenTracing是开放式分布式追踪规范，为分布式追踪创建更标准的API和工具，不过CNCF并不是官方标准机构，

### Zipkin

- Docker

```
docker run -d -p 9411:9411 openzipkin/zipkin
```

- Java

```
curl -sSL https://zipkin.io/quickstart.sh | bash -s
java -jar zipkin.jar
```

- Running from Source

```
# get the latest source
git clone https://github.com/openzipkin/zipkin
cd zipkin
# Build the server and also make its dependencies
./mvnw -DskipTests --also-make -pl zipkin-server clean install
# Run the server
java -jar ./zipkin-server/target/zipkin-server-*exec.jar
```

### Jaeger

- Docker

```
docker run -d --name jaeger \
  -e COLLECTOR_ZIPKIN_HOST_PORT=:9411 \
  -p 5775:5775/udp \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 14268:14268 \
  -p 14250:14250 \
  -p 9411:9411 \
  jaegertracing/all-in-one:1.25
```

- Running with all-in-one binary

```
https://www.jaegertracing.io/download/

jaeger-all-in-one --collector.zipkin.host-port=:9411

访问 http://localhost:16686 查看Jaeger UI
```

### Jaeger各端口作用

```
Port	Protocol	Component	Function
5775	UDP	        agent	    accept zipkin.thrift over compact thrift protocol (deprecated, used by legacy clients only)
6831	UDP	        agent	    accept jaeger.thrift over compact thrift protocol
6832	UDP	        agent	    accept jaeger.thrift over binary thrift protocol
5778	HTTP	    agent	    serve configs
16686	HTTP	    query	    serve frontend
14268	HTTP	    collector	accept jaeger.thrift directly from clients
14250	HTTP	    collector	accept model.proto
9411	HTTP	    collector	Zipkin compatible endpoint (optional)
```

### Kratos Tracing

> Kratos框架提供的tracing基于OpenTelemetry
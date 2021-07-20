### 链路追踪

#### OpenTracing
> OpenTracing是开放式分布式追踪规范，为分布式追踪创建更标准的API和工具，不过CNCF并不是官方标准机构，

#### Zipkin
> 分布式追踪系统
> 默认端口9411

#### Zipkin Run

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
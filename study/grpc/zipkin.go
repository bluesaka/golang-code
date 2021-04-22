/**
@link https://segmentfault.com/a/1190000016677230

Zipkin 分布式追踪系统

Start with jar
```
curl -sSL https://zipkin.io/quickstart.sh | bash -s
java -jar zipkin.jar
```

Start with docker
```
# Note: this is mirrored as ghcr.io/openzipkin/zipkin
docker run -d -p 9411:9411 openzipkin/zipkin
```

http://localhost:9411/zipkin/



 */

package grpc

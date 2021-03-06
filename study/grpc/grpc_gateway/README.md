### GRPC Gateway

#### 安装protoc
```
go install \
    github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
    github.com/golang/protobuf/protoc-gen-go
```

#### gRPC
```
# 生成pb文件
cd proto/simple && protoc --go_out=plugins=grpc:. hello.proto

# 启动gRPC服务
go run grpc_simple_server.go

# grpcui调用gRPC服务
grpcui -plaintext localhost:8080
```

#### gRPC Gateway
```
cd proto

# 生成pb文件
protoc --go_out=plugins=grpc:. ./gateway/hello_grpc_gateway.proto

在google/api文件夹下添加annotations.proto和http.proto
# 生成pb.gw文件
protoc --grpc-gateway_out=logtostderr=true:. ./gateway/hello_grpc_gateway.proto

# 开启gRPC服务
go run grpc_gateway_grpc_server.go

# 开启HTTP服务 (依赖gRPC服务)
go run grpc_gateway_http_server.go

# 或者直接开启gRPC+HTTP服务
go run grpc_gateway_server.go

# 请求接口
curl -d '{"name":"hello"}' http://localhost:8081/v1/echo

# import "google/api/annotations.proto" 标红
不影响编译和运行，Goland可以在Preferences --> Languages & Framewokrs --> Protocol Buffers 添加相关thrid_party目录，如：
.../pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis
```

#### Kratos参考
```
cd proto

protoc --proto_path=. \
       --proto_path=./third_party \
       --go_out=paths=source_relative:. \
       --go-http_out=paths=source_relative:. \
       --go-grpc_out=paths=source_relative:. \
       --validate_out=paths=source_relative,lang=go:. \
       --openapiv2_out . \
        kratos_proto/hello_grpc_gateway.proto

生成的源码文件为：
xxx.pb.go
xxx.pb.validate.go
xxx.pb.swagger.json
xxx_grpc.pb.go
xxx_http.pb.go

# 问题
grpc.SupportPackageIsVersion7 undefined
之前为了兼容etcd v3.3，grpc的版本较低，可升级到etcd3.5和grpc版本
```

#### Swagger
```
# 生成swagger.json文件
protoc --swagger_out=logtostderr=true:. ./gateway/hello_grpc_gateway.proto

# 安装swagger
brew tap go-swagger/go-swagger
brew install go-swagger

# 启动swagger服务
swagger serve --host=0.0.0.0 --port=9000 --no-open  hello_grpc_gateway.swagger.json

# 访问swagger
http://localhost:9000/docs
```
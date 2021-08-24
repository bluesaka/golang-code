### grpcui

> grpcui、grpcurl 是grpc调用工具，类似于curl对http的调用

#### 安装
```
// install with go mod
go install github.com/fullstorydev/grpcui/cmd/grpcui
go install github.com/fullstorydev/grpcui/cmd/grpcurl
```

#### 使用
```
grpcui -help

grpcui -plaintext localhost:8088
grpcurl -plaintext -d '{"user_id":1}' 'localhost:8088' shop.User/user_list

报错：Failed to compute set of methods to expose: server does not support the reflection API
解决：reflection.Register(s)
```


### bloomrpc
```
https://github.com/uw-labs/bloomrpc.git
brew cask install bloomrpc
```
/**
protoc是一款用C++编写的工具，其可以将proto文件编译为指定语言的代码，不过并不支持go语言
protoc-gen-go 是protobuf编译插件系统中的Go版本

安装protoc：
进入 https://github.com/protocolbuffers/protobuf/releases
选择合适的压缩包文件，如 protoc-3.14.0-osx-x86_64.zip
解压后将 protoc-3.14.0-osx-x86_64/bin/protoc 移动到 $GOPATH/bin 下
mv protoc $GOPATH/bin
或者使用brew install protobuf

安装protoc-gen-go
go get -u github.com/golang/protobuf/protoc-gen-go@v1.3.2

生成 pb.go 文件
protoc --go_out=. helloworld.proto

很多坑都来自于protoc和protoc-gen-go的版本问题
如protoc-gen-go v1.26.0版本执行上面的命令，会报以下错误：
```
protoc-gen-go: invalid Go import path "." for "helloworld.proto"
The import path must contain at least one forward slash ('/') character.
See https://developers.google.com/protocol-buffers/docs/reference/go-generated#package for more information.

```

各种奇怪的报错，建议使用合适的版本进行规避

 */
package main

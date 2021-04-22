/**
一、GOPATH 和 GOROOT
golang中没有项目，只有包的概念
- GOROOT：GOROOT是Go的安装目录（类似于java的JDK）
- GOPATH：GOPATH是工作空间,保存go项目代码和第三方依赖包

$ ls $GOROOT/bin
go    godoc gofmt

$ ls $GOPATH/bin
demo          dlv           goctl         golint        gopls         protoc-gen-go update

GOPATH可以设置多个，其中第一个是默认的包目录
- src 存放源代码文件
- pkg 存放编译后的文件
- bin 存放编译后的可执行文件
- go get 和 go install下载的包会在第一个path中的src目录
- import包时的搜索路径


$ go env 查看go环境
GO111MODULE="on"
GOARCH="amd64"
...

$ go version
go version go1.15.5 darwin/amd64

go module：
GO111MODULE=on或者auto下，go get会栽下载到$GOPATH/pkg/mod下

go build 命令：
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o outputname x.go
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o outputname x.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o outputname x.go

// 编译的时候加上-ldflags "-s -w"参数去掉符号表和调试信息，一般能减小20%的大小。
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ./bin/main

go install命令与go build类似，go install知识将编译的中间文件放在 $GOPATH 的 pkg 目录下，并将编译结果放在 $GOPATH 的 bin 目录下
 */

package gopath

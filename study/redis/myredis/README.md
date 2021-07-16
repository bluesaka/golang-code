### Go编写Redis服务器

> [链接](https://juejin.cn/post/6974560368026714119)

- 如何编写 Go 语言 TCP 服务器
- 设计并实现安全可靠的通信协议（redis 协议）
- 如何使用 Go 语言开发高并发程序
- 设计和实现分布式集群以及分布式事务
- 熟悉链表、哈希表、跳表以及时间轮等常用数据结构

### Redis协议

自 Redis 2.0 以后的通信统一为 RESP 协议（Redis Serialization Protocol)

RESP 是一个二进制安全的文本协议，工作于 TCP 协议上。RESP 以行作为单位，客户端和服务器发送的命令或数据一律以 \r\n（CRLF）作为换行符。

二进制安全是指允许协议中出现任意字符而不会导致故障。比如 C 语言的字符串以 \0 作为结尾不允许字符串中间出现 \0，而 Go 语言的 string 则允许出现 \0，我们说 Go 语言的 string 是二进制安全的，而 C 语言字符串不是二进制安全的。

#### RESP 5种格式：

- 简单字符串（Simple String）： 服务器用来返回简单的结果，比如 "OK" 非二进制安全，且不允许换行
- 错误信息（Error）：服务器用来返回简单的错误信息，比如 "ERR Invalid Synatx" 非二进制安全，且不允许换行
- 整数（Integer）：llen、scard 等命令的返回值，64 位有符号整数
- 字符串（Bulk String）：二进制安全字符串，比如 get 等命令的返回值
- 数组（Array，又称 Multi Bulk Strings）：Bulk String 数组，客户端发送指令以及 lrange 等命令响应的格式

#### RESP 通过第一个字符来表示格式

- 简单字符串：以"+"开始，如："+OK\r\n"
- 错误：以"-" 开始，如："-ERR Invalid Synatx\r\n"
- 整数：以":"开始，如：":123\r\n"
- 字符串：以 $ 开始
- 数组：以 * 开始
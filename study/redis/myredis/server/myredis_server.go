package main

import (
	"bufio"
	"io"
	"log"
	"net"
)

func main() {
	ListenAnsServe(":8000")
}

func ListenAnsServe(address string) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}
	defer lis.Close()
	log.Println("server start on", address)

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Println("server accept error:", err)
			continue
		}
		go Handle(conn)
	}
}

func Handle(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		// ReadString 会一直阻塞直到遇到分隔符 '\n'
		// 遇到分隔符后 ReadString 会返回上次遇到分隔符到现在收到的所有数据
		// 若在遇到分隔符之前发生异常, ReadString 会返回已收到的数据和错误信息
		msg, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				log.Println("connection closed")
			} else {
				log.Println("read error:", err)
			}
			return
		}

		// 将收到的信息发送给客户端
		conn.Write([]byte(msg))
	}
}

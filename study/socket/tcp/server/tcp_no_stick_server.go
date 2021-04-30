package main

import (
	"bufio"
	"go-code/study/socket/tcp/schema"
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	log.Println("server starting at 127.0.0.1:20000")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("accept failed, err:", err)
			continue
		}

		go process2(conn)
	}

}

func process2(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		msg, err := schema.Decode(reader)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("read from client failed, err:", err)
			break
		}

		log.Println("receive data from client:", msg)
	}
}

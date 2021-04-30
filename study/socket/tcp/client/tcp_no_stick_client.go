package main

import (
	"go-code/study/socket/tcp/schema"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	log.Println("client starting at 127.0.0.1:20000")

	for i := 0; i < 20; i++ {
		data, err := schema.Encode(`Hello there, how are you?`)
		if err != nil {
			log.Println("client encode msg failed, err", err)
		}
		conn.Write(data)
	}
}

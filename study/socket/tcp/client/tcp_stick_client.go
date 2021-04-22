package main

import (
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

	/**
	出现粘包，多条数据粘在一起
	2021/04/19 13:36:27 receive data from client: Hello there, how are you?Hello there, how are you?Hello there, how are you?Hello there, how are you?Hello there, how are you?Hello there, how are you?Hello there, how are you?Hello there, how are you?Hello there, how are you?Hello there, how are you?
	2021/04/19 13:36:27 receive data from client: Hello there, how are you?Hello there, how are you?Hello there, how are you?Hello there, how are you?Hello there, how are you?Hello there, how are you?Hello there, how are you?Hello there, how are you?Hello there, how are you?Hello there, how are you?
	*/
	for i := 0; i < 20; i++ {
		conn.Write([]byte(`Hello there, how are you?`))
	}
}

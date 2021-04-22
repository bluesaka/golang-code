package main

import (
	"bufio"
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

		go process(conn)
	}

}

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	var buf [1024]byte

	for {
		n, err := reader.Read(buf[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("read from client failed, err:", err)
			break
		}

		b := buf[:n]
		log.Println("receive data from client:", string(b))
		conn.Write(b)
	}
}

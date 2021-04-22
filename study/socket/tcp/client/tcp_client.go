package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		panic(err)
	}

	log.Println("client starting at 127.0.0.1:20000")

	reader := bufio.NewReader(os.Stdin)
	for {
		s, _ := reader.ReadString('\n')
		s = strings.TrimSpace(s)
		if strings.ToUpper(s) == "Q" {
			return
		}

		_, err := conn.Write([]byte(s))
		if err != nil {
			log.Println("send failed, err:", err)
			return
		}

		var buf [1024]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			log.Println("read failed, err:", err)
			return
		}

		log.Println("receive data from server:", string(buf[:n]))
	}
}

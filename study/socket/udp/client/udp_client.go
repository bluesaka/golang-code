package main

import (
	"bufio"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 30000,
	})
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	log.Println("client starting at 127.0.0.1:30000")

	reader := bufio.NewReader(os.Stdin)
	for {
		s, _ := reader.ReadString('\n')
		_, err := conn.Write([]byte(s))
		if err != nil {
			log.Println("send to server failed, err:", err)
			return
		}

		var buf [1024]byte
		n, addr, err := conn.ReadFromUDP(buf[:])
		if err != nil {
			log.Println("read from server failed, err:", err)
			return
		}

		log.Printf("read from %v, msg:%v\n", addr, string(buf[:n]))
	}

}

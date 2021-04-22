package main

import (
	"log"
	"net"
)

func main() {
	listener, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 30000,
	})
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	log.Println("server starting at 127.0.0.1:30000")

	for {
		var buf [1024]byte
		n, addr, err := listener.ReadFromUDP(buf[:])
		if err != nil {
			log.Println("read from udp failed, err:", err)
			continue
		}

		_, err = listener.WriteToUDP(buf[:n], addr)
		if err != nil {
			log.Printf("write to %v failed, err:%v\n", addr, err)
			continue
		}
	}

}

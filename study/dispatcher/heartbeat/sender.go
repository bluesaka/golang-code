package heartbeat

import (
	"errors"
	"log"
	"net"
	"time"
)

const (
	queueName = "heartbeat-test-queue"
)

func Start() {
	mq := NewRabbitMQ()
	defer mq.Close()

	addr, err := getIP()
	if err != nil {
		log.Fatal(err)
	}
	for {
		mq.Produce(queueName, addr)
		time.Sleep(time.Second)
	}
}

func getIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, addr := range addrs {
		if ip, ok := addr.(*net.IPNet); ok && !ip.IP.IsLoopback() && ip.IP.To4() != nil {
			return ip.IP.String(), nil
		}
	}
	return "", errors.New("no ip")
}
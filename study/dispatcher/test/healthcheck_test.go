package test

import (
	"fmt"
	"go-code/study/dispatcher/healthcheck"
	"testing"
	"time"
)

func TestHealthCheck(t *testing.T) {
	healthcheck.AddAddr("http://www.baidu.com", "http://www.qq.com", "http://www.not")
	fmt.Println("addrList", healthcheck.GetAliveAddrList())

	go healthcheck.Start()

	time.Sleep(time.Second * 30)
	fmt.Println("addrList", healthcheck.GetAliveAddrList())
}

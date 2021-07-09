package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func main() {
	// prometheus metrics接口
	http.Handle("/metrics", promhttp.Handler())

	// 访问 localhost:9001/metrics，会返回goroutine、thread、gc、堆栈等信息
	http.ListenAndServe(":9001", nil)
}

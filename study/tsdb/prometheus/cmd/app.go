package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"time"
)

var (
	cp = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_test",
		Help: "myapp test desc",
	})
)

func main() {
	go recordMetrics()

	// 访问 localhost:9000，可以看到 myapp_test 的信息
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":9000", nil)
}

func recordMetrics() {
	for {
		cp.Inc()
		time.Sleep(time.Second * 2)
	}
}

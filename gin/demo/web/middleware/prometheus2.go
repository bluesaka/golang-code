package middleware

import (
	"github.com/gin-gonic/gin"
	prom "github.com/prometheus/client_golang/prometheus"
	"strconv"
	"time"
)

const serverNamespace = "editor_test_http"

var (
	durationHistogram *prom.HistogramVec
	reqCounter        *prom.CounterVec
	registerCollector []prom.Collector
)

var (
	durationHistogramCfg = promCfg{
		Namespace: serverNamespace,
		Subsystem: "request",
		Name:      "duration_ms",
		Help:      "http request duration(ms).",
		Labels:    []string{"path", "method", "code"},
		// 统计的柱状图指标(单位毫秒)， (-,10], (10, 100], (100, 500], (500, 1000], (1000, 2000], (2000, 5000]
		Buckets:   []float64{10, 100, 500, 1000, 2000, 5000},
	}

	reqCounterCfg = promCfg{
		Namespace: serverNamespace,
		Subsystem: "request",
		Name:      "total",
		Help:      "http server request count.",
		Labels:    []string{"path", "code"},
	}
)

type promCfg struct {
	Namespace string
	Subsystem string
	Name      string
	Help      string
	Labels    []string
	Buckets   []float64
}

func init() {
	// durationHistogram
	durationHistogram = prom.NewHistogramVec(prom.HistogramOpts{
		Namespace: durationHistogramCfg.Namespace,
		Subsystem: durationHistogramCfg.Subsystem,
		Name:      durationHistogramCfg.Name,
		Help:      durationHistogramCfg.Help,
		Buckets:   durationHistogramCfg.Buckets,
	}, durationHistogramCfg.Labels)
	prom.MustRegister(durationHistogram)
	registerCollector = append(registerCollector, durationHistogram)

	// reqCounter
	reqCounter = prom.NewCounterVec(prom.CounterOpts{
		Namespace: reqCounterCfg.Namespace,
		Subsystem: reqCounterCfg.Subsystem,
		Name:      reqCounterCfg.Name,
		Help:      reqCounterCfg.Help,
	}, reqCounterCfg.Labels)
	prom.MustRegister(reqCounter)
	registerCollector = append(registerCollector, reqCounter)
}

// Prometheus2 prometheus rest middleware
func Prometheus2(c *gin.Context) {
	startTime := time.Now()
	c.Next()
	if c.Request.Method == "OPTIONS" {
		return
	}

	durationHistogram.WithLabelValues(c.Request.URL.Path, c.Request.Method, strconv.Itoa(c.Writer.Status())).
		Observe(float64(time.Since(startTime) / time.Millisecond))

	reqCounter.WithLabelValues(c.Request.URL.Path, strconv.Itoa(c.Writer.Status())).Inc()
}

// ClosePrometheus unregister prometheus collector
func ClosePrometheus() {
	for _, c := range registerCollector {
		prom.Unregister(c)
	}
}

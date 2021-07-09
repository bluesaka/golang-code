package prometheus

import (
	prom "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var (
	durationHistogram *prom.HistogramVec
)

type config struct {
	Namespace string
	Subsystem string
	Name      string
	Help      string
	Labels    []string
	Buckets   []float64
}

func HistogramTest() {
	http.Handle("/metrics", promhttp.Handler())
	initHistogram()
	go histogram()
	defer prom.Unregister(durationHistogram)
	http.ListenAndServe(":8091", nil)

}

func initHistogram() {
	// config
	cfg := config{
		Namespace: "z_http",
		Subsystem: "request",
		Name:      "duration_ms",
		Help:      "http request duration(ms).",
		Labels:    []string{"path", "code"},
		Buckets:   []float64{5, 10, 50, 100, 500, 1000},
	}

	// create HistogramVec
	durationHistogram = prom.NewHistogramVec(prom.HistogramOpts{
		Namespace: cfg.Namespace,
		Subsystem: cfg.Subsystem,
		Name:      cfg.Name,
		Help:      cfg.Help,
		Buckets:   cfg.Buckets,
	}, cfg.Labels)

	// register
	prom.MustRegister(durationHistogram)
}

func histogram() {
	// observe each route
	for i := 1; i <= 10; i++ {
		// run 10 times for each route
		for j := 0; j < 10; j++ {
			rand.Seed(time.Now().UnixNano())
			code := "200"
			randNum := rand.Intn(2)
			randMs := rand.Intn(100)
			if randNum == 1 {
				code = "500"
			}
			path := "route_" + strconv.Itoa(i)

			durationHistogram.WithLabelValues(path, code).Observe(float64(randMs))
		}
	}
}

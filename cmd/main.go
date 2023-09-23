package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/catinapoke/availability/internal/ping"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	delay     = time.Second * 2
	parameter = "address"
)

var (
	GaugeOptions = prometheus.GaugeOpts{
		Name: "ping_time",
		Help: "Page load time for specific url",
	}
	GaugeVec = promauto.NewGaugeVec(
		GaugeOptions, []string{parameter},
	)
)

func main() {
	ctx := context.TODO()
	urls := []string{"https://vk.com", "http://192.168.0.1", "https://1.1.1.1"}

	for _, url := range urls {
		log.Printf("Start ping for %s\n", url)

		pinger := ping.NewPinger(url, delay, GaugeVec.With(prometheus.Labels{parameter: url}))
		pinger.StartAsync(ctx)
	}

	log.Printf("Start listening prometheus\n")
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}

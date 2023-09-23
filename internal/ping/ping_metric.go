package ping

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type Pinger struct {
	address string
	delay   time.Duration
	gauge   prometheus.Gauge
}

func NewPinger(address string, delay time.Duration, metric prometheus.Gauge) *Pinger {
	return &Pinger{
		address: address,
		delay:   delay,
		gauge:   metric,
	}
}

func (p *Pinger) StartAsync(ctx context.Context) {
	go func() {
		tick := time.NewTicker(p.delay)
		defer tick.Stop()
		defer log.Printf("stop metric for %s", p.address)

		for {
			select {
			case <-ctx.Done():
				return
			case <-tick.C:
				seconds, err := p.request(p.address)

				if err != nil {
					p.setMetric(0)
				} else {
					p.setMetric(seconds)
				}
			}
		}
	}()
}

func (p *Pinger) setMetric(seconds float64) {
	p.gauge.Set(seconds)
}

func (Pinger) request(url string) (float64, error) {
	start := time.Now()

	_, err := http.Get(url)
	if err != nil {
		return -1, err
	}

	return time.Since(start).Seconds(), nil
}

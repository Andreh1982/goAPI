package metric

import (
	"github.com/prometheus/client_golang/prometheus"
)

type service struct {
	httpRequestHistogram *prometheus.HistogramVec
}

func NewPrometheusService() (*service, error) {
	http := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "http",
		Name:      "request_duration_seconds",
		Help:      "The latency of the HTTP requests.",
		Buckets:   prometheus.DefBuckets,
	}, []string{"handler", "method", "code"})

	s := &service{
		httpRequestHistogram: http,
	}

	err := prometheus.Register(s.httpRequestHistogram)

	if err != nil && err.Error() != "duplicate metrics collector registration attempted" {
		return nil, err
	}
	return s, nil
}

func (s *service) SaveHTTP(h *HTTP) {
	s.httpRequestHistogram.WithLabelValues(h.Handler, h.Method, h.StatusCode).Observe(h.Duration)
}

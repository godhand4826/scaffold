package restful

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func RegisterMetricsHandler(router *chi.Mux) {
	router.Handle("/metrics", promhttp.Handler())
}

func Metrics() func(h http.Handler) http.Handler {
	var httpDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "Duration of HTTP requests.",
		Buckets: prometheus.DefBuckets,
	}, []string{"path", "method", "status"})

	prometheus.MustRegister(httpDuration)

	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			next.ServeHTTP(ww, r)

			httpDuration.WithLabelValues(
				chi.RouteContext(r.Context()).RoutePattern(),
				r.Method,
				strconv.Itoa(ww.Status()),
			).Observe(time.Since(start).Seconds())
		}
		return http.HandlerFunc(fn)
	}
}

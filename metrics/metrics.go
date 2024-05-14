package metrics

import (
	"github.com/golang/glog"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

const (
	MetricsNamespace = "httpserver"
)

type ExecutionTimer struct {
	histo *prometheus.HistogramVec
	start time.Time
	last  time.Time
}

func NewTimer() *ExecutionTimer {
	return NewExecutionTimer(functionLatency)
}

func (t *ExecutionTimer) ObserveTotal() {
	(*t.histo).WithLabelValues("total").Observe(time.Now().Sub(t.start).Seconds())
}

func NewExecutionTimer(latency *prometheus.HistogramVec) *ExecutionTimer {
	now := time.Now()
	return &ExecutionTimer{
		histo: latency,
		start: now,
		last:  now,
	}
}

var functionLatency = CreateExecutionTimeMetrics(MetricsNamespace, "Time Spent.")

func CreateExecutionTimeMetrics(namespace string, metrics string) *prometheus.HistogramVec {
	return prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: namespace,
		Name:      "execution_latency_seconds",
		Help:      metrics,
		Buckets:   prometheus.ExponentialBuckets(0.001, 2, 15),
	}, []string{"step"})
}

func Register() {
	err := prometheus.Register(functionLatency)
	if err != nil {
		glog.Fatal("prometheus register failed", err)
	}
}

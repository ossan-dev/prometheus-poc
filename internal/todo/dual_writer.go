package todo

import (
	"fmt"
	"os"

	"github.com/prometheus/client_golang/prometheus"
)

type fnRegister func() (err error)

var metricReconciliation map[string]prometheus.Counter

func registerMetric(functions ...fnRegister) error {
	for _, fn := range functions {
		err := fn()
		if err != nil {
			return err
		}
	}
	return nil
}

func GetTodosInc() {
	if err := writeMetric("get_todos_request_count", 1.0); err != nil {
		fmt.Fprintf(os.Stderr, "failed to write metric: %v\n", err.Error())
	}
	getTodosCounter.Inc()
}

func GenericMetricInc(metricName string) {
	if err := writeMetric(metricName, 1.0); err != nil {
		fmt.Fprintf(os.Stderr, "failed to write metric: %v\n", err.Error())
	}
	counter, isFound := metricReconciliation[metricName]
	if !isFound {
		fmt.Fprintf(os.Stderr, "no matching prometheus metric with name: %s\n", metricName)
		return
	}
	counter.Inc()
}

func SetupDualWriter() error {
	metricReconciliation = make(map[string]prometheus.Counter, 0)
	return registerMetric(registerPrometheusMetrics, setGraphiteConn)
}

package todo

import "github.com/prometheus/client_golang/prometheus"

var (
	getTodosCounter, createTodoCounter prometheus.Counter
)

// registerPrometheusMetrics registers Prometheus metrics
func registerPrometheusMetrics() (err error) {
	getTodosCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "get_todos_request_count",
			Help: "No of request handled by the GetTodos handler",
		},
	)
	createTodoCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "create_todo_request_count",
			Help: "No of request handled by the CreateTodo handler",
		},
	)
	metricReconciliation["get_todos_request_count"] = getTodosCounter
	prometheus.MustRegister(getTodosCounter)
	prometheus.MustRegister(createTodoCounter)
	return nil
}

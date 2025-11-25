package main

import (
	"net/http"
	"sync"

	"github.com/ossan-dev/prometheuspoc/internal/todo"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	var once sync.Once
	once.Do(todo.RegisterPrometheusMetrics())
	http.HandleFunc("GET /api/todos", todo.GetTodos)
	http.HandleFunc("POST /api/todos", todo.CreateTodo)
	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

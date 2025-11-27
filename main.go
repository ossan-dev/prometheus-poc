package main

import (
	"net/http"

	"github.com/ossan-dev/prometheuspoc/internal/todo"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	if err := todo.SetupDualWriter(); err != nil {
		panic(err)
	}
	http.HandleFunc("GET /api/todos", todo.GetTodos)
	http.HandleFunc("POST /api/todos", todo.CreateTodo)
	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}

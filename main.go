package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type todo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}

var todos []todo = []todo{
	{1, "first"},
	{2, "second"},
}

func postTodo(w http.ResponseWriter, r *http.Request) {
	var todo todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	todos = append(todos, todo)
	w.WriteHeader(http.StatusCreated)
	w.Write(fmt.Appendf([]byte{}, `{"id": %d}`, todo.ID))
}

func main() {
	http.HandleFunc("GET /api/todos", func(w http.ResponseWriter, r *http.Request) {
		data, err := json.MarshalIndent(todos, "", "\t")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.Write(data)
	})
	http.HandleFunc("POST /api/todos", postTodo)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

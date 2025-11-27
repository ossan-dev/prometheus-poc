package todo

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetTodos returns saved todos
func GetTodos(w http.ResponseWriter, r *http.Request) {
	// GetTodosInc()  // old version
	GenericMetricInc("get_todos_request_count")
	data, err := json.MarshalIndent(todos, "", "\t")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(data)
}

// CreateTodo creates a Todo and returns the id
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	createTodoCounter.Inc()
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

var todos []todo = []todo{
	{1, "first"},
	{2, "second"},
}

type todo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}

package main

import (
	"net/http"

	"todo-list/app/handlers/todoHandler"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", todoHandler.TodoIndex).Methods("GET")
	r.HandleFunc("/todos", todoHandler.TodoIndex).Methods("GET")
	// r.HandleFunc("/todos/{todoId}", TodoShow).Methods("GET")
	r.HandleFunc("/todos/new", todoHandler.TodoNew).Methods("GET")
	r.HandleFunc("/todos", todoHandler.TodoCreate).Methods("POST")
	r.HandleFunc("/todos/{todoId}/edit", todoHandler.TodoEdit).Methods("GET")
	r.HandleFunc("/todos/{todoId}/update", todoHandler.TodoUpdate).Methods("POST")
	r.HandleFunc("/todos/{todoId}/done", todoHandler.TodoDone).Methods("POST")
	r.HandleFunc("/todos/{todoId}/undone", todoHandler.TodoUndone).Methods("POST")
	r.HandleFunc("/todos/{todoId}/delete", todoHandler.TodoDelete).Methods("POST")

	http.ListenAndServe(":80", r)
}

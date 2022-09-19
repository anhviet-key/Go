package main

import (
	// "fmt"
	"log"
	// "text/template"
	"net/http"
	"todo-app/package/handler"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
		http.FileServer(http.Dir("/Go/todo-app/static/"))))
	r.HandleFunc("/todo", handler.GetAllTodo).Methods(http.MethodGet)
	r.HandleFunc("/api/todo", handler.GetApiTodo).Methods(http.MethodGet)
	r.HandleFunc("/api/todo/{id}", handler.GetTodoById).Methods(http.MethodGet)
	r.HandleFunc("/api/todo", handler.CreateTodo).Methods(http.MethodPost)
	r.HandleFunc("/api/todo/{id}", handler.UpdateTodo).Methods(http.MethodPut)
	r.HandleFunc("/api/todo/{id}", handler.DeleteTodo).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":8080", r))
}

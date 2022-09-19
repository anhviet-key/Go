package data

import "todo-app/package/todo"

var Todos []todo.Todo

func init() {
	Todos = []todo.Todo{
		{ID: 1, Title: "Task 1", Status: true},
	}
}

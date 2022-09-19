package handler

import (
	"encoding/json"
	"html/template"
	"net/http"
	"sort"
	"strconv"
	"todo-app/package/data"
	"todo-app/package/todo"

	"github.com/gorilla/mux"
)

func GetAllTodo(writer http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/html/layout.html"))
	datas := data.Todos

	sort.Slice(datas, func(i, j int) bool {
		return datas[i].ID > datas[j].ID
	})
	tmpl.Execute(writer, datas)

}
func GetApiTodo(writer http.ResponseWriter, request *http.Request) {
	responseWithJson(writer, http.StatusOK, data.Todos)
}

func GetTodoById(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		responseWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Invalid todo id"})
		return
	}

	for _, todo := range data.Todos {
		if todo.ID == id {
			responseWithJson(writer, http.StatusOK, todo)
			return
		}
	}

	responseWithJson(writer, http.StatusNotFound, map[string]string{"message": "Todo not found"})
}

func CreateTodo(writer http.ResponseWriter, request *http.Request) {
	var newTodo todo.Todo

	if request.Method != "POST" {
		http.Redirect(writer, request, "/", http.StatusSeeOther)
		return
	}
	fInput := request.FormValue("todoInput")
	if fInput == "" {
		return
	}
	newTodo.ID = generateId(data.Todos)
	newTodo.Title = fInput
	newTodo.Status = true
	data.Todos = append(data.Todos, newTodo)
	http.Redirect(writer, request, "/todo", http.StatusSeeOther)
	// if err := json.NewDecoder(request.Body).Decode(&newTodo); err != nil {
	// 	responseWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
	// 	return
	// }
	// newTodo.ID = generateId(data.Todos)
	// data.Todos = append(data.Todos, newTodo)
	// responseWithJson(writer, http.StatusCreated, newTodo)
}

func UpdateTodo(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		responseWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Invalid todo id"})
		return
	}

	var updateTodo todo.Todo
	if err := json.NewDecoder(request.Body).Decode(&updateTodo); err != nil {
		responseWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}
	updateTodo.ID = id

	for i, todo := range data.Todos {
		if todo.ID == id {
			data.Todos[i] = updateTodo
			responseWithJson(writer, http.StatusOK, updateTodo)
			return
		}
	}

	responseWithJson(writer, http.StatusNotFound, map[string]string{"message": "Todo not found"})
}

func DeleteTodo(writer http.ResponseWriter, request *http.Request) {

	// request.Method = "DELETE"
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		responseWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Invalid todo id"})
		return
	}

	for i, todo := range data.Todos {
		if todo.ID == id {
			data.Todos = append(data.Todos[:i], data.Todos[i+1:]...)
			responseWithJson(writer, http.StatusOK, map[string]string{"message": "Todo was deleted"})
			return
		}
	}
	responseWithJson(writer, http.StatusNotFound, map[string]string{"message": "Todo not found"})
}

func responseWithJson(writer http.ResponseWriter, status int, object interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	json.NewEncoder(writer).Encode(object)
}

func generateId(todos []todo.Todo) int {
	var maxId int
	for _, todo := range todos {
		if todo.ID > maxId {
			maxId = todo.ID
		}
	}
	return maxId + 1
}

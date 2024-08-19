package controller

import (
	"TodoList/model"
	"encoding/json"
	"log"
	"net/http"
)

func ApiTodoViewController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		todos, err := model.GetAllTodo()
		if err != nil {
			log.Fatalln("Not Found Todos", err)
		}
		todo, _ := json.MarshalIndent(todos, "", "\t\t")
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Write(todo)
	}
}

package controller

import (
	"TodoList/model"
	"log"
	"net/http"
	"strconv"
)

func TodoHandler(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	page := vals.Get("page")
	todos, err := model.GetAllTodo()
	if err != nil {
		log.Fatalln("Get All Todos Error", err)
	}
	var paginator interface{}
	if page == "" {
		paginator = model.Paginator(todos, 1, 5)
	} else {
		p, _ := strconv.Atoi(page)
		paginator = model.Paginator(todos, p, 5)
	}
	cookie, _ := ReadCookieServer(r)
	if cookie != "" {
		Render(w, paginator, "default", "index", "home", "nav.private", "footer")
	} else {
		Render(w, paginator, "default", "index", "home", "nav", "footer")
	}
}
func TodoView(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	id, _ := strconv.Atoi(vals.Get("todoId"))
	todo, err := model.GetTodo(id)
	if err != nil {
		log.Fatalln("Get All Todos Error", err)
	}
	cookie, _ := ReadCookieServer(r)
	if cookie != "" {
		Render(w, todo, "default", "index", "view", "nav.private", "footer")
	} else {
		Render(w, todo, "default", "index", "view", "nav", "footer")
	}
}

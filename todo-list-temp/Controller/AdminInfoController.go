package controller

import (
	"TodoList/model"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
)

func AdminInfoController(w http.ResponseWriter, r *http.Request) {
	cookie, err := ReadCookieServer(r)
	if err != nil {
		log.Fatalln("Get Cookie Error", err)
		http.Redirect(w, r, "/login", 302)
	}
	if cookie == "" {
		log.Fatalln("Get Cookie Error", err)
		http.Redirect(w, r, "/login", 302)
	} else {
		vals := r.URL.Query()
		pagenow, _ := strconv.Atoi(vals.Get("page"))
		user, err := model.GetUser(cookie)
		if err != nil {
			log.Fatalln("Get User Error", err)
		}
		todos, err := model.GetUserTodo(user.Id)
		if err != nil {
			log.Fatalln("Get Todos Error", err)
		}
		paginator := model.Paginator(todos, pagenow, 10)
		Render(w, paginator, "admin", "default", "todolist", "nav.private")
	}
}

func AdminUserController(w http.ResponseWriter, r *http.Request) {
	cookie, err := ReadCookieServer(r)
	if err != nil {
		log.Fatalln("Get Cookie Error", err)
		http.Redirect(w, r, "/login", 302)
	}
	if cookie == "" {
		log.Fatalln("Get Cookie Error", err)
		http.Redirect(w, r, "/login", 302)
	} else {
		user, err := model.GetUser(cookie)
		if err != nil {
			log.Fatalln("Found User Error", err)
		}
		Render(w, user, "admin", "default", "user", "nav.private")
	}
}
func AdminCreateController(w http.ResponseWriter, r *http.Request) {
	cookie, err := ReadCookieServer(r)
	if err != nil {
		log.Fatalln("Get Cookie Error", err)
		http.Redirect(w, r, "/login", 302)
	}
	if cookie == "" {
		log.Fatalln("Get Cookie Error", err)
		http.Redirect(w, r, "/login", 302)
	} else {
		user, err := model.GetUser(cookie)
		if err != nil {
			log.Fatalln("Found User Error", err)
		}
		if r.Method == "GET" {
			Render(w, nil, "admin", "default", "createtodo", "nav.private")
		} else if r.Method == "POST" {
			title := r.FormValue("title")
			content := r.FormValue("content")
			err = user.CreateTodo(title, content)
			if err != nil {
				log.Fatalln("Create Todo Error", err)
			}
			http.Redirect(w, r, "/admin/todolist/", 302)
		}
	}
}
func AdminEditUserController(w http.ResponseWriter, r *http.Request) {
	cookie, err := ReadCookieServer(r)
	if err != nil {
		log.Fatalln("Get Cookie Error", err)
		http.Redirect(w, r, "/login", 302)
	}
	if cookie == "" {
		log.Fatalln("Get Cookie Error", err)
		http.Redirect(w, r, "/login", 302)
	} else {
		user, err := model.GetUser(cookie)
		if err != nil {
			log.Fatalln("Found User Error", err)
		}
		if r.Method == "GET" {
			Render(w, user, "admin", "default", "edituser", "nav.private")
		} else if r.Method == "POST" {
			user.Username = r.FormValue("username")
			user.Password = r.FormValue("password")
			err = user.UpdateUser()
			if err != nil {
				log.Fatalln("Create Todo Error", err)
			}
			http.Redirect(w, r, "/login", 302)
		}
	}
}
func AdminDoneController(w http.ResponseWriter, r *http.Request) {
	cookie, err := ReadCookieServer(r)
	if err != nil {
		log.Fatalln("Get Cookie Error", err)
		http.Redirect(w, r, "/login", 302)
	}
	if cookie == "" {
		log.Fatalln("Get Cookie Error", err)
		http.Redirect(w, r, "/login", 302)
	} else {
		id, _ := strconv.Atoi(filepath.Base(r.URL.Path))
		todo, _ := model.GetTodo(id)
		err = todo.UpdateTodo()
		if err != nil {
			w.WriteHeader(404)
		}
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
	}
}
func AdminDeleteController(w http.ResponseWriter, r *http.Request) {
	cookie, err := ReadCookieServer(r)
	if err != nil {
		log.Fatalln("Get Cookie Error", err)
		http.Redirect(w, r, "/login", 302)
	}
	if cookie == "" {
		log.Fatalln("Get Cookie Error", err)
		http.Redirect(w, r, "/login", 302)
	} else {
		id, _ := strconv.Atoi(filepath.Base(r.URL.Path))
		todo, _ := model.GetTodo(id)
		err = todo.DeleteTodo()
		if err != nil {
			w.WriteHeader(404)
		}
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
	}
}

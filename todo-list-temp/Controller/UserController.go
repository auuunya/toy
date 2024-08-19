package controller

import (
	"TodoList/model"
	"log"
	"net/http"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	cookie, err := ReadCookieServer(r)
	if err != nil && cookie != "" {
		http.Redirect(w, r, "/admin", 302)
	} else {
		if r.Method == "GET" {
			Render(w, nil, "default", "index", "nav", "signup", "footer")
		}
		if r.Method == "POST" {
			r.ParseForm()
			if r.FormValue("username") == "" || r.FormValue("password") == "" || r.FormValue("email") == "" {
				log.Fatalln("You cannot have null values")
			} else {
				user := &model.User{
					Username: r.FormValue("username"),
					Password: r.FormValue("password"),
					Email:    r.FormValue("email"),
				}
				err := user.CreateUser()
				if err != nil {
					log.Fatalln("Create User Error:", err)
				}
				http.Redirect(w, r, "/login", 302)
			}
		}
	}
}
func Login(w http.ResponseWriter, r *http.Request) {
	cookie, err := ReadCookieServer(r)
	if err != nil && cookie != "" {
		http.Redirect(w, r, "/admin", 302)
	} else {
		if r.Method == "GET" {
			Render(w, nil, "default", "index", "nav", "login", "footer")
		}
		if r.Method == "POST" {
			r.ParseForm()
			email := r.FormValue("email")
			password := r.FormValue("password")
			user, err := model.CheckPass(email, password)
			if err != nil {
				log.Fatalln("Username or Password Error", err)
			}
			WriteCookieServer(w, user.Username, "/")
			http.Redirect(w, r, "/admin/", 302)
		}
	}
}
func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, _ := ReadCookieServer(r)
	DeleteCookieServer(w, cookie, "/")
	http.Redirect(w, r, "/", 302)
}

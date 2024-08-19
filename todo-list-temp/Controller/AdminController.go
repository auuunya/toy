package controller

import (
	"TodoList/model"
	"net/http"
)

func AdminController(w http.ResponseWriter, r *http.Request) {
	cookie, _ := ReadCookieServer(r)
	if cookie == "" {
		http.Redirect(w, r, "/login/", 302)
	} else {
		user, err := model.GetUser(cookie)
		if err != nil {
			http.Redirect(w, r, "/login/", 302)
		} else {
			Render(w, user, "admin", "default", "nav.private", "home")
		}
	}
}

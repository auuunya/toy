package controller

import (
	"net/http"
	"time"
)

func ReadCookieServer(r *http.Request) (string, error) {
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}

func WriteCookieServer(w http.ResponseWriter, value string, path string) {
	cookie := http.Cookie{
		Name:     "_cookie",
		Value:    value,
		MaxAge:   86400,
		HttpOnly: true,
		Path:     path,
	}
	http.SetCookie(w, &cookie)
}

func DeleteCookieServer(w http.ResponseWriter, value string, path string) {
	cookie := http.Cookie{
		Name:     "_cookie",
		HttpOnly: true,
		Value:    value,
		Path:     path,
		MaxAge:   -1,
		Expires:  time.Unix(1, 0),
	}
	http.SetCookie(w, &cookie)
}

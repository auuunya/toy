package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func InitializeStatic(theme string) {
	fs := http.FileServer(http.Dir(strings.Join([]string{"View", theme, "static"}, "/")))
	http.Handle(strings.Join([]string{"", theme, "static", ""}, "/"),
		http.StripPrefix(strings.Join([]string{"", theme, "static"}, "/"), fs))
}

func Render(w http.ResponseWriter, data interface{}, theme string, filenames ...string) {
	var files []string
	for _, filename := range filenames {
		files = append(files, fmt.Sprintf("View/"+theme+"/%s.html", filename))
	}
	tmp := template.Must(template.ParseFiles(files...))
	tmp.ExecuteTemplate(w, "default", data)
}

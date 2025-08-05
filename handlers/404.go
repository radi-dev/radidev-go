package handlers

import (
	"html/template"
)

func NotFound() (tmpl *template.Template) {
	tmpl = template.Must(template.ParseFiles("templates/404.html"))

	return

}

package handlers

import (
	"html/template"
)

func Homepage() (tmpl *template.Template) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		tmpl = NotFound()
	}

	return

}

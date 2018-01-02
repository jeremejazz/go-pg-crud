package lib

import "html/template"

type Page struct {
	Title       string
	Description string
}

var Tmpl = template.Must(template.ParseGlob("templates/*"))

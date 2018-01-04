package lib

import "html/template"

var Tmpl = template.Must(template.ParseGlob("templates/*"))

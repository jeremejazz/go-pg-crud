package lib

import (
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {

	page := &Page{Title: "Home"}
	Tmpl.ExecuteTemplate(w, "index", page)
}

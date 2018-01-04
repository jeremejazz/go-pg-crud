package lib

import (
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {

	p := &Page{
		Title: "Home",
		Persons: []*Person{
			&Person{
				Name:    "Timmy",
				Address: "22 b",
			},
			&Person{
				Name:    "JJ ABrams",
				Address: "In a Galaxy Far Far away",
			},
		},
	}

	Tmpl.ExecuteTemplate(w, "index", p)
}

func HandleList(w http.ResponseWriter, r *http.Request) {

}

func HandleNew(w http.ResponseWriter, r *http.Request) {
	//
}

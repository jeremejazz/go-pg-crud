package main

import (
	"fmt"
	"go-pg-crud/lib"
	"net/http"
)

func main() {
	fmt.Println("Starting server :3000")
	http.HandleFunc("/", lib.Index)

	http.ListenAndServe(":3000", nil)

}

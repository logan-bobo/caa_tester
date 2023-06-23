package main

import (

	"net/http"
)

func root(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "templates/index.html")
}

func main() {
	http.HandleFunc("/", root)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	

	http.ListenAndServe(":8090", nil)
}

package main

import (

	"net/http"
)

func home(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "templates/index.html")
}

func main() {
	http.HandleFunc("/", home)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/node_modules/", http.StripPrefix("/node_modules/", http.FileServer(http.Dir("node_modules"))))
	

	http.ListenAndServe(":8090", nil)
}

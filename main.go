package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl = template.Must(template.ParseFiles("index.html")) // creates the index.html page

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil) // display template
}

func main() {
	fmt.Println("App Started")

	mux := http.NewServeMux() // helps to call the correct handler based on the URL

	mux.HandleFunc("/", indexHandler) // what function to call on main page
	http.ListenAndServe(":3000", mux) // start a local server to run files
}
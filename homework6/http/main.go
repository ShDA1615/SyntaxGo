package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func StartPage(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "static/index.html")

}

func Hello1(w http.ResponseWriter, r *http.Request) {

	Name := r.FormValue("name")
	tmpl, _ := template.ParseFiles("static/hello.html")
	tmpl.Execute(w, Name)

}
func main() {

	http.HandleFunc("/", StartPage)
	http.HandleFunc("/hello", Hello1)
	fmt.Println("Сервер запущен на 8181 порте.")
	http.ListenAndServe(":8181", nil)

}

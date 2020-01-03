package main

import (
	"net/http"
	// "github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"html/template"
)

var templates *template.Template

func main(){
	r := mux.NewRouter()
	templates = template.Must(template.ParseGlob("templates/*.html"))
	r.HandleFunc("/", shortifyGet).Methods("GET")
	r.HandleFunc("/", shortifyPost).Methods("POST")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func shortifyGet(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
	return
}

func shortifyPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	url := r.PostForm.Get("url")
	w.Write([]byte(url))
	return
}
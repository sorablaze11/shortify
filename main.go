package main

import (
	"net/http"
	// "github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"html/template"
	"strings"
	"github.com/sorablaze11/shortify/kgs"
	// "fmt"
)

var templates *template.Template

func main(){
	kgs.Init()
	r := mux.NewRouter()
	templates = template.Must(template.ParseGlob("templates/*.html"))
	r.HandleFunc("/", ShortifyGet).Methods("GET")
	r.HandleFunc("/{shortUrl}", ShortUrlGet).Methods("GET")
	r.HandleFunc("/", ShortifyPost).Methods("POST")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func ShortUrlGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// fmt.Println(vars["shortUrl"])
	redirecUrl, err := kgs.GetUrl(vars["shortUrl"])
	if err != nil {
		templates.ExecuteTemplate(w, "error.html", nil)
	}else {
		hasHttps := strings.HasPrefix(redirecUrl, "https://")
		hasHttps = hasHttps || strings.HasPrefix(redirecUrl, "http://")
		if !hasHttps {
			redirecUrl = "https://" + redirecUrl
		}
		// fmt.Println(redirecUrl)
		http.Redirect(w, r, redirecUrl, 301)
	}
	return
}

func ShortifyGet(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
	return
}

func ShortifyPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	url := r.PostForm.Get("url")
	// fmt.Println(url)
	shortUrl := kgs.ReturnShortUrl(url);
	// fmt.Println(shortUrl)
	w.Write([]byte(shortUrl))
	return
}
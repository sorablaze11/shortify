package main

import (
	"net/http"
	// "github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/", index).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to our url shortner service"))
	return
}
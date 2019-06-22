package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kazu1029/go-jwt/auth"
)

type post struct {
	Title string `json:"title"`
	Tag   string `json:"tag"`
	URL   string `json:"url"`
}

func main() {
	r := mux.NewRouter()

	r.Handle("/public", public)
	r.Handle("/private", auth.JwtMiddleware.Handler(private))
	r.Handle("/auth", auth.GetTokenHandler)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("ListenAndServe:", nil)
	}
}

var public = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	post := &post{
		Title: "Sample Title",
		Tag:   "JWT",
		URL:   "https://qiita.com",
	}
	json.NewEncoder(w).Encode(post)
})

var private = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	post := &post{
		Title: "Title Private",
		Tag:   "Go",
		URL:   "https://qiita.com/timeline",
	}
	json.NewEncoder(w).Encode(post)
})

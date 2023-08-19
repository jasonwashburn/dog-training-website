package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jasonwashburn/dog-training-website/internal/handlers"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/register", handlers.RegisterUserHandler)
	r.HandleFunc("/login", handlers.GetLoginHandler).Methods("GET")
	r.HandleFunc("/login", handlers.PostLoginHandler).Methods("POST")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}

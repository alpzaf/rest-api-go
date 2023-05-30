package main

import (
	"log"
	"net/http"
	"rest-go/pkg/pkg/db"
	"rest-go/pkg/pkg/handlers"

	"github.com/gorilla/mux"
)

func main() {
	DB := db.Init()
	h := handlers.New(DB)
	router := mux.NewRouter()

	router.HandleFunc("/users", h.GetAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/users", h.AddUser).Methods(http.MethodPost)

	log.Println("API is running")
	http.ListenAndServe(":4000", router)
}
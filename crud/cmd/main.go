package main

import (
	"crud/pkg/db"
	"crud/pkg/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	DB := db.Init() //return by Init func &a ->> mem loca
	h := handlers.New(DB)

	router := mux.NewRouter()

	router.HandleFunc("/books", h.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/books", h.AddBook).Methods(http.MethodPost)
	router.HandleFunc("/book/{id}", h.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/book/{id}", h.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/book/{id}", h.DeleteBook).Methods(http.MethodDelete)

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)
}

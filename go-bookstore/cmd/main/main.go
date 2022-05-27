package main

import (
	"fmt"
	"go-bookstore/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("started at port 9010")
	log.Fatal(http.ListenAndServe(":9010", r))
}

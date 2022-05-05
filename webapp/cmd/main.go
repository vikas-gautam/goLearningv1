//https://golangbyexample.com/json-request-body-golang-http/

package main

import (
	"log"
	"net/http"
	"webapp/pkg/handlers"

	"github.com/gorilla/mux"
)

// //var bookings = make([]UserData, 0) //2:48   both works
// var bookings []UserData

func main() {

	//mock data

	router := mux.NewRouter()

	// http.HandleFunc("/", helloWorld)
	router.HandleFunc("/conference/name", handlers.ConferenceName).Methods(http.MethodGet)

	router.HandleFunc("/conference/tickets", handlers.AvailableTickets).Methods(http.MethodGet)

	router.HandleFunc("/conference/book-ticket", handlers.BookTicket).Methods(http.MethodPost)

	router.HandleFunc("/conference/getuserdata", handlers.GetUserData).Methods(http.MethodGet)

	router.HandleFunc("/conference/getuser/{id}", handlers.GetUser).Methods(http.MethodGet)

	log.Println("API is running!")

	log.Fatal(http.ListenAndServe(":9090", router))

}
//curl -X POST -H "Content-Type: application/json" \
//-d '{"FirstName":"abc","LastName":"abc","Email":"abc","UserTickets":"abc"}' \
//http://localhost:8080/conference/book-ticket

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const conferenceTickets = 50

//var remainingTickets uint = 50

type UserData struct {
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	Email           string `json:"email"`
	NumberOfTickets uint   `json:"numberOfTickets"`
}

var bookings = make([]UserData, 0)

func main() {

	//mock data

	route := mux.NewRouter()

	// http.HandleFunc("/", helloWorld)
	route.HandleFunc("/conference/name", conferenceName).Methods("GET")
	route.HandleFunc("/conference/ticket", availableTickets).Methods("GET")
	route.HandleFunc("/conference/book-ticket", bookTicket).Methods("POST")
	route.HandleFunc("/conference/getuserdata", getUserData).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", route))
}

func conferenceName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "welcome to GO conference")
	log.Println("serving /conference/name")

}

func availableTickets(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "we have total of %v tickets available\n", conferenceTickets)
	log.Println("serving /conference/tickets")

}

func getUserData(w http.ResponseWriter, r *http.Request) {

	bookings = append(bookings, UserData{FirstName: "vikas", LastName: "gautam", Email: "v@com", NumberOfTickets: 5})

	fmt.Fprintln(w, "Please enter your details via POST method  as per given response")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookings)

	log.Println("serving /conference/getUserData")

}

func bookTicket(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Please enter your details via POST method")

	log.Println("serving /conference/book-ticket")

}

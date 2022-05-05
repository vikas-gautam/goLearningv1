package handlers

import (
	"fmt"
	"log"
	"net/http"
)

const conferenceTickets = 10

func AvailableTickets(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "we have total of %v tickets available\n", conferenceTickets)
	log.Println("serving /conference/tickets")
}

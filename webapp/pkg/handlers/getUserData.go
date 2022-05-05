package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"webapp/pkg/mocks"
)

func GetUserData(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Please enter your details via POST method  as per given response")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mocks.Bookings)
	log.Println("serving /conference/getUserData")

}

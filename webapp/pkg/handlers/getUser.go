package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/pkg/mocks"
	"webapp/pkg/models"

	"github.com/gorilla/mux"
)


func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Gets params
	fmt.Println(params)

	for _, item := range mocks.Bookings {

		if strconv.Itoa(item.ID) == params["id"] {
			json.NewEncoder(w).Encode(item)
			return //very imp
		}
	}
	json.NewEncoder(w).Encode(&models.UserData{})
}

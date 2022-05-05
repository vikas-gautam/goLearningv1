package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"webapp/pkg/mocks"
	"webapp/pkg/models"
)

var remainingTickets uint = 10
var sortedList = []int{}
var randNumber int

func notContains(sortedList []int, randNumber int) bool {
	for _, n := range sortedList {
		if n == randNumber {
			return false
		}
	}
	return true
}

func bookingids() int {
	randNumber = rand.Intn(10)
	if len(sortedList) == 0 {
		sortedList = append(sortedList, randNumber)
		return randNumber
	} else if len(sortedList) == 9 {
		return -1
	}
	
	for {
		for _, num := range sortedList {
			fmt.Println(num)
			randNumber = rand.Intn(10)
			fmt.Printf("random number generated :%v\n", randNumber)
			if notContains(sortedList, randNumber) && randNumber != 0 {
				sortedList = append(sortedList, randNumber)
				fmt.Println(sortedList)
				//randNumber
				return randNumber

			} else {
				fmt.Println("sorted list not appended bcz num already exists")
				fmt.Println(randNumber)
			}
		}
	}
}

func BookTicket(w http.ResponseWriter, r *http.Request) {

	log.Println("serving /conference/book-ticket")

	var userdatafromcurl models.UserData
	_ = json.NewDecoder(r.Body).Decode(&userdatafromcurl)

	isValidNAme := len(userdatafromcurl.FirstName) >= 2 && len(userdatafromcurl.LastName) >= 2
	isValidEmail := strings.Contains(userdatafromcurl.Email, "@")
	isValidTicketNumber := userdatafromcurl.UserTickets > 0 && userdatafromcurl.UserTickets <= remainingTickets
    
	if isValidNAme && isValidEmail && isValidTicketNumber {
		remainingTickets = remainingTickets - userdatafromcurl.UserTickets

		// randNumber value has been returned by func
		randNumber := bookingids()
        if randNumber != -1{
			// assign randNumber to ID
			userdatafromcurl.ID = randNumber
			// userdatafromcurl will get ID value here
			mocks.Bookings = append(mocks.Bookings, userdatafromcurl)

			response := fmt.Sprintf(`{"message": "%v"}`, mocks.Bookings)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, "response from app  %v \n", response)

			w.Write([]byte(fmt.Sprintf(`{"message": "Hi %v %v. Thanks for booking the %v ticket and your booking id is %v for Go Conference. We will send out yur booking information to %v and remaining tickets are: %v"}`, userdatafromcurl.FirstName, userdatafromcurl.LastName, userdatafromcurl.UserTickets, userdatafromcurl.ID, userdatafromcurl.Email, remainingTickets)))

		} else {
			fmt.Println("booking ids exhausted")
			w.Write([]byte(fmt.Sprintln(`{"message": "booking ids exhausted"}`)))

		}
		
	} else if !isValidTicketNumber {
		//end program
		fmt.Println("Conference has been booked out, please try next time")
		w.Write([]byte(fmt.Sprintln(`{"message": "Hi . conference has been booked out"}`)))

	} else {

		fmt.Println("provided payload is not correct please reverify and request again")

	}
}

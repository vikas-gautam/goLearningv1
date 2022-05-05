package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets = 50

var conferenceName = "Go Conference" //immediately assigned value
var remainingTickets uint = 50

//var bookings = make([]map[string]string, 0)     // creating slice using make
var bookings = make([]UserData, 0) // empty list of UserData struct

type UserData struct { // create a type called UserData based on struct of firstNAme,lastName
	// infact u can create a custom type based on any other data type like string, int etc
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	greetingUsers()
	firstName, lastName, email, userTickets := getUserInput()
	isValidNAme, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)
	if isValidNAme && isValidEmail && isValidTicketNumber {

		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)
		//go doSomehtingelse

		firstNames := getFirstNames()

		fmt.Printf("List of bookings are %v \n", bookings)
		fmt.Printf("All the bookings with firstnames are %v\n: ", firstNames)
		fmt.Printf("Thankyou  %v have booked %v tickets and remaining tickets are %v \n", firstNames[0], userTickets, remainingTickets)

		if remainingTickets == 0 {
			//end program
			fmt.Println("Conference has been booked out, please try next time")
			//break
		}
	} else {
		if !isValidNAme {
			fmt.Println("name is too short")
		}
		if !isValidEmail {
			fmt.Println("email doesn't contain @")
		}

	}
	wg.Wait()
}

func greetingUsers() { // confName string means var confName string
	fmt.Printf("welcome to %v booking application\n", conferenceName)
	//fmt.Printf("conferenceName is %T, conferenceTickets is %T, remaingTickets is %T \n", conferenceName,  conferenceTickets, remaningTickets)
	fmt.Printf("we have total of %v tickets and %v tickets are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("get your tickets here to attend")
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter number of tickets you want to book: ")
	fmt.Scan(&userTickets)
	fmt.Print("Enter the eamil address: ")
	fmt.Scan(&email)

	return firstName, lastName, email, userTickets
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
		//var names = strings.Fields(booking)    // strings.Fields() - splits the strings with whitespaces
		//fmt.Printf("slice after using field separator: %v\n", names)
		//firstNames = append(firstNames, booking["firstName"])

		fmt.Printf("value of firstNames %v\n", firstNames)
	}
	return firstNames
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) (uint, []UserData) {
	remainingTickets = remainingTickets - userTickets

	//create a map for user || map is datatype so we create a variable

	//var userData = make(map[string]string)     // all keys have same datatype and all values have same datatype
	// make function will create empty map by wrapping with make()
	// nameOfthemap["keyname"] = value

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// // userTickets is uint so convert into string == "30"
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	return remainingTickets, bookings
}

//just a simulation

func sendTicket(userTickets uint, firstName string, lastName string, email string) {

	var ticket = fmt.Sprintf("%v ticket for %v %v\n", userTickets, firstName, lastName)
	fmt.Println("#######################")
	time.Sleep(10 * time.Second)
	fmt.Printf("sending ticket: %v to email address %v\n", ticket, email)
	fmt.Println("######################")
	wg.Done()
}

// variables and datatypes

//var userName string
//var userTickets uint
//ask the user for some details
//fmt.Scan(userName)   // here userName has got empty value, it didn't asked for value input
//fix is use "&"" pointer
//fmt.Print("Enter your first name: ")
//fmt.Scan(&userName)

//fmt.Print("Enter number of tickets you want to book: ")
//fmt.Scan(&userTickets)

//TILL NOW WE ONLY HAVE ONE USER DOING BOOKING AT ONE TIME and app exits BUT assume it is a web application and multiple USERS
// ARE DOING BOOKING SIMULTANEOUSLY

//loops in GO

//fmt.Printf("We have only have %v tickets left, u can't book %v tickets\n", remainingTickets, userTickets)
//fmt.Println("input data is invalid")

//isValidCity := city == Singapore || city == London   // postive check

// array and slices
// var bookings [50]string          // array type = combination of size and data type

// bookings[0] = firstName + " " + lastName

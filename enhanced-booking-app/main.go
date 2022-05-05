package main

import (
	"fmt"
//	"net/http"
	"os"
	"log"
	"io/ioutil"
	"strconv"
)

//type UserData struct {                 // create a type called UserData based on struct of firstNAme,lastName 
	                                   // infact u can create a custom type based on any other data type like string, int etc
// firstName string
// lastName string
// email string
// numberOfTickets uint
// }

var firstName string
var lastName string
var email string
var numberOfTickets uint

func main() {
	firstName, lastName, email, numberOfTickets := getUserInput()
	fName := createFile(firstName, lastName, email, numberOfTickets)
	readFile(fName)
}











//function defined here

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

func createFile(firstName string, lastName string, email string, userTickets uint) (string){
	var fileName string
	myFile, err := os.Create("userInput.txt")
	if err != nil {
		log.Fatal("ERROR! ", err)
	}
	log.Println("Empty file created successfully. ", myFile)

    defer myFile.Close()
	var numberOftickets = strconv.FormatUint(uint64(userTickets), 10)
   
   //numberOfTickets = strconv.FormatUint(uint64(userTickets), 10)
    
	if _, err := myFile.WriteString("FirstName: " + firstName + "\n" + "lastname: " + lastName + "\n" + "email: " + email + "\n" + "numberOftickets: " + numberOftickets + "\n"); err != nil {
        log.Fatal("ERROR ", err)
    }
	log.Println("File has been appended with data ", myFile)
	fileName  = myFile.Name()
	fmt.Printf("Name of file: %s\n", myFile.Name())
	return fileName
}

func readFile(fileName string) {
    
	data, err := ioutil.ReadFile(fileName)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("***** Appended data is:  ******\n %s", data)

}
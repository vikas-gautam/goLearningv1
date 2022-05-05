package handlers

import (
	"crud/pkg/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (h handler) AddBook(w http.ResponseWriter, r *http.Request) {
	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}
	var book models.Book
	json.Unmarshal(body, &book)

	fmt.Println(book)

	// Append to the Book mocks
	// book.Id = rand.Intn(100)
	// fmt.Println(book)

	// mocks.Books = append(mocks.Books, book)

	// fmt.Println(mocks.Books)

	// Append to the Books table

	if result := h.DB.Create(&book); result.Error != nil {
		fmt.Println(result.Error)
	}

	
	// Send a 201 created response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
	json.NewEncoder(w).Encode(book)

}

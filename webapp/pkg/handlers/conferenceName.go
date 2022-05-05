package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func ConferenceName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "welcome to GO conference")
	log.Println("serving API /conference/name")

}

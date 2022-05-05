package models

type UserData struct {
	ID          int    `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	UserTickets uint   `json:"userTickets"`
}
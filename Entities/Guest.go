package Entities

type Guest struct {
	GuestID   int    `json:"guest_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	PhoneNo   int    `json:"phone_no"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

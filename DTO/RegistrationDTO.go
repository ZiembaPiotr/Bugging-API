package DTO

type RegistrationDTO struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	PhoneNo   int    `json:"phone_no"`
	Password  string `json:"password"`
	IsAdmin   bool   `json:"is_admin"`
}

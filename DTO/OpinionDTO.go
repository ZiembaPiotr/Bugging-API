package DTO

type OpinionDTO struct {
	ReservationID int    `json:"reservation_id"`
	HotelID       int    `json:"hotel_id"`
	Email         string `json:"email"`
	Rate          int    `json:"rate"`
	Opinion       string `json:"opinion"`
}

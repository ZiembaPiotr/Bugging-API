package DTO

import "backend-BD/Entities"

type HotelDetailsDTO struct {
	HotelID int              `json:"hotel_id"`
	Name    string           `json:"name"`
	Country string           `json:"country"`
	City    string           `json:"city"`
	Address string           `json:"address"`
	Stars   int              `json:"stars"`
	Rooms   []Entities.Rooms `json:"rooms"`
}

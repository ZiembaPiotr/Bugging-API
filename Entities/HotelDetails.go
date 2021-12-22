package Entities

type HotelDetails struct {
	HotelID int    `gorm:"primaryKey" json:"hotel_id"`
	Name    string `json:"name"`
	Country string `json:"country"`
	City    string `json:"city"`
	Address string `json:"address"`
	Stars   int    `json:"stars"`
}

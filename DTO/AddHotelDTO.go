package DTO

type AddHotelDTO struct {
	AdminID int    `json:"admin_id"`
	Name    string `json:"name"`
	Country string `json:"country"`
	City    string `json:"city"`
	Address string `json:"address"`
	Stars   int    `json:"stars"`
}

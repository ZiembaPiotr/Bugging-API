package DTO

type AddRoomDTO struct {
	AdminID  int    `json:"admin_id"`
	Capacity int    `json:"capacity"`
	Photos   string `json:"photos"`
	Price    int    `json:"price"`
}

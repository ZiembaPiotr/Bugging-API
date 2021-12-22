package Entities

type Rooms struct {
	HotelID  int    `json:"hotel_ID"`
	RoomID   int    `json:"room_id"`
	Capacity int    `json:"capacity"`
	Photos   string `json:"photos"`
	Price    int    `json:"price"`
}

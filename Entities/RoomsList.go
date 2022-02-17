package Entities

type Rooms struct {
	RoomID   int          `gorm:"primaryKey" json:"room_id"`
	HotelID  int          `json:"hotel_id"`
	Hotel    HotelDetails `gorm:"foreignKey:HotelID" json:"hotel"`
	Capacity int          `json:"capacity"`
	Photos   string       `json:"photos"`
	Price    int          `json:"price"`
}

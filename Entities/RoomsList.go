package Entities

type Rooms struct {
	HotelID  int          `gorm:"primaryKey" json:"hotel_ID"`
	Hotel    HotelDetails `gorm:"foreignKey:HotelID" json:"hotel"`
	RoomID   int          `gorm:"primaryKey" json:"room_id"`
	Capacity int          `json:"capacity"`
	Photos   string       `json:"photos"`
	Price    int          `json:"price"`
}

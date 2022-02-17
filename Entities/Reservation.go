package Entities

type Reservation struct {
	ReservationID int    `gorm:"primaryKey" json:"reservation_id"`
	HotelID       int    `json:"hotel_id"`
	RoomID        int    `json:"room_id"`
	GuestID       int    `json:"guest_id"`
	Guest         Guest  `gorm:"foreignKey:GuestID" json:"guest"`
	CheckInDate   string `json:"check_in_date"`
	CheckOutDate  string `json:"check_out_date"`
	Opinion       string `gorm:"type:text" json:"opinion"`
	Mark          int    `json:"mark"`
	CreatedOn     string `json:"created_on"`
}

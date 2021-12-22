package Entities

type Reservation struct {
	ReservationID int    `gorm:"primaryKey" json:"reservation_id"`
	HotelID       int    `json:"hotel_id"`
	RoomID        int    `json:"room_id"`
	GuestID       int    `json:"guest_id"`
	CheckInDate   string `json:"check_in_date"`
	CheckOutDate  string `json:"check_out_date"`
	OpinionID     int    `json:"opinion_id"`
	Mark          int    `json:"mark"`
	CreatedOn     string `json:"created_on"`
}

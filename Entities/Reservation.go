package Entities

type Reservation struct {
	ReservationID int          `gorm:"primaryKey" json:"reservation_id"`
	HotelID       int          `json:"hotel_id"`
	Hotel         HotelDetails `gorm:"foreignKey:HotelID" json:"hotel"`
	RoomID        int          `json:"room_id"`
	Room          Rooms        `gorm:"foreignKey:RoomID" json:"room"`
	GuestID       int          `json:"guest_id"`
	Guest         Guest        `gorm:"foreignKey:GuestID" json:"guest"`
	CheckInDate   string       `json:"check_in_date"`
	CheckOutDate  string       `json:"check_out_date"`
	OpinionID     int          `json:"opinion_id"`
	Mark          int          `json:"mark"`
	CreatedOn     string       `json:"created_on"`
}

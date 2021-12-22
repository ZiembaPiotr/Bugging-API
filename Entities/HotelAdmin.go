package Entities

type HotelAdmin struct {
	HotelID      int    `gorm:"primaryKey" json:"hotel_id"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	CreatedOn    string `json:"created_on"`
	LastModifyOn string `json:"last_modify_on"`
	LastLoginIP  string `json:"last_login_ip"`
}

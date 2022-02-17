package Entities

type HotelAdmin struct {
	AdminID      int          `gorm:"primaryKey" json:"admin_id"`
	FirstName    string       `json:"first_name"`
	LastName     string       `json:"last_name"`
	Email        string       `json:"email"`
	Password     string       `json:"password"`
	PhoneNo      int          `json:"phone_no"`
	HotelID      int          `json:"hotel_id"`
	Hotel        HotelDetails `gorm:"foreignKey:HotelID" json:"hotel"`
	CreatedOn    string       `json:"created_on"`
	LastModifyOn string       `json:"last_modify_on"`
	LastLoginIP  string       `json:"last_login_ip"`
}

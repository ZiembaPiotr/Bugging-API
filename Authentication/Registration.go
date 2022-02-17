package Authentication

import (
	"backend-BD/Crypt"
	"backend-BD/DTO"
	"backend-BD/Database"
	"backend-BD/Entities"
	"encoding/json"
	"log"
	"net/http"
)

func Registration() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var newUser DTO.RegistrationDTO

		if err := json.NewDecoder(request.Body).Decode(&newUser); err != nil {
			writer.WriteHeader(400)
			if _, err := writer.Write([]byte("Wrong body!")); err != nil {
				log.Println(err)
			}
			log.Println(err)
			return
		}

		tableName := selectTable(&newUser)

		if alreadyExist := checkIfAlreadyExist(tableName, newUser.Email); alreadyExist {
			writer.WriteHeader(409)
			if _, err := writer.Write([]byte("User already exist")); err != nil {
				log.Println("User already exist")
			}
			return
		}

		db := Database.SetUpDBConnection()

		hashPassword, err := Crypt.HashPassword(newUser.Password)
		if err != nil {
			log.Println("Failed to hash password!")
			return
		}

		if tableName == "guests" {
			//query = `INSERT INTO guests(guest_id, first_name, last_name, phone_no, email, password) VALUES(?, ?, ?, ?, ?, ?)`

			db.Create(&Entities.Guest{
				FirstName: newUser.FirstName,
				LastName:  newUser.LastName,
				PhoneNo:   newUser.PhoneNo,
				Email:     newUser.Email,
				Password:  hashPassword,
			})

		} else {
			//query = `INSERT INTO guests(hotel_admin_id, email, password, phone_no, hotel_id,first_name, last_name created_on, last_modify_on, last_login_ip) VALUES(?, ?, ?, ?, ?, ?)`

			db.Create(&Entities.HotelAdmin{
				FirstName:    newUser.FirstName,
				LastName:     newUser.LastName,
				Email:        newUser.Email,
				Password:     hashPassword,
				PhoneNo:      newUser.PhoneNo,
				HotelID:      0,
				Hotel:        Entities.HotelDetails{},
				CreatedOn:    "",
				LastModifyOn: "",
				LastLoginIP:  request.RemoteAddr,
			})
		}

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(201)
		if _, err := writer.Write([]byte("Registration went successful")); err != nil {
			log.Println(err)
		}

	}
}

func checkIfAlreadyExist(tableName string, newEmail string) bool {
	var email string
	var query string

	db := Database.SetUpDBConnection()

	if tableName == "guests" {
		query = `SELECT email FROM guests WHERE email = ?`
	} else {
		query = `SELECT email FROM hotel_admins WHERE email = ?`
	}

	db.Raw(query, newEmail).Scan(&email)

	if email != "" {
		return true
	}

	return false
}

func selectTable(newUser *DTO.RegistrationDTO) string {
	if newUser.IsAdmin {
		return "hotel_admins"
	}

	return "guests"
}

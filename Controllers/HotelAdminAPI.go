package Controllers

import (
	"backend-BD/DTO"
	"backend-BD/Database"
	"backend-BD/Entities"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func RemoveAdmin() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		idToRemove := mux.Vars(request)["id"]

		idToRemoveInt, err := strconv.Atoi(idToRemove)
		if err != nil {
			log.Println(err)
		}

		if exist := checkIfAdminExists(idToRemoveInt); !exist {
			writer.WriteHeader(204)
			return
		}

		db := Database.SetUpDBConnection()

		query := `DELETE FROM hotel_admins WHERE admin_id = ?`

		db.Exec(query, idToRemove)

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(202)
	}
}

func AddHotel() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var newHotelDTO DTO.AddHotelDTO
		var newHotel Entities.HotelDetails

		if err := json.NewDecoder(request.Body).Decode(&newHotelDTO); err != nil {
			log.Println(err)
			return
		}

		db := Database.SetUpDBConnection()

		//query = `INSERT INTO hotel_details(hotel_id, name, country, city, address, stars) VALUES (?, ? , ?, ?, ?)`

		db.Create(&Entities.HotelDetails{
			Name:    newHotelDTO.Name,
			Country: newHotelDTO.Country,
			City:    newHotelDTO.City,
			Address: newHotelDTO.Address,
			Stars:   newHotelDTO.Stars,
		})

		query := `UPDATE hotel_admins SET hotel_id = (SELECT hotel_id FROM hotel_details ORDER BY hotel_id DESC LIMIT 1) WHERE admin_id = ?`

		db.Exec(query, newHotelDTO.AdminID)

		query = `SELECT * FROM hotel_details ORDER BY hotel_id DESC LIMIT 1`
		db.Raw(query).Scan(&newHotel)

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(200)
		if err := json.NewEncoder(writer).Encode(&newHotel); err != nil {
			log.Println(err)
		}
	}
}

func AddRoom() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var newRoomDTO DTO.AddRoomDTO
		var hotelID int

		if err := json.NewDecoder(request.Body).Decode(&newRoomDTO); err != nil {
			log.Println(err)
			return
		}

		db := Database.SetUpDBConnection()

		query := `SELECT hotel_id FROM hotel_admins WHERE admin_id = ?`

		db.Raw(query, newRoomDTO.AdminID).Scan(&hotelID)

		db.Create(&Entities.Rooms{
			HotelID:  hotelID,
			Hotel:    Entities.HotelDetails{},
			Capacity: newRoomDTO.Capacity,
			Photos:   newRoomDTO.Photos,
			Price:    newRoomDTO.Price,
		})

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(200)
		if err := json.NewEncoder(writer).Encode(&newRoomDTO); err != nil {
			log.Println(err)
		}
	}
}

func checkIfAdminExists(ID int) bool {
	var count int

	db := Database.SetUpDBConnection()

	query := `SELECT COUNT(*) FROM hotel_admins WHERE admin_id = ?`

	db.Raw(query, ID).Scan(&count)

	return count != 0
}

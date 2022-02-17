package Controllers

import (
	"backend-BD/DTO"
	"backend-BD/Database"
	"backend-BD/Entities"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HotelList() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var hotelList []Entities.HotelDetails

		db := Database.SetUpDBConnection()

		query := `SELECT * FROM hotel_details;`

		db.Raw(query).Scan(&hotelList)

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(200)
		if err := json.NewEncoder(writer).Encode(&hotelList); err != nil {
			log.Println(err)
		}
	}
}

func HotelDetails() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var hotelDetailsDTO DTO.HotelDetailsDTO
		var hotelDetails Entities.HotelDetails
		var roomsList []Entities.Rooms

		hotelID := mux.Vars(request)["id"]

		db := Database.SetUpDBConnection()

		hotelQuery := `SELECT * FROM hotel_details WHERE hotel_id = ?;`

		db.Raw(hotelQuery, hotelID).Scan(&hotelDetails)

		roomsQuery := `SELECT * FROM rooms WHERE hotel_id = ?;`

		db.Raw(roomsQuery, hotelID).Scan(&roomsList)

		hotelDetailsDTO.HotelID = hotelDetails.HotelID
		hotelDetailsDTO.Name = hotelDetails.Name
		hotelDetailsDTO.Country = hotelDetails.Country
		hotelDetailsDTO.City = hotelDetails.City
		hotelDetailsDTO.Address = hotelDetails.Address
		hotelDetailsDTO.Stars = hotelDetails.Stars
		hotelDetailsDTO.Rooms = roomsList

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(200)
		if err := json.NewEncoder(writer).Encode(&hotelDetailsDTO); err != nil {
			log.Println(err)
		}
	}
}

package Controllers

import (
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

func Hotel() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var hotel []Entities.HotelDetails

		hotelID := mux.Vars(request)["id"]

		db := Database.SetUpDBConnection()

		query := `SELECT * FROM hotel_details WHERE hotel_id = ?`

		db.Raw(query, hotelID).Scan(&hotel)

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(200)
		if err := json.NewEncoder(writer).Encode(&hotel); err != nil {
			log.Println(err)
		}
	}
}

func HotelRooms() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var roomsList []Entities.Rooms

		hotelID := mux.Vars(request)["id"]

		db := Database.SetUpDBConnection()

		roomsQuery := `SELECT * FROM rooms WHERE hotel_id = ?;`

		db.Raw(roomsQuery, hotelID).Scan(&roomsList)

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(200)
		if err := json.NewEncoder(writer).Encode(&roomsList); err != nil {
			log.Println(err)
		}
	}
}

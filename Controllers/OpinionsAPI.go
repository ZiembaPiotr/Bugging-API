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

func HotelOpinions() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var reservationList []Entities.Reservation
		var opinionDTO DTO.OpinionDTO
		var opinionList []DTO.OpinionDTO
		var email string

		hotelID := mux.Vars(request)["id"]
		hotelIDint, err := strconv.Atoi(hotelID)
		if err != nil {
			log.Println("Failed to convert")
		}

		db := Database.SetUpDBConnection()

		query := `SELECT r.*, ha.email FROM reservations r, hotel_admins ha WHERE r.hotel_id = ? AND ha.hotel_id = r.hotel_id;`

		db.Raw(query, hotelID).Scan(&reservationList)

		for _, reservation := range reservationList {

			query := `SELECT email FROM guests WHERE guest_id = ?;`

			db.Raw(query, reservation.GuestID).Scan(&email)

			log.Println(email)

			opinionDTO.Email = email
			opinionDTO.HotelID = hotelIDint
			opinionDTO.Opinion = reservation.Opinion
			opinionDTO.Rate = reservation.Mark
			opinionDTO.ReservationID = reservation.ReservationID

			opinionList = append(opinionList, opinionDTO)
		}

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(200)
		if err := json.NewEncoder(writer).Encode(&opinionList); err != nil {
			log.Println(err)
		}
	}
}

func AddNewOpinion() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var newOpinionDTO DTO.OpinionDTO
		var guestID int

		if err := json.NewDecoder(request.Body).Decode(&newOpinionDTO); err != nil {
			log.Println(err)
			return
		}

		db := Database.SetUpDBConnection()

		query := `SELECT guest_id FROM guests WHERE email = ?`

		db.Raw(query, newOpinionDTO.Email).Scan(&guestID)

		//if guestID ...

		db.Create(&Entities.Reservation{
			RoomID:       0,
			HotelID:      newOpinionDTO.HotelID,
			GuestID:      guestID,
			Guest:        Entities.Guest{},
			CheckInDate:  "",
			CheckOutDate: "",
			Opinion:      newOpinionDTO.Opinion,
			Mark:         newOpinionDTO.Rate,
			CreatedOn:    "",
		})

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(200)
		if err := json.NewEncoder(writer).Encode(&newOpinionDTO); err != nil {
			log.Println(err)
		}
	}
}

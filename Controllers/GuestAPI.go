package Controllers

import (
	"backend-BD/Database"
	"backend-BD/Entities"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func GetAllGuests() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		db := Database.SetUpDBConnection()

		query := `SELECT * FROM guests;`

		var guests []Entities.Guest

		db.Raw(query).Scan(&guests)

		for _, guest := range guests {
			guest.Password = "********************"
		}

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(200)
		if err := json.NewEncoder(writer).Encode(guests); err != nil {
			writer.WriteHeader(404)
			log.Fatal(err)
		}
	}
}

func GetGuestById() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var guest Entities.Guest

		db := Database.SetUpDBConnection()

		searchedID := mux.Vars(request)["id"]

		query := `SELECT * FROM guests WHERE guest_id = ?`

		db.Raw(query, searchedID).Scan(&guest)

		guest.Password = "*****************"

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(200)
		if err := json.NewEncoder(writer).Encode(guest); err != nil {
			log.Println("Failed to jsonfy user")
			return
		}
	}
}

func DeleteGuest() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		idToRemove := mux.Vars(request)["id"]

		idToRemoveInt, err := strconv.Atoi(idToRemove)
		if err != nil {
			log.Println(err)
		}

		if exist := checkIfExists(idToRemoveInt); !exist {
			writer.WriteHeader(204)
			return
		}

		db := Database.SetUpDBConnection()

		query := `DELETE FROM guests WHERE guest_id = ? ;`

		db.Exec(query, idToRemoveInt)

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(202)
	}
}

func checkIfExists(ID int) bool {
	var count int

	db := Database.SetUpDBConnection()

	query := `SELECT COUNT(*) FROM guests WHERE guest_id = ?`

	db.Raw(query, ID).Scan(&count)

	return count != 0
}

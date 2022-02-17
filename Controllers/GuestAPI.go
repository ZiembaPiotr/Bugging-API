package Controllers

import (
	"backend-BD/Database"
	"backend-BD/Entities"
	"encoding/json"
	"log"
	"net/http"
)

func GetAllGuests() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db := Database.SetUpDBConnection()

		query := `SELECT * FROM guests;`

		var guests []Entities.Guest

		db.Raw(query).Scan(&guests)

		if err := json.NewEncoder(w).Encode(guests); err != nil {
			w.WriteHeader(404)
			log.Fatal(err)
		}
	}
}

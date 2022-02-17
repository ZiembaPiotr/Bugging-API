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

func LogInGuests() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var loggedUser DTO.LogInDTO
		var guest Entities.Guest

		if err := json.NewDecoder(request.Body).Decode(&loggedUser); err != nil {
			log.Println(err)
			return
		}

		if loggedUser.IsAdmin {
			writer.WriteHeader(401)
			if _, err := writer.Write([]byte("Unauthorized")); err != nil {
				log.Println(err)
			}
			return
		}

		db := Database.SetUpDBConnection()

		query := `SELECT * FROM guests WHERE email = ?`

		db.Raw(query, loggedUser.Email).Scan(&guest)

		if isValid := Crypt.CheckPasswordHash(loggedUser.Password, guest.Password); !isValid {
			writer.WriteHeader(401)
			if _, err := writer.Write([]byte("Wrong password")); err != nil {
				log.Println(err)
			}
			log.Println("Wrong password")
			return
		}

		guest.Password = "****************"

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(200)
		if err := json.NewEncoder(writer).Encode(guest); err != nil {
			log.Println(err)
		}
	}
}

func LogInAdmin() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var loggedUser DTO.LogInDTO
		var admin Entities.Guest

		if err := json.NewDecoder(request.Body).Decode(&loggedUser); err != nil {
			log.Println(err)
			return
		}

		if !loggedUser.IsAdmin {
			writer.WriteHeader(401)
			if _, err := writer.Write([]byte("Unauthorized")); err != nil {
				log.Println(err)
			}
			return
		}

		db := Database.SetUpDBConnection()

		query := `SELECT * FROM hotel_admins WHERE email = ?`

		db.Raw(query, loggedUser.Email).Scan(&admin)

		if isValid := Crypt.CheckPasswordHash(loggedUser.Password, admin.Password); !isValid {
			writer.WriteHeader(401)
			if _, err := writer.Write([]byte("Wrong password")); err != nil {
				log.Println(err)
			}
			log.Println("Wrong password")
			return
		}

		admin.Password = "****************"

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(200)
		if err := json.NewEncoder(writer).Encode(admin); err != nil {
			log.Println(err)
		}
	}
}

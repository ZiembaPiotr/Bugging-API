package Router

import (
	"backend-BD/Authentication"
	"backend-BD/Controllers"
	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter()

	authenticationRouter(router)
	guestsRouter(router)
	adminRouter(router)
	hotelDetailsRouter(router)
	opinionsRouter(router)

	return router
}

func authenticationRouter(router *mux.Router) {
	authenticationRouter := router.PathPrefix("/authentication").Subrouter()

	authenticationRouter.HandleFunc("/registration", Authentication.Registration()).Methods("POST")
	authenticationRouter.HandleFunc("/log-in-guest", Authentication.LogInGuests()).Methods("POST")
	authenticationRouter.HandleFunc("/log-in-admin", Authentication.LogInAdmin()).Methods("POST")
}

func guestsRouter(router *mux.Router) {
	guestsRouter := router.PathPrefix("/guests").Subrouter()

	guestsRouter.HandleFunc("/get-all", Controllers.GetAllGuests()).Methods("GET")
	guestsRouter.HandleFunc("/get-by-id/{id}", Controllers.GetGuestById()).Methods("GET")
	guestsRouter.HandleFunc("/delete/{id}", Controllers.DeleteGuest()).Methods("DELETE")
}

func adminRouter(router *mux.Router) {
	adminRouter := router.PathPrefix("/admins").Subrouter()

	adminRouter.HandleFunc("/remove/{id}", Controllers.RemoveAdmin()).Methods("DELETE")
	adminRouter.HandleFunc("/add-hotel", Controllers.AddHotel()).Methods("POST")
	adminRouter.HandleFunc("/add-room", Controllers.AddRoom()).Methods("POST")
}

func hotelDetailsRouter(router *mux.Router) {
	hotelDetails := router.PathPrefix("/hotel-details").Subrouter()

	hotelDetails.HandleFunc("/all", Controllers.HotelList()).Methods("GET")
	hotelDetails.HandleFunc("/{id}", Controllers.Hotel()).Methods("GET")
	hotelDetails.HandleFunc("/hotel/{id}", Controllers.HotelRooms()).Methods("GET")
}

func opinionsRouter(router *mux.Router) {
	opinionsRouter := router.PathPrefix("/opinions").Subrouter()

	opinionsRouter.HandleFunc("/{id}", Controllers.HotelOpinions()).Methods("GET")
	opinionsRouter.HandleFunc("/add-new", Controllers.AddNewOpinion()).Methods("POST")
}

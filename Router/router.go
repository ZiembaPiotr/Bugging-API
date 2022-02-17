package Router

import (
	"backend-BD/Authentication"
	"backend-BD/Controllers"
	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter()

	authenticationRouter(router)

	return router
}

func authenticationRouter(router *mux.Router) {
	authenticationRouter := router.PathPrefix("/authentication").Subrouter()

	authenticationRouter.HandleFunc("/registration", Authentication.Registration()).Methods("POST")
	authenticationRouter.HandleFunc("/log-in-guest", Authentication.LogInGuests()).Methods("POST")
	authenticationRouter.HandleFunc("/log-in-admin", Authentication.LogInAdmin()).Methods("POST")
}

func guestsRouter(router *mux.Router) {
	guestsRouter := router.PathPrefix("guests").Subrouter()

	guestsRouter.HandleFunc("/get-all", Controllers.GetAllGuests()).Methods("GET")
}

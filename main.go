package main

import (
	"backend-BD/Database"
	"backend-BD/Router"
	"log"
	"net/http"
)

func main() {
	r := Router.CreateRouter()

	connTest := Database.SetUpDBConnection()

	connTest.Database.Raw("SELECT id, name, age FROM users WHERE name = 3")

	if err := http.ListenAndServe(":2115", r); err != nil {
		log.Fatal(err)
	}
}

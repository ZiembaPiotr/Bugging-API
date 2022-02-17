package main

import (
	"backend-BD/Router"
	"log"
	"net/http"
)

func main() {
	r := Router.CreateRouter()
	
	if err := http.ListenAndServe(":2115", r); err != nil {
		log.Fatal(err)
	}
}

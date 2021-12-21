package Router

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/test", Test())

	return router
}

func Test() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
		if err != nil {
			log.Fatal(err)
		}
	}
}

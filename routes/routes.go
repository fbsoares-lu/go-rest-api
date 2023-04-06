package routes

import (
	"log"
	"net/http"
	"rest-go-api/controllers"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalities", controllers.FindAll).Methods("Get")
	r.HandleFunc("/api/personalities/{id}", controllers.FindOne).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", r))
}

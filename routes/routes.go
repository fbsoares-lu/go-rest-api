package routes

import (
	"log"
	"net/http"
	"rest-go-api/controllers"
	"rest-go-api/middleware"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalities", controllers.FindAll).Methods("Get")
	r.HandleFunc("/api/personalities/{id}", controllers.FindOne).Methods("Get")
	r.HandleFunc("/api/personalities", controllers.Create).Methods("Post")
	r.HandleFunc("/api/personalities/{id}", controllers.Delete).Methods("Delete")
	r.HandleFunc("/api/personalities/{id}", controllers.Update).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}

package controllers

import (
	"encoding/json"
	"net/http"
	"rest-go-api/models"
	"strconv"

	"github.com/gorilla/mux"
)

func FindAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalities)
}

func FindOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, personality := range models.Personalities {
		if strconv.Itoa(personality.Id) == id {
			json.NewEncoder(w).Encode(personality)
		}
	}
}
